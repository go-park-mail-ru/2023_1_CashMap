package entities

// User entity info
//
//	@Description	User account information
type User struct {
	ID         uint   `json:"id"          db:"id"`
	Email      string `json:"email"       db:"email"`
	Link       string `json:"link"        db:"link"`
	Password   string `json:"-"           db:"password"`
	FirstName  string `json:"first_name"  db:"first_name"`
	LastName   string `json:"last_name"   db:"last_name"`
	Avatar     string `json:"avatar"      db:"avatar"`
	Sex        string `json:"sex"         db:"sex"`
	Status     string `json:"status"      db:"status"`
	Bio        string `json:"bio"         db:"bio"`
	BirthDay   string `json:"birthday"    db:"birthday"`
	DateJoined string `json:"date_joined" db:"date_joined"`
}
