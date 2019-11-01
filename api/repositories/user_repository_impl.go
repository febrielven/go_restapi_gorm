package repositories

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/febrielven/go_restapi_gorm/api/config"
	"github.com/febrielven/go_restapi_gorm/api/models"
	"github.com/jinzhu/gorm"
)

func NewUserRepository() *userRepository {
	db := config.GetDB()

	return &userRepository{
		db: db,
	}
}

type userRepository struct {
	db *gorm.DB
}

func (server *userRepository) Prepare(user *models.User) {

	user.ID = 0
	user.Nickname = html.EscapeString(strings.TrimSpace(user.Nickname))
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}

func (server *userRepository) Validate(user *models.User, action string) error {

	switch strings.ToLower(action) {
	case "update":
		if user.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if user.Password == "" {
			return errors.New("Required Password")
		}
		if user.Email == "" {
			return errors.New("Required Email")
		}

		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if user.Password == "" {
			return errors.New("Required Password")
		}
		if user.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if user.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if user.Password == "" {
			return errors.New("Required Password")
		}
		if user.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}

}

// func FindAllUser(u *models.User){
// }

func (server *userRepository) FindAllUser() (*[]models.User, error) {
	var err error
	users := []models.User{}
	//SELECT * FROM users limit 100
	err = server.db.Debug().Model(&models.User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]models.User{}, err
	}

	return &users, err
}

func (server *userRepository) SaveUser(user *models.User) (*models.User, error) {
	var err error
	err = server.db.Debug().Create(&user).Error

	if err != nil {
		return &models.User{}, err
	}
	return user, nil

}

func (server *userRepository) FindUserByID(user *models.User, uid uint32) (*models.User, error) {
	var err error
	//SELECT * FROM users where id = ? LIMIT 1;
	err = server.db.Debug().Model(models.User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &models.User{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &models.User{}, errors.New("User Not Found")
	}

	return user, err
}
