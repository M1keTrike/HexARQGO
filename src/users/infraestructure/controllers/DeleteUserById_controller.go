package infraestructure

import (
	"demo/src/users/application"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteUserByIDRequest struct {
	Id int `json:"id" binding:"required"`
}

type DeleteUserByIDController struct{
	du application.DeleteUserByID
}

func NewDeleteUserByIDController(du application.DeleteUserByID) *DeleteUserByIDController {
	return &DeleteUserByIDController{du: du}
}

func (du_c *DeleteUserByIDController) Execute(c *gin.Context) {
	var req DeleteUserByIDRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	idToDelete := req.Id
	fmt.Println(idToDelete)

	err := du_c.du.Execute(idToDelete)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user", "id": idToDelete})
		return
	}

	c.JSON(http.StatusOK, gin.H{"successful": "Deleted user", "id": idToDelete})
}
