package response

import "depeche/static/entities"

type LoadFileResponse struct {
	Body   LoadFileBody `json:"body"`
	Status int          `json:"status"`
}

type LoadFileBody struct {
	Form []*entities.UserFile `json:"form"`
}

type DeleteFileRequest struct {
	Body *entities.UserFile `json:"body"`
}
