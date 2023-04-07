package handlers

import (
	"bytes"
	"depeche/internal/delivery/middleware"
	storage "depeche/internal/repository/local_storage"
	sessionStorage "depeche/internal/session/repository/local_storage"
	authService "depeche/internal/session/service"
	"depeche/internal/usecase/service"
	"depeche/pkg/apperror"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	authTestcases = map[string]struct {
		url    string
		method string
		body   gin.H
		code   int
		err    error
		cookie string
	}{
		"Sign-Up 200 Success": {
			url:    "/auth/sign-up",
			method: http.MethodPost,
			body: gin.H{
				"body": gin.H{
					"email":    "example@example.com",
					"password": "Qwerty123!",
				},
			},
			code:   http.StatusOK,
			err:    nil,
			cookie: "expected",
		},
		"Sign-Up 400 Invalid Json": {
			url:    "/auth/sign-up",
			method: http.MethodPost,
			body: gin.H{
				"body": "invalid json body",
			},
			code: http.StatusBadRequest,
			err:  apperror.BadRequest,
		},
		"Sign-Up 409 Already exists": {
			url:    "/auth/sign-up",
			method: http.MethodPost,
			body: gin.H{
				"body": gin.H{
					"email":    "user1@mail.ru",
					"password": "Qwerty123!",
				},
			},
			code: http.StatusConflict,
			err:  apperror.UserAlreadyExists,
		},
		"Sign-In 200 success": {
			url:    "/auth/sign-in",
			method: http.MethodPost,
			body: gin.H{
				"body": gin.H{
					"email":    "user1@mail.ru",
					"password": "Qwerty123!",
				},
			},
			code:   http.StatusOK,
			cookie: "expected",
		},
		"Sign-In 400 BadRequest": {
			url:    "/auth/sign-in",
			method: http.MethodPost,
			body: gin.H{
				"body": "invalid json body",
			},
			code: http.StatusBadRequest,
			err:  apperror.BadRequest,
		},
		"Sign-In 404 not found": {
			url:    "/auth/sign-in",
			method: http.MethodPost,
			body: gin.H{
				"body": gin.H{
					"email":    "notfound@mail.ru",
					"password": "Qwerty123!",
				},
			},
			code: http.StatusNotFound,
			err:  apperror.UserNotFound,
		},
		"Sign-In 401 incorrect credentials": {
			url:    "/auth/sign-in",
			method: http.MethodPost,
			body: gin.H{
				"body": gin.H{
					"email":    "user1@mail.ru",
					"password": "Password123!",
				},
			},
			code: http.StatusUnauthorized,
			err:  apperror.IncorrectCredentials,
		},
	}
)

var router = gin.Default()

func TestMain(m *testing.M) {
	userStorage := storage.NewUserStorage()
	sessionStorage := sessionStorage.NewMemorySessionStorage()
	feedStorage := storage.NewFeedStorage()

	userService := service.NewUserService(userStorage)
	authService := authService.NewAuthService(sessionStorage)
	feedService := service.NewFeedService(feedStorage)

	userHandler := NewUserHandler(userService, authService)
	feedHandler := NewFeedHandler(feedService)

	handler := NewHandler(userHandler, feedHandler, nil, nil)

	router.Use(middleware.ErrorMiddleware())
	router.Use(func(ctx *gin.Context) {
		ctx.Set("email", "user1@mail.ru")
	})

	authEndpointsGroup := router.Group("/auth")
	{
		authEndpointsGroup.POST("/sign-in", handler.SignIn)
		authEndpointsGroup.POST("/sign-up", handler.SignUp)
		authEndpointsGroup.POST("/logout", handler.LogOut)
		authEndpointsGroup.GET("/check", handler.CheckAuth)
	}

	apiEndpointsGroup := router.Group("/api")
	{
		apiEndpointsGroup.GET("/feed", handler.GetFeed)
	}

	os.Exit(m.Run())
}

func TestUserHandler(t *testing.T) {
	for name, test := range authTestcases {
		t.Run(name, func(t *testing.T) {

			request, err := request(test.method, test.url, test.body)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, request)

			require.Equal(t, test.code, rr.Code)

			if test.err != nil {
				body, err := json.Marshal(gin.H{
					"status":  middleware.Errors[test.err].Code,
					"message": middleware.Errors[test.err].Message,
				})
				require.NoError(t, err)
				require.Equal(t, body, rr.Body.Bytes())
			} else {
				require.NotNil(t, rr.Result().Cookies())
				require.NotNil(t, rr.Result().Cookies()[0])
			}
		})
	}
}

func request(method string, url string, data gin.H) (*http.Request, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return http.NewRequest(method, url, bytes.NewReader(body))
}
