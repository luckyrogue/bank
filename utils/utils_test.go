package utils

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseJSONBody_Success(t *testing.T) {
	jsonData := `{"account_type": "Current"}`
	req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewBuffer([]byte(jsonData)))

	var body struct {
		AccountType string `json:"account_type"`
	}

	err := ParseJSONBody(httptest.NewRecorder(), req, &body)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if body.AccountType != "Current" {
		t.Errorf("expected account_type to be 'Current', got '%s'", body.AccountType)
	}
}

func TestParseJSONBody_InvalidJSON(t *testing.T) {
	invalidJSON := `{"account_type": "Current"`
	req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewBuffer([]byte(invalidJSON)))

	var body struct {
		AccountType string `json:"account_type"`
	}

	err := ParseJSONBody(httptest.NewRecorder(), req, &body)
	if err == nil {
		t.Fatal("expected error, got none")
	}
}

func TestGenerateID(t *testing.T) {
	id := GenerateID()

	if len(id) == 0 {
		t.Error("expected non-empty UUID, got empty string")
	}
}
