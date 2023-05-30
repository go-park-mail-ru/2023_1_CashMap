package entities

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
	Link string `json:"link"`
	Role string `json:"role"`
}
