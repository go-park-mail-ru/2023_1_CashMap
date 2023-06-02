package entities

//go:generate easyjson --all group.go

type Group struct {
	ID           uint              `json:"-"                    db:"id"`
	Link         string            `json:"link"                 db:"link"`
	Title        string            `json:"title"                db:"title"`
	Info         string            `json:"info"                 db:"info"`
	Privacy      string            `json:"privacy"              db:"privacy"`
	Avatar       string            `json:"avatar_url"           db:"avatar"`
	MembersCount int               `json:"subscribers"          db:"subscribers"`
	CreationDate string            `json:"creation_date"        db:"creation_date"`
	OwnerLink    string            `json:"-"                    db:"owner_link"`
	HideOwner    bool              `json:"hide_owner"           db:"hide_author"`
	IsDeleted    bool              `json:"is_deleted"           db:"is_deleted"`
	Management   []GroupManagement `json:"management"`
}

type GroupManagement struct {
	Link      string `json:"link" db:"link"`
	Role      string `json:"role"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Avatar    string `json:"avatar" db:"avatar"`
}
