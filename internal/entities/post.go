package entities

// Post entity info
//
//	@Description	All post information
type Post struct {
	ID               uint           `db:"id" json:"id"`
	AuthorLink       string         `db:"author_link" json:"author_link"`
	CommunityLink    *string        `db:"community_link" json:"-"`
	OwnerLink        *string        `db:"owner_link" json:"-"`
	CommunityInfo    *CommunityInfo `db:"community_info" json:"community_info"`
	OwnerInfo        *OwnerInfo     `db:"owner_info" json:"owner_info"`
	ShouldShowAuthor bool           `db:"show_author" json:"show_author"`
	Text             string         `db:"text_content" json:"text_content"`
	Likes            int            `db:"likes_amount" json:"likes_amount"`
	CreationDate     string         `db:"creation_date" json:"creation_date"`
	ChangeDate       string         `db:"change_date" json:"change_date"`
	IsDeleted        bool           `db:"is_deleted" json:"is_deleted"`
	Attachments      []string       `db:"attachments" json:"attachments"`
}

type OwnerInfo struct {
	FirstName *string `db:"first_name" json:"first_name"`
	LastName  *string `db:"last_name" json:"last_name"`
	AvatarUrl *string `db:"url" json:"url"`
}

type CommunityInfo struct {
	Title     *string `db:"title" json:"title"`
	AvatarUrl *string `db:"url" json:"url"`
}
