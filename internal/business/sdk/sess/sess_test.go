package sess

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/sessions"
)

const secret = "EyaC2BPcJtNqU3tjEHy+c+Wmqc1yihYIbUWEl/jk0Ga73kWBclmuSFd9HuJKwJw/Wdsh1XnjY2Bw1HBVph6WOw==" // nolint: gosec

const testSessionID = "test_session"

func TestLifecycle(t *testing.T) {
	t.Parallel()

	m, err := NewManager(
		newMockStore(),
		sessions.Options{
			MaxAge: 86400 * 30,
		},
		[]byte(secret),
	)
	if err != nil {
		t.Fatalf("failed to create manager: %v", err)
	}

	r, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	var session *sessions.Session
	var cookie *http.Cookie

	// New session
	{
		session, err = m.New(r, testSessionID)
		if err != nil {
			t.Fatalf("m.New() error = %v", err)
		}
		if !session.IsNew {
			t.Fatal("expected new session")
		}
	}

	// Save session
	{
		rr := httptest.NewRecorder()
		if err := m.Save(r, rr, session); err != nil {
			t.Fatalf("m.Save() error = %v", err)
		}
		res := rr.Result()
		defer res.Body.Close()
		cookies := res.Cookies()
		if v := len(cookies); v != 1 {
			t.Fatalf("expected 1 cookie, got %d", v)
		}
		cookie = cookies[0]
	}

	// Existing session
	{
		r.AddCookie(cookie)
		got, err := m.New(r, testSessionID)
		if err != nil {
			t.Fatalf("m.New() error = %v", err)
		}
		if got.IsNew {
			t.Fatalf("expected existing session, got new")
		}
	}

	// Delete session
	{
		session.Options.MaxAge = -1

		rr := httptest.NewRecorder()
		if err := m.Save(r, rr, session); err != nil {
			t.Fatalf("m.Save() error = %v", err)
		}
		res := rr.Result()
		defer res.Body.Close()
		cookies := res.Cookies()
		if v := len(cookies); v != 1 {
			t.Fatalf("expected 1 cookie, got %d", v)
		}
		if cookie := cookies[0]; cookie.Value != "" {
			t.Fatalf("expected cookie to be cleared, got %s", cookie.Value)
		}
	}

	// Verify session is deleted
	{
		got, err := m.New(r, testSessionID)
		if err != nil {
			t.Fatalf("m.New() error = %v", err)
		}
		if !got.IsNew {
			t.Fatal("expected new session, got existing")
		}
	}
}

func TestSessionOptionsAreUniquePerSession(t *testing.T) {
	t.Parallel()

	m, err := NewManager(
		newMockStore(),
		sessions.Options{
			MaxAge: 86400 * 30,
		},
		[]byte(secret),
	)
	if err != nil {
		t.Fatalf("failed to create manager: %v", err)
	}

	r, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	session, err := m.Get(r, testSessionID)
	if err != nil {
		t.Fatalf("failed to create session: %v", err)
	}
	session.Options.MaxAge = -1

	if m.options.MaxAge != 86400*30 {
		t.Fatal("expected session options to be unique per session")
	}
}

// ============================================================================================

type mockStore struct {
	mu sync.Mutex
	m  map[string]Session
}

func newMockStore() *mockStore {
	return &mockStore{
		m: make(map[string]Session),
	}
}

var _ Storer = (*mockStore)(nil)

func (ms *mockStore) QueryByID(_ context.Context, id string) (Session, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	s, ok := ms.m[id]
	if !ok {
		return Session{}, ErrNotFound
	}

	return s, nil
}

func (ms *mockStore) Create(_ context.Context, s Session) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.m[s.ID] = s
	return nil
}

func (ms *mockStore) Update(_ context.Context, s Session) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.m[s.ID] = s
	return nil
}

func (ms *mockStore) Delete(_ context.Context, id string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	delete(ms.m, id)
	return nil
}

func (ms *mockStore) DeleteExpired(_ context.Context) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	for k, v := range ms.m {
		if v.ExpiresAt.Before(time.Now()) {
			delete(ms.m, k)
		}
	}

	return nil
}
