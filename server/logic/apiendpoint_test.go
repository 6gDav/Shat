package logic

import (
	"encoding/json"
	"hosting_login_page/history"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestServePage(t *testing.T) {
	_ = os.Chdir("../..")

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	ServePage(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected 200, got %d. Body: %s", rec.Code, rec.Body.String())
	}
}

func TestRedirectUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.RemoteAddr = "192.0.2.1:1234" //TEST-NET-1 an IP address for tests

	rec := httptest.NewRecorder()

	RedirectUser(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected 200, got %d. Body: %s", rec.Code, rec.Body.String())
	}

	var got struct {
		Redirect  bool   `json:"redirect"`
		UserName  string `json:"userName"`
		IpAddress string `json:"ipAddress"`
	}

	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("Failed to decode response JSON: %v", err)
	}

	if got.Redirect != false {
		t.Errorf("Expected redirect to be false, got %v", got.Redirect)
	}
}

func TestRestoreChatHistory(t *testing.T) {
	testIP := "192.0.2.1"

	history.ChatStoreMu.Lock()
	history.ChatHistory = map[string][]history.ChatMessage{
		"Group": {
			{
				Type:     "group",
				FromIp:   "10.0.0.1",
				Message:  "Group lorem ipsum",
				RoomId:   "Group",
				UserName: "Jhon Doe",
			},
		},
		"192.168.1.200_192.168.1.201": {
			{
				Type:     "private",
				FromIp:   "192.168.1.200",
				Message:  "Private lorem ipsum",
				RoomId:   "192.168.1.200_192.168.1.201",
				UserName: "Jhon Doe",
			},
		},
	}
	history.ChatStoreMu.Unlock()

	req := httptest.NewRequest(http.MethodGet, "/restore-history", nil)

	req.RemoteAddr = testIP + ":12345"

	rec := httptest.NewRecorder()

	RestoreChatHistory(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected 200, got %d. Body: %s", rec.Code, rec.Body.String())
	}

	var responseList []history.ChatMessage
	err := json.Unmarshal(rec.Body.Bytes(), &responseList)
	if err != nil {
		t.Fatalf("The response is not a the valid form: %v", err)
	}
}
