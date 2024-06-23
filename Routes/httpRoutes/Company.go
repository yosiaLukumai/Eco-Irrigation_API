package httproutes

import (
	handlers "TEST_SERVER/Handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func CompanyRoutes(app *mux.Router) {
	r := app.PathPrefix("/company").Subrouter()
	r.HandleFunc("/signUp", handlers.SignUp).Methods(http.MethodPost)

	// adding the auth
	r.HandleFunc("/add/role", handlers.AddRoleCompany).Methods(http.MethodPost)
	r.HandleFunc("/add/system", handlers.AddPump).Methods(http.MethodPost)
	r.HandleFunc("/add/package", handlers.AddPackage).Methods(http.MethodPost)
	r.HandleFunc("/find/packages/names", handlers.FindPackagesNames).Methods(http.MethodGet)
	r.HandleFunc("/find/systems", handlers.FindSystems).Methods(http.MethodPost)
	r.HandleFunc("/find/unassigned/systems", handlers.FindUnassigned).Methods(http.MethodGet)

}
