package dto

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
	CommunityLink    *string  `form:"community_link" json:"community_link"`
	OwnerLink        *string  `form:"owner_link" json:"owner_link"`
	UserLink         string   `form:"author_link" json:"author_link"`
	ShouldShowAuthor bool     `form:"show_author" json:"show_author"`
	Text             string   `form:"text" json:"text"`
	Attachments      []string `form:"attachments" json:"attachments"`
}

type PostDelete struct {
	PostID *uint `json:"post_id" valid:"-"`
}

type PostUpdate struct {
	PostID           *uint              `json:"post_id" valid:"-"`
	ShouldShowAuthor *bool              `json:"show_author" valid:"-"`
	Text             *string            `json:"text" valid:"-"`
	Attachments      *UpdateAttachments `json:"attachments"`
	ChangeDate       string             `json:"-" json:"-"`
}

type UpdateAttachments struct {
	Deleted []string `json:"deleted"`
	Added   []string `json:"added"`
}

type LikeDTO struct {
	PostID *uint `json:"post_id" valid:"-"`
}
