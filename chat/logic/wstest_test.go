package logic

import (
	"hosting_login_page/chat/client"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestHandShake_RegisteredClient(t *testing.T) {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		HandShake(w, r, upgrader)
	}))
	defer server.Close()

	serverURL, _ := url.Parse(server.URL)
	ip := strings.Split(serverURL.Host, ":")[0]

	client.ClientsMu.Lock()
	client.Clients = make(map[string]*client.Client)
	client.Clients[ip] = &client.Client{
		IP:   ip,
		Name: "TestUser",
	}
	client.ClientsMu.Unlock()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")

	wsConn, resp, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to connect to WebSocket: %v", err)
	}
	defer wsConn.Close()

	if resp.StatusCode != http.StatusSwitchingProtocols {
		t.Errorf("Exptet state 101 got: %d", resp.StatusCode)
	}

	time.Sleep(50 * time.Millisecond)

	client.ClientsMu.Lock()
	user, exists := client.Clients[ip]
	client.ClientsMu.Unlock()

	if !exists {
		t.Errorf("Client not found in th Map")
	}
	if user.Conn == nil {
		t.Errorf("The user.Conn did not got setted")
	}
}

func TestHandShake_UnregisteredClient(t *testing.T) {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		HandShake(w, r, upgrader)
	}))
	defer server.Close()

	client.ClientsMu.Lock()
	client.Clients = make(map[string]*client.Client)
	client.ClientsMu.Unlock()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")

	wsConn, resp, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("HansShake must be succesfull befor closeing %v", err)
	}
	defer wsConn.Close()
	defer resp.Body.Close()

	var errPayload map[string]string
	if err := wsConn.ReadJSON(&errPayload); err != nil {
		t.Fatalf("Failed to read the JSON message: %v", err)
	}
	if errPayload["message"] != "Please submit your name first!" {
		t.Errorf("Exepted message: %s", errPayload["message"])
	}

	_, _, err = wsConn.ReadMessage()
	if err == nil {
		t.Fatalf("The connect must be closed")
	}

	closeErr, ok := err.(*websocket.CloseError)
	if !ok {
		t.Fatalf("Unexepted error type: %v (CloseError-t exepted)", err)
	}

	if closeErr.Code != 4001 {
		t.Errorf("Waited close code 4001, got: %d", closeErr.Code)
	}
}
