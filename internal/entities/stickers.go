package entities

type Sticker struct {
	ID            uint   `json:"id"             db:"id"`
	Url           string `json:"url"            db:"url"`
	StickerpackID uint   `json:"stickerpack_id" db:"stickerpack_id"`
}

type StickerPack struct {
	ID              uint      `json:"id"               db:"id"`
	Title           string    `json:"title"            db:"title"`
	Author          string    `json:"author"           db:"author"`
	DepecheAuthored bool      `json:"depeche_authored" db:"depeche_authored"`
	Cover           string    `json:"cover"            db:"cover"`
	CreationDate    string    `json:"creation_date"    db:"creation_date"`
	Description     string    `json:"description"      db:"description"`
	Stickers        []Sticker `json:"stickers"`
}
