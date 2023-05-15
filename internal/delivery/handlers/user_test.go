package handlers

import (
	"bytes"
	"depeche/authorization_ms/authEntities"
	mock_service "depeche/authorization_ms/service/mocks"
	"depeche/internal/delivery/dto"
	"depeche/internal/delivery/middleware"
	"depeche/internal/entities"
	mock_usecase "depeche/internal/usecase/mocks"
	"depeche/pkg/apperror"
	middleware2 "depeche/pkg/middleware"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_SignUp(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		body         gin.H
		dto          *dto.SignUp
		cookie       string
		csrfToken    string
		checkCookie  bool
		expectedBody gin.H
		withBodyRes  bool
		expectedCode int

		setupMock func(service *mock_usecase.MockUser,
			authService *mock_service.MockAuth,
			csrfService *mock_service.MockCSRFUsecase,
			in *dto.SignUp,
			cookie, csrfToken string)
	}{
		{
			name:   "Success",
			method: "POST",
			body: gin.H{
				"body": gin.H{
					"email":      "e.larkin@mail.ru",
					"password":   "Qwerty123",
					"first_name": "Егор",
					"last_name":  "Ларкин",
				},
			},
			dto: &dto.SignUp{
				Email:     "e.larkin@mail.ru",
				Password:  "Qwerty123",
				FirstName: "Егор",
				LastName:  "Ларкин",
			},
			cookie:       "test_cookie_value",
			checkCookie:  true,
			withBodyRes:  false,
			expectedBody: nil,
			expectedCode: http.StatusOK,

			setupMock: func(service *mock_usecase.MockUser, authService *mock_service.MockAuth, csrfService *mock_service.MockCSRFUsecase, in *dto.SignUp, cookie, csrfToken string) {
				service.EXPECT().SignUp(in).Return(&entities.User{}, nil)
				authService.EXPECT().Authenticate(in.Email).Return(cookie, nil)
				csrfService.EXPECT().CreateCSRFToken(in.Email).Return(csrfToken, nil)
			},
		},
		{
			name:   "User already exists",
			method: "POST",
			body: gin.H{
				"body": gin.H{
					"email":      "e.larkin@mail.ru",
					"password":   "Qwerty123",
					"first_name": "Егор",
					"last_name":  "Ларкин",
				},
			},
			dto: &dto.SignUp{
				Email:     "e.larkin@mail.ru",
				Password:  "Qwerty123",
				FirstName: "Егор",
				LastName:  "Ларкин",
			},
			expectedCode: http.StatusConflict,
			withBodyRes:  true,
			expectedBody: gin.H{
				"status":  middleware.Errors[apperror.UserAlreadyExists].Code,
				"message": middleware.Errors[apperror.UserAlreadyExists].Message,
			},
			setupMock: func(service *mock_usecase.MockUser, authService *mock_service.MockAuth, csrfService *mock_service.MockCSRFUsecase, in *dto.SignUp, cookie, csrfToken string) {
				service.EXPECT().SignUp(in).Return(nil, apperror.NewServerError(apperror.UserAlreadyExists, nil))
			},
		},
		{
			name:   "Auth Service Error",
			method: "POST",
			body: gin.H{
				"body": gin.H{
					"email":      "e.larkin@mail.ru",
					"password":   "Qwerty123",
					"first_name": "Егор",
					"last_name":  "Ларкин",
				},
			},
			dto: &dto.SignUp{
				Email:     "e.larkin@mail.ru",
				Password:  "Qwerty123",
				FirstName: "Егор",
				LastName:  "Ларкин",
			},
			expectedCode: http.StatusUnauthorized,
			withBodyRes:  true,
			expectedBody: gin.H{
				"status":  middleware.Errors[apperror.NoAuth].Code,
				"message": middleware.Errors[apperror.NoAuth].Message,
			},
			setupMock: func(service *mock_usecase.MockUser, authService *mock_service.MockAuth, csrfService *mock_service.MockCSRFUsecase, in *dto.SignUp, cookie, csrfToken string) {
				service.EXPECT().SignUp(in).Return(&entities.User{}, nil)
				authService.EXPECT().Authenticate(in.Email).Return("", apperror.NewServerError(apperror.NoAuth, nil))
			},
		},
		{
			name:   "CSRF Service Error",
			method: "POST",
			body: gin.H{
				"body": gin.H{
					"email":      "e.larkin@mail.ru",
					"password":   "Qwerty123",
					"first_name": "Егор",
					"last_name":  "Ларкин",
				},
			},
			dto: &dto.SignUp{
				Email:     "e.larkin@mail.ru",
				Password:  "Qwerty123",
				FirstName: "Егор",
				LastName:  "Ларкин",
			},
			expectedCode: http.StatusUnauthorized,
			withBodyRes:  true,
			expectedBody: gin.H{
				"status":  middleware.Errors[apperror.NoAuth].Code,
				"message": middleware.Errors[apperror.NoAuth].Message,
			},
			setupMock: func(service *mock_usecase.MockUser, authService *mock_service.MockAuth, csrfService *mock_service.MockCSRFUsecase, in *dto.SignUp, cookie, csrfToken string) {
				service.EXPECT().SignUp(in).Return(&entities.User{}, nil)
				authService.EXPECT().Authenticate(in.Email).Return("cookie", nil)
				csrfService.EXPECT().CreateCSRFToken(in.Email).Return("", apperror.NewServerError(apperror.NoAuth, nil))
			},
		},
		{
			name:   "BadRequest",
			method: "POST",
			body: gin.H{
				"request": gin.H{
					"password":   "Qwerty123",
					"first_name": "Егор",
					"last_name":  "Ларкин",
				},
			},
			expectedCode: http.StatusBadRequest,
			withBodyRes:  true,
			expectedBody: gin.H{
				"status":  middleware.Errors[apperror.BadRequest].Code,
				"message": middleware.Errors[apperror.BadRequest].Message,
			},
			setupMock: func(service *mock_usecase.MockUser, authService *mock_service.MockAuth, csrfService *mock_service.MockCSRFUsecase, in *dto.SignUp, cookie, csrfToken string) {

			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_usecase.NewMockUser(ctrl)
			mockAuth := mock_service.NewMockAuth(ctrl)
			mockCSRF := mock_service.NewMockCSRFUsecase(ctrl)

			userHandler := UserHandler{
				service:     mockService,
				authService: mockAuth,
				csrfService: mockCSRF,
			}
			test.setupMock(mockService, mockAuth, mockCSRF, test.dto, test.cookie, test.csrfToken)

			router := gin.New()
			router.Use(middleware2.ErrorMiddleware())
			router.POST("/", userHandler.SignUp)
			req, err := request(test.method, "/", test.body)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			if test.withBodyRes {
				body, err := json.Marshal(test.expectedBody)
				require.NoError(t, err)
				require.Equal(t, body, rr.Body.Bytes())
			}

			require.Equal(t, test.expectedCode, rr.Code)

			if test.checkCookie {
				require.NotEmpty(t, rr.Result().Cookies())
				require.Equal(t, test.cookie, rr.Result().Cookies()[0].Value)
				require.Equal(t, test.csrfToken, rr.Result().Header.Get("X-Csrf-Token"))
			}
		})
	}
}

