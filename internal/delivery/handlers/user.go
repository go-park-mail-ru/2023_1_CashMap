package handlers

import (
	"depeche/authorization_ms/authEntities"
	auth "depeche/authorization_ms/service"
	"depeche/internal/delivery/dto"
	"depeche/internal/delivery/utils"
	"depeche/internal/entities"
	_ "depeche/internal/entities/doc"
	"depeche/internal/entities/response"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
	"errors"
	"fmt"
	"github.com/mailru/easyjson"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service     usecase.User
	authService auth.Auth
	csrfService auth.CSRFUsecase
}

func NewUserHandler(userService usecase.User, authService auth.Auth, csrfService auth.CSRFUsecase) *UserHandler {
	return &UserHandler{
		service:     userService,
		authService: authService,
		csrfService: csrfService,
	}
}

// SignIn godoc
//
//	@Summary		Sign in
//	@Description	Authorize client with credentials (login and password).
//	@Tags			Auth
//	@Accept			json
//	@Param			request	body	doc.SignIn	true	"User credentials"
//	@Success		200
//	@Failure		400	{object}	middleware.ErrorResponse
//	@Failure		404	{object}	middleware.ErrorResponse
//	@Failure		409	{object}	middleware.ErrorResponse
//	@Failure		500	{object}	middleware.ErrorResponse
//	@Router			/auth/sign-in [post]
func (uh *UserHandler) SignIn(ctx *gin.Context) {
	inputDTO := new(response.SignInRequest)
	if err := easyjson.UnmarshalFromReader(ctx.Request.Body, inputDTO); err != nil || inputDTO.Body == nil {
		_ = ctx.Error(apperror.NewBadRequest())
		return
	}

	_, err := uh.service.SignIn(inputDTO.Body)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	token, err := uh.authService.Authenticate(inputDTO.Body.Email)
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

	csrfToken, err := uh.csrfService.CreateCSRFToken(inputDTO.Body.Email)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Header("X-Csrf-Token", csrfToken)

	ctx.Status(http.StatusOK)
}

// SignUp godoc
//
//	@Summary		Sign up
//	@Description	Register client with credentials and other user info.
//	@Tags			Auth
//	@Accept			json
//	@Param			email	body		doc.SignUp	true	"Required register fields"
//	@Failure		400		{object}	middleware.ErrorResponse
//	@Failure		401		{object}	middleware.ErrorResponse
//	@Failure		404		{object}	middleware.ErrorResponse
//	@Failure		500		{object}	middleware.ErrorResponse
//	@Router			/auth/sign-up [post]
func (uh *UserHandler) SignUp(ctx *gin.Context) {
	inputDTO := new(response.SignUpRequest)
	if err := easyjson.UnmarshalFromReader(ctx.Request.Body, inputDTO); err != nil || inputDTO.Body == nil {
		_ = ctx.Error(apperror.NewBadRequest())
		return
	}

	_, err := uh.service.SignUp(inputDTO.Body)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	token, err := uh.authService.Authenticate(inputDTO.Body.Email)
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

	csrfToken, err := uh.csrfService.CreateCSRFToken(inputDTO.Body.Email)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Header("X-Csrf-Token", csrfToken)

	ctx.Status(http.StatusOK)

}

