package middlewares

import (
	"TEST_SERVER/session"
	"TEST_SERVER/utils"
	"fmt"
	"log"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check if the user is saved in the store
		_, ok := session.Get(r, "_id").(string)

		if !ok {
			utils.CreateOutput(w, fmt.Errorf("login to access the resource"), false, nil)
			return
		}
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