func TestUserHandler_SignIn(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		body         gin.H
		dto          *dto.SignIn
		cookie       string
		csrfToken    string
		checkCookie  bool
		expectedBody gin.H
		withBodyRes  bool
		expectedCode int

		setupMock func(service *mock_usecase.MockUser,
			authService *mock_service.MockAuth,
			csrfService *mock_service.MockCSRFUsecase,
			in *dto.SignIn,
			cookie, csrfToken string)
	}{
		{
			name:   "Success",
			method: "POST",
			body: gin.H{
				"body": gin.H{
					"email":    "e.larkin@mail.ru",
					"password": "Qwerty123",
				},
			},
			dto: &dto.SignIn{
				Email:    "e.larkin@mail.ru",
				Password: "Qwerty123",
			},
			cookie:       "test_cookie_value",
			checkCookie:  true,
			withBodyRes:  false,
			expectedBody: nil,
			expectedCode: http.StatusOK,

			setupMock: func(service *mock_usecase.MockUser, authService *mock_service.MockAuth, csrfService *mock_service.MockCSRFUsecase, in *dto.SignIn, cookie, csrfToken string) {
				service.EXPECT().SignIn(in).Return(&entities.User{}, nil)
				authService.EXPECT().Authenticate(in.Email).Return(cookie, nil)
				csrfService.EXPECT().CreateCSRFToken(in.Email).Return(csrfToken, nil)
			},
		},
		{
			name:   "Incorrect credentials",
			method: "POST",
			body: gin.H{
				"body": gin.H{
					"email":    "e.larkin@mail.ru",
					"password": "Qwerty123",
				},
			},
			dto: &dto.SignIn{
				Email:    "e.larkin@mail.ru",
				Password: "Qwerty123",
			},
			checkCookie: false,
			withBodyRes: true,
			expectedBody: gin.H{
				"status":  middleware.Errors[apperror.IncorrectCredentials].Code,
				"message": middleware.Errors[apperror.IncorrectCredentials].Message,
			},
			expectedCode: http.StatusUnauthorized,

			setupMock: func(service *mock_usecase.MockUser, authService *mock_service.MockAuth, csrfService *mock_service.MockCSRFUsecase, in *dto.SignIn, cookie, csrfToken string) {
				service.EXPECT().SignIn(in).Return(nil, apperror.NewServerError(apperror.IncorrectCredentials, nil))
			},
		},
		{
			name:   "Auth Service Error",
			method: "POST",
			body: gin.H{
				"body": gin.H{
					"email":    "e.larkin@mail.ru",
					"password": "Qwerty123",
				},
			},
			dto: &dto.SignIn{
				Email:    "e.larkin@mail.ru",
				Password: "Qwerty123",
			},
			expectedCode: http.StatusUnauthorized,
			withBodyRes:  true,
			expectedBody: gin.H{
				"status":  middleware.Errors[apperror.NoAuth].Code,
				"message": middleware.Errors[apperror.NoAuth].Message,
			},
			setupMock: func(service *mock_usecase.MockUser, authService *mock_service.MockAuth, csrfService *mock_service.MockCSRFUsecase, in *dto.SignIn, cookie, csrfToken string) {
				service.EXPECT().SignIn(in).Return(&entities.User{}, nil)
				authService.EXPECT().Authenticate(in.Email).Return("", apperror.NewServerError(apperror.NoAuth, nil))
			},
		},
		{
			name:   "CSRF Service Error",
			method: "POST",
			body: gin.H{
				"body": gin.H{
					"email":    "e.larkin@mail.ru",
					"password": "Qwerty123",
				},
			},
			dto: &dto.SignIn{
				Email:    "e.larkin@mail.ru",
				Password: "Qwerty123",
			},
			expectedCode: http.StatusUnauthorized,
			withBodyRes:  true,
			expectedBody: gin.H{
				"status":  middleware.Errors[apperror.NoAuth].Code,
				"message": middleware.Errors[apperror.NoAuth].Message,
			},
			setupMock: func(service *mock_usecase.MockUser, authService *mock_service.MockAuth, csrfService *mock_service.MockCSRFUsecase, in *dto.SignIn, cookie, csrfToken string) {
				service.EXPECT().SignIn(in).Return(&entities.User{}, nil)
				authService.EXPECT().Authenticate(in.Email).Return("cookie", nil)
				csrfService.EXPECT().CreateCSRFToken(in.Email).Return("", apperror.NewServerError(apperror.NoAuth, nil))
			},
		},
		{
			name:   "BadRequest",
			method: "POST",
			body: gin.H{
				"request": gin.H{
					"password":   "Qwerty123",
					"first_name": "Егор",
				},
			},
			expectedCode: http.StatusBadRequest,
			withBodyRes:  true,
			expectedBody: gin.H{
				"status":  middleware.Errors[apperror.BadRequest].Code,
				"message": middleware.Errors[apperror.BadRequest].Message,
			},
			setupMock: func(service *mock_usecase.MockUser, authService *mock_service.MockAuth, csrfService *mock_service.MockCSRFUsecase, in *dto.SignIn, cookie, csrfToken string) {

			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_usecase.NewMockUser(ctrl)
			mockAuth := mock_service.NewMockAuth(ctrl)
			mockCSRF := mock_service.NewMockCSRFUsecase(ctrl)

			userHandler := UserHandler{
				service:     mockService,
				authService: mockAuth,
				csrfService: mockCSRF,
			}
			test.setupMock(mockService, mockAuth, mockCSRF, test.dto, test.cookie, test.csrfToken)

			router := gin.New()
			router.Use(middleware2.ErrorMiddleware())
			router.POST("/", userHandler.SignIn)
			req, err := request(test.method, "/", test.body)
			require.NoError(t, err)

			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)
			if test.withBodyRes {
				body, err := json.Marshal(test.expectedBody)
				require.NoError(t, err)
				require.Equal(t, body, rr.Body.Bytes())
			}

			require.Equal(t, test.expectedCode, rr.Code)

			if test.checkCookie {
				require.NotEmpty(t, rr.Result().Cookies())
				require.Equal(t, test.cookie, rr.Result().Cookies()[0].Value)
				require.Equal(t, test.csrfToken, rr.Result().Header.Get("X-Csrf-Token"))
			}
		})
	}
}

