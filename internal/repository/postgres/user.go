package postgres

import (
	"database/sql"
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	utildb "depeche/internal/repository/utils"
	"depeche/internal/utils"
	"depeche/pkg/apperror"
	"fmt"
	"github.com/arbovm/levenshtein"
	"github.com/jackc/pgerrcode"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"strings"

	"github.com/fatih/structs"
)

type UserRepository struct {
	DB *sqlx.DB
}

// TODO specify errors (replace 500 with smth) and find a way to wrap an sql error
// TODO remove println

func NewPostgresUserRepo(db *sqlx.DB) repository.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) GetUser(query string, args ...interface{}) (*entities.User, error) {
	user := &entities.User{}
	row := ur.DB.QueryRowx(query, args)
	err := row.StructScan(user)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperror.NewServerError(apperror.UserNotFound, err)
		}
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}
	return user, nil
}

func (ur *UserRepository) GetUserById(id uint) (*entities.User, error) {
	user := &entities.User{}
	row := ur.DB.QueryRowx(UserById, id)
	err := row.StructScan(user)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return nil, apperror.NewServerError(apperror.UserNotFound, err)
		}
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}
	return user, nil
}

func (ur *UserRepository) GetUserByLink(link string) (*entities.User, error) {
	user := &entities.User{}
	row := ur.DB.QueryRowx(UserByLink, link)
	err := row.StructScan(user)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return nil, apperror.NewServerError(apperror.UserNotFound, err)
		}
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}
	return user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	user := &entities.User{}
	row := ur.DB.QueryRowx(UserByEmail, email)
	err := row.StructScan(user)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperror.NewServerError(apperror.UserNotFound, err)
		}
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}
	return user, nil
}

