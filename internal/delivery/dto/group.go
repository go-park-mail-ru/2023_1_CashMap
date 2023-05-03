package dto

type Group struct {
	Title     string `json:"title"`
	Info      string `json:"group_info"`
	Avatar    string `json:"avatar_url"`
	Privacy   string `json:"privacy"`
	HideOwner bool   `json:"hide_owner"`
}

type UpdateGroup struct {
	Title     *string `json:"title"`
	Info      *string `structs:"group_info" json:"group_info"`
	Link      *string `json:"link"`
	Avatar    *string `json:"avatar_url"`
	Privacy   *string `json:"privacy"`
	HideOwner *bool   `json:"hide_owner"`
}
type Requests struct {
	AcceptEmail string `json:"link"`
}

type AddManager struct {
}
