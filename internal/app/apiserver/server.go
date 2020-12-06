package apiserver

import (
	"context"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	// sessionName        = "gopherschool"
	// ctxKeyUser ctxKey = iota
	// ctxKeyRequestID

	ctxKeyRequestID = iota
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  interface{}
	// store        store.Store
	sessionStore interface{}
	// sessionStore sessions.Store
}

// type ctxKey int8

func newServer(store interface{}, sessionStore interface{}) *server {
	s := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        store,
		sessionStore: sessionStore,
	}

	fmt.Println("Server configs established")

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)

	staticDir := "/assets/"
	s.router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	s.router.HandleFunc("/", s.handleIndex()).Methods(http.MethodGet)
	// s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	// s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	// s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")

	// private := s.router.PathPrefix("/private").Subrouter()
	// private.Use(s.authenticateUser)
	// private.HandleFunc("/whoami", s.handleWhoami()).Methods("GET")
}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		next.ServeHTTP(w, r)

		logger.Infof("completed in %v", time.Now().Sub(start))
	})
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		// myvar := map[string]interface{}{"MyHeader": "this is index page"}

		t, _ := template.ParseFiles("./views/layout.html", "./views/index.html")

		t.ExecuteTemplate(w, "layout", struct {
			MyHeader string
		}{
			MyHeader: "this is index page",
		})
	}
}
