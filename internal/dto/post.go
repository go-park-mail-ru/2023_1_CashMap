package dto

type PostDTO struct {
	SenderLink    string `json:"sender_link"`
	PostID        uint   `json:"post_id"`
	CommunityLink string `json:"community_link"`
	UserLink      string `json:"user_link"`
}
