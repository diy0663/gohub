package auth

import (
	"fmt"
	"net/http"

	v1 "github.com/diy0663/gohub/app/http/controllers/api/v1"
	"github.com/diy0663/gohub/app/models/user"
	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	type IsPhoneExistRequest struct {
		Phone string `json:"phone"`
	}
	request := IsPhoneExistRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})

}
