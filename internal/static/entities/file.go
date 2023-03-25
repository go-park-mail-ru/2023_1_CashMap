package entities

type FileType int64

const (
	IMAGE    FileType = 1
	DOCUMENT FileType = 2
)

type UserFile struct {
	Path string
	Type FileType
}
