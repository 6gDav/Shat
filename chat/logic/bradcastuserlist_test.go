package logic

import (
	"Shat/chat/client"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

type BroadcastPayload struct {
	Type string            `json:"type"`
	Data map[string]string `json:"data"`
}

func TestBroadcast_UserNameList(t *testing.T) {
	var serverConn *websocket.Conn
	var connWg sync.WaitGroup
	connWg.Add(1)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatalf("Failed to upgrade the connection: %v", err)
		}
		serverConn = conn
		connWg.Done()
	}))
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")
	clientConn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("Client connection failure: %v", err)
	}
	defer clientConn.Close()

	connWg.Wait()
	defer serverConn.Close()

	client.ClientsMu.Lock()
	client.Clients = map[string]*client.Client{
		"192.168.1.10": {
			IP:   "192.168.1.10",
			Name: "Jhon",
			Conn: serverConn,
		},
		"192.168.1.20": {
			IP:   "192.168.1.20",
			Name: "Doe",
			Conn: nil,
		},
	}
	client.ClientsMu.Unlock()

	BroadcastUserNameList()

	_ = clientConn.SetReadDeadline(time.Now().Add(2 * time.Second))
	messageType, messageData, err := clientConn.ReadMessage()
	if err != nil {
		t.Fatalf("Message didnt arrived %v", err)
	}

	if messageType != websocket.TextMessage {
		t.Errorf("Exepted message type (%d), got: %d", websocket.TextMessage, messageType)
	}

	var payload BroadcastPayload
	if err := json.Unmarshal(messageData, &payload); err != nil {
		t.Fatalf("JSON decode error: %v", err)
	}

	if payload.Type != "USER_NAME_LIST" {
		t.Errorf("Exepted type 'USER_NAME_LIST', got: '%s'", payload.Type)
	}

	if len(payload.Data) != 1 {
		t.Fatalf("Exepted adata lenght %d", len(payload.Data))
	}

	if name, ok := payload.Data["192.168.1.10"]; !ok || name != "Jhon" {
		t.Errorf("Not proper user data: %+v", payload.Data)
	}
}
