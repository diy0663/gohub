package auth

import (
	"fmt"
	"net/http"

	v1 "github.com/diy0663/gohub/app/http/controllers/v1"
	"github.com/diy0663/gohub/app/models/user"
	"github.com/gin-gonic/gin"
)

// // 本API必须要实现的功能
// type SignupInterface interface {
// 	IsPhoneExist(c *gin.Context)
// 	IsEmailExist(c *gin.Context)
// }

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	type PhoneExistsRequest struct {
		Phone string `json:"phone,omitempty" valid:"phone"`
	}

	request := PhoneExistsRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"exists": user.IsPhoneExists(request.Phone),
	})

}
