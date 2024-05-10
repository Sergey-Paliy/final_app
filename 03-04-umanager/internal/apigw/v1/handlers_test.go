package v1_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"final_app/03-04-umanager/internal/apigw/v1"
)

type mockServerInterface struct{}

func (m *mockServerInterface) PostContactsHandler(w http.ResponseWriter, r *http.Request) {

}

func (m *mockServerInterface) DeleteContactsHandler(w http.ResponseWriter, r *http.Request) {

}

func TestPostContactsHandler(t *testing.T) {
	// Создаем фэйковый запрос методом POST
	reqBody := []byte(`{"name": "John", "email": "john@example.com"}`)
	req, err := http.NewRequest("POST", "/v1/contacts", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Создаем экземпляр обработчика
	handler := v1.New(&mockServerInterface{}, &mockServerInterface{})

	handler.PostContactsHandler(rr, req)

	// Проверяем статус код ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestDeleteContactsHandler(t *testing.T) {
	// Создаем фэйковый запрос методом DELETE
	req, err := http.NewRequest("DELETE", "/v1/contacts/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := v1.New(&mockServerInterface{}, &mockServerInterface{})

	handler.DeleteContactsHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
