// Code originated from github.com/antonlindstrom/pgstore
//
// Copyright 2017 Anton Lindström
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is furnished to do
// so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package sess

import (
	"context"
	"encoding/base32"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var ErrNotFound = errors.New("session not found")

type Storer interface {
	QueryByID(ctx context.Context, id string) (Session, error)
	Create(ctx context.Context, s Session) error
	Update(ctx context.Context, s Session) error
	Delete(ctx context.Context, id string) error
	DeleteExpired(ctx context.Context) error
}

type Manager struct {
	storer  Storer
	codecs  []securecookie.Codec
	options sessions.Options
}

func NewManager(storer Storer, opts sessions.Options, keyPairs ...[]byte) (*Manager, error) {
	// if no key pairs are provided, we can't encrypt the session data
	if len(keyPairs) == 0 || len(keyPairs[0]) == 0 {
		return nil, errors.New("missing key pairs")
	}

	m := &Manager{
		storer:  storer,
		codecs:  securecookie.CodecsFromPairs(keyPairs...),
		options: opts,
	}

	m.MaxAge(m.options.MaxAge)

	return m, nil
}

var _ sessions.Store = (*Manager)(nil)

// Get Fetches a session for a given name after it has been added to the
// registry.
func (m *Manager) Get(r *http.Request, name string) (*sessions.Session, error) {
	return sessions.GetRegistry(r).Get(m, name)
}

// GetContext is a wrapper around Get that can be used with context.Context.
func (m *Manager) GetContext(ctx context.Context, r *http.Request, name string) (*sessions.Session, error) {
	r = r.WithContext(ctx)
	return m.Get(r, name)
}

// New returns a new session for the given name without adding it to the registry.
func (m *Manager) New(r *http.Request, name string) (*sessions.Session, error) {
	session := sessions.NewSession(m, name)

	opts := m.options
	session.Options = &opts
	session.IsNew = true

	co, err := r.Cookie(name)
	if err != nil {
		return session, nil
	}

	err = securecookie.DecodeMulti(name, co.Value, &session.ID, m.codecs...)
	if err != nil {
		return session, nil
	}

	ctx := r.Context()
	err = m.load(ctx, session)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return session, nil
		}
		return session, fmt.Errorf("load: %w", err)
	}

	session.IsNew = false

	return session, nil
}

// NewContext is a wrapper around New that can be used with context.Context.
func (m *Manager) NewContext(ctx context.Context, r *http.Request, name string) (*sessions.Session, error) {
	r = r.WithContext(ctx)
	return m.New(r, name)
}

// Save saves the given session into the database and deletes cookies if needed
func (m *Manager) Save(r *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	ctx := r.Context()

	// Set delete if max-age is < 0
	if session.Options.MaxAge < 0 {
		if err := m.storer.Delete(ctx, session.ID); err != nil {
			return fmt.Errorf("delete: sessionID[%s]: %w", session.ID, err)
		}
		http.SetCookie(w, sessions.NewCookie(session.Name(), "", session.Options))
		return nil
	}

	if session.ID == "" {
		// Generate a random session ID key suitable for storage in the DB
		session.ID = strings.TrimRight(
			base32.StdEncoding.EncodeToString(
				securecookie.GenerateRandomKey(32),
			), "=")
	}

	if err := m.save(ctx, session); err != nil {
		return fmt.Errorf("unable to save session to the database, err: %w", err)
	}

	// Keep the session ID key in a cookie so that it can be looked up in DB later.
	encoded, err := securecookie.EncodeMulti(session.Name(), session.ID, m.codecs...)
	if err != nil {
		return fmt.Errorf("unable to encode session ID, err: %w", err)
	}

	http.SetCookie(w, sessions.NewCookie(session.Name(), encoded, session.Options))
	return nil
}

// SaveContext is a wrapper around Save that can be used with context.Context.
func (m *Manager) SaveContext(ctx context.Context, w http.ResponseWriter, r *http.Request, session *sessions.Session) error {
	r = r.WithContext(ctx)
	return m.Save(r, w, session)
}

// MaxLength restricts the maximum length of new sessions to l.
// If l is 0 there is no limit to the size of a session, use with caution.
// The default for a new PGStore is 4096. PostgreSQL allows for max
// value sizes of up to 1GB (http://www.postgresql.org/docs/current/interactive/datatype-character.html)
func (m *Manager) MaxLength(l int) {
	for _, c := range m.codecs {
		if codec, ok := c.(*securecookie.SecureCookie); ok {
			codec.MaxLength(l)
		}
	}
}

// MaxAge sets the maximum age for the store and the underlying cookie
// implementation. Individual sessions can be deleted by setting Options.MaxAge
// = -1 for that session.
func (m *Manager) MaxAge(age int) {
	m.options.MaxAge = age

	// Set the maxAge for each securecookie instance.
	for _, codec := range m.codecs {
		if sc, ok := codec.(*securecookie.SecureCookie); ok {
			sc.MaxAge(age)
		}
	}
}

// Delete removes the session from the database and deletes the session cookie.
func (m *Manager) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request, session *sessions.Session) error {
	session.Options.MaxAge = -1
	return m.SaveContext(ctx, w, r, session)
}

// load fetches a session by ID from the database and decodes its content
// into session.Values.
func (m *Manager) load(ctx context.Context, session *sessions.Session) error {
	sess, err := m.storer.QueryByID(ctx, session.ID)
	if err != nil {
		return fmt.Errorf("query: id[%s]: %w", session.ID, err)
	}

	err = securecookie.DecodeMulti(session.Name(), sess.Data, &session.Values, m.codecs...)
	if err != nil {
		return fmt.Errorf("unable to decode session data, err: %w", err)
	}

	return nil
}

// save writes encoded session.Values to a database record.
// writes to http_sessions table by default.
func (m *Manager) save(ctx context.Context, session *sessions.Session) error {
	encoded, err := securecookie.EncodeMulti(session.Name(), session.Values, m.codecs...)
	if err != nil {
		return err
	}

	crAt := session.Values["created_at"]
	exAt := session.Values["expires_at"]

	var expiresAt time.Time

	createdAt, ok := crAt.(time.Time)
	if !ok {
		createdAt = time.Now()
	}

	if exAt == nil {
		expiresAt = time.Now().Add(time.Second * time.Duration(session.Options.MaxAge))
	} else {
		expiresAt = exAt.(time.Time)
		if expiresAt.Sub(time.Now().Add(time.Second*time.Duration(session.Options.MaxAge))) < 0 {
			expiresAt = time.Now().Add(time.Second * time.Duration(session.Options.MaxAge))
		}
	}

	sess := Session{
		ID:        session.ID,
		Data:      encoded,
		CreatedAt: createdAt,
		UpdatedAt: time.Now(),
		ExpiresAt: expiresAt,
	}

	if session.IsNew {
		if err := m.storer.Create(ctx, sess); err != nil {
			return fmt.Errorf("create: %w", err)
		}
		return nil
	}

	return m.storer.Update(ctx, sess)
}
