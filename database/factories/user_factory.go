package factories

import (
	"github.com/bxcodec/faker/v3"
	"github.com/diy0663/go_project_packages/hash"
	"github.com/diy0663/gohub/app/models/user"
)

func MakeUsers(count int) []user.User {

	var objs []user.User
	for i := 0; i < count; i++ {
		objs = append(objs, user.User{

			Name:     faker.Name(),
			Email:    faker.Email(),
			Phone:    faker.Phonenumber(),
			Password: hash.BcryptHash("123456"),
		})
	}
	return objs
}
