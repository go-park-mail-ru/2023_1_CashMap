package doc

import "depeche/internal/delivery/dto"

type Profile struct {
	Body dto.Profile `json:"body"`
}

type SignIn struct {
	Body dto.SignIn `json:"body"`
}

type SignUp struct {
	Body dto.SignUp `json:"body"`
}

type EditProfile struct {
	Body dto.EditProfile `json:"body"`
}

type Subscribes struct {
	Body dto.Subscribes `json:"body"`
}

type ProfileArray struct {
	Body []dto.Profile `json:"body"`
}
