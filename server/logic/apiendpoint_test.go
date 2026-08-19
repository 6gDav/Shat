package logic

import (
	"encoding/json"
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
	req.RemoteAddr = "192.0.2.1:1234"

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
