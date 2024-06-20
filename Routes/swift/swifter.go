package swift

import "github.com/gorilla/websocket"

type WebSocketRouter struct {
    routes map[string]func(*websocket.Conn, []byte)
}

func NewWebSocketRouter() *WebSocketRouter {
    return &WebSocketRouter{
        routes: make(map[string]func(*websocket.Conn, []byte)),
    }
}

func (router *WebSocketRouter) On(path string, handlerFunc func(*websocket.Conn, []byte)) {
    router.routes[path] = handlerFunc
}

func (router *WebSocketRouter) Handle(conn *websocket.Conn, messageType int, message []byte) {
    routeHandler, ok := router.routes[string(message)]
    if !ok {
        return 
    }
    routeHandler(conn, message)
}
