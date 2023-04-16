package policies

import (
	"github.com/diy0663/gohub/app/models/topic"
	"github.com/diy0663/gohub/pkg/auth"
	"github.com/gin-gonic/gin"
)

// 自己才能修改自己的
func CanModifyTopic(c *gin.Context, topicModel topic.Topic) bool {
	return auth.CurrentUID(c) == topicModel.UserID
}

// func CanViewTopic(c *gin.Context, topicModel topic.Topic) bool {}
// func CanCreateTopic(c *gin.Context, topicModel topic.Topic) bool {}
// func CanUpdateTopic(c *gin.Context, topicModel topic.Topic) bool {}
// func CanDeleteTopic(c *gin.Context, topicModel topic.Topic) bool {}
