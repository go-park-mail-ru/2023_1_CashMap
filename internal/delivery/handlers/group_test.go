package handlers

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/delivery/middleware"
	"depeche/internal/entities"
	mock_usecase "depeche/internal/usecase/mocks"
	"depeche/pkg/apperror"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGroupHandler_GetGroup(t *testing.T) {
	tests := []struct {
		name string
		link string

		expectedBody gin.H
		expectedCode int
		setupMock    func(service *mock_usecase.MockGroup, link string)
	}{
		{
			name: "Success",
			link: "id123",

			expectedBody: gin.H{
				"body": gin.H{
					"group": entities.Group{
						Link:  "id123",
						Title: "Group",
					},
				},
			},
			expectedCode: http.StatusOK,
			setupMock: func(service *mock_usecase.MockGroup, link string) {
				group := &entities.Group{
					Link:  "id123",
					Title: "Group",
				}
				service.EXPECT().GetGroup(link).Return(group, nil)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_usecase.NewMockGroup(ctrl)

			groupHandler := GroupHandler{
				service: mockService,
			}
			test.setupMock(mockService, test.link)

			router := gin.New()
			router.Use(middleware.ErrorMiddleware())

			router.GET("/:link", groupHandler.GetGroup)

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

func TestGroupHandler_GetGroups(t *testing.T) {
	tests := []struct {
		name   string
		limit  int
		offset int
		email  string

		expectedBody gin.H
		expectedCode int

		setupMock func(service *mock_usecase.MockGroup, email string, limit, offset int)
	}{
		{
			name:   "Success",
			limit:  2,
			offset: 0,
			email:  "e.larkin@mail.ru",

			expectedBody: gin.H{
				"body": gin.H{
					"groups": []entities.Group{
						{
							Link:  "id1234",
							Title: "Group#1234",
						},
						{
							Link:  "id10",
							Title: "Group#10",
						},
					},
				},
			},
			expectedCode: http.StatusOK,
			setupMock: func(service *mock_usecase.MockGroup, email string, limit, offset int) {
				groups := []*entities.Group{
					{
						Link:  "id1234",
						Title: "Group#1234",
					},
					{
						Link:  "id10",
						Title: "Group#10",
					},
				}
				service.EXPECT().GetUserGroupsByEmail(email, limit, offset).Return(groups, nil)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_usecase.NewMockGroup(ctrl)

			groupHandler := GroupHandler{
				service: mockService,
			}
			test.setupMock(mockService, test.email, test.limit, test.offset)

			router := gin.New()
			router.Use(middleware.ErrorMiddleware())
			if test.email != "" {
				router.Use(func(context *gin.Context) {
					context.Set("email", test.email)
				})
			}
			router.GET("/", groupHandler.GetGroups)
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

func TestGroupHandler_GetUserGroups(t *testing.T) {
	tests := []struct {
		name   string
		limit  int
		offset int
		link   string

		expectedBody gin.H
		expectedCode int

		setupMock func(service *mock_usecase.MockGroup, link string, limit, offset int)
	}{
		{
			name:   "Success",
			limit:  2,
			offset: 0,
			link:   "id100",

			expectedBody: gin.H{
				"body": gin.H{
					"groups": []entities.Group{
						{
							Link:  "id1234",
							Title: "Group#1234",
						},
						{
							Link:  "id10",
							Title: "Group#10",
						},
					},
				},
			},
			expectedCode: http.StatusOK,
			setupMock: func(service *mock_usecase.MockGroup, link string, limit, offset int) {
				groups := []*entities.Group{
					{
						Link:  "id1234",
						Title: "Group#1234",
					},
					{
						Link:  "id10",
						Title: "Group#10",
					},
				}
				service.EXPECT().GetUserGroupsByLink(link, limit, offset).Return(groups, nil)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_usecase.NewMockGroup(ctrl)

			groupHandler := GroupHandler{
				service: mockService,
			}
			test.setupMock(mockService, test.link, test.limit, test.offset)

			router := gin.New()
			router.Use(middleware.ErrorMiddleware())

			router.GET("/", groupHandler.GetUserGroups)
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

func TestGroupHandler_GetPopularGroups(t *testing.T) {
	tests := []struct {
		name   string
		limit  int
		offset int
		email  string

		expectedBody gin.H
		expectedCode int

		setupMock func(service *mock_usecase.MockGroup, email string, limit, offset int)
	}{
		{
			name:   "Success",
			limit:  2,
			offset: 0,
			email:  "e.larkin@mail.ru",

			expectedBody: gin.H{
				"body": gin.H{
					"groups": []entities.Group{
						{
							Link:  "id1234",
							Title: "Group#1234",
						},
						{
							Link:  "id10",
							Title: "Group#10",
						},
					},
				},
			},
			expectedCode: http.StatusOK,
			setupMock: func(service *mock_usecase.MockGroup, email string, limit, offset int) {
				groups := []*entities.Group{
					{
						Link:  "id1234",
						Title: "Group#1234",
					},
					{
						Link:  "id10",
						Title: "Group#10",
					},
				}
				service.EXPECT().GetPopularGroups(email, limit, offset).Return(groups, nil)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_usecase.NewMockGroup(ctrl)

			groupHandler := GroupHandler{
				service: mockService,
			}
			test.setupMock(mockService, test.email, test.limit, test.offset)

			router := gin.New()
			router.Use(middleware.ErrorMiddleware())
			if test.email != "" {
				router.Use(func(context *gin.Context) {
					context.Set("email", test.email)
				})
			}
			router.GET("/", groupHandler.GetPopularGroups)
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

func TestGroupHandler_GetManagedGroups(t *testing.T) {
	tests := []struct {
		name   string
		limit  int
		offset int
		email  string

		expectedBody gin.H
		expectedCode int

		setupMock func(service *mock_usecase.MockGroup, email string, limit, offset int)
	}{
		{
			name:   "Success",
			limit:  2,
			offset: 0,
			email:  "e.larkin@mail.ru",

			expectedBody: gin.H{
				"body": gin.H{
					"groups": []entities.Group{
						{
							Link:  "id1234",
							Title: "Group#1234",
						},
						{
							Link:  "id10",
							Title: "Group#10",
						},
					},
				},
			},
			expectedCode: http.StatusOK,
			setupMock: func(service *mock_usecase.MockGroup, email string, limit, offset int) {
				groups := []*entities.Group{
					{
						Link:  "id1234",
						Title: "Group#1234",
					},
					{
						Link:  "id10",
						Title: "Group#10",
					},
				}
				service.EXPECT().GetManagedGroups(email, limit, offset).Return(groups, nil)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_usecase.NewMockGroup(ctrl)

			groupHandler := GroupHandler{
				service: mockService,
			}
			test.setupMock(mockService, test.email, test.limit, test.offset)

			router := gin.New()
			router.Use(middleware.ErrorMiddleware())
			if test.email != "" {
				router.Use(func(context *gin.Context) {
					context.Set("email", test.email)
				})
			}
			router.GET("/", groupHandler.GetManagedGroups)
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

func TestGroupHandler_GetSubscribers(t *testing.T) {
	tests := []struct {
		name   string
		limit  int
		offset int
		link   string
		email  string

		expectedBody gin.H
		expectedCode int

		setupMock func(service *mock_usecase.MockGroup, link string, limit, offset int)
	}{
		{
			name:   "Success",
			limit:  2,
			offset: 0,
			link:   "id1234",
			email:  "e.larkin@mail.ru",

			expectedBody: gin.H{
				"body": gin.H{
					"profiles": []dto.Profile{
						{
							Link:      "id1234",
							FirstName: "Egor",
						},
						{
							Link:      "id10",
							FirstName: "Pavel",
						},
					},
				},
			},
			expectedCode: http.StatusOK,
			setupMock: func(service *mock_usecase.MockGroup, link string, limit, offset int) {
				users := []*entities.User{
					{
						Link:      "id1234",
						FirstName: "Egor",
					},
					{
						Link:      "id10",
						FirstName: "Pavel",
					},
				}
				service.EXPECT().GetSubscribers(link, limit, offset).Return(users, nil)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_usecase.NewMockGroup(ctrl)

			groupHandler := GroupHandler{
				service: mockService,
			}
			test.setupMock(mockService, test.link, test.limit, test.offset)

			router := gin.New()
			router.Use(middleware.ErrorMiddleware())
			if test.email != "" {
				router.Use(func(context *gin.Context) {
					context.Set("email", test.email)
				})
			}
			router.GET("/:link", groupHandler.GetSubscribers)
			query := fmt.Sprintf("%s?limit=%d&offset=%d", test.link, test.limit, test.offset)
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

func TestGroupHandler_PendingGroupRequests(t *testing.T) {
	tests := []struct {
		name   string
		limit  int
		offset int
		link   string
		email  string

		expectedBody gin.H
		expectedCode int

		setupMock func(service *mock_usecase.MockGroup, email, link string, limit, offset int)
	}{
		{
			name:   "Success",
			limit:  2,
			offset: 0,
			link:   "id1234",
			email:  "e.larkin@mail.ru",

			expectedBody: gin.H{
				"body": gin.H{
					"profiles": []dto.Profile{
						{
							Link:      "id1234",
							FirstName: "Egor",
						},
						{
							Link:      "id10",
							FirstName: "Pavel",
						},
					},
				},
			},
			expectedCode: http.StatusOK,
			setupMock: func(service *mock_usecase.MockGroup, email, link string, limit, offset int) {
				users := []*entities.User{
					{
						Link:      "id1234",
						FirstName: "Egor",
					},
					{
						Link:      "id10",
						FirstName: "Pavel",
					},
				}
				service.EXPECT().GetPendingRequests(email, link, limit, offset).Return(users, nil)
			},
		},
		{
			name:  "Bad request",
			limit: -1,
			link:  "id1234",
			email: "e.larkin@mail.ru",

			expectedCode: http.StatusBadRequest,
			expectedBody: gin.H{
				"status":  middleware.Errors[apperror.BadRequest].Code,
				"message": middleware.Errors[apperror.BadRequest].Message,
			},

			setupMock: func(service *mock_usecase.MockGroup, email, link string, limit, offset int) {

			},
		},
		{
			name:   "no auth",
			link:   "id1234",
			limit:  2,
			offset: 0,

			expectedCode: http.StatusUnauthorized,
			expectedBody: gin.H{
				"status":  middleware.Errors[apperror.NoAuth].Code,
				"message": middleware.Errors[apperror.NoAuth].Message,
			},

			setupMock: func(service *mock_usecase.MockGroup, email, link string, limit, offset int) {
			},
		},
		{
			name:         "internal error",
			limit:        2,
			offset:       0,
			link:         "id1234",
			email:        "e.larkin@mail.ru",
			expectedCode: http.StatusInternalServerError,
			expectedBody: gin.H{
				"status":  middleware.Errors[apperror.InternalServerError].Code,
				"message": middleware.Errors[apperror.InternalServerError].Message,
			},
			setupMock: func(service *mock_usecase.MockGroup, email, link string, limit, offset int) {
				service.EXPECT().GetPendingRequests(email, link, limit, offset).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_usecase.NewMockGroup(ctrl)

			groupHandler := GroupHandler{
				service: mockService,
			}
			test.setupMock(mockService, test.email, test.link, test.limit, test.offset)

			router := gin.New()
			router.Use(middleware.ErrorMiddleware())
			if test.email != "" {
				router.Use(func(context *gin.Context) {
					context.Set("email", test.email)
				})
			}
			router.GET("/:link", groupHandler.PendingGroupRequests)
			query := fmt.Sprintf("%s?limit=%d&offset=%d", test.link, test.limit, test.offset)
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