func (ur *UserRepository) GetFriends(user *entities.User, limit, offset int) ([]*entities.User, error) {
	var users []*entities.User
	rows, err := ur.DB.Queryx(FriendsById, user.ID, offset, limit)
	defer utildb.CloseRows(rows)
	if err != nil {
		if err == sql.ErrNoRows {
			return users, nil
		}
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}
	for rows.Next() {
		var user = &entities.User{}
		if err := rows.StructScan(user); err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) GetSubscribes(user *entities.User, limit, offset int) ([]*entities.User, error) {
	var users []*entities.User
	rows, err := ur.DB.Queryx(SubscribesById, user.ID, offset, limit)
	defer utildb.CloseRows(rows)
	if err != nil {
		if err == sql.ErrNoRows {
			return users, nil
		}
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}
	for rows.Next() {
		var user = &entities.User{}
		if err := rows.StructScan(user); err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) GetSubscribers(user *entities.User, limit, offset int) ([]*entities.User, error) {
	var users []*entities.User
	rows, err := ur.DB.Queryx(SubscribersById, user.ID, offset, limit)
	defer utildb.CloseRows(rows)
	if err != nil {
		if err == sql.ErrNoRows {
			return users, nil
		}
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}
	for rows.Next() {
		var user = &entities.User{}
		if err := rows.StructScan(user); err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) GetUsers(email string, limit, offset int) ([]*entities.User, error) {
	var users []*entities.User
	rows, err := ur.DB.Queryx(RandomUsers, email, offset, limit)
	defer utildb.CloseRows(rows)
	if err != nil {
		if err == sql.ErrNoRows {
			return users, nil
		}
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}
	for rows.Next() {
		var user = &entities.User{}
		if err := rows.StructScan(user); err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) UpdateAvatar(email string, url string) error {
	err := ur.DB.QueryRowx(UpdateAvatar, url, email).Scan()
	if err != nil {
		if err != sql.ErrNoRows {
			return apperror.NewServerError(apperror.InternalServerError, err)
		}
	}
	return nil
}

func (ur *UserRepository) CheckLinkExists(link string) (bool, error) {
	var exists bool
	err := ur.DB.QueryRowx(CheckLink, link).Scan(&exists)
	if err != nil {
		return false, apperror.NewServerError(apperror.InternalServerError, err)
	}

	return exists, nil
}

func (ur *UserRepository) Subscribe(subEmail, targetLink, requestTime string) (bool, error) {
	_, err := ur.DB.Exec(Subscribe, subEmail, targetLink, requestTime)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			switch err.Code {
			case pgerrcode.UniqueViolation:
				return false, apperror.NewServerError(apperror.RepeatedSubscribe, err)
			default:
				return false, apperror.NewServerError(apperror.InternalServerError, err)
			}
		}
		return false, apperror.NewServerError(apperror.InternalServerError, err)
	}
	return false, nil
}

func (ur *UserRepository) Unsubscribe(userEmail, targetLink string) (bool, error) {
	rows, err := ur.DB.Queryx(Unsubscribe, userEmail, targetLink)
	defer utildb.CloseRows(rows)
	if err != nil {
		fmt.Println(err)
		// TODO check subscribe conflict (repeated unsubscribe)
		return false, apperror.NewServerError(apperror.InternalServerError, err)
	}
	err = rows.Close()
	if err != nil {
		return false, apperror.NewServerError(apperror.InternalServerError, err)
	}
	return false, nil
}

func (ur *UserRepository) RejectFriendRequest(userEmail, targetLink string) error {
	rows, err := ur.DB.Queryx(Unsubscribe, userEmail, targetLink)
	defer utildb.CloseRows(rows)
	if err != nil {
		// TODO check
		return apperror.NewServerError(apperror.InternalServerError, err)

	}
	err = rows.Close()
	if err != nil {
		return apperror.NewServerError(apperror.InternalServerError, err)
	}
	return nil
}

func (ur *UserRepository) CreateUser(user *entities.User) (*entities.User, error) {

	tx, err := ur.DB.Beginx()
	if err != nil {

		// TODO check
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	var id uint
	err = tx.QueryRowx(CreateUser, user.Email, user.Password, user.FirstName, user.LastName, utils.CurrentTimeString()).Scan(&id)
	if err != nil {
		// TODO check
		return nil, apperror.NewServerError(apperror.InternalServerError, err)

	}

	_, err = tx.Exec(UpdateUserLink, fmt.Sprintf("id%d", id), id)
	if err != nil {
		// TODO check
		return nil, apperror.NewServerError(apperror.InternalServerError, err)

	}
	err = tx.Commit()
	if err != nil {
		// TODO check
		return nil, apperror.NewServerError(apperror.InternalServerError, err)

	}
	return user, nil
}

func (ur *UserRepository) UpdateUser(email string, user *dto.EditProfile) (*entities.User, error) {
	query := "update userprofile set "
	var fields []interface{}
	for name, el := range structs.Map(user) {
		field, ok := el.(*string)
		if !ok {
			continue
		}
		if dbName, exists := mapNames[name]; field != nil && exists {
			fields = append(fields, *field)
			query += fmt.Sprintf("%s = $%d, ", dbName, len(fields))
		}
	}
	query = strings.TrimSuffix(query, ", ")
	query += fmt.Sprintf(" where email = $%d", len(fields)+1)
	fields = append(fields, email)
	rows, err := ur.DB.Queryx(query, fields...)
	defer utildb.CloseRows(rows)
	if err != nil {
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	err = rows.Close()
	if err != nil {
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	return nil, nil
}

var mapNames = map[string]string{
	"Email":       "email",
	"NewPassword": "password",
	"FirstName":   "first_name",
	"LastName":    "last_name",
	"Link":        "link",
	"Sex":         "sex",
	"Status":      "status",
	"Bio":         "bio",
	"Birthday":    "birthday",
}

func (ur *UserRepository) DeleteUser(email string) error {
	var id int
	err := ur.DB.QueryRowx(DeleteUser, email).Scan(&id)
	if err != nil {
		return apperror.NewServerError(apperror.InternalServerError, err)
	}
	return nil
}

func (ur *UserRepository) GetPendingFriendRequests(user *entities.User, limit, offset int) ([]*entities.User, error) {
	rows, err := ur.DB.Queryx(PendingFriendRequestsById, user.ID, offset, limit)
	defer utildb.CloseRows(rows)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	var profiles []*entities.User
	for rows.Next() {
		profile := &entities.User{}
		err := rows.StructScan(profile)
		if err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
		profiles = append(profiles, profile)
	}
	return profiles, nil

}

//// IsFriend returns true when user is subscribed on target and vice versa
//func (ur *UserRepository) IsFriend(user, target *entities.User) (bool, error) {
//	var isFriend bool
//	err := ur.DB.QueryRowx(IsFriend, user.ID, target.ID).Scan(&isFriend)
//	if err != nil {
//		return false, apperror.InternalServerError
//	}
//	return isFriend, nil
//}
//
//// IsSubscriber returns true when user is subscribed on target
//func (ur *UserRepository) IsSubscriber(user, target *entities.User) (bool, error) {
//	var isSub bool
//	err := ur.DB.QueryRowx(IsSubscriber, user.ID, target.ID).Scan(&isSub)
//	if err != nil {
//		return false, apperror.InternalServerError
//	}
//	return isSub, nil
//}
//
//// IsSubscribed returns true when target is subscribed on user (rejected request)
//func (ur *UserRepository) IsSubscribed(user, target *entities.User) (bool, error) {
//	var isSub bool
//	err := ur.DB.QueryRowx(IsSubscribed, user.ID, target.ID).Scan(&isSub)
//	if err != nil {
//		return false, apperror.InternalServerError
//	}
//	return isSub, nil
//}

// HasPendingRequest returns true when target is subscribed on user (unseen yet request)
//func (ur *UserRepository) HasPendingRequest(user, target *entities.User) (bool, error) {
//	var pending bool
//	err := ur.DB.QueryRowx(HasPendingRequest, user.ID, target.ID).Scan(&pending)
//	if err != nil {
//		return false, apperror.InternalServerError
//	}
//	return pending, nil
//}

func (ur *UserRepository) SearchUserByName(searchDTO *dto.GlobalSearchDTO) ([]*entities.UserInfo, error) {
	// TODO: сначала выбрать из друзей
	searchName := utils.NormalizeString(*searchDTO.SearchQuery)
	splitSearchName := strings.Split(searchName, " ")

	rows, err := ur.DB.Queryx("SELECT id, first_name, last_name FROM UserProfile AS profile")
	if err != nil {
		return nil, err
	}

	distances := make([]int, len(splitSearchName))

	user := userForSearch{}
	usersBatch := make([]userForSearch, 0, *searchDTO.BatchSize)
	var usersAmount uint

	maxLevenshteinDistance := float64(utils.GetMaxLength(splitSearchName...)) * 0.5

MAIN:
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			return nil, err
		}

		// считаем минимальное расстояние между искомым именем и именем из базы
		for i := 0; i < len(splitSearchName); i++ {
			firstDistance := levenshtein.Distance(strings.ToLower(user.FirstName), splitSearchName[i])
			secondDistance := levenshtein.Distance(strings.ToLower(user.LastName), splitSearchName[i])

			distances[i] = utils.Min(firstDistance, secondDistance)
		}

		user.Distance = float64(utils.SliceMin(distances))
		if user.Distance > maxLevenshteinDistance {
			continue
		}

		// обновляем список с наилучшими мэтчами по имени
		var i uint
		for i = 0; i < usersAmount; i++ {
			if user.Distance < usersBatch[i].Distance {
				if usersAmount < *searchDTO.BatchSize+*searchDTO.Offset {
					usersBatch = append(usersBatch, userForSearch{})
				}

				utils.ShiftRight(usersBatch, int(i), 1)
				usersBatch[i] = user
				usersAmount++
				continue MAIN
			}
		}

		if usersAmount < *searchDTO.BatchSize {
			usersBatch = append(usersBatch, user)
			usersAmount++
		}
	}

	var users = make([]*entities.UserInfo, 0, *searchDTO.BatchSize)

	if *searchDTO.Offset >= uint(len(usersBatch)) {
		return nil, nil
	}
	for ind, info := range usersBatch[*searchDTO.Offset:] {
		users = append(users, nil)
		users[ind] = new(entities.UserInfo)
		err = ur.DB.Get(users[ind], "SELECT first_name, last_name, link,  url FROM UserProfile AS profile LEFT JOIN Photo ON profile.avatar_id = photo.id WHERE profile.id = $1", info.ID)
		if err != nil {
			return nil, err
		}
	}

	return users, nil
}

type userForSearch struct {
	ID        uint   `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Distance  float64
}
