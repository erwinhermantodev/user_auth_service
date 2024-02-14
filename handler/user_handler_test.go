package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/erwinhermantodev/user_auth_service/repository"
	"github.com/erwinhermantodev/user_auth_service/repository/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_Login(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryInterface)
	handler := NewUserHandler(mockRepo)

	tests := []struct {
		name           string
		phoneNumber    string
		password       string
		mockFindResult *repository.User
		mockFindErr    error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Invalid credentials",
			phoneNumber:    "+621234567890",
			password:       "wrongpassword",
			mockFindResult: nil,
			mockFindErr:    nil,
			expectedStatus: http.StatusUnauthorized, // Corrected status code
			expectedBody:   `{"error":"Invalid phone number or password"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.On("FindUserByPhone", tc.phoneNumber).Return(tc.mockFindResult, tc.mockFindErr)
			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"phone_number":"`+tc.phoneNumber+`","password":"`+tc.password+`"}`))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := echo.New().NewContext(req, rec)

			if assert.NoError(t, handler.Login(c)) {
				assert.Equal(t, tc.expectedStatus, rec.Code)
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
