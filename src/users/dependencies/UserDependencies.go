package dependencies

import (
	"database/sql"
	"demo/src/users/application"
	userInfra "demo/src/users/infraestructure/controllers"
	userRouters "demo/src/users/infraestructure/interfaces/http/routers"
	userRepositories "demo/src/users/infraestructure/persistence"

	"github.com/gin-gonic/gin"
)

type UsersDependencies struct {
	DB *sql.DB
}

func NewUsersDependencies(db *sql.DB) *UsersDependencies {
	return &UsersDependencies{DB: db}
}

func (d *UsersDependencies) Execute(r *gin.Engine) {
	userRepository := userRepositories.NewUserRepository(d.DB)

	createUserUseCase := application.NewCreateUser(*userRepository)
	createUserController := userInfra.NewCreateUserController(*createUserUseCase)

	deleteUserUseCase := application.NewDeleteUserByID(*userRepository)
	deleteUserByIDController := userInfra.NewDeleteUserByIDController(*deleteUserUseCase)

	getAllUserUseCase := application.NewGetAllUsers(*userRepository)
	getAllUserController := userInfra.NewGetAllUsersController(*getAllUserUseCase)

	updateUserByIDUseCase := application.NewEditUser(*userRepository)
	updateUserByIDController := userInfra.NewUpdateUserByIDController(*updateUserByIDUseCase)

	userRouters.AttachUserRoutes(
		r,
		createUserController,
		deleteUserByIDController,
		getAllUserController,
		updateUserByIDController,
	)
}