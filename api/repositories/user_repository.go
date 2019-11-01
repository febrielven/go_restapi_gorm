package repositories

import (
	"github.com/febrielven/go_restapi_gorm/api/models"
)

type UserRepository interface {
	Prepare(*models.User)
	FindAllUsers() ([]*models.User, error)
	Validate(*models.User)
	SaveUser(*models.User) (*models.User, error)
	FindUserByID(*models.User, uint32) (models.User, error)
}
