package localStorage

import (
	"depeche/internal/entities"
	"errors"
	"sort"
	"time"
)

type FeedStorage struct {
	groups map[string]*entities.Group
	users  map[string]*entities.User
}

func GetFriendsPosts(user *entities.User, filterDateTime time.Time, postsNumber int) ([]entities.Post, error) {
	friendsPosts := make([]entities.Post, 0, postsNumber)
	for _, friend := range user.Friends {
		for _, post := range friend.Posts {
			if post.Date.Before(filterDateTime) {
				friendsPosts = append(friendsPosts, post)
			}
		}
	}

	sort.Slice(friendsPosts, func(i int, j int) bool {
		return friendsPosts[i].Date.Before(friendsPosts[j].Date)
	})

	if len(friendsPosts) >= postsNumber {
		return friendsPosts[:postsNumber], nil
	}

	return friendsPosts[:len(friendsPosts)], nil
}

func GetGroupsPosts(user *entities.User, filterDateTime time.Time, postsNumber int) ([]entities.Post, error) {
	groupsPosts := make([]entities.Post, 0, postsNumber)
	for _, group := range user.Groups {
		for _, post := range group.Posts {
			if post.Date.Before(filterDateTime) {
				groupsPosts = append(groupsPosts, post)
			}
		}
	}

	sort.Slice(groupsPosts, func(i int, j int) bool {
		return groupsPosts[i].Date.Before(groupsPosts[j].Date)
	})

	if len(groupsPosts) >= postsNumber {
		return groupsPosts[:postsNumber], nil
	}

	return groupsPosts[:len(groupsPosts)], nil
}

type UserStorage struct {
	user map[string]*entities.User
}

func NewUserStorage() *UserStorage {
	return &UserStorage{
		user: mockUsers,
	}
}

func (lc *UserStorage) GetUserById(id uint) (*entities.User, error) {
	return nil, nil
}

func (lc *UserStorage) GetUserByEmail(email string) (*entities.User, error) {
	user := lc.user[email]
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (lc *UserStorage) GetUserFriends(user *entities.User) ([]*entities.User, error) {
	return nil, nil
}

func (lc *UserStorage) CreateUser(user *entities.User) (*entities.User, error) {
	if lc.user[user.Email] != nil {
		return nil, errors.New("user already exists")
	}
	lc.user[user.Email] = user
	return user, nil
}

var mockUsers = map[string]*entities.User{
	"user1@mail.ru": {
		Email:     "user1@mail.ru",
		Password:  "some_hash",
		FirstName: "Vladimir",
		LastName:  "Mayakovsky",
	},
	"user2@mail.ru": {
		Email:     "user1@mail.ru",
		Password:  "some_hash",
		FirstName: "Sergei",
		LastName:  "Esenin",
	},
	"user3@mail.ru": {
		Email:     "user1@mail.ru",
		Password:  "some_hash",
		FirstName: "Fedor",
		LastName:  "Tutchev",
	},
	"user4@mail.ru": {
		Email:     "user1@mail.ru",
		Password:  "some_hash",
		FirstName: "Michail",
		LastName:  "Lermontov",
	},
	"user5@mail.ru": {
		Email:     "user1@mail.ru",
		Password:  "some_hash",
		FirstName: "Alexandr",
		LastName:  "Pushkin",
	},
}

var mockGroups = map[string]*entities.Group{}

var mockPosts = []entities.Post{}