// LogOut godoc
//
//	@Summary		Log out
//	@Description	Delete user authorization_ms and invalidate authorization_ms cookie
//	@Tags			Auth
//	@Success		200
//	@Failure		401	{object}	middleware.ErrorResponse
//	@Failure		404	{object}	middleware.ErrorResponse
//	@Router			/auth/logout [post]
func (uh *UserHandler) LogOut(ctx *gin.Context) {
	token, err := ctx.Cookie("session_id")
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.NoAuth, nil))
		return
	}

	// Игнорим ошибку ибо дальше все равно логаут
	userSession, err := uh.authService.CheckSession(token)
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.NoAuth, nil))
		return
	}

	csrfToken := ctx.Request.Header.Get("X-Csrf-Token")
	if csrfToken != "" {
		csrf := &authEntities.CSRF{
			Token: csrfToken,
			Email: userSession.Email,
		}
		err = uh.csrfService.InvalidateCSRFToken(csrf)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
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
		_ = ctx.Error(apperror.NewServerError(apperror.NoAuth, nil))
		return
	}

	userSession, err := uh.authService.CheckSession(token)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	csrfToken := ctx.Request.Header.Get("X-Csrf-Token")
	if csrfToken == "" {
		_ = ctx.Error(apperror.NewServerError(apperror.Forbidden, nil))
		return
	}

	csrf := &authEntities.CSRF{
		Token: csrfToken,
		Email: userSession.Email,
	}

	success, err := uh.csrfService.ValidateCSRFToken(csrf)
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.Forbidden, nil))
		return
	}

	if !success {
		newCsrfToken, err := uh.csrfService.CreateCSRFToken(userSession.Email)
		if err != nil {
			_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, nil))
			return
		}

		ctx.Header("X-Csrf-Token", newCsrfToken)
	}

	ctx.Status(http.StatusOK)
}

