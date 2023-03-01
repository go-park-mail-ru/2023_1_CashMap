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

func NewFeedStorage() *FeedStorage {
	return &FeedStorage{
		groups: mockGroups,
		users:  mockUsers,
	}
}

func (storage *FeedStorage) GetFriendsPosts(user *entities.User, filterDateTime time.Time, postsNumber int) ([]entities.Post, error) {
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

func (storage *FeedStorage) GetGroupsPosts(user *entities.User, filterDateTime time.Time, postsNumber int) ([]entities.Post, error) {
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

var mockPosts = []entities.Post{
	{SenderName: "VK Образование",
		Text: "Сегодня VK Образование совместно с сервисом «Другое Дело» запустили квест «VK: Крафтим карьеру в IT»" +
			" из восьми мини-игр в формате метавселенной Minecraft.\n\nИгра поможет студентам и школьникам познакомиться" +
			" с представителями IT-профессий в формате популярной игры Minecraft. Ведущий и герои в интерактивном формате" +
			" расскажут ребятам о представителях разных IT-профессий: фронтенд-разработчике, дизайнере интерфейсов," +
			" бэкенд-разработчике, игровом разработчике, QA-специалисте, ML-инженере, UX-исследователе, таргетологе." +
			" А ещё проведут по миру карьерных и образовательных проектов VK, в которых в будущем игроки сами смогут" +
			" поучаствовать.\n\nБаллы, полученные за успешное выполнение всех заданий в игре, можно обменять на ценные" +
			" призы в «Другом Деле»: мерч, подписки, технику, билеты на концерты или поездки по стране.\n\nПройти квест" +
			" можно по ссылке: vk.cc/clIuYy.",
		Date:  time.Date(2023, time.February, 27, 17, 16, 3, 0, time.Local),
		Likes: 500,
		Comments: []entities.Comment{
			{
				ID:      1,
				Sender:  "Alexandr Pushkin",
				Date:    time.Date(2023, time.February, 28, 12, 6, 3, 0, time.Local),
				Text:    "И это прекрасно",
				ReplyTo: -1,
			},

			{
				ID:     1,
				Sender: "Michail Lermontov",
				Date:   time.Date(2023, time.February, 28, 12, 6, 3, 0, time.Local),
				Text: "Уважения заслуживают те люди, которые независимо от ситуации," +
					" времени и места, остаются такими же, какие они есть на самом деле.",
				ReplyTo: -1,
			},
		},
	},

	{SenderName: "VK Образование",
		Text: "Учи.ру теперь с нами на 100% 🤓\nУже два года VK активно взаимодействует с крупнейшей образовательной" +
			" платформой для школьников: запускает олимпиады, проводит конференции для педагогов и родителей." +
			" Сегодня Учи.ру присоединилась к семье наших проектов.\n Раньше VK принадлежало 25% платформы и её сервисов," +
			" а теперь мы вместе на все 100%. Эта сделка усилит позиции компании в сегменте дополнительного школьного образования 🎓",
		Date:  time.Date(2023, time.February, 20, 14, 44, 1, 0, time.Local),
		Likes: 404,
		Comments: []entities.Comment{
			{
				ID:      1,
				Sender:  "Fedor Tutchev",
				Date:    time.Date(2023, time.February, 20, 16, 12, 10, 0, time.Local),
				Text:    "Умом Россию не понять, в Россию можно только верить...",
				ReplyTo: -1,
			},

			{
				ID:      2,
				Sender:  "Michail Lermontov",
				Date:    time.Date(2023, time.February, 21, 1, 2, 3, 0, time.Local),
				Text:    "Полностью согласен с вами, Федор Иванович!",
				ReplyTo: 1,
			},

			{
				ID:      3,
				Sender:  "Vladimir Mayakovsky",
				Date:    time.Date(2023, time.February, 24, 22, 24, 4, 0, time.Local),
				Text:    "Позвольте не согласиться",
				ReplyTo: 1,
			},
		},
	},

	{SenderName: "VK Образование",
		Text: "Учи.ру теперь с нами на 100% 🤓\nУже два года VK активно взаимодействует с крупнейшей образовательной" +
			" платформой для школьников: запускает олимпиады, проводит конференции для педагогов и родителей." +
			" Сегодня Учи.ру присоединилась к семье наших проектов.\n Раньше VK принадлежало 25% платформы и её сервисов," +
			" а теперь мы вместе на все 100%. Эта сделка усилит позиции компании в сегменте дополнительного школьного образования 🎓",
		Date:  time.Date(2023, time.February, 20, 14, 44, 1, 0, time.Local),
		Likes: 404,
		Comments: []entities.Comment{
			{
				ID:      1,
				Sender:  "Fedor Tutchev",
				Date:    time.Date(2023, time.February, 20, 16, 12, 10, 0, time.Local),
				Text:    "Умом Россию не понять, в Россию можно только верить...",
				ReplyTo: -1,
			},

			{
				ID:      2,
				Sender:  "Michail Lermontov",
				Date:    time.Date(2023, time.February, 21, 1, 2, 3, 0, time.Local),
				Text:    "Полностью согласен с вами, Федор Иванович!",
				ReplyTo: 1,
			},

			{
				ID:      3,
				Sender:  "Vladimir Mayakovsky",
				Date:    time.Date(2023, time.February, 24, 22, 24, 4, 0, time.Local),
				Text:    "Позвольте не согласиться",
				ReplyTo: 1,
			},
		},
	},

	{SenderName: "VK Образование",
		Text: "Учи.ру теперь с нами на 100% 🤓\nУже два года VK активно взаимодействует с крупнейшей образовательной" +
			" платформой для школьников: запускает олимпиады, проводит конференции для педагогов и родителей." +
			" Сегодня Учи.ру присоединилась к семье наших проектов.\n Раньше VK принадлежало 25% платформы и её сервисов," +
			" а теперь мы вместе на все 100%. Эта сделка усилит позиции компании в сегменте дополнительного школьного образования 🎓",
		Date:  time.Date(2023, time.February, 20, 14, 44, 1, 0, time.Local),
		Likes: 404,
		Comments: []entities.Comment{
			{
				ID:      1,
				Sender:  "Fedor Tutchev",
				Date:    time.Date(2023, time.February, 20, 16, 12, 10, 0, time.Local),
				Text:    "Умом Россию не понять, в Россию можно только верить...",
				ReplyTo: -1,
			},

			{
				ID:      2,
				Sender:  "Michail Lermontov",
				Date:    time.Date(2023, time.February, 21, 1, 2, 3, 0, time.Local),
				Text:    "Полностью согласен с вами, Федор Иванович!",
				ReplyTo: 1,
			},

			{
				ID:      3,
				Sender:  "Vladimir Mayakovsky",
				Date:    time.Date(2023, time.February, 24, 22, 24, 4, 0, time.Local),
				Text:    "Позвольте не согласиться",
				ReplyTo: 1,
			},
		},
	},
}
