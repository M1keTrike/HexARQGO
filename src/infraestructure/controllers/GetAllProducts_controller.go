package infraestructure

import (
	"demo/src/application"

	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllProductController struct {
	ga application.GetAllProducts
}

func NewGetAllProductsController(ga application.GetAllProducts) *GetAllProductController {
	return &GetAllProductController{ga:ga}
}

func (ga_c *GetAllProductController) Execute(c *gin.Context) {
	

	res, err := ga_c.ga.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrive all products"})
		return
	}

	
	c.JSON(http.StatusOK, gin.H{"Products": res})
}
