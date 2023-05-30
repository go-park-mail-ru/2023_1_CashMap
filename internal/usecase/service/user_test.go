package service

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	mock_repository "depeche/internal/repository/mocks"
	"depeche/internal/utils"
	"depeche/internal/utils/testUtils"
	"depeche/pkg/apperror"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserService_SignIn(t *testing.T) {
	tests := []struct {
		name          string
		inputDTO      *dto.SignIn
		expectedUser  *entities.User
		expectedError error

		setupMock func(repo *mock_repository.MockUserRepository, in *dto.SignIn)
	}{
		{
			name: "Successful",
			inputDTO: &dto.SignIn{
				Email:    "e-larkin@mail.ru",
				Password: "Qwerty123",
			},
			expectedUser:  &entities.User{},
			expectedError: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, in *dto.SignIn) {
				repo.EXPECT().GetUserByEmail(in.Email).Return(&entities.User{Password: utils.Hash("Qwerty123")}, nil)
			},
		},
		{
			name: "User doesn't exist",
			inputDTO: &dto.SignIn{
				Email:    "e-larkin@mail.ru",
				Password: "Qwerty123",
			},
			expectedUser:  nil,
			expectedError: apperror.UserNotFound,

			setupMock: func(repo *mock_repository.MockUserRepository, in *dto.SignIn) {
				repo.EXPECT().GetUserByEmail(in.Email).Return(nil, apperror.NewServerError(apperror.UserNotFound, nil))
			},
		},
		{
			name: "Incorrect password",
			inputDTO: &dto.SignIn{
				Email:    "e-larkin@mail.ru",
				Password: "Qwerty123",
			},
			expectedUser:  nil,
			expectedError: apperror.IncorrectCredentials,

			setupMock: func(repo *mock_repository.MockUserRepository, in *dto.SignIn) {
				repo.EXPECT().GetUserByEmail(in.Email).Return(&entities.User{Password: "Another Password"}, nil)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockUserRepository(ctrl)

			userService := UserService{
				repo: mockRepository,
			}
			test.setupMock(mockRepository, test.inputDTO)
			_, err := userService.SignIn(test.inputDTO)
			if test.expectedError != nil {
				fmt.Println(err)
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.expectedError, err)
			}

		})
	}
}

