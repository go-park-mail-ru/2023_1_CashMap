package entities

type FileType string

const (
	IMAGE    FileType = "img"
	DOCUMENT FileType = "doc"
	STICKER  FileType = "sticker"
)

type UserFile struct {
	Name string   `form:"name" json:"name"`
	Type FileType `form:"type" json:"type"`
}
