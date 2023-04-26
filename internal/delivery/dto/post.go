package dto

import "io"

type PostGetByID struct {
	PostID uint `form:"post_id" json:"post_id"`
}

type PostsGetByLink struct {
	CommunityLink *string `form:"community_link" json:"community_link"`
	OwnerLink     *string `form:"owner_link" json:"owner_link"`
	BatchSize     uint    `form:"batch_size" json:"batch_size"`
	LastPostDate  string  `form:"last_post_date" json:"last_post_date"`
}

type PostCreate struct {
	CommunityLink    *string         `form:"community_link" json:"community_link"`
	OwnerLink        *string         `form:"owner_link" json:"owner_link"`
	UserLink         string          `form:"author_link" json:"author_link"`
	ShouldShowAuthor bool            `form:"show_author" json:"show_author"`
	Text             string          `form:"text" json:"text"`
	Attachments      []io.ReadCloser `form:"attachments" json:"attachments"`
}

type PostDelete struct {
	PostID *uint `json:"post_id" valid:"-"`
}

type PostUpdate struct {
	PostID              *uint            `form:"post_id" valid:"-"`
	ShouldShowAuthor    *bool            `form:"show_author" valid:"-"`
	Text                *string          `form:"text" valid:"-"`
	Attachments         *[]io.ReadCloser `form:"attachments" valid:"-"`
	AttachmentsToRemove *[]string        `form:"attachments_to_remove" valid:"-"`
	ChangeDate          string           `form:"-" json:"-"`
}

type LikeDTO struct {
	PostID *uint `json:"post_id" valid:"-"`
}
