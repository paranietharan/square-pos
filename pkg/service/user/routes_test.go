package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"square-pos/pkg/types"
	"testing"

	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the email is empty", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "Paranie",
			LastName:  "Tharan",
			Password:  "1245P",
		}

		m, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(m))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister).Methods(http.MethodPost)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}

		var response map[string]string
		if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
			t.Fatalf("could not parse response: %v", err)
		}

		if response["error"] != "email is required" {
			t.Errorf("expected error message 'email is required', got '%v'", response["error"])
		}
	})
}

type mockUserStore struct{}

func (m *mockUserStore) UpdateUser(u types.User) error {
	return nil
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return &types.User{}, nil
}

func (m *mockUserStore) CreateUser(u *types.User) error {
	return nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return &types.User{}, nil
}
