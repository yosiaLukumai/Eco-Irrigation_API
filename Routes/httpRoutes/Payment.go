package httproutes

import (
	"net/http"

	handlers "TEST_SERVER/Handlers"

	"github.com/gorilla/mux"
)

func PaymentRoutes(app *mux.Router) {
	r := app.PathPrefix("/payment").Subrouter()
	r.HandleFunc("/payment", handlers.SavePayement).Methods(http.MethodPost)
	r.HandleFunc("/callBack", handlers.PaymentCallBack).Methods(http.MethodPost)
}