// Subscribe godoc
//
//	@Summary		Subscribe
//	@Description	Subscribe to other user
//	@Tags			Subscribes
//	@Accepts		json
//	@Param			request	body	doc.Subscribes	true	"Link to user to subscribe."
//	@Success		200
//	@Failure		400	{object}	middleware.ErrorResponse
//	@Failure		401	{object}	middleware.ErrorResponse
//	@Failure		409	{object}	middleware.ErrorResponse
//	@Failure		500	{object}	middleware.ErrorResponse
//	@Router			/api/user/sub [post]
func (uh *UserHandler) Subscribe(ctx *gin.Context) {
	inputDTO := new(response.SubscribeRequest)
	if err := easyjson.UnmarshalFromReader(ctx.Request.Body, inputDTO); err != nil || inputDTO.Body == nil {
		_ = ctx.Error(apperror.NewBadRequest())
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	err = uh.service.Subscribe(email, inputDTO.Body.Link)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
}

// Unsubscribe godoc
//
//	@Summary		Unsubscribe
//	@Description	Unsubscribe from other user
//	@Tags			Subscribes
//	@Accepts		json
//	@Param			request	body	doc.Subscribes	true	"Link to user to unsubscribe."
//	@Success		200
//	@Failure		400	{object}	middleware.ErrorResponse
//	@Failure		401	{object}	middleware.ErrorResponse
//	@Failure		409	{object}	middleware.ErrorResponse
//	@Failure		500	{object}	middleware.ErrorResponse
//	@Router			/api/user/unsub [post]
func (uh *UserHandler) Unsubscribe(ctx *gin.Context) {
	inputDTO := new(response.UnsubscribeRequest)
	if err := easyjson.UnmarshalFromReader(ctx.Request.Body, inputDTO); err != nil || inputDTO.Body == nil {
		_ = ctx.Error(apperror.NewBadRequest())
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	err = uh.service.Unsubscribe(email, inputDTO.Body.Link)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
}

// Reject godoc
//
//	@Summary		Reject
//	@Description	Reject friend request
//	@Tags			Subscribes
//	@Accepts		json
//	@Param			request	body	doc.Subscribes	true	"Link to user to reject friend request."
//	@Success		200
//	@Failure		400	{object}	middleware.ErrorResponse
//	@Failure		401	{object}	middleware.ErrorResponse
//	@Failure		409	{object}	middleware.ErrorResponse
//	@Failure		500	{object}	middleware.ErrorResponse
//	@Router			/api/user/reject [post]
func (uh *UserHandler) Reject(ctx *gin.Context) {
	inputDTO := new(response.RejectRequest)
	if err := easyjson.UnmarshalFromReader(ctx.Request.Body, inputDTO); err != nil || inputDTO.Body == nil {
		_ = ctx.Error(apperror.NewBadRequest())
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = uh.service.Reject(email, inputDTO.Body.Link)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
}

// Profile godoc
//
//	@Summary		Profile
//	@Description	Get profile by link
//	@Tags			Profiles
//	@Param			link	path		string	true	"link to requested profile"
//	@Success		200		{object}	doc.Profile
//	@Failure		400		{object}	middleware.ErrorResponse
//	@Failure		403		{object}	middleware.ErrorResponse
//	@Failure		404		{object}	middleware.ErrorResponse
//	@Failure		500		{object}	middleware.ErrorResponse
//	@Router			/api/user/profile/link [get]
func (uh *UserHandler) Profile(ctx *gin.Context) {
	link := ctx.Param("link")

	user, err := uh.service.GetProfileByLink("", link)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	profile := dto.NewProfileFromUser(user)

	_response := response.ProfileResponse{
		Body: response.ProfileBody{
			Profile: profile,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

// Self godoc
//
//	@Summary		Self
//	@Description	Get self profile
//	@Tags			Profiles
//	@Success		200	{object}	doc.Profile
//	@Failure		400	{object}	middleware.ErrorResponse
//	@Failure		403	{object}	middleware.ErrorResponse
//	@Failure		404	{object}	middleware.ErrorResponse
//	@Failure		500	{object}	middleware.ErrorResponse
//	@Router			/api/user/profile [get]
func (uh *UserHandler) Self(ctx *gin.Context) {
	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	profile, err := uh.service.GetProfileByEmail(email)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	_response := response.SelfResponse{
		Body: response.SelfBody{
			Profile: profile,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

// EditProfile godoc
//
//	@Summary		EditProfile
//	@Description	Edit profile
//	@Tags			Profiles
//	@Accept			json
//	@Param			request	body	doc.EditProfile	false	"Edited fields"
//	@Success		200
//	@Failure		400	{object}	middleware.ErrorResponse
//	@Failure		401	{object}	middleware.ErrorResponse
//	@Failure		404	{object}	middleware.ErrorResponse
//	@Failure		500	{object}	middleware.ErrorResponse
//	@Router			/api/user/profile/edit [patch]
func (uh *UserHandler) EditProfile(ctx *gin.Context) {
	inputDTO := new(response.EditProfileRequest)
	if err := easyjson.UnmarshalFromReader(ctx.Request.Body, inputDTO); err != nil || inputDTO.Body == nil {
		_ = ctx.Error(apperror.NewBadRequest())
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	err = uh.service.EditProfile(email, inputDTO.Body)
	if err != nil {
		_ = ctx.Error(err)
	}
}

// Friends godoc
//
//	@Summary		Friends
//	@Description	Get friends
//	@Tags			Profiles
//	@Success		200		{object}	doc.ProfileArray
//	@Param			link	query		string	true	"link to requested profile"
//	@Param			limit	query		int		true	"amount of profiles"
//	@Param			offset	query		int		true	"number of batch"
//	@Failure		400		{object}	middleware.ErrorResponse
//	@Failure		401		{object}	middleware.ErrorResponse
//	@Failure		404		{object}	middleware.ErrorResponse
//	@Failure		500		{object}	middleware.ErrorResponse
//	@Router			/api/user/friends [get]
func (uh *UserHandler) Friends(ctx *gin.Context) {
	link := ctx.Query("link")
	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	limit, offset, err := utils.GetLimitOffset(ctx)
	if err != nil {
		_ = ctx.Error(err)
	}

	users, err := uh.service.GetFriendsByLink(email, link, limit, offset)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	friends := make([]*dto.Profile, 0, len(users))

	for _, user := range users {
		friends = append(friends, dto.NewProfileFromUser(user))
	}

	_response := response.FriendsResponse{
		Body: response.FriendsBody{
			Friends: friends,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

// Subscribes godoc
//
//	@Summary		Subscribes
//	@Description	Get subscribes or subscribers for requested user
//	@Tags			Profiles
//	@Success		200		{object}	doc.ProfileArray
//	@Param			type	query		string	true	"in/out for subscribers/subscribes"
//	@Param			link	query		string	true	"link to requested profile"
//	@Param			limit	query		int		true	"amount of profiles"
//	@Param			offset	query		int		true	"number of batch"
//	@Failure		400		{object}	middleware.ErrorResponse
//	@Failure		401		{object}	middleware.ErrorResponse
//	@Failure		404		{object}	middleware.ErrorResponse
//	@Failure		500		{object}	middleware.ErrorResponse
//	@Router			/api/user/sub [get]
func (uh *UserHandler) Subscribes(ctx *gin.Context) {
	subType := ctx.Query("type")
	link := ctx.Query("link")

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	limit, offset, err := utils.GetLimitOffset(ctx)
	if err != nil {
		_ = ctx.Error(err)
	}

	var users []*entities.User

	switch subType {
	case "in":
		users, err = uh.service.GetSubscribersByLink(email, link, limit, offset)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
	case "out":
		users, err = uh.service.GetSubscribesByLink(email, link, limit, offset)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
	default:
		_ = ctx.Error(apperror.NewBadRequest())
		return
	}

	subs := make([]*dto.Profile, 0, len(users))

	for _, user := range users {
		subs = append(subs, dto.NewProfileFromUser(user))
	}

	_response := response.SubscribesResponse{
		Body: response.SubscribesBody{
			Subs: subs,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

func (uh *UserHandler) RandomUsers(ctx *gin.Context) {
	limit, offset, err := utils.GetLimitOffset(ctx)
	if err != nil {
		_ = ctx.Error(err)
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	users, err := uh.service.GetAllUsers(email, limit, offset)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	profiles := make([]*dto.Profile, 0, len(users))

	for _, user := range users {
		profiles = append(profiles, dto.NewProfileFromUser(user))
	}

	_response := response.RandomUsersResponse{
		Body: response.RandomUsersBody{
			Profiles: profiles,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

func (uh *UserHandler) PendingRequests(ctx *gin.Context) {
	limit, offset, err := utils.GetLimitOffset(ctx)
	if err != nil {
		_ = ctx.Error(err)
	}
	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	users, err := uh.service.GetPendingRequestsByEmail(email, limit, offset)
	if err != nil {
		return
	}

	profiles := make([]*dto.Profile, 0, len(users))

	for _, user := range users {
		profiles = append(profiles, dto.NewProfileFromUser(user))
	}

	_response := response.RandomUsersResponse{
		Body: response.RandomUsersBody{
			Profiles: profiles,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

func (uh *UserHandler) GetGlobalSearchResult(ctx *gin.Context) {
	dto := &dto.GlobalSearchDTO{}

	err := ctx.ShouldBind(dto)
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New("failed to parse search dto from query")))
		return
	}
	fmt.Println(*dto.SearchQuery)
	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	userSearchResult, communitiesSearchResult, err := uh.service.GlobalSearch(email, dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	_response := response.GetGlobalSearchResultResponse{
		Body: response.GetGlobalSearchResultBody{
			Users:       userSearchResult,
			Communities: communitiesSearchResult,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

func (uh *UserHandler) UserStatus(ctx *gin.Context) {
	link := ctx.Query("link")
	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
	}

	status, err := uh.service.UserStatus(email, link)
	if err != nil {
		_ = ctx.Error(err)
	}

	_response := response.UserStatusResponse{
		Body: response.UserStatusBody{
			Status: dto.StatusToString[status],
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

//nolint:unused
func (uh *UserHandler) getSession(ctx *gin.Context) (*authEntities.Session, error) {
	token, err := ctx.Cookie("session_id")
	if err != nil {

		return nil, apperror.NewServerError(apperror.NoAuth, nil)
	}
	stored, err := uh.authService.CheckSession(token)
	if err != nil {
		return nil, err
	}
	return stored, nil
}
