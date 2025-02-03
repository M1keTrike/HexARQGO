package infraestructure

import (
	"demo/src/users/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllUsersController struct {
	ga application.GetAllUsers
}

func NewGetAllUsersController(ga application.GetAllUsers) *GetAllUsersController {
	return &GetAllUsersController{ga: ga}
}

func (ga_c *GetAllUsersController) Execute(c *gin.Context) {
	res, err := ga_c.ga.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve all users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Users": res})
}