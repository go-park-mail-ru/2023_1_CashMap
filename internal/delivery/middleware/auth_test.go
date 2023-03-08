package middleware

import (
	"depeche/internal/entities"
	storage "depeche/internal/repository/local_storage"
	"depeche/internal/session"
	"depeche/pkg/apperror"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var authTestcases = map[string]struct {
	cookieName  string
	cookieValue string
	code        int
	err         error
}{
	"with cookie": {
		cookieName:  "session_id",
		cookieValue: "expected",
		code:        http.StatusOK,
	},
	"without cookie": {
		cookieName: "invalid",
		code:       http.StatusUnauthorized,
		err:        apperror.NoAuth,
	},
}

var router = gin.Default()

func TestMain(m *testing.M) {
	storage.NewUserStorage()

	authMW := AuthMiddleware{
		service: &TestAuthService{
			session: sessions,
		},
	}
	router.Use(authMW.Middleware())
	router.Use(ErrorMiddleware())
	router.GET("/", func(context *gin.Context) {

	})
	os.Exit(m.Run())
}
func TestErrorMiddleware(t *testing.T) {
	for name, test := range authTestcases {
		t.Run(name, func(t *testing.T) {

			request := httptest.NewRequest(http.MethodGet, "/", nil)
			request.AddCookie(&http.Cookie{
				Name:  test.cookieName,
				Value: test.cookieValue,
			})
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, request)

			require.Equal(t, test.code, rr.Code)

			if test.err != nil {
				body, err := json.Marshal(gin.H{
					"status":  Errors[test.err].Code,
					"message": Errors[test.err].Message,
				})
				require.NoError(t, err)
				require.Equal(t, body, rr.Body.Bytes())
			}
		})
	}
}

type TestAuthService struct {
	session map[string]*session.Session
}

func (t *TestAuthService) CheckSession(token string) (*session.Session, error) {
	stored := t.session[token]
	if stored == nil {
		return nil, apperror.NoAuth
	}
	return stored, nil
}

func (t *TestAuthService) Authenticate(user *entities.User) (string, error) {
	return "", nil
}

func (t *TestAuthService) LogOut(token string) error {
	return nil
}

var sessions = map[string]*session.Session{
	"expected": {
		Email:     "user1@email.com",
		ExpiresAt: time.Date(2023, time.March, 30, 20, 20, 20, 20, time.Local),
	},
}
