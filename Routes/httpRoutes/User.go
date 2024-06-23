package httproutes

import (
	handlers "TEST_SERVER/Handlers"
	"TEST_SERVER/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func UserRoutes(app *mux.Router) {
	r := app.PathPrefix("/user").Subrouter()
	r.HandleFunc("/signIn", handlers.SignIn).Methods(http.MethodPost)
	r.HandleFunc("/register", handlers.SignUp)
	// r.HandleFunc("/verify/admin/{token}", handlers.VerifyAdmin).Methods(http.MethodGet)
	// r.HandleFunc("/recover", handlers.RecoverAcc).Methods(http.MethodPost)
	r.HandleFunc("/verify/email", handlers.VerifyEmail).Methods(http.MethodPost)

	routerA := app.PathPrefix("/users/auth").Subrouter()
	routerA.Use(middlewares.Auth)
	// routerA.HandleFunc("/create/one", handlers.CreateUser).Methods("POST")
	// r.HandleFunc("/register", svc.RegisterUser).Methods(http.MethodPost)
	// r.HandleFunc("/find", svc.FindUser).Methods(http.MethodPost)
	// r.HandleFunc("/update", svc.UpdateUser).Methods(http.MethodPost)
	// r.HandleFunc("/findAll", svc.FindAllUsers).Methods(http.MethodGet)
}
