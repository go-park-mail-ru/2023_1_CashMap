package dto

type Group struct {
	Title     string `json:"title"`
	Info      string `json:"group_info"`
	Avatar    string `json:"avatar"`
	Privacy   string `json:"privacy"`
	HideOwner bool   `json:"hide_owner"`
}

type UpdateGroup struct {
	Title     *string `json:"title"`
	Info      *string `json:"group_info"`
	Avatar    *string `json:"avatar"`
	Privacy   *string `json:"privacy"`
	HideOwner *bool   `json:"hide_owner"`
}

type AddManager struct {
}
