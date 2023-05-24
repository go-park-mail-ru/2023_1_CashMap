package handlers

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/delivery/utils"
	"depeche/internal/entities/response"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson"
	"net/http"
)

type GroupHandler struct {
	service usecase.Group
}

func NewGroupHandler(service usecase.Group) *GroupHandler {
	return &GroupHandler{
		service: service,
	}
}

func (gh *GroupHandler) GetGroup(ctx *gin.Context) {
	link := ctx.Param("link")

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
	}
	group, err := gh.service.GetGroup(link)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	isSub, err := gh.service.CheckSub(email, link)
	if err != nil {
		_ = ctx.Error(err)
	}
	isAdmin, err := gh.service.CheckAdmin(email, link)
	if err != nil {
		_ = ctx.Error(err)
	}

	_response := response.GetGroupResponse{
		Body: response.GetGroupBody{
			GroupInfo: group,
			IsAdmin:   isAdmin,
			IsSub:     isSub,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

func (gh *GroupHandler) GetGroups(ctx *gin.Context) {
	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	limit, offset, err := utils.GetLimitOffset(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	groups, err := gh.service.GetUserGroupsByEmail(email, limit, offset)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	_response := response.GetGroupsResponse{
		Body: response.GetGroupsBody{
			Groups: groups,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

func (gh *GroupHandler) GetUserGroups(ctx *gin.Context) {
	link := ctx.Query("link")

	limit, offset, err := utils.GetLimitOffset(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	groups, err := gh.service.GetUserGroupsByLink(link, limit, offset)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	_response := response.GetUserGroupsResponse{
		Body: response.GetUserGroupsBody{
			Groups: groups,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

func (gh *GroupHandler) GetPopularGroups(ctx *gin.Context) {
	limit, offset, err := utils.GetLimitOffset(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	groups, err := gh.service.GetPopularGroups(email, limit, offset)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	_response := response.GetPopularGroupsResponse{
		Body: response.GetPopularGroupsBody{
			Groups: groups,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

func (gh *GroupHandler) GetManagedGroups(ctx *gin.Context) {
	limit, offset, err := utils.GetLimitOffset(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	groups, err := gh.service.GetManagedGroups(email, limit, offset)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	_response := response.GetManagedGroupsResponse{
		Body: response.GetManagedGroupsBody{
			Groups: groups,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

func (gh *GroupHandler) CreateGroup(ctx *gin.Context) {
	inputDTO := new(response.CreateGroupRequest)
	if err := easyjson.UnmarshalFromReader(ctx.Request.Body, inputDTO); err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New("failed to parse struct")))
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	err = gh.service.CreateGroup(email, inputDTO.Body)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
}

func (gh *GroupHandler) UpdateGroup(ctx *gin.Context) {
	link := ctx.Param("link")

	inputDTO := new(response.UpdateGroupRequest)
	if err := easyjson.UnmarshalFromReader(ctx.Request.Body, inputDTO); err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New("failed to parse struct")))
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	err = gh.service.UpdateGroup(link, email, inputDTO.Body)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
}

func (gh *GroupHandler) DeleteGroup(ctx *gin.Context) {
	link := ctx.Param("link")

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = gh.service.DeleteGroup(email, link)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
}

func (gh *GroupHandler) GetSubscribers(ctx *gin.Context) {
	link := ctx.Param("link")

	limit, offset, err := utils.GetLimitOffset(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	users, err := gh.service.GetSubscribers(link, limit, offset)
	if err != nil {
		return
	}

	profiles := make([]*dto.Profile, 0, len(users))
	for _, user := range users {
		profiles = append(profiles, dto.NewProfileFromUser(user))
	}

	_response := response.GetSubscribersResponse{
		Body: response.GetSubscribersBody{
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

func (gh *GroupHandler) SubscribeGroup(ctx *gin.Context) {
	link := ctx.Param("link")
	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = gh.service.Subscribe(email, link)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
}

func (gh *GroupHandler) UnsubscribeGroup(ctx *gin.Context) {
	link := ctx.Param("link")
	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = gh.service.Unsubscribe(email, link)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
}

func (gh *GroupHandler) AcceptRequest(ctx *gin.Context) {
	link := ctx.Param("link")

	inputDTO := new(response.AcceptRequestRequest)
	if err := easyjson.UnmarshalFromReader(ctx.Request.Body, inputDTO); err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New("failed to parse struct")))
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	err = gh.service.AcceptRequest(email, link, inputDTO.Body.AcceptEmail)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
}

func (gh *GroupHandler) AcceptAllRequests(ctx *gin.Context) {
	link := ctx.Param("link")
	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = gh.service.AcceptAllRequests(email, link)
	if err != nil {
		_ = ctx.Error(err)
	}
}

func (gh *GroupHandler) DeclineRequest(ctx *gin.Context) {
	link := ctx.Param("link")
	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	inputDTO := new(response.AcceptRequestRequest)
	if err := easyjson.UnmarshalFromReader(ctx.Request.Body, inputDTO); err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New("failed to parse struct")))
		return
	}

	err = gh.service.DeclineRequest(email, inputDTO.Body.AcceptEmail, link)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
}

func (gh *GroupHandler) PendingGroupRequests(ctx *gin.Context) {
	link := ctx.Param("link")
	limit, offset, err := utils.GetLimitOffset(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	users, err := gh.service.GetPendingRequests(email, link, limit, offset)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	profiles := make([]*dto.Profile, 0, len(users))
	for _, user := range users {
		profiles = append(profiles, dto.NewProfileFromUser(user))
	}

	_response := response.PendingGroupRequestsResponse{
		Body: response.PendingGroupRequestsBody{
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
