package entities

// Post entity info
//
//	@Description	All post information
type Post struct {
	ID               uint     `db:"id" json:"id"`
	AuthorLink       string   `db:"author_link" json:"author_link"`
	CommunityLink    *string  `db:"community_link" json:"community_link"`
	OwnerLink        *string  `db:"owner_link" json:"owner_link"`
	ShouldShowAuthor bool     `db:"show_author" json:"show_author"`
	Text             string   `db:"text_content" json:"text_content"`
	Likes            int      `db:"likes_amount" json:"likes_amount"`
	CreationDate     string   `db:"creation_date" json:"creation_date"`
	ChangeDate       string   `db:"change_date" json:"change_date"`
	IsDeleted        bool     `db:"is_deleted" json:"is_deleted"`
	Attachments      []string `db:"attachments" json:"attachments"`
}