func TestUserHandler_LogOut(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		cookie       string
		csrfToken    string
		setCookie    bool
		email        string
		expectedBody gin.H
		withBodyRes  bool
		expectedCode int

		setupMock func(service *mock_usecase.MockUser,
			authService *mock_service.MockAuth,
			csrfService *mock_service.MockCSRFUsecase,
			cookie, csrfToken, email string)
	}{
		{
			name:         "Success",
			method:       "POST",
			cookie:       "test_cookie_value",
			csrfToken:    "test_csrf_token_value",
			email:        "e.larkin@mail.ru",
			withBodyRes:  false,
			setCookie:    true,
			expectedCode: http.StatusOK,
			setupMock: func(service *mock_usecase.MockUser, authService *mock_service.MockAuth, csrfService *mock_service.MockCSRFUsecase, cookie, csrfToken, email string) {
				userSession := &authEntities.Session{
					Email: email,
				}
				authService.EXPECT().CheckSession(cookie).Return(userSession, nil)
				csrf := &authEntities.CSRF{
					Email: email,
					Token: csrfToken,
				}
				csrfService.EXPECT().InvalidateCSRFToken(csrf).Return(nil)
				authService.EXPECT().LogOut(cookie).Return(nil)
			},
		},
		{
			name:        "No cookie",
			method:      "POST",
			withBodyRes: true,
			expectedBody: gin.H{
				"status":  middleware.Errors[apperror.NoAuth].Code,
				"message": middleware.Errors[apperror.NoAuth].Message,
			},
			expectedCode: http.StatusUnauthorized,
			setupMock: func(service *mock_usecase.MockUser, authService *mock_service.MockAuth, csrfService *mock_service.MockCSRFUsecase, cookie, csrfToken, email string) {

			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_usecase.NewMockUser(ctrl)
			mockAuth := mock_service.NewMockAuth(ctrl)
			mockCSRF := mock_service.NewMockCSRFUsecase(ctrl)

			userHandler := UserHandler{
				service:     mockService,
				authService: mockAuth,
				csrfService: mockCSRF,
			}
			test.setupMock(mockService, mockAuth, mockCSRF, test.cookie, test.csrfToken, test.email)

			router := gin.New()
			router.Use(middleware2.ErrorMiddleware())
			router.POST("/", userHandler.LogOut)
			req, err := request(test.method, "/", nil)
			require.NoError(t, err)
			if test.setCookie {
				cookie := &http.Cookie{
					Value: test.cookie,
					Name:  "session_id",
				}
				req.AddCookie(cookie)
				req.Header.Set("X-Csrf-Token", test.csrfToken)
			}

			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)
			if test.withBodyRes {
				body, err := json.Marshal(test.expectedBody)
				require.NoError(t, err)
				require.Equal(t, body, rr.Body.Bytes())
			}

			require.Equal(t, test.expectedCode, rr.Code)
		})
	}
}

