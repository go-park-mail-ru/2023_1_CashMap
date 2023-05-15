package handlers

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	mock_usecase "depeche/internal/usecase/mocks"
	"depeche/pkg/middleware"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMessageHandler_Send(t *testing.T) {
	tests := []struct {
		name         string
		email        string
		body         gin.H
		expectedBody gin.H
		expectedCode int
		dto          *dto.NewMessageDTO
		setupMock    func(service *mock_usecase.MockMessageUsecase, email string, dto *dto.NewMessageDTO)
	}{
		{
			name:  "Success",
			email: "e.larkin@mail.ru",
			body: gin.H{
				"body": gin.H{
					"chat_id":              1,
					"message_content_type": "text",
					"text_content":         "msg text",
				},
			},
			dto: &dto.NewMessageDTO{
				ChatId:      1,
				ContentType: "text",
				Text:        "msg text",
			},
			expectedCode: http.StatusOK,
			setupMock: func(service *mock_usecase.MockMessageUsecase, email string, dto *dto.NewMessageDTO) {
				service.EXPECT().Send(email, dto).Return(&entities.Message{}, nil)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_usecase.NewMockMessageUsecase(ctrl)

			msgHandler := MessageHandler{
				mockService,
			}
			test.setupMock(mockService, test.email, test.dto)

			router := gin.New()
			router.Use(middleware.ErrorMiddleware())
			if test.email != "" {
				router.Use(func(context *gin.Context) {
					context.Set("email", test.email)
				})
			}
			router.POST("/", msgHandler.Send)

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

var (
	testLink1  = "id1"
	testLink2  = "id2"
	testBatch  = uint(1)
	testOffset = uint(0)
)

func TestMessageHandler_NewChat(t *testing.T) {
	tests := []struct {
		name         string
		email        string
		body         gin.H
		expectedBody gin.H
		expectedCode int
		dto          *dto.CreateChatDTO
		setupMock    func(service *mock_usecase.MockMessageUsecase, email string, dto *dto.CreateChatDTO)
	}{
		{
			name:  "Success",
			email: "e.larkin@mail.ru",
			body: gin.H{
				"body": gin.H{
					"user_links": []string{"id1"},
				},
			},
			dto: &dto.CreateChatDTO{
				UserLinks: []string{"id1"},
			},
			expectedCode: http.StatusOK,
			expectedBody: gin.H{
				"body": gin.H{
					"chat": entities.Chat{
						ChatID: 1,
						Users: []*entities.UserInfo{
							{
								Link: &testLink1,
							},
							{
								Link: &testLink2,
							},
						},
					},
				},
			},
			setupMock: func(service *mock_usecase.MockMessageUsecase, email string, dto *dto.CreateChatDTO) {
				chat := &entities.Chat{
					ChatID: 1,
					Users: []*entities.UserInfo{
						{
							Link: &testLink1,
						},
						{
							Link: &testLink2,
						},
					},
				}
				service.EXPECT().CreateChat(email, dto).Return(chat, nil)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_usecase.NewMockMessageUsecase(ctrl)

			msgHandler := MessageHandler{
				mockService,
			}
			test.setupMock(mockService, test.email, test.dto)

			router := gin.New()
			router.Use(middleware.ErrorMiddleware())
			if test.email != "" {
				router.Use(func(context *gin.Context) {
					context.Set("email", test.email)
				})
			}
			router.POST("/", msgHandler.NewChat)

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

func TestMessageHandler_GetChats(t *testing.T) {
	tests := []struct {
		name         string
		email        string
		dto          *dto.GetChatsDTO
		expectedBody gin.H
		expectedCode int
		setupMock    func(service *mock_usecase.MockMessageUsecase, email string, dto *dto.GetChatsDTO)
	}{
		{
			name:  "Success",
			email: "e.larkin@mail.ru",
			dto: &dto.GetChatsDTO{
				BatchSize: &testBatch,
				Offset:    &testOffset,
			},
			expectedBody: gin.H{
				"body": gin.H{
					"chats": []entities.Chat{
						{
							0,
							[]*entities.UserInfo{
								{
									Link: &testLink1,
								},
								{
									Link: &testLink2,
								},
							},
						},
					},
				},
			},
			expectedCode: http.StatusOK,
			setupMock: func(service *mock_usecase.MockMessageUsecase, email string, dto *dto.GetChatsDTO) {
				chats := []*entities.Chat{
					{
						ChatID: 0,
						Users: []*entities.UserInfo{
							{
								Link: &testLink1,
							},
							{
								Link: &testLink2,
							},
						},
					},
				}
				service.EXPECT().GetChatsList(email, dto).Return(chats, nil)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_usecase.NewMockMessageUsecase(ctrl)

			msgHandler := MessageHandler{
				mockService,
			}
			test.setupMock(mockService, test.email, test.dto)

			router := gin.New()
			router.Use(middleware.ErrorMiddleware())
			if test.email != "" {
				router.Use(func(context *gin.Context) {
					context.Set("email", test.email)
				})
			}
			router.GET("/", msgHandler.GetChats)
			query := fmt.Sprintf("?batch_size=%d&offset=%d", test.dto.BatchSize, test.dto.Offset)

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
