package logic

import (
	"hosting_login_page/chat/client"
	"hosting_login_page/chat/helper"
	"hosting_login_page/history"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func setupTestWS(t *testing.T) (*websocket.Conn, *websocket.Conn, func()) {
	t.Helper()
	var serverConn *websocket.Conn

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		serverConn, err = upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatalf("Failed to upgrade WS: %v", err)
		}
	}))

	wsURL := "ws" + strings.TrimPrefix(s.URL, "http")
	clientConn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to dial WS: %v", err)
	}

	cleanup := func() {
		clientConn.Close()
		if serverConn != nil {
			serverConn.Close()
		}
		s.Close()
	}

	time.Sleep(20 * time.Millisecond)

	return serverConn, clientConn, cleanup
}

func TestChat_PrivateChat(t *testing.T) {
	history.ChatHistory = make(map[string][]history.ChatMessage)
	client.Clients = make(map[string]*client.Client)

	senderIP := "192.168.1.2"
	targetIP := "192.168.1.3"
	username := "TestSender"
	msgText := "Private lorem ipsum message"
	saveChat := true

	targetServerConn, targetClientConn, cleanupTarget := setupTestWS(t)
	defer cleanupTarget()

	senderServerConn, senderClientConn, cleanupSender := setupTestWS(t)
	defer cleanupSender()

	client.ClientsMu.Lock()
	client.Clients[senderIP] = &client.Client{Conn: senderServerConn}
	client.Clients[targetIP] = &client.Client{Conn: targetServerConn}
	client.ClientsMu.Unlock()

	defer func() {
		client.ClientsMu.Lock()
		delete(client.Clients, senderIP)
		delete(client.Clients, targetIP)
		client.ClientsMu.Unlock()
	}()

	cs := &ManageChat{IP: senderIP}
	cs.PrivateChat(targetIP, msgText, username, saveChat)

	var receivedTargetMsg ChatMessage
	_ = targetClientConn.SetReadDeadline(time.Now().Add(1 * time.Second))
	if err := targetClientConn.ReadJSON(&receivedTargetMsg); err != nil {
		t.Fatalf("Target did not receive message: %v", err)
	}

	expectedRoomID := helper.GenerateRoomID(senderIP, targetIP)
	if receivedTargetMsg.Text != msgText || receivedTargetMsg.From != senderIP || receivedTargetMsg.RoomID != expectedRoomID {
		t.Errorf("Target received invalid message structure: %+v", receivedTargetMsg)
	}

	var receivedSenderMsg ChatMessage
	_ = senderClientConn.SetReadDeadline(time.Now().Add(1 * time.Second))
	if err := senderClientConn.ReadJSON(&receivedSenderMsg); err != nil {
		t.Fatalf("Sender did not receive echo message: %v", err)
	}
	if receivedSenderMsg.Text != msgText {
		t.Errorf("Sender echo text mismatch. Got: %s, Want: %s", receivedSenderMsg.Text, msgText)
	}

	history.ChatStoreMu.Lock()
	savedMessages, exists := history.ChatHistory[expectedRoomID]
	history.ChatStoreMu.Unlock()

	if !exists || len(savedMessages) == 0 {
		t.Fatalf("Expected message to be stored in history for room %s, but found none", expectedRoomID)
	}

	if savedMessages[0].Message != msgText || savedMessages[0].FromIp != senderIP {
		t.Errorf("Saved history entry mismatch: %+v", savedMessages[0])
	}
}

func TestChat_GroupChat(t *testing.T) {
	history.ChatHistory = make(map[string][]history.ChatMessage)
	client.Clients = make(map[string]*client.Client)

	senderIP := "192.168.1.10"
	member1IP := "192.168.1.20"
	member2IP := "192.168.1.30"
	username := "SenderUser"
	msgText := "Group lorem ipsum message"
	saveChat := true

	senderServerConn, senderClientConn, cleanupSender := setupTestWS(t)
	defer cleanupSender()

	m1ServerConn, m1ClientConn, cleanupM1 := setupTestWS(t)
	defer cleanupM1()

	m2ServerConn, m2ClientConn, cleanupM2 := setupTestWS(t)
	defer cleanupM2()

	client.ClientsMu.Lock()
	client.Clients[senderIP] = &client.Client{IP: senderIP, Conn: senderServerConn}
	client.Clients[member1IP] = &client.Client{IP: member1IP, Conn: m1ServerConn}
	client.Clients[member2IP] = &client.Client{IP: member2IP, Conn: m2ServerConn}
	client.ClientsMu.Unlock()

	defer func() {
		client.ClientsMu.Lock()
		delete(client.Clients, senderIP)
		delete(client.Clients, member1IP)
		delete(client.Clients, member2IP)
		client.ClientsMu.Unlock()
	}()

	cs := &ManageChat{IP: senderIP}
	cs.GroupChat(msgText, username, saveChat)

	receivers := []struct {
		name       string
		clientConn *websocket.Conn
		expectedTo string
	}{
		{"Sender", senderClientConn, senderIP},
		{"Member 1", m1ClientConn, member1IP},
		{"Member 2", m2ClientConn, member2IP},
	}

	for _, r := range receivers {
		var receivedMsg ChatMessage
		_ = r.clientConn.SetReadDeadline(time.Now().Add(1 * time.Second))
		if err := r.clientConn.ReadJSON(&receivedMsg); err != nil {
			t.Fatalf("%s did not receive group message: %v", r.name, err)
		}

		if receivedMsg.Type != "group" || receivedMsg.Text != msgText || receivedMsg.From != senderIP || receivedMsg.To != r.expectedTo || receivedMsg.RoomID != "Group" {
			t.Errorf("%s received invalid group message payload: %+v", r.name, receivedMsg)
		}
	}

	history.ChatStoreMu.Lock()
	savedMessages, exists := history.ChatHistory["Group"]
	history.ChatStoreMu.Unlock()

	if !exists || len(savedMessages) != 1 {
		t.Fatalf("Expected exactly 1 group history message, found %d", len(savedMessages))
	}

	if savedMessages[0].Message != msgText || savedMessages[0].FromIp != senderIP || savedMessages[0].RoomId != "Group" {
		t.Errorf("Group history entry mismatch: %+v", savedMessages[0])
	}
}
