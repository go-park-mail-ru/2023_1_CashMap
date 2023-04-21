package entities

type Group struct {
	ID           uint              `json:"-"                    db:"id"`
	Link         string            `json:"link"                 db:"link"`
	Title        string            `json:"title"                db:"title"`
	Avatar       string            `json:"avatar,omitempty"     db:"avatar"`
	MembersCount int               `json:"members_count"        db:"members_count"`
	CreationDate string            `json:"creation_date"        db:"creation_date"`
	OwnerLink    string            `json:"owner_link,omitempty" db:"owner_link"`
	HideOwner    bool              `json:"hide_author"          db:"hide_owner"`
	Management   []GroupManagement `json:"management"`
}

type GroupManagement struct {
	Link string `json:"link"`
	Role string `json:"role"`
}
