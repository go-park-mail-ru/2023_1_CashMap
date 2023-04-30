package entities

// User entity info
//
//	@Description	User account information
type User struct {
	ID         uint   `json:"-"           db:"id"`
	Email      string `json:"email"       db:"email"`
	Link       string `json:"user_link"   db:"link"`
	Password   string `json:"-"           db:"password"`
	FirstName  string `json:"first_name"  db:"first_name"`
	LastName   string `json:"last_name"   db:"last_name"`
	Avatar     string `json:"avatar"      db:"avatar"`
	Sex        string `json:"sex"         db:"sex"`
	Status     string `json:"status"      db:"status"`
	Bio        string `json:"bio"         db:"bio"`
	BirthDay   string `json:"birthday"    db:"birthday"`
	DateJoined string `json:"date_joined" db:"date_joined"`
	LastActive string `json:"last_active" db:"last_active"`
	Private    bool   `json:"private"     db:"private"`
}

type UserInfo struct {
	FirstName *string `db:"first_name" json:"first_name"`
	LastName  *string `db:"last_name" json:"last_name"`
	AvatarUrl *string `db:"url" json:"avatar_url"`
	Link      *string `db:"link" json:"link"`
}
