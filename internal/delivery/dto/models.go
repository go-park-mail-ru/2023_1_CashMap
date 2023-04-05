package dto

// [INCOMING]

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUp struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type EditProfile struct {
	Email     *string `json:"email"`
	Password  *string `json:"password"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Link      *string `json:"link"`
	Sex       *string `json:"sex"`
	Status    *string `json:"status"`
	Bio       *string `json:"bio"`
	Birthday  *string `json:"birthday"`
}

type Subscribes struct {
	Link string `json:"link"`
}

// [OUTGOING]

type Profile struct {
	Link      string `json:"link"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
	Status    string `json:"status"`
	Bio       string `json:"bio"`
	BirthDate string `json:"birth_date"`
	Private   bool   `json:"private"`
	Sex       string `json:"sex"`
}

func (si *SignIn) AuthEmail() string {
	return si.Email
}

func (su *SignUp) AuthEmail() string {
	return su.Email
}
