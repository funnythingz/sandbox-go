package main

import(
    "log"
    "net"
    "net/http"
    "sync"

    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    "github.com/gorilla/websocket"
)

var ActiveClients = make(map[ClientConn]int)
var ActiveClientsRWMutex sync.RWMutex

type ClientConn struct {

    websocket *websocket.Conn
    clientIP net.Addr
}

func addClient(cc ClientConn) {

    ActiveClientsRWMutex.Lock()
    ActiveClients[cc] = 0
    ActiveClientsRWMutex.Unlock()
}

func deleteClient(cc ClientConn) {

    ActiveClientsRWMutex.Lock()
    delete(ActiveClients, cc)
    ActiveClientsRWMutex.Unlock()
}

func broadcastMessage(messageType int, message []byte) {

    ActiveClientsRWMutex.RLock()

    defer ActiveClientsRWMutex.RUnlock()

    for client, _ := range ActiveClients {
        if err := client.websocket.WriteMessage(messageType, message); err != nil {
            return
        }
    }
}

func main() {

    m := martini.Classic()

    m.Use(render.Renderer(render.Options{
        Layout: "layout",
    }))
    m.Use(martini.Static("assets"))

    m.Get("/", Index)
    m.Get("/about", About)
    m.Get("/ws", WebSocket)

    m.Run()
}

func Index(r render.Render) {
    r.HTML(200, "index", nil)
}

func About(r render.Render) {
    r.HTML(200, "about", nil)
}

func WebSocket(w http.ResponseWriter, r *http.Request) {

    ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)

    if _, ok := err.(websocket.HandshakeError); ok {
        http.Error(w, "Not a websocket handshake", 400)
        return
    } else if err != nil {
        log.Println(err)
        return
    }

    client := ws.RemoteAddr()
    clientConn := ClientConn{ws, client}
    addClient(clientConn)

    for {
        messageType, p, err := ws.ReadMessage()

        if err != nil {
            deleteClient(clientConn)
            log.Println(err)
            return
        }

        broadcastMessage(messageType, p)
    }
}
