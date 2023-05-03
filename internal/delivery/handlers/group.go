package handlers

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/delivery/utils"
	"depeche/internal/usecase"
	"github.com/gin-gonic/gin"
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
	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"group_info": group,
			"is_sub":     isSub,
			"is_admin":   isAdmin,
		},
	})
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
	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"groups": groups,
		},
	})
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

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"groups": groups,
		},
	})
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

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"groups": groups,
		},
	})
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

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"groups": groups,
		},
	})
}

func (gh *GroupHandler) CreateGroup(ctx *gin.Context) {
	body, err := utils.GetBody[dto.Group](ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	err = gh.service.CreateGroup(email, body)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
}

func (gh *GroupHandler) UpdateGroup(ctx *gin.Context) {
	link := ctx.Param("link")

	body, err := utils.GetBody[dto.UpdateGroup](ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	err = gh.service.UpdateGroup(link, email, body)
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
	err = gh.service.DeleteGroup(link, email)
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

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"profiles": profiles,
		},
	})
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
	body, err := utils.GetBody[dto.Requests](ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	err = gh.service.AcceptRequest(email, link, body.AcceptEmail)
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
	body, err := utils.GetBody[dto.Requests](ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = gh.service.DeclineRequest(email, body.AcceptEmail, link)
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

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"profiles": profiles,
		},
	})
}
