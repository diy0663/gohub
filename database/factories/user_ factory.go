package factories

import (
	"github.com/bxcodec/faker/v3"
	"github.com/diy0663/gohub/app/models/user"
	"github.com/diy0663/gohub/pkg/helper"
)

func MakeUsers(times int) []user.User {
	var objs []user.User

	faker.SetGenerateUniqueValues(true)
	for i := 0; i < times; i++ {
		model := user.User{

			Name:     faker.Username(),
			Email:    faker.Email(),
			Phone:    helper.RandomNumber(11),
			Password: "123456",
		}
		objs = append(objs, model)
	}
	return objs
}
