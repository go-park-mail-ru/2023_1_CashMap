package handlers

import (
	"depeche/internal/entities"
	"depeche/internal/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service     usecase.User
	authService usecase.Auth
}

func NewUserHandler(userService usecase.User, authService usecase.Auth) *UserHandler {
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
	var user entities.User

	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	_, err = uh.service.SignIn(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	token, err := uh.authService.Authenticate(&user)
	sessionCookie := &http.Cookie{
		Name:     "session_id",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 1000), // ЗАХАРДКОЖЕНО ВРЕМЯ ЭКСПИРАЦИИ
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
	var user entities.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// TODO: уточнить код ответа (сейчас летит 400)
	_, err = uh.service.SignUp(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
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
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = uh.authService.LogOut(token)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newCookie := &http.Cookie{
		Name:     "session_id",
		Value:    token,
		Expires:  time.Now().Add(-time.Hour), // ЗАХАРДКОЖЕНО ВРЕМЯ ЭКСПИРАЦИИ
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	}

	http.SetCookie(ctx.Writer, newCookie)
	ctx.Status(http.StatusOK)
}

func (uh *UserHandler) CheckAuth(ctx *gin.Context) {
	token, err := ctx.Cookie("session_id")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	_, err = uh.authService.CheckSession(token)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	ctx.Status(http.StatusOK)
}
