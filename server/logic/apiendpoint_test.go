package logic

import (
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
