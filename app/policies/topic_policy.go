package policies

import (
	"github.com/diy0663/gohub/app/models/topic"
	"github.com/diy0663/gohub/pkg/auth"
	"github.com/gin-gonic/gin"
)

// todo 保存之后自动引入 其他的package
// 针对下面的各个授权策略做修改

func CanModifyTopic(c *gin.Context, topicModel topic.Topic) bool {
	return auth.CurrentUID(c) == topicModel.UserID
}

// func CanViewTopic(c *gin.Context, topicModel topic.Topic) bool {}
// func CanCreateTopic(c *gin.Context, topicModel topic.Topic) bool {}
// func CanUpdateTopic(c *gin.Context, topicModel topic.Topic) bool {}
func CanDeleteTopic(c *gin.Context, topicModel topic.Topic) bool {
	// 暂时设定只能自己删自己 的
	return auth.CurrentUID(c) == topicModel.UserID
}
