package debug

import (
	"expvar"
	"net/http"
	"net/http/pprof"
	"strings"

	"github.com/arl/statsviz"
)

// Mux registers all the debug routes from the standard library into a new mux
// bypassing the use of the DefaultServerMux. Using the DefaultServerMux would
// be a security risk since a dependency could inject a handler into our service
// without us knowing it.
func Mux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, cleanPath(r.RequestURI+"/debug/pprof/"), http.StatusMovedPermanently)
	})

	mux.HandleFunc("GET /debug/pprof/", pprof.Index)
	mux.HandleFunc("GET /debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("GET /debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("GET /debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("GET /debug/pprof/trace", pprof.Trace)
	mux.Handle("GET /debug/vars/", expvar.Handler())

	mux.Handle("GET /debug/pprof/goroutine", pprof.Handler("goroutine"))
	mux.Handle("GET /debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	mux.Handle("GET /debug/pprof/mutex", pprof.Handler("mutex"))
	mux.Handle("GET /debug/pprof/heap", pprof.Handler("heap"))
	mux.Handle("GET /debug/pprof/block", pprof.Handler("block"))
	mux.Handle("GET /debug/pprof/allocs", pprof.Handler("allocs"))

	_ = statsviz.Register(mux)

	return mux
}

// cleanPath removes the double slashes from the path to make sure the path is
// clean. Otherwise, the path might not match the route.
func cleanPath(path string) string {
	if strings.HasPrefix(path, "//") {
		return path[1:]
	}
	return path
}
