package infraestructure

import (
	"demo/src/users/application"
	"demo/src/users/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateUserByIDController struct {
	up application.EditUser
}

type UpdateUserRequest struct {
	Id       int  `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     int    `json:"role" binding:"required"`
}

func NewUpdateUserByIDController(up application.EditUser) *UpdateUserByIDController {
	return &UpdateUserByIDController{up: up}
}

func (up_c *UpdateUserByIDController) Execute(c *gin.Context) {
	var req UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updateUser := entities.User{
		Id:       req.Id,
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	}

	err := up_c.up.Execute(&updateUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user", "id": updateUser.Id})
		return
	}

	c.JSON(http.StatusOK, gin.H{"successful": "Updated user", "id": req.Id})
}
