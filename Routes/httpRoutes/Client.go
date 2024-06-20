package httproutes

import (
	handlers "TEST_SERVER/Handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func ClientRoutes(app *mux.Router) {
	r := app.PathPrefix("/client").Subrouter()

	r.HandleFunc("/register", handlers.RegisterClient).Methods(http.MethodPost)
	r.HandleFunc("/find", handlers.FindFarmers).Methods(http.MethodPost)
	// r.HandleFunc("/find", svc.FindClient).Methods(http.MethodPost)
	// r.HandleFunc("/update", svc.UpdateClient).Methods(http.MethodPost)
	// r.HandleFunc("/findAll", svc.FindAllClients).Methods(http.MethodGet)
	// r.HandleFunc("/delete", svc.Delete).Methods(http.MethodPost)
	// r.HandleFunc("/deleteAll", svc.DeleteAll).Methods(http.MethodGet)
	// r.HandleFunc("/signIn", svc.SignIn).Methods(http.MethodPost)
}