func TestUserService_SignUp(t *testing.T) {
	tests := []struct {
		name          string
		inputDTO      *dto.SignUp
		expectedUser  *entities.User
		expectedError error

		setupMock func(repo *mock_repository.MockUserRepository, in *dto.SignUp)
	}{
		{
			name: "Successful",
			inputDTO: &dto.SignUp{
				Email:     "e-larkin@mail.ru",
				Password:  "Qwerty123",
				FirstName: "Egor",
				LastName:  "Larkin",
			},
			expectedUser:  nil,
			expectedError: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, in *dto.SignUp) {
				repo.EXPECT().GetUserByEmail(in.Email).Return(nil, apperror.NewServerError(apperror.UserNotFound, nil))
				repo.EXPECT().CreateUser(&entities.User{
					Email:     in.Email,
					Password:  utils.Hash(in.Password),
					FirstName: in.FirstName,
					LastName:  in.LastName,
				}).Return(&entities.User{}, nil)
				repo.EXPECT().SubscribeOnDefaultGroup(in.Email)
			},
		},
		{
			name: "User already exists",
			inputDTO: &dto.SignUp{
				Email:     "e-larkin@mail.ru",
				Password:  "Qwerty123",
				FirstName: "Egor",
				LastName:  "Larkin",
			},
			expectedUser:  nil,
			expectedError: apperror.UserAlreadyExists,

			setupMock: func(repo *mock_repository.MockUserRepository, in *dto.SignUp) {
				repo.EXPECT().GetUserByEmail(in.Email).Return(&entities.User{}, nil)
			},
		},
		{
			name: "Internal error get by email",
			inputDTO: &dto.SignUp{
				Email:     "e-larkin@mail.ru",
				Password:  "Qwerty123",
				FirstName: "Egor",
				LastName:  "Larkin",
			},
			expectedUser:  nil,
			expectedError: apperror.InternalServerError,

			setupMock: func(repo *mock_repository.MockUserRepository, in *dto.SignUp) {
				repo.EXPECT().GetUserByEmail(in.Email).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
		{
			name: "Internal error create",
			inputDTO: &dto.SignUp{
				Email:     "e-larkin@mail.ru",
				Password:  "Qwerty123",
				FirstName: "Egor",
				LastName:  "Larkin",
			},
			expectedUser:  nil,
			expectedError: apperror.InternalServerError,

			setupMock: func(repo *mock_repository.MockUserRepository, in *dto.SignUp) {
				repo.EXPECT().GetUserByEmail(in.Email).Return(nil, apperror.NewServerError(apperror.UserNotFound, nil))
				repo.EXPECT().CreateUser(&entities.User{
					Email:     in.Email,
					Password:  utils.Hash(in.Password),
					FirstName: in.FirstName,
					LastName:  in.LastName,
				}).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockUserRepository(ctrl)

			userService := UserService{
				repo: mockRepository,
			}
			test.setupMock(mockRepository, test.inputDTO)
			_, err := userService.SignUp(test.inputDTO)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestUserService_GetProfileByEmail(t *testing.T) {
	tests := []struct {
		name          string
		email         string
		expectedUser  *entities.User
		expectedError error

		setupMock func(repo *mock_repository.MockUserRepository, email string)
	}{
		{
			name:          "Found",
			email:         "e-larkin@mail.ru",
			expectedUser:  nil,
			expectedError: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, email string) {
				repo.EXPECT().GetUserByEmail(email).Return(&entities.User{}, nil)
			},
		},
		{
			name:          "Not found",
			email:         "e-larkin@mail.ru",
			expectedUser:  nil,
			expectedError: apperror.UserNotFound,

			setupMock: func(repo *mock_repository.MockUserRepository, email string) {
				repo.EXPECT().GetUserByEmail(email).Return(nil, apperror.NewServerError(apperror.UserNotFound, nil))
			},
		},
		{
			name:          "Internal error",
			email:         "e-larkin@mail.ru",
			expectedUser:  nil,
			expectedError: apperror.InternalServerError,

			setupMock: func(repo *mock_repository.MockUserRepository, email string) {
				repo.EXPECT().GetUserByEmail(email).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockUserRepository(ctrl)

			userService := UserService{
				repo: mockRepository,
			}
			test.setupMock(mockRepository, test.email)
			_, err := userService.GetProfileByEmail(test.email)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestUserService_GetProfileByLink(t *testing.T) {
	tests := []struct {
		name          string
		email         string
		link          string
		id            uint
		expectedUser  *entities.User
		expectedError error

		numericLink bool

		setupMockNumericLink func(repo *mock_repository.MockUserRepository, id uint)
		setupMockCustomLink  func(repo *mock_repository.MockUserRepository, link string)
	}{
		{
			name:          "Found with 'id<number>' link",
			link:          "id1234",
			id:            1234,
			expectedUser:  nil,
			expectedError: nil,

			numericLink: true,

			setupMockNumericLink: func(repo *mock_repository.MockUserRepository, id uint) {
				repo.EXPECT().GetUserById(id).Return(&entities.User{}, nil)
			},
		},
		{
			name:          "Invalid 'id<number>' link",
			link:          "idabc",
			expectedUser:  nil,
			expectedError: apperror.BadRequest,

			numericLink: true,

			setupMockNumericLink: func(repo *mock_repository.MockUserRepository, id uint) {

			},
		},
		{
			name:          "Found with custom link",
			link:          "egor123",
			expectedError: nil,
			numericLink:   false,

			setupMockCustomLink: func(repo *mock_repository.MockUserRepository, link string) {
				repo.EXPECT().GetUserByLink(link).Return(&entities.User{}, nil)
			},
		},
		{
			name:          "Not Found with custom link",
			link:          "egor123",
			expectedError: apperror.UserNotFound,
			numericLink:   false,

			setupMockCustomLink: func(repo *mock_repository.MockUserRepository, link string) {
				repo.EXPECT().GetUserByLink(link).Return(nil, apperror.NewServerError(apperror.UserNotFound, nil))
			},
		},
		{
			name:          "Not Found with 'id<number>' link",
			link:          "id1234",
			id:            1234,
			expectedUser:  nil,
			expectedError: apperror.UserNotFound,

			numericLink: true,

			setupMockNumericLink: func(repo *mock_repository.MockUserRepository, id uint) {
				repo.EXPECT().GetUserById(id).Return(nil, apperror.NewServerError(apperror.UserNotFound, nil))
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockUserRepository(ctrl)

			userService := UserService{
				repo: mockRepository,
			}

			if test.numericLink {
				test.setupMockNumericLink(mockRepository, test.id)
				_, err := userService.GetProfileByLink(test.email, test.link)
				if test.expectedError != nil {
					uerr, ok := err.(*apperror.ServerError)
					require.Equal(t, true, ok)
					require.Equal(t, test.expectedError, uerr.UserErr)
				} else {
					require.Equal(t, test.expectedError, err)
				}
				return
			}
			test.setupMockCustomLink(mockRepository, test.link)
			_, err := userService.GetProfileByLink(test.email, test.link)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestUserService_GetAllUsers(t *testing.T) {
	tests := []struct {
		name   string
		email  string
		limit  int
		offset int

		expectedUsers []*entities.User
		expectedError error

		setupMock func(repo *mock_repository.MockUserRepository, email string, limit, offset int)
	}{
		{
			name:   "Successful",
			email:  "e-larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedError: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, email string, limit, offset int) {
				repo.EXPECT().GetUsers(email, limit, offset).Return([]*entities.User{
					{
						Email: "random1@email.com",
					},
					{
						Email: "random2@email.com",
					},
				}, nil)
			},
		},
		{
			name:   "Internal Error",
			email:  "e-larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedError: apperror.InternalServerError,

			setupMock: func(repo *mock_repository.MockUserRepository, email string, limit, offset int) {
				repo.EXPECT().GetUsers(email, limit, offset).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockUserRepository(ctrl)

			userService := UserService{
				repo: mockRepository,
			}
			test.setupMock(mockRepository, test.email, test.limit, test.offset)
			users, err := userService.GetAllUsers(test.email, test.limit, test.offset)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.expectedError, err)
				require.Equal(t, test.limit, len(users))
			}
		})
	}
}

func TestUserService_EditProfile(t *testing.T) {

	tests := []struct {
		name    string
		email   string
		link    string
		avatar  string
		profile *dto.EditProfile

		expectedError error

		setupMock func(repo *mock_repository.MockUserRepository,
			email, avatar, link string, profile *dto.EditProfile)
	}{
		{
			name:    "Basic info success",
			email:   "e-larkin@mail.ru",
			profile: testUtils.InitProfileBasic("Egor", "Larkin", "Bio"),

			expectedError: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, email, avatar, link string, profile *dto.EditProfile) {
				repo.EXPECT().UpdateUser(email, profile).Return(&entities.User{}, nil)
			},
		},
		{
			name:    "Basic internal error",
			email:   "e-larkin@mail.ru",
			profile: testUtils.InitProfileBasic("Egor", "Larkin", "Bio"),

			expectedError: apperror.InternalServerError,

			setupMock: func(repo *mock_repository.MockUserRepository, email, avatar, link string, profile *dto.EditProfile) {
				repo.EXPECT().UpdateUser(email, profile).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
		{
			name:    "Avatar success",
			email:   "e-larkin@mail.ru",
			avatar:  "static/avatar/url/12345",
			profile: testUtils.InitProfileAvatar("static/avatar/url/12345"),

			expectedError: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, email, avatar, link string, profile *dto.EditProfile) {
				repo.EXPECT().UpdateAvatar(email, avatar).Return(nil)
				repo.EXPECT().UpdateUser(email, profile).Return(&entities.User{}, nil)
			},
		},
		{
			name:    "Avatar update internal error",
			email:   "e-larkin@mail.ru",
			avatar:  "static/avatar/url/12345",
			profile: testUtils.InitProfileAvatar("static/avatar/url/12345"),

			expectedError: apperror.InternalServerError,

			setupMock: func(repo *mock_repository.MockUserRepository, email, avatar, link string, profile *dto.EditProfile) {
				repo.EXPECT().UpdateAvatar(email, avatar).Return(apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
		{
			name:    "Password success",
			email:   "e-larkin@mail.ru",
			profile: testUtils.InitProfilePasswordWithPrev("oldPassword", "newPassword"),

			expectedError: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, email, avatar, link string, profile *dto.EditProfile) {
				repo.EXPECT().GetUserByEmail(email).Return(&entities.User{Password: utils.Hash("oldPassword")}, nil)
				repo.EXPECT().UpdateUser(email, profile).Return(&entities.User{}, nil)
			},
		},
		{
			name:    "Incorrect password",
			email:   "e-larkin@mail.ru",
			profile: testUtils.InitProfilePasswordWithPrev("oldPasswordIncorrect", "newPassword"),

			expectedError: apperror.Forbidden,

			setupMock: func(repo *mock_repository.MockUserRepository, email, avatar, link string, profile *dto.EditProfile) {
				repo.EXPECT().GetUserByEmail(email).Return(&entities.User{Password: utils.Hash("oldPassword")}, nil)
			},
		},
		{
			name:    "Password update internal error",
			email:   "e-larkin@mail.ru",
			profile: testUtils.InitProfilePasswordWithPrev("oldPasswordIncorrect", "newPassword"),

			expectedError: apperror.InternalServerError,

			setupMock: func(repo *mock_repository.MockUserRepository, email, avatar, link string, profile *dto.EditProfile) {
				repo.EXPECT().GetUserByEmail(email).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
		{
			name:    "Previous password not sent",
			email:   "e-larkin@mail.ru",
			profile: testUtils.InitProfilePasswordFail("newPassword"),

			expectedError: apperror.Forbidden,

			setupMock: func(repo *mock_repository.MockUserRepository, email, avatar, link string, profile *dto.EditProfile) {
				repo.EXPECT().GetUserByEmail(email).Return(&entities.User{Password: utils.Hash("oldPassword")}, nil)
			},
		},
		{
			name:    "Link success",
			email:   "e-larkin@mail.ru",
			link:    "newLink",
			profile: testUtils.InitProfileLink("newLink"),

			expectedError: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, email, avatar, link string, profile *dto.EditProfile) {
				repo.EXPECT().CheckLinkExists(link).Return(false, nil)
				repo.EXPECT().UpdateUser(email, profile).Return(&entities.User{}, nil)
			},
		},
		{
			name:    "Link exists",
			email:   "e-larkin@mail.ru",
			link:    "newLink",
			profile: testUtils.InitProfileLink("newLink"),

			expectedError: apperror.UserAlreadyExists,

			setupMock: func(repo *mock_repository.MockUserRepository, email, avatar, link string, profile *dto.EditProfile) {
				repo.EXPECT().CheckLinkExists(link).Return(true, nil)
			},
		},
		{
			name:    "Link Internal Error",
			email:   "e-larkin@mail.ru",
			link:    "newLink",
			profile: testUtils.InitProfileLink("newLink"),

			expectedError: apperror.InternalServerError,

			setupMock: func(repo *mock_repository.MockUserRepository, email, avatar, link string, profile *dto.EditProfile) {
				repo.EXPECT().CheckLinkExists(link).Return(false, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockUserRepository(ctrl)

			userService := UserService{
				repo: mockRepository,
			}

			test.setupMock(mockRepository, test.email, test.avatar, test.link, test.profile)
			err := userService.EditProfile(test.email, test.profile)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestUserService_Subscribe(t *testing.T) {
	tests := []struct {
		name       string
		email      string
		followLink string
		reqTime    string

		expectedError error

		setupMock func(repo *mock_repository.MockUserRepository, subEmail, followLink, reqTime string)
	}{
		{
			name:       "Success",
			email:      "e-larkin@mail.ru",
			followLink: "id100",
			reqTime:    utils.CurrentTimeString(),

			expectedError: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, subEmail, followLink, reqTime string) {
				repo.EXPECT().Subscribe(subEmail, followLink, reqTime).Return(true, nil)
			},
		},
		{
			name:       "Internal error",
			email:      "e-larkin@mail.ru",
			followLink: "id100",
			reqTime:    utils.CurrentTimeString(),

			expectedError: apperror.InternalServerError,

			setupMock: func(repo *mock_repository.MockUserRepository, subEmail, followLink, reqTime string) {
				repo.EXPECT().Subscribe(subEmail, followLink, reqTime).Return(false, apperror.InternalServerError)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockUserRepository(ctrl)

			userService := UserService{
				repo: mockRepository,
			}

			test.setupMock(mockRepository, test.email, test.followLink, test.reqTime)
			err := userService.Subscribe(test.email, test.followLink)
			require.Equal(t, test.expectedError, err)
		})
	}
}

func TestUserService_Unsubscribe(t *testing.T) {
	tests := []struct {
		name       string
		email      string
		followLink string

		expectedError error

		setupMock func(repo *mock_repository.MockUserRepository, subEmail, followLink string)
	}{
		{
			name:       "Success",
			email:      "e-larkin@mail.ru",
			followLink: "id100",

			expectedError: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, subEmail, followLink string) {
				repo.EXPECT().Unsubscribe(subEmail, followLink).Return(true, nil)
			},
		},
		{
			name:       "Internal error",
			email:      "e-larkin@mail.ru",
			followLink: "id100",

			expectedError: apperror.InternalServerError,

			setupMock: func(repo *mock_repository.MockUserRepository, subEmail, followLink string) {
				repo.EXPECT().Unsubscribe(subEmail, followLink).Return(false, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockUserRepository(ctrl)

			userService := UserService{
				repo: mockRepository,
			}

			test.setupMock(mockRepository, test.email, test.followLink)
			err := userService.Unsubscribe(test.email, test.followLink)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestUserService_Reject(t *testing.T) {
	tests := []struct {
		name         string
		rejectEmail  string
		followerLink string

		expectedError error

		setupMock func(repo *mock_repository.MockUserRepository, rejectEmail, followerLink string)
	}{
		{
			name:         "Success",
			rejectEmail:  "email-to-reject@mail.ru",
			followerLink: "id1",

			expectedError: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, rejectEmail, followerLink string) {
				repo.EXPECT().RejectFriendRequest(rejectEmail, followerLink).Return(nil)
			},
		},
		{
			name:         "Internal Error",
			rejectEmail:  "email-to-reject@mail.ru",
			followerLink: "id1",

			expectedError: apperror.InternalServerError,

			setupMock: func(repo *mock_repository.MockUserRepository, rejectEmail, followerLink string) {
				repo.EXPECT().RejectFriendRequest(rejectEmail, followerLink).Return(apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockUserRepository(ctrl)

			userService := UserService{
				repo: mockRepository,
			}

			test.setupMock(mockRepository, test.rejectEmail, test.followerLink)
			err := userService.Reject(test.rejectEmail, test.followerLink)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestUserService_GetFriendsByEmail(t *testing.T) {
	tests := []struct {
		name   string
		email  string
		limit  int
		offset int

		expectedUsers []*entities.User
		expectedError error

		setupMock func(repo *mock_repository.MockUserRepository, email string, limit, offset int, expected []*entities.User)
	}{
		{
			name:   "Success",
			email:  "e-larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedError: nil,
			expectedUsers: []*entities.User{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
			},

			setupMock: func(repo *mock_repository.MockUserRepository, email string, limit, offset int, expectedUsers []*entities.User) {
				returnedUser := &entities.User{
					Email: email,
				}
				repo.EXPECT().GetUserByEmail(email).Return(returnedUser, nil)
				repo.EXPECT().GetFriends(returnedUser, limit, offset).Return(expectedUsers, nil)
			},
		},
		{
			name:   "Error get user",
			email:  "e-larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedError: apperror.InternalServerError,
			expectedUsers: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, email string, limit, offset int, expectedUsers []*entities.User) {
				repo.EXPECT().GetUserByEmail(email).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
		{
			name:   "Error get friends",
			email:  "e-larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedError: apperror.InternalServerError,
			expectedUsers: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, email string, limit, offset int, expectedUsers []*entities.User) {
				returnedUser := &entities.User{
					Email: email,
				}
				repo.EXPECT().GetUserByEmail(email).Return(returnedUser, nil)
				repo.EXPECT().GetFriends(returnedUser, limit, offset).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockUserRepository(ctrl)

			userService := UserService{
				repo: mockRepository,
			}

			test.setupMock(mockRepository, test.email, test.limit, test.offset, test.expectedUsers)
			users, err := userService.GetFriendsByEmail(test.email, test.limit, test.offset)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.expectedError, err)
			}
			if test.expectedUsers != nil {
				require.Equal(t, test.limit, len(users))
			}
		})
	}
}

func TestUserService_GetSubscribesByEmail(t *testing.T) {
	tests := []struct {
		name   string
		email  string
		limit  int
		offset int

		expectedUsers []*entities.User
		expectedError error

		setupMock func(repo *mock_repository.MockUserRepository, email string, limit, offset int, expected []*entities.User)
	}{
		{
			name:   "Success",
			email:  "e-larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedError: nil,
			expectedUsers: []*entities.User{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
			},

			setupMock: func(repo *mock_repository.MockUserRepository, email string, limit, offset int, expectedUsers []*entities.User) {
				returnedUser := &entities.User{
					Email: email,
				}
				repo.EXPECT().GetUserByEmail(email).Return(returnedUser, nil)
				repo.EXPECT().GetSubscribes(returnedUser, limit, offset).Return(expectedUsers, nil)
			},
		},
		{
			name:   "Error get user",
			email:  "e-larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedError: apperror.InternalServerError,
			expectedUsers: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, email string, limit, offset int, expectedUsers []*entities.User) {
				repo.EXPECT().GetUserByEmail(email).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
		{
			name:   "Error get friends",
			email:  "e-larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedError: apperror.InternalServerError,
			expectedUsers: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, email string, limit, offset int, expectedUsers []*entities.User) {
				returnedUser := &entities.User{
					Email: email,
				}
				repo.EXPECT().GetUserByEmail(email).Return(returnedUser, nil)
				repo.EXPECT().GetSubscribes(returnedUser, limit, offset).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockUserRepository(ctrl)

			userService := UserService{
				repo: mockRepository,
			}

			test.setupMock(mockRepository, test.email, test.limit, test.offset, test.expectedUsers)
			users, err := userService.GetSubscribesByEmail(test.email, test.limit, test.offset)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.expectedError, err)
			}
			if test.expectedUsers != nil {
				require.Equal(t, test.limit, len(users))
			}
		})
	}
}

func TestUserService_GetSubscribersByEmail(t *testing.T) {
	tests := []struct {
		name   string
		email  string
		limit  int
		offset int

		expectedUsers []*entities.User
		expectedError error

		setupMock func(repo *mock_repository.MockUserRepository, email string, limit, offset int, expected []*entities.User)
	}{
		{
			name:   "Success",
			email:  "e-larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedError: nil,
			expectedUsers: []*entities.User{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
			},

			setupMock: func(repo *mock_repository.MockUserRepository, email string, limit, offset int, expectedUsers []*entities.User) {
				returnedUser := &entities.User{
					Email: email,
				}
				repo.EXPECT().GetUserByEmail(email).Return(returnedUser, nil)
				repo.EXPECT().GetSubscribers(returnedUser, limit, offset).Return(expectedUsers, nil)
			},
		},
		{
			name:   "Error get user",
			email:  "e-larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedError: apperror.InternalServerError,
			expectedUsers: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, email string, limit, offset int, expectedUsers []*entities.User) {
				repo.EXPECT().GetUserByEmail(email).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
		{
			name:   "Error get friends",
			email:  "e-larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedError: apperror.InternalServerError,
			expectedUsers: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, email string, limit, offset int, expectedUsers []*entities.User) {
				returnedUser := &entities.User{
					Email: email,
				}
				repo.EXPECT().GetUserByEmail(email).Return(returnedUser, nil)
				repo.EXPECT().GetSubscribers(returnedUser, limit, offset).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockUserRepository(ctrl)

			userService := UserService{
				repo: mockRepository,
			}

			test.setupMock(mockRepository, test.email, test.limit, test.offset, test.expectedUsers)
			users, err := userService.GetSubscribersByEmail(test.email, test.limit, test.offset)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.expectedError, err)
			}
			if test.expectedUsers != nil {
				require.Equal(t, test.limit, len(users))
			}
		})
	}
}

func TestUserService_GetPendingRequestsByEmail(t *testing.T) {
	tests := []struct {
		name   string
		email  string
		limit  int
		offset int

		expectedUsers []*entities.User
		expectedError error

		setupMock func(repo *mock_repository.MockUserRepository, email string, limit, offset int, expected []*entities.User)
	}{
		{
			name:   "Success",
			email:  "e-larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedError: nil,
			expectedUsers: []*entities.User{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
			},

			setupMock: func(repo *mock_repository.MockUserRepository, email string, limit, offset int, expectedUsers []*entities.User) {
				returnedUser := &entities.User{
					Email: email,
				}
				repo.EXPECT().GetUserByEmail(email).Return(returnedUser, nil)
				repo.EXPECT().GetPendingFriendRequests(returnedUser, limit, offset).Return(expectedUsers, nil)
			},
		},
		{
			name:   "Error get user",
			email:  "e-larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedError: apperror.InternalServerError,
			expectedUsers: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, email string, limit, offset int, expectedUsers []*entities.User) {
				repo.EXPECT().GetUserByEmail(email).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
		{
			name:   "Error get friends",
			email:  "e-larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedError: apperror.InternalServerError,
			expectedUsers: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, email string, limit, offset int, expectedUsers []*entities.User) {
				returnedUser := &entities.User{
					Email: email,
				}
				repo.EXPECT().GetUserByEmail(email).Return(returnedUser, nil)
				repo.EXPECT().GetPendingFriendRequests(returnedUser, limit, offset).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockUserRepository(ctrl)

			userService := UserService{
				repo: mockRepository,
			}

			test.setupMock(mockRepository, test.email, test.limit, test.offset, test.expectedUsers)
			users, err := userService.GetPendingRequestsByEmail(test.email, test.limit, test.offset)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.expectedError, err)
			}
			if test.expectedUsers != nil {
				require.Equal(t, test.limit, len(users))
			}
		})
	}
}

func TestUserService_GetFriendsByLink(t *testing.T) {
	tests := []struct {
		name         string
		requestEmail string
		targetLink   string

		limit  int
		offset int

		isNumericLink bool
		id            uint

		expectedError error
		expectedUsers []*entities.User

		setupMock func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int,
			isNumeric bool, id uint,
			expectedUsers []*entities.User)
	}{
		{
			name:         "Success numeric link",
			requestEmail: "e-larkin@mail.ru",
			targetLink:   "id123",
			limit:        2,
			offset:       0,

			isNumericLink: true,
			id:            123,
			expectedError: nil,
			expectedUsers: []*entities.User{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
			},
			setupMock: func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int, isNumeric bool, id uint, expectedUsers []*entities.User) {
				returnedUser := &entities.User{Link: targetLink}
				repo.EXPECT().GetUserById(id).Return(returnedUser, nil)
				repo.EXPECT().GetFriends(returnedUser, limit, offset).Return(expectedUsers, nil)
			},
		},
		{
			name:         "BadRequest numeric link",
			requestEmail: "e-larkin@mail.ru",
			targetLink:   "idaaaa",
			limit:        2,
			offset:       0,

			isNumericLink: true,
			expectedError: apperror.BadRequest,
			expectedUsers: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int, isNumeric bool, id uint, expectedUsers []*entities.User) {

			},
		},
		{
			name:         "Error Get User numeric link",
			requestEmail: "e-larkin@mail.ru",
			targetLink:   "id123",
			limit:        2,
			offset:       0,

			isNumericLink: true,
			id:            123,
			expectedError: apperror.InternalServerError,
			expectedUsers: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int, isNumeric bool, id uint, expectedUsers []*entities.User) {
				repo.EXPECT().GetUserById(id).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
		{
			name:         "Error GetFriends numeric link",
			requestEmail: "e-larkin@mail.ru",
			targetLink:   "id123",
			limit:        2,
			offset:       0,

			isNumericLink: true,
			id:            123,
			expectedError: apperror.InternalServerError,
			expectedUsers: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int, isNumeric bool, id uint, expectedUsers []*entities.User) {
				returnedUser := &entities.User{Link: targetLink}
				repo.EXPECT().GetUserById(id).Return(returnedUser, nil)
				repo.EXPECT().GetFriends(returnedUser, limit, offset).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
		{
			name:         "Success custom link",
			requestEmail: "e-larkin@mail.ru",
			targetLink:   "egor_larkin",
			limit:        2,
			offset:       0,

			isNumericLink: false,
			expectedError: nil,
			expectedUsers: []*entities.User{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
			},
			setupMock: func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int, isNumeric bool, id uint, expectedUsers []*entities.User) {
				returnedUser := &entities.User{Link: targetLink}
				repo.EXPECT().GetUserByLink(targetLink).Return(returnedUser, nil)
				repo.EXPECT().GetFriends(returnedUser, limit, offset).Return(expectedUsers, nil)
			},
		},
		{
			name:         "Error GetUser custom link",
			requestEmail: "e-larkin@mail.ru",
			targetLink:   "egor_larkin",
			limit:        2,
			offset:       0,

			isNumericLink: false,
			expectedError: apperror.InternalServerError,
			expectedUsers: nil,
			setupMock: func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int, isNumeric bool, id uint, expectedUsers []*entities.User) {
				repo.EXPECT().GetUserByLink(targetLink).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
		{
			name:         "Error GetFriends custom link",
			requestEmail: "e-larkin@mail.ru",
			targetLink:   "egor_larkin",
			limit:        2,
			offset:       0,

			isNumericLink: false,
			expectedError: apperror.InternalServerError,
			expectedUsers: nil,
			setupMock: func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int, isNumeric bool, id uint, expectedUsers []*entities.User) {
				returnedUser := &entities.User{Link: targetLink}
				repo.EXPECT().GetUserByLink(targetLink).Return(returnedUser, nil)
				repo.EXPECT().GetFriends(returnedUser, limit, offset).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test := test
			t.Run(test.name, func(t *testing.T) {
				t.Parallel()
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				mockRepository := mock_repository.NewMockUserRepository(ctrl)

				userService := UserService{
					repo: mockRepository,
				}

				test.setupMock(mockRepository, test.requestEmail, test.targetLink, test.limit, test.offset,
					test.isNumericLink, test.id,
					test.expectedUsers)
				users, err := userService.GetFriendsByLink(test.requestEmail, test.targetLink, test.limit, test.offset)
				if test.expectedError != nil {
					uerr, ok := err.(*apperror.ServerError)
					require.Equal(t, true, ok)
					require.Equal(t, test.expectedError, uerr.UserErr)
				} else {
					require.Equal(t, test.expectedError, err)
				}
				if test.expectedUsers != nil {
					require.Equal(t, test.limit, len(users))
				}
			})
		})
	}
}

func TestUserService_GetSubscribesByLink(t *testing.T) {
	tests := []struct {
		name         string
		requestEmail string
		targetLink   string

		limit  int
		offset int

		isNumericLink bool
		id            uint

		expectedError error
		expectedUsers []*entities.User

		setupMock func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int,
			isNumeric bool, id uint,
			expectedUsers []*entities.User)
	}{
		{
			name:         "Success numeric link",
			requestEmail: "e-larkin@mail.ru",
			targetLink:   "id123",
			limit:        2,
			offset:       0,

			isNumericLink: true,
			id:            123,
			expectedError: nil,
			expectedUsers: []*entities.User{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
			},
			setupMock: func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int, isNumeric bool, id uint, expectedUsers []*entities.User) {
				returnedUser := &entities.User{Link: targetLink}
				repo.EXPECT().GetUserById(id).Return(returnedUser, nil)
				repo.EXPECT().GetSubscribes(returnedUser, limit, offset).Return(expectedUsers, nil)
			},
		},
		{
			name:         "BadRequest numeric link",
			requestEmail: "e-larkin@mail.ru",
			targetLink:   "idaaaa",
			limit:        2,
			offset:       0,

			isNumericLink: true,
			expectedError: apperror.BadRequest,
			expectedUsers: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int, isNumeric bool, id uint, expectedUsers []*entities.User) {

			},
		},
		{
			name:         "Error Get User numeric link",
			requestEmail: "e-larkin@mail.ru",
			targetLink:   "id123",
			limit:        2,
			offset:       0,

			isNumericLink: true,
			id:            123,
			expectedError: apperror.InternalServerError,
			expectedUsers: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int, isNumeric bool, id uint, expectedUsers []*entities.User) {
				repo.EXPECT().GetUserById(id).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
		{
			name:         "Error GetSubscribes numeric link",
			requestEmail: "e-larkin@mail.ru",
			targetLink:   "id123",
			limit:        2,
			offset:       0,

			isNumericLink: true,
			id:            123,
			expectedError: apperror.InternalServerError,
			expectedUsers: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int, isNumeric bool, id uint, expectedUsers []*entities.User) {
				returnedUser := &entities.User{Link: targetLink}
				repo.EXPECT().GetUserById(id).Return(returnedUser, nil)
				repo.EXPECT().GetSubscribes(returnedUser, limit, offset).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
		{
			name:         "Success custom link",
			requestEmail: "e-larkin@mail.ru",
			targetLink:   "egor_larkin",
			limit:        2,
			offset:       0,

			isNumericLink: false,
			expectedError: nil,
			expectedUsers: []*entities.User{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
			},
			setupMock: func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int, isNumeric bool, id uint, expectedUsers []*entities.User) {
				returnedUser := &entities.User{Link: targetLink}
				repo.EXPECT().GetUserByLink(targetLink).Return(returnedUser, nil)
				repo.EXPECT().GetSubscribes(returnedUser, limit, offset).Return(expectedUsers, nil)
			},
		},
		{
			name:         "Error GetUser custom link",
			requestEmail: "e-larkin@mail.ru",
			targetLink:   "egor_larkin",
			limit:        2,
			offset:       0,

			isNumericLink: false,
			expectedError: apperror.InternalServerError,
			expectedUsers: nil,
			setupMock: func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int, isNumeric bool, id uint, expectedUsers []*entities.User) {
				repo.EXPECT().GetUserByLink(targetLink).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
		{
			name:         "Error GetSubscribes custom link",
			requestEmail: "e-larkin@mail.ru",
			targetLink:   "egor_larkin",
			limit:        2,
			offset:       0,

			isNumericLink: false,
			expectedError: apperror.InternalServerError,
			expectedUsers: nil,
			setupMock: func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int, isNumeric bool, id uint, expectedUsers []*entities.User) {
				returnedUser := &entities.User{Link: targetLink}
				repo.EXPECT().GetUserByLink(targetLink).Return(returnedUser, nil)
				repo.EXPECT().GetSubscribes(returnedUser, limit, offset).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test := test
			t.Run(test.name, func(t *testing.T) {
				t.Parallel()
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				mockRepository := mock_repository.NewMockUserRepository(ctrl)

				userService := UserService{
					repo: mockRepository,
				}

				test.setupMock(mockRepository, test.requestEmail, test.targetLink, test.limit, test.offset,
					test.isNumericLink, test.id,
					test.expectedUsers)
				users, err := userService.GetSubscribesByLink(test.requestEmail, test.targetLink, test.limit, test.offset)
				if test.expectedError != nil {
					uerr, ok := err.(*apperror.ServerError)
					require.Equal(t, true, ok)
					require.Equal(t, test.expectedError, uerr.UserErr)
				} else {
					require.Equal(t, test.expectedError, err)
				}
				if test.expectedUsers != nil {
					require.Equal(t, test.limit, len(users))
				}
			})
		})
	}
}

func TestUserService_GetSubscribersByLink(t *testing.T) {
	tests := []struct {
		name         string
		requestEmail string
		targetLink   string

		limit  int
		offset int

		isNumericLink bool
		id            uint

		expectedError error
		expectedUsers []*entities.User

		setupMock func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int,
			isNumeric bool, id uint,
			expectedUsers []*entities.User)
	}{
		{
			name:         "Success numeric link",
			requestEmail: "e-larkin@mail.ru",
			targetLink:   "id123",
			limit:        2,
			offset:       0,

			isNumericLink: true,
			id:            123,
			expectedError: nil,
			expectedUsers: []*entities.User{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
			},
			setupMock: func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int, isNumeric bool, id uint, expectedUsers []*entities.User) {
				returnedUser := &entities.User{Link: targetLink}
				repo.EXPECT().GetUserById(id).Return(returnedUser, nil)
				repo.EXPECT().GetSubscribers(returnedUser, limit, offset).Return(expectedUsers, nil)
			},
		},
		{
			name:         "BadRequest numeric link",
			requestEmail: "e-larkin@mail.ru",
			targetLink:   "idaaaa",
			limit:        2,
			offset:       0,

			isNumericLink: true,
			expectedError: apperror.BadRequest,
			expectedUsers: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int, isNumeric bool, id uint, expectedUsers []*entities.User) {

			},
		},
		{
			name:         "Error Get User numeric link",
			requestEmail: "e-larkin@mail.ru",
			targetLink:   "id123",
			limit:        2,
			offset:       0,

			isNumericLink: true,
			id:            123,
			expectedError: apperror.InternalServerError,
			expectedUsers: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int, isNumeric bool, id uint, expectedUsers []*entities.User) {
				repo.EXPECT().GetUserById(id).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
		{
			name:         "Error GetSubscribers numeric link",
			requestEmail: "e-larkin@mail.ru",
			targetLink:   "id123",
			limit:        2,
			offset:       0,

			isNumericLink: true,
			id:            123,
			expectedError: apperror.InternalServerError,
			expectedUsers: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int, isNumeric bool, id uint, expectedUsers []*entities.User) {
				returnedUser := &entities.User{Link: targetLink}
				repo.EXPECT().GetUserById(id).Return(returnedUser, nil)
				repo.EXPECT().GetSubscribers(returnedUser, limit, offset).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
		{
			name:         "Success custom link",
			requestEmail: "e-larkin@mail.ru",
			targetLink:   "egor_larkin",
			limit:        2,
			offset:       0,

			isNumericLink: false,
			expectedError: nil,
			expectedUsers: []*entities.User{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
			},
			setupMock: func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int, isNumeric bool, id uint, expectedUsers []*entities.User) {
				returnedUser := &entities.User{Link: targetLink}
				repo.EXPECT().GetUserByLink(targetLink).Return(returnedUser, nil)
				repo.EXPECT().GetSubscribers(returnedUser, limit, offset).Return(expectedUsers, nil)
			},
		},
		{
			name:         "Error GetUser custom link",
			requestEmail: "e-larkin@mail.ru",
			targetLink:   "egor_larkin",
			limit:        2,
			offset:       0,

			isNumericLink: false,
			expectedError: apperror.InternalServerError,
			expectedUsers: nil,
			setupMock: func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int, isNumeric bool, id uint, expectedUsers []*entities.User) {
				repo.EXPECT().GetUserByLink(targetLink).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
		{
			name:         "Error GetSubscribers custom link",
			requestEmail: "e-larkin@mail.ru",
			targetLink:   "egor_larkin",
			limit:        2,
			offset:       0,

			isNumericLink: false,
			expectedError: apperror.InternalServerError,
			expectedUsers: nil,
			setupMock: func(repo *mock_repository.MockUserRepository, requestEmail, targetLink string, limit, offset int, isNumeric bool, id uint, expectedUsers []*entities.User) {
				returnedUser := &entities.User{Link: targetLink}
				repo.EXPECT().GetUserByLink(targetLink).Return(returnedUser, nil)
				repo.EXPECT().GetSubscribers(returnedUser, limit, offset).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test := test
			t.Run(test.name, func(t *testing.T) {
				t.Parallel()
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				mockRepository := mock_repository.NewMockUserRepository(ctrl)

				userService := UserService{
					repo: mockRepository,
				}

				test.setupMock(mockRepository, test.requestEmail, test.targetLink, test.limit, test.offset,
					test.isNumericLink, test.id,
					test.expectedUsers)
				users, err := userService.GetSubscribersByLink(test.requestEmail, test.targetLink, test.limit, test.offset)
				if test.expectedError != nil {
					uerr, ok := err.(*apperror.ServerError)
					require.Equal(t, true, ok)
					require.Equal(t, test.expectedError, uerr.UserErr)
				} else {
					require.Equal(t, test.expectedError, err)
				}
				if test.expectedUsers != nil {
					require.Equal(t, test.limit, len(users))
				}
			})
		})
	}
}
