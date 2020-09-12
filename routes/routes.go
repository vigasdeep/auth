package routes

import (
	"auth/controllers"
	"auth/utils/auth"
	"net/http"

	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)
	r.Methods("OPTIONS").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	r.HandleFunc("/", controllers.TestAPI).Methods("GET")
	r.HandleFunc("/api", controllers.TestAPI).Methods("GET")
	r.HandleFunc("/register", controllers.CreateUser).Methods("OPTIONS")

	r.HandleFunc("/register", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("OPTIONS")

	r.HandleFunc("/login", controllers.Login).Methods("POST")

	// Auth route
	s := r.PathPrefix("/auth").Subrouter()
	s.Use(auth.JwtVerify)
	// s.HandleFunc("/user", controllers.FetchUsers).Methods("OPTIONS")

	s.HandleFunc("/user", controllers.FetchUsers).Methods("GET")
	s.HandleFunc("/request_subdomain", controllers.RequestSubdomain).Methods("POST")
	s.HandleFunc("/subdomains", controllers.FetchSubdomains).Methods("GET")
	s.HandleFunc("/subdomain/{id}", controllers.DeleteSubdomain).Methods("DELETE")

	s.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	s.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	s.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")

	a := r.PathPrefix("/admin").Subrouter()
	a.Use(auth.JwtVerifyAdmin)
	a.HandleFunc("/approve_subdomain", controllers.RequestSubdomain).Methods("POST")
	return r
}

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
