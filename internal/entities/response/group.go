package response

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type GetGroupResponse struct {
	Body GetGroupBody `json:"body"`
}

type GetGroupBody struct {
	GroupInfo *entities.Group `json:"group_info"`
	IsSub     bool            `json:"is_sub"`
	IsAdmin   bool            `json:"is_admin"`
}

type GetGroupsResponse struct {
	Body GetGroupsBody `json:"body"`
}

type GetGroupsBody struct {
	Groups []*entities.Group `json:"groups"`
}

type GetUserGroupsResponse struct {
	Body GetUserGroupsBody `json:"body"`
}

type GetUserGroupsBody struct {
	Groups []*entities.Group `json:"groups"`
}

type GetPopularGroupsResponse struct {
	Body GetPopularGroupsBody `json:"body"`
}

type GetPopularGroupsBody struct {
	Groups []*entities.Group `json:"groups"`
}

type GetManagedGroupsResponse struct {
	Body GetManagedGroupsBody `json:"body"`
}

type GetManagedGroupsBody struct {
	Groups []*entities.Group `json:"groups"`
}

type CreateGroupRequest struct {
	Body *dto.Group `json:"body"`
}

type UpdateGroupRequest struct {
	Body *dto.UpdateGroup `json:"body"`
}

type GetSubscribersResponse struct {
	Body GetSubscribersBody `json:"body"`
}

type GetSubscribersBody struct {
	Profiles []*dto.Profile `json:"profiles"`
}

type AcceptRequestRequest struct {
	Body *dto.Requests `json:"body"`
}

type PendingGroupRequestsResponse struct {
	Body PendingGroupRequestsBody `json:"body"`
}

type PendingGroupRequestsBody struct {
	Profiles []*dto.Profile `json:"profiles"`
}
