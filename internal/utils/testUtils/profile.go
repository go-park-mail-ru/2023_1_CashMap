package testUtils

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

func InitProfilePasswordFail(new string) *dto.EditProfile {
	return &dto.EditProfile{
		NewPassword: &new,
	}
}

func InitProfilePasswordWithPrev(old, new string) *dto.EditProfile {
	return &dto.EditProfile{
		PreviousPassword: &old,
		NewPassword:      &new,
	}
}

func InitProfileLink(link string) *dto.EditProfile {
	return &dto.EditProfile{
		Link: &link,
	}
}

func InitProfileAvatar(avatar string) *dto.EditProfile {
	return &dto.EditProfile{
		Avatar: &avatar,
	}
}

func InitProfileBasic(fName, lName, bio string) *dto.EditProfile {
	return &dto.EditProfile{
		FirstName: &fName,
		LastName:  &lName,
		Bio:       &bio,
	}
}

func InitUserProfile(firstName, lastName string) *entities.UserInfo {
	return &entities.UserInfo{
		FirstName: &firstName,
		LastName:  &lastName,
	}
}

func InitPostCreateDtoForCommunity(communityLink string) *dto.PostCreate {
	return &dto.PostCreate{
		CommunityLink: &communityLink,
		OwnerLink:     nil,
	}
}

func InitPostCreateDtoForUser(ownerLink string) *dto.PostCreate {
	return &dto.PostCreate{
		CommunityLink: nil,
		OwnerLink:     &ownerLink,
	}
}
