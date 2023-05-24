package response

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type SignInRequest struct {
	Body *dto.SignIn `json:"body"`
}

type SignUpRequest struct {
	Body *dto.SignUp `json:"body"`
}

type SubscribeRequest struct {
	Body *dto.Subscribes `json:"body"`
}

type UnsubscribeRequest struct {
	Body *dto.Subscribes `json:"body"`
}

type RejectRequest struct {
	Body *dto.Subscribes `json:"body"`
}

type EditProfileRequest struct {
	Body *dto.EditProfile `json:"body"`
}

type ProfileResponse struct {
	Body ProfileBody `json:"body"`
}

type ProfileBody struct {
	Profile *dto.Profile `json:"profile"`
}

type SelfResponse struct {
	Body SelfBody `json:"body"`
}

type SelfBody struct {
	Profile *entities.User `json:"profile"`
}

type FriendsResponse struct {
	Body FriendsBody `json:"body"`
}

type FriendsBody struct {
	Friends []*dto.Profile `json:"friends"`
}

type SubscribesResponse struct {
	Body SubscribesBody `json:"body"`
}

type SubscribesBody struct {
	Subs []*dto.Profile `json:"subs"`
}

type RandomUsersResponse struct {
	Body RandomUsersBody `json:"body"`
}

type RandomUsersBody struct {
	Profiles []*dto.Profile `json:"profiles"`
}

type PendingRequestsResponse struct {
	Body PendingRequestsBody `json:"body"`
}

type PendingRequestsBody struct {
	Profiles []*dto.Profile `json:"profiles"`
}

type GetGlobalSearchResultResponse struct {
	Body GetGlobalSearchResultBody `json:"body"`
}

type GetGlobalSearchResultBody struct {
	Users       []*entities.UserInfo      `json:"users"`
	Communities []*entities.CommunityInfo `json:"communitites"`
}

type UserStatusResponse struct {
	Body UserStatusBody `json:"body"`
}

type UserStatusBody struct {
	Status string `json:"status"`
}
