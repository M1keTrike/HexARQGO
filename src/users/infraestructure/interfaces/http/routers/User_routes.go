package routers

import (
	infrastructure "demo/src/users/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func AttachUserRoutes(
	r *gin.Engine,
	createUserController *infrastructure.CreateUserController,
	deleteUserByIDController *infrastructure.DeleteUserByIDController,
	getAllUsersController *infrastructure.GetAllUsersController,
	updateUserByIDController *infrastructure.UpdateUserByIDController,
) {
	usersGroup := r.Group("/users")
	{
		usersGroup.POST("", createUserController.Execute)
		usersGroup.DELETE("", deleteUserByIDController.Execute)
		usersGroup.GET("", getAllUsersController.Execute)
		usersGroup.PUT("", updateUserByIDController.Execute)
	}
}
