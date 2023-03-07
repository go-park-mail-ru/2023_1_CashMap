package handlers

import (
	"bytes"
	"depeche/internal/delivery/middleware"
	"depeche/internal/mocks/usecase"
	"depeche/pkg/apperror"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	signUpTestcases = map[string]struct {
		request *http.Request
		code    int
		err     error
		cookie  string
	}{
		"200 Success #1": {
			request: request(http.MethodPost, "/auth/sign-up",
				`{
					"body":{
						"email": "example@example.com",
						"password": "Qwerty123!"
					}
				}`,
			),
			code:   http.StatusOK,
			err:    nil,
			cookie: "expected",
		},
		"400 Invalid Json": {
			request: request(http.MethodPost, "/auth/sign-up",
				`invalid json body`,
			),
			code: http.StatusBadRequest,
			err:  apperror.BadRequest,
		},
	}

	signInTestcases = map[string]struct {
	}{}
	logoutTestcases = map[string]struct {
	}{}
	checkAuthTestcases = map[string]struct {
	}{}
)

func TestUserHandlerSignUp(t *testing.T) {
	for name, test := range signUpTestcases {
		t.Run(name, func(t *testing.T) {
			mockUserService := new(usecase.MockUserService)
			mockUserService.On("SignUp", mock.AnythingOfType("*entities.User")).Return(nil, nil)

			mockAuthService := new(usecase.MockAuthService)
			mockAuthService.On("Authenticate", mock.AnythingOfType("*entities.User")).Return(test.cookie, nil)

			handler := NewUserHandler(mockUserService, mockAuthService)

			router := gin.Default()
			router.Use(middleware.ErrorMiddleware())
			router.POST("/auth/sign-up", handler.SignUp)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, test.request)

			require.Equal(t, test.code, rr.Code)

			if test.err != nil {
				body, err := json.Marshal(gin.H{
					"status":  middleware.Errors[test.err].Code,
					"message": middleware.Errors[test.err].Message,
				})
				require.NoError(t, err)
				require.Equal(t, body, rr.Body.Bytes())
			} else {
				require.Equal(t, test.cookie, rr.Result().Cookies()[0].Value)
			}
			//mockUserService.AssertExpectations(t)
			//mockAuthService.AssertExpectations(t)
		})
	}
}

func TestUserHandlerSignIn(t *testing.T) {

}

func TestUserHandlerLogOut(t *testing.T) {

}

func TestUserHandlerCheckAuth(t *testing.T) {

}

func request(method string, url string, data string) *http.Request {
	req, _ := http.NewRequest(method, url, body([]byte(data)))
	return req
}
func body(data []byte) io.Reader {
	reader := bytes.NewReader(data)
	return reader
}
