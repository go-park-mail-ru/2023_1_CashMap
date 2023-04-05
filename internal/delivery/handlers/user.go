package handlers

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/session"
	authService "depeche/internal/session/service"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
	"net/http"
	"strconv"
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
		Data dto.SignIn `json:"body"`
	}{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}

	_, err = uh.service.SignIn(&request.Data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	token, err := uh.authService.Authenticate(&request.Data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	sessionCookie := &http.Cookie{
		Name:     "session_id",
		Value:    token,
		Expires:  time.Now().Add(time.Second * 86400),
		MaxAge:   86400,
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
		Data dto.SignUp `json:"body"`
	}{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}

	_, err = uh.service.SignUp(&request.Data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	token, err := uh.authService.Authenticate(&request.Data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	sessionCookie := &http.Cookie{
		Name:     "session_id",
		Value:    token,
		Expires:  time.Now().Add(time.Second * 86400),
		MaxAge:   86400,
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

func (uh *UserHandler) SubscribeHandler(subType int) gin.HandlerFunc {

	var useCase func(string, string) error
	switch subType {
	case Subscribe:
		useCase = uh.service.Subscribe
	case Unsubscribe:
		useCase = uh.service.Unsubscribe
	case Reject:
		useCase = uh.service.Reject
	}

	return func(ctx *gin.Context) {

		stored, err := uh.getSession(ctx)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		var request = struct {
			Data dto.Subscribes `json:"body"`
		}{}

		err = ctx.ShouldBindJSON(&request)
		if err != nil {
			_ = ctx.Error(apperror.BadRequest)
			return
		}
		email := stored.Email
		err = useCase(email, request.Data.Link)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
	}
}

const (
	Subscribe = iota
	Unsubscribe
	Reject
)

func (uh *UserHandler) Profile(ctx *gin.Context) {
	link := ctx.Param("link")

	profile, err := uh.service.GetProfileByLink("", link)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"profile": profile,
		},
	})
}

func (uh *UserHandler) Self(ctx *gin.Context) {
	stored, err := uh.getSession(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	email := stored.Email

	profile, err := uh.service.GetProfileByEmail(email)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"profile": profile,
		},
	})
}

func (uh *UserHandler) EditProfile(ctx *gin.Context) {
	stored, err := uh.getSession(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	email := stored.Email
	request := struct {
		Data dto.EditProfile `json:"body"`
	}{}

	err = uh.service.EditProfile(email, &request.Data)
	if err != nil {
	}
}

func (uh *UserHandler) Friends(ctx *gin.Context) {
	link := ctx.Query("link")
	stored, err := uh.getSession(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	email := stored.Email

	limitQ := ctx.Query("limit")
	offsetQ := ctx.Query("offset")

	limit, err := strconv.Atoi(limitQ)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}
	offset, err := strconv.Atoi(offsetQ)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}

	friends, err := uh.service.GetFriendsByLink(email, link, limit, offset)
	if err != nil {
		_ = ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"friends": friends,
		},
	})
}

func (uh *UserHandler) Subscribes(ctx *gin.Context) {
	subType := ctx.Query("type")
	link := ctx.Query("link")

	stored, err := uh.getSession(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	email := stored.Email
	limitQ := ctx.Query("limit")
	offsetQ := ctx.Query("offset")

	limit, err := strconv.Atoi(limitQ)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}
	offset, err := strconv.Atoi(offsetQ)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}

	var subs []*entities.User

	switch subType {
	case "in":
		subs, err = uh.service.GetSubscribersByLink(email, link, limit, offset)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
	case "out":
		subs, err = uh.service.GetSubscribesByLink(email, link, limit, offset)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
	default:
		_ = ctx.Error(apperror.BadRequest)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"subs": subs,
		},
	})
}

func (uh *UserHandler) getSession(ctx *gin.Context) (*session.Session, error) {
	token, err := ctx.Cookie("session_id")
	if err != nil {

		return nil, apperror.NoAuth
	}
	stored, err := uh.authService.CheckSession(token)
	if err != nil {
		return nil, err
	}
	return stored, nil
}
