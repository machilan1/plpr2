package sess

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/sessions"
)

func TestCleanup(t *testing.T) {
	t.Parallel()

	m, err := NewManager(
		newMockStore(),
		sessions.Options{},
		[]byte(secret),
	)
	if err != nil {
		t.Fatalf("failed to create manager: %v", err)
	}
	// Start the cleanup goroutine.
	defer m.StopCleanup(m.Cleanup(time.Millisecond * 500))

	r, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	session, err := m.New(r, testSessionID)
	if err != nil {
		t.Fatalf("m.New() error = %v", err)
	}

	// Expire the session.
	session.Options.MaxAge = 1 // 1 second

	// Save the session.
	rr := httptest.NewRecorder()
	if err := m.Save(r, rr, session); err != nil {
		t.Fatalf("m.Save() error = %v", err)
	}
	res := rr.Result()
	defer res.Body.Close()
	cookies := res.Cookies()
	rr.Result().Body.Close()
	if v := len(cookies); v != 1 {
		t.Fatalf("expected 1 cookie, got %d", v)
	}
	cookie := cookies[0]

	// Wait for the cleanup goroutine to run.
	time.Sleep(time.Millisecond * 2000)

	// Attempt to fetch the session.
	r.AddCookie(cookie)
	session, err = m.Get(r, testSessionID)
	if err != nil {
		t.Fatalf("m.Get() error = %v", err)
	}
	if !session.IsNew {
		t.Fatal("expected new session")
	}
}
