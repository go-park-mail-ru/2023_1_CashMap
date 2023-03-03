package entities

type Group struct {
	ID           uint
	Title        string
	HeaderImage  string
	MembersCount int
	Owners       []User
	Posts        []Post
	// TODO: доделать поля модели для будущих потребностей
}
