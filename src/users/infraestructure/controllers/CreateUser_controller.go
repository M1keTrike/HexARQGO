package infraestructure

import (
	"demo/src/users/application"
	"demo/src/users/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     int    `json:"role" binding:"required"`
}

type CreateUserController struct {
	cu application.CreateUser
}

func NewCreateUserController(cu application.CreateUser) *CreateUserController {
	return &CreateUserController{cu: cu}
}

func (cu_c *CreateUserController) Execute(c *gin.Context) {
	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user := entities.User{
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	}

	err := cu_c.cu.Execute(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"username": user.Username, "role": user.Role})
}
