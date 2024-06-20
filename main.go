package main

import (
	swifthandlers "TEST_SERVER/Handlers/swiftHandlers"
	httproutes "TEST_SERVER/Routes/httpRoutes"
	"TEST_SERVER/Routes/swift"
	"TEST_SERVER/database"

	// "TEST_SERVER/database"
	// "TEST_SERVER/session"

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

// initializing the socket programmig for various functionalities

var Upgrader = websocket.Upgrader{
	WriteBufferSize: 2010,
	ReadBufferSize:  2010,
}

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Env loading error \n")
	}
	log.Println("Succecfully loaded envi")
	database.InitDatabase()

	// Initializing the session too....
	//session.NewSession("afm_tecv1")

	// payment_access
	// utils.InitAzamPay()

	// intializing the swift_websocket implementation...

}
func printing(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Hello test from server")
	w.Write([]byte("Welcome back"))

}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	wsRouter := swift.NewWebSocketRouter()
	// Register WebSocket routes
	wsRouter.On("/fetch/data", swifthandlers.FetchData)
	// Add more routes here

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// Handle incoming messages
		wsRouter.Handle(conn, messageType, message)
	}
}

func main() {

	app := mux.NewRouter()
	app.HandleFunc("/test", printing)

	// http routes....
	httproutes.CompanyRoutes(app)
	httproutes.UserRoutes(app)
	httproutes.ClientRoutes(app)

	http.HandleFunc("/swift", handleWebSocket)

	http.Handle("/", app)
	// adding cors for (CROSS ORIGIN)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
		// Debug:            true,
	})

	handlers := c.Handler(app)
	http.ListenAndServe(":3400", handlers)

}
