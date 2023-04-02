package entities

// Post entity info
//
//	@Description	All post information
type Post struct {
	ID               uint   `db:"id"`
	AuthorID         uint   `db:"author_id"`
	communityID      uint   `db:"community_id"`
	ShouldShowAuthor bool   `db:"show_author"`
	Text             string `db:"text_content"`
	Likes            int    `db:"likes_amount"`
	CreationDate     string `db:"creation_date"`
	ChangeDate       string `db:"change_date"`
	IsDeleted        bool   `db:"is_deleted"`
}