func TestUserHandler_Subscribe(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		email    string
		setEmail bool
		body     gin.H
		dto      *dto.Subscribes

		expectedBody gin.H
		withBodyRes  bool
		expectedCode int
		setupMock    func(service *mock_usecase.MockUser, email, link string)
	}{
		{
			name:   "Success",
			method: "POST",
			email:  "e.larkin@mail.ru",
			dto: &dto.Subscribes{
				Link: "id1234567",
			},
			setEmail: true,
			body: gin.H{
				"body": gin.H{
					"user_link": "id1234567",
				},
			},
			expectedCode: http.StatusOK,
			setupMock: func(service *mock_usecase.MockUser, email, link string) {
				service.EXPECT().Subscribe(email, link).Return(nil)
			},
		},
		{
			name:     "No auth",
			method:   "POST",
			setEmail: false,
			dto: &dto.Subscribes{
				Link: "id1234567",
			},
			body: gin.H{
				"body": gin.H{
					"user_link": "id1234567",
				},
			},

			withBodyRes:  true,
			expectedCode: http.StatusUnauthorized,

			expectedBody: gin.H{
				"status":  middleware.Errors[apperror.NoAuth].Code,
				"message": middleware.Errors[apperror.NoAuth].Message,
			},

			setupMock: func(service *mock_usecase.MockUser, email, link string) {

			},
		},
		{
			name:     "Subscribe error",
			method:   "POST",
			setEmail: true,
			email:    "e.larkin@mail.ru",
			dto: &dto.Subscribes{
				Link: "id1234567",
			},
			body: gin.H{
				"body": gin.H{
					"user_link": "id1234567",
				},
			},

			withBodyRes:  true,
			expectedCode: http.StatusInternalServerError,

			expectedBody: gin.H{
				"status":  middleware.Errors[apperror.InternalServerError].Code,
				"message": middleware.Errors[apperror.InternalServerError].Message,
			},

			setupMock: func(service *mock_usecase.MockUser, email, link string) {
				service.EXPECT().Subscribe(email, link).Return(apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
		{
			name:     "Bad request",
			method:   "POST",
			setEmail: true,
			email:    "e.larkin@mail.ru",
			dto: &dto.Subscribes{
				Link: "id1234567",
			},
			body: gin.H{
				"request": gin.H{
					"user_link": "id1234567",
				},
			},

			withBodyRes:  true,
			expectedCode: http.StatusBadRequest,

			expectedBody: gin.H{
				"status":  middleware.Errors[apperror.BadRequest].Code,
				"message": middleware.Errors[apperror.BadRequest].Message,
			},

			setupMock: func(service *mock_usecase.MockUser, email, link string) {

			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_usecase.NewMockUser(ctrl)
			mockAuth := mock_service.NewMockAuth(ctrl)
			mockCSRF := mock_service.NewMockCSRFUsecase(ctrl)

			userHandler := UserHandler{
				service:     mockService,
				authService: mockAuth,
				csrfService: mockCSRF,
			}
			test.setupMock(mockService, test.email, test.dto.Link)

			router := gin.New()
			router.Use(middleware2.ErrorMiddleware())
			if test.setEmail {
				router.Use(func(context *gin.Context) {
					context.Set("email", test.email)
				})
			}
			router.POST("/", userHandler.Subscribe)
			req, err := request(test.method, "/", test.body)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			if test.withBodyRes {
				body, err := json.Marshal(test.expectedBody)
				require.NoError(t, err)
				require.Equal(t, body, rr.Body.Bytes())
			}

			require.Equal(t, test.expectedCode, rr.Code)
		})
	}
}

func TestUserHandler_Friends(t *testing.T) {
	tests := []struct {
		name   string
		link   string
		limit  int
		offset int

		email string

		expectedBody gin.H
		expectedCode int

		setupMock func(service *mock_usecase.MockUser,
			email, link string, limit, offset int)
	}{
		{
			name:   "Success",
			link:   "id1234",
			limit:  2,
			offset: 0,

			email: "e.larkin@mail.ru",

			expectedBody: gin.H{
				"body": gin.H{
					"friends": []dto.Profile{
						{
							Link:      "id1",
							FirstName: "Pavel",
							LastName:  "Repin",
						},
						{
							Link:      "id2",
							FirstName: "Egor",
							LastName:  "Larkin",
						},
					},
				},
			},
			expectedCode: http.StatusOK,

			setupMock: func(service *mock_usecase.MockUser, email, link string, limit, offset int) {
				returned := []*entities.User{
					{
						Link:      "id1",
						FirstName: "Pavel",
						LastName:  "Repin",
					},
					{
						Link:      "id2",
						FirstName: "Egor",
						LastName:  "Larkin",
					},
				}
				service.EXPECT().GetFriendsByLink(email, link, limit, offset).Return(returned, nil)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_usecase.NewMockUser(ctrl)

			userHandler := UserHandler{
				service: mockService,
			}
			test.setupMock(mockService, test.email, test.link, test.limit, test.offset)

			router := gin.New()
			router.Use(middleware2.ErrorMiddleware())
			if test.email != "" {
				router.Use(func(context *gin.Context) {
					context.Set("email", test.email)
				})
			}

			router.GET("/", userHandler.Friends)
			query := fmt.Sprintf("?link=%s&limit=%d&offset=%d", test.link, test.limit, test.offset)
			req, err := request("GET", "/"+query, nil)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			if test.expectedBody != nil {
				body, err := json.Marshal(test.expectedBody)
				require.NoError(t, err)
				require.Equal(t, body, rr.Body.Bytes())
			}

			require.Equal(t, test.expectedCode, rr.Code)
		})
	}
}

func TestUserHandler_Subscribes(t *testing.T) {
	tests := []struct {
		name    string
		link    string
		limit   int
		offset  int
		subType string
		email   string

		expectedBody gin.H
		expectedCode int

		setupMock func(service *mock_usecase.MockUser,
			email, link string, limit, offset int)
	}{
		{
			name:    "Success type=in",
			link:    "id1234",
			subType: "in",
			limit:   2,
			offset:  0,

			email: "e.larkin@mail.ru",

			expectedBody: gin.H{
				"body": gin.H{
					"subs": []dto.Profile{
						{
							Link:      "id1",
							FirstName: "Pavel",
							LastName:  "Repin",
						},
						{
							Link:      "id2",
							FirstName: "Egor",
							LastName:  "Larkin",
						},
					},
				},
			},
			expectedCode: http.StatusOK,

			setupMock: func(service *mock_usecase.MockUser, email, link string, limit, offset int) {
				returned := []*entities.User{
					{
						Link:      "id1",
						FirstName: "Pavel",
						LastName:  "Repin",
					},
					{
						Link:      "id2",
						FirstName: "Egor",
						LastName:  "Larkin",
					},
				}
				service.EXPECT().GetSubscribersByLink(email, link, limit, offset).Return(returned, nil)
			},
		},
		{
			name:    "Success type=out",
			link:    "id1234",
			subType: "out",
			limit:   2,
			offset:  0,

			email: "e.larkin@mail.ru",

			expectedBody: gin.H{
				"body": gin.H{
					"subs": []dto.Profile{
						{
							Link:      "id1",
							FirstName: "Pavel",
							LastName:  "Repin",
						},
						{
							Link:      "id2",
							FirstName: "Egor",
							LastName:  "Larkin",
						},
					},
				},
			},
			expectedCode: http.StatusOK,

			setupMock: func(service *mock_usecase.MockUser, email, link string, limit, offset int) {
				returned := []*entities.User{
					{
						Link:      "id1",
						FirstName: "Pavel",
						LastName:  "Repin",
					},
					{
						Link:      "id2",
						FirstName: "Egor",
						LastName:  "Larkin",
					},
				}
				service.EXPECT().GetSubscribesByLink(email, link, limit, offset).Return(returned, nil)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_usecase.NewMockUser(ctrl)

			userHandler := UserHandler{
				service: mockService,
			}
			test.setupMock(mockService, test.email, test.link, test.limit, test.offset)

			router := gin.New()
			router.Use(middleware2.ErrorMiddleware())
			if test.email != "" {
				router.Use(func(context *gin.Context) {
					context.Set("email", test.email)
				})
			}

			router.GET("/", userHandler.Subscribes)
			query := fmt.Sprintf("?link=%s&limit=%d&offset=%d&type=%s", test.link, test.limit, test.offset, test.subType)
			req, err := request("GET", "/"+query, nil)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			if test.expectedBody != nil {
				body, err := json.Marshal(test.expectedBody)
				require.NoError(t, err)
				require.Equal(t, body, rr.Body.Bytes())
			}

			require.Equal(t, test.expectedCode, rr.Code)
		})
	}
}

func TestUserHandler_RandomUsers(t *testing.T) {
	tests := []struct {
		name   string
		limit  int
		offset int
		email  string

		expectedBody gin.H
		expectedCode int

		setupMock func(service *mock_usecase.MockUser,
			email string, limit, offset int)
	}{
		{
			name:   "Success",
			limit:  2,
			offset: 0,

			email: "e.larkin@mail.ru",

			expectedBody: gin.H{
				"body": gin.H{
					"profiles": []dto.Profile{
						{
							Link:      "id1",
							FirstName: "Pavel",
							LastName:  "Repin",
						},
						{
							Link:      "id2",
							FirstName: "Egor",
							LastName:  "Larkin",
						},
					},
				},
			},
			expectedCode: http.StatusOK,

			setupMock: func(service *mock_usecase.MockUser, email string, limit, offset int) {
				returned := []*entities.User{
					{
						Link:      "id1",
						FirstName: "Pavel",
						LastName:  "Repin",
					},
					{
						Link:      "id2",
						FirstName: "Egor",
						LastName:  "Larkin",
					},
				}
				service.EXPECT().GetAllUsers(email, limit, offset).Return(returned, nil)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_usecase.NewMockUser(ctrl)

			userHandler := UserHandler{
				service: mockService,
			}
			test.setupMock(mockService, test.email, test.limit, test.offset)

			router := gin.New()
			router.Use(middleware2.ErrorMiddleware())
			if test.email != "" {
				router.Use(func(context *gin.Context) {
					context.Set("email", test.email)
				})
			}

			router.GET("/", userHandler.RandomUsers)
			query := fmt.Sprintf("?limit=%d&offset=%d", test.limit, test.offset)
			req, err := request("GET", "/"+query, nil)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			if test.expectedBody != nil {
				body, err := json.Marshal(test.expectedBody)
				require.NoError(t, err)
				require.Equal(t, body, rr.Body.Bytes())
			}

			require.Equal(t, test.expectedCode, rr.Code)
		})
	}
}

func TestUserHandler_PendingRequests(t *testing.T) {
	tests := []struct {
		name   string
		limit  int
		offset int
		email  string

		expectedBody gin.H
		expectedCode int

		setupMock func(service *mock_usecase.MockUser,
			email string, limit, offset int)
	}{
		{
			name:   "Success",
			limit:  2,
			offset: 0,

			email: "e.larkin@mail.ru",

			expectedBody: gin.H{
				"body": gin.H{
					"profiles": []dto.Profile{
						{
							Link:      "id1",
							FirstName: "Pavel",
							LastName:  "Repin",
						},
						{
							Link:      "id2",
							FirstName: "Egor",
							LastName:  "Larkin",
						},
					},
				},
			},
			expectedCode: http.StatusOK,

			setupMock: func(service *mock_usecase.MockUser, email string, limit, offset int) {
				returned := []*entities.User{
					{
						Link:      "id1",
						FirstName: "Pavel",
						LastName:  "Repin",
					},
					{
						Link:      "id2",
						FirstName: "Egor",
						LastName:  "Larkin",
					},
				}
				service.EXPECT().GetPendingRequestsByEmail(email, limit, offset).Return(returned, nil)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_usecase.NewMockUser(ctrl)

			userHandler := UserHandler{
				service: mockService,
			}
			test.setupMock(mockService, test.email, test.limit, test.offset)

			router := gin.New()
			router.Use(middleware2.ErrorMiddleware())
			if test.email != "" {
				router.Use(func(context *gin.Context) {
					context.Set("email", test.email)
				})
			}

			router.GET("/", userHandler.PendingRequests)
			query := fmt.Sprintf("?limit=%d&offset=%d", test.limit, test.offset)
			req, err := request("GET", "/"+query, nil)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			if test.expectedBody != nil {
				body, err := json.Marshal(test.expectedBody)
				require.NoError(t, err)
				require.Equal(t, body, rr.Body.Bytes())
			}

			require.Equal(t, test.expectedCode, rr.Code)
		})
	}
}

func TestUserHandler_Unsubscribe(t *testing.T) {
	tests := []struct {
		name  string
		email string
		body  gin.H
		dto   *dto.Subscribes

		expectedBody gin.H
		expectedCode int

		setupMock func(service *mock_usecase.MockUser, email string, dto *dto.Subscribes)
	}{
		{
			name:  "Success",
			email: "e.larkin@mail.ru",
			dto: &dto.Subscribes{
				Link: "id123",
			},

			body: gin.H{
				"body": gin.H{
					"user_link": "id123",
				},
			},

			expectedCode: http.StatusOK,
			setupMock: func(service *mock_usecase.MockUser, email string, dto *dto.Subscribes) {
				service.EXPECT().Unsubscribe(email, dto.Link).Return(nil)
			},
		},
		{
			name: "Unauthorized",
			dto: &dto.Subscribes{
				Link: "id123",
			},

			body: gin.H{
				"body": gin.H{
					"user_link": "id123",
				},
			},
			expectedCode: http.StatusUnauthorized,

			setupMock: func(service *mock_usecase.MockUser, email string, dto *dto.Subscribes) {

			},
		},
		{
			name:         "Bad request",
			email:        "e.larkin@mail.ru",
			expectedCode: http.StatusBadRequest,
			setupMock: func(service *mock_usecase.MockUser, email string, dto *dto.Subscribes) {

			},
		},
		{
			name:  "Internal error",
			email: "e.larkin@mail.ru",
			dto: &dto.Subscribes{
				Link: "id123",
			},

			body: gin.H{
				"body": gin.H{
					"user_link": "id123",
				},
			},
			expectedCode: http.StatusInternalServerError,
			setupMock: func(service *mock_usecase.MockUser, email string, dto *dto.Subscribes) {
				service.EXPECT().Unsubscribe(email, dto.Link).Return(apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_usecase.NewMockUser(ctrl)

			userHandler := UserHandler{
				service: mockService,
			}
			test.setupMock(mockService, test.email, test.dto)

			router := gin.New()
			router.Use(middleware2.ErrorMiddleware())
			if test.email != "" {
				router.Use(func(context *gin.Context) {
					context.Set("email", test.email)
				})
			}
			router.POST("/", userHandler.Unsubscribe)
			req, err := request("POST", "/", test.body)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			if test.expectedBody != nil {
				body, err := json.Marshal(test.expectedBody)
				require.NoError(t, err)
				require.Equal(t, body, rr.Body.Bytes())
			}
			require.Equal(t, test.expectedCode, rr.Code)
		})
	}
}

func TestUserHandler_Profile(t *testing.T) {
	tests := []struct {
		name string
		link string

		expectedBody gin.H
		expectedCode int

		setupMock func(service *mock_usecase.MockUser, link string)
	}{
		{
			name: "Success",
			link: "id123",

			expectedBody: gin.H{
				"body": gin.H{
					"profile": dto.Profile{
						Link:      "id123",
						FirstName: "Pavel",
						LastName:  "Repin",
					},
				},
			},
			expectedCode: http.StatusOK,
			setupMock: func(service *mock_usecase.MockUser, link string) {
				profile := &entities.User{
					Link:      "id123",
					FirstName: "Pavel",
					LastName:  "Repin",
				}
				service.EXPECT().GetProfileByLink("", link).Return(profile, nil)
			},
		},
		{
			name: "Not found",
			link: "id123",

			expectedBody: gin.H{
				"status":  middleware.Errors[apperror.UserNotFound].Code,
				"message": middleware.Errors[apperror.UserNotFound].Message,
			},

			expectedCode: http.StatusNotFound,

			setupMock: func(service *mock_usecase.MockUser, link string) {
				service.EXPECT().GetProfileByLink("", link).Return(nil, apperror.NewServerError(apperror.UserNotFound, nil))
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_usecase.NewMockUser(ctrl)

			userHandler := UserHandler{
				service: mockService,
			}
			test.setupMock(mockService, test.link)

			router := gin.New()
			router.Use(middleware2.ErrorMiddleware())
			router.GET("/:link", userHandler.Profile)

			req, err := request("GET", "/"+test.link, nil)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			if test.expectedBody != nil {
				body, err := json.Marshal(test.expectedBody)
				require.NoError(t, err)
				require.Equal(t, body, rr.Body.Bytes())
			}

			require.Equal(t, test.expectedCode, rr.Code)
		})
	}
}

func TestUserHandler_Reject(t *testing.T) {
	tests := []struct {
		name  string
		email string

		body gin.H
		dto  *dto.Subscribes

		expectedBody gin.H

		expectedCode int
		setupMock    func(service *mock_usecase.MockUser, email, link string)
	}{
		{
			name:  "Success",
			email: "e.larkin@mail.ru",

			body: gin.H{
				"body": gin.H{
					"user_link": "id1",
				},
			},
			dto: &dto.Subscribes{
				Link: "id1",
			},

			expectedCode: http.StatusOK,

			setupMock: func(service *mock_usecase.MockUser, email, link string) {
				service.EXPECT().Reject(email, link).Return(nil)
			},
		},
		{
			name: "No auth",
			body: gin.H{
				"body": gin.H{
					"user_link": "id1",
				},
			},
			dto: &dto.Subscribes{
				Link: "id1",
			},
			expectedCode: http.StatusUnauthorized,

			expectedBody: gin.H{
				"status":  middleware.Errors[apperror.NoAuth].Code,
				"message": middleware.Errors[apperror.NoAuth].Message,
			},

			setupMock: func(service *mock_usecase.MockUser, email, link string) {

			},
		},
		{
			name: "Bad Request",
			body: gin.H{
				"request": "link",
			},
			dto: &dto.Subscribes{
				Link: "link",
			},

			expectedBody: gin.H{
				"status":  middleware.Errors[apperror.BadRequest].Code,
				"message": middleware.Errors[apperror.BadRequest].Message,
			},
			expectedCode: http.StatusBadRequest,
			setupMock: func(service *mock_usecase.MockUser, email, link string) {

			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_usecase.NewMockUser(ctrl)

			userHandler := UserHandler{
				service: mockService,
			}

			test.setupMock(mockService, test.email, test.dto.Link)

			router := gin.New()
			router.Use(middleware2.ErrorMiddleware())
			if test.email != "" {
				router.Use(func(context *gin.Context) {
					context.Set("email", test.email)
				})
			}
			router.POST("/", userHandler.Reject)
			req, err := request("POST", "/", test.body)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			if test.expectedBody != nil {
				body, err := json.Marshal(test.expectedBody)
				require.NoError(t, err)
				require.Equal(t, body, rr.Body.Bytes())
			}

			require.Equal(t, test.expectedCode, rr.Code)
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
