package handlers

import (
	"depeche/internal/entities"
	authService "depeche/internal/session/service"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service     usecase.User
	authService authService.Auth
}

func NewUserHandler(userService usecase.User, authService authService.Auth) *UserHandler {
	return &UserHandler{
		service:     userService,
		authService: authService,
	}
}

// SignIn godoc
//
//	@Summary		Sign in
//	@Description	Authorize client with credentials (login and password).
//	@Tags			signin
//	@Accept			json
//	@Param			login		body	string	true	"User login"
//	@Param			password	body	string	true	"User password"
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/auth/sign-in [post]
func (uh *UserHandler) SignIn(ctx *gin.Context) {
	var request = struct {
		User entities.User `json:"body"`
	}{}

	err := ctx.BindJSON(&request)
	fmt.Print(request.User)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}

	_, err = uh.service.SignIn(&request.User)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	token, err := uh.authService.Authenticate(&request.User)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	sessionCookie := &http.Cookie{
		Name:     "session_id",
		Value:    token,
		Expires:  time.Now().Add(time.Second * 86400),
		MaxAge:   0,
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	}

	http.SetCookie(ctx.Writer, sessionCookie)
	ctx.Status(http.StatusOK)
}

// SignUp godoc
//
//	@Summary		Sign up
//	@Description	Register client with credentials and other user info.
//	@Tags			signup
//	@Accept			json
//	@Param			email		body	string	true	"User email"
//	@Param			password	body	string	true	"User password"
//	@Param			first_name	body	string	true	"User first name"
//	@Param			last_name	body	string	true	"User last name"
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/auth/sign-up [post]
func (uh *UserHandler) SignUp(ctx *gin.Context) {
	var request = struct {
		User entities.User `json:"body"`
	}{}
	err := ctx.BindJSON(&request)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}

	_, err = uh.service.SignUp(&request.User)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	token, err := uh.authService.Authenticate(&request.User)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	sessionCookie := &http.Cookie{
		Name:     "session_id",
		Value:    token,
		Expires:  time.Now().Add(time.Second * 86400),
		MaxAge:   0,
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	}

	http.SetCookie(ctx.Writer, sessionCookie)
	ctx.Status(http.StatusOK)

}

// LogOut godoc
// @Summary		Log out
// @Description	Delete user session and invalidate session cookie
// @Tags			logout
// @Success		200
// @Failure		400
// @Failure		500
// @Router			/auth/logout [post]
func (uh *UserHandler) LogOut(ctx *gin.Context) {
	token, err := ctx.Cookie("session_id")
	if err != nil {
		_ = ctx.Error(apperror.NoAuth)
		return
	}
	err = uh.authService.LogOut(token)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	expired := &http.Cookie{
		Name:     "session_id",
		Value:    token,
		Expires:  time.Now().Add(-time.Hour),
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	}

	http.SetCookie(ctx.Writer, expired)
	ctx.Status(http.StatusOK)
}

func (uh *UserHandler) CheckAuth(ctx *gin.Context) {
	token, err := ctx.Cookie("session_id")
	if err != nil {
		_ = ctx.Error(apperror.NoAuth)
		return
	}
	_, err = uh.authService.CheckSession(token)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
