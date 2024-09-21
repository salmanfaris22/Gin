package handlers

import (
	"github.come/salmanfaris22/Skill-map/manegers"
)

type UserHadler struct {
	grupName   string
	userManger *manegers.UserMangege
}

func NewUserHandler(userManger *manegers.UserMangege) *UserHadler {
	return &UserHadler{
		"user",
		userManger,
	}
}

func Register() *UserHadler {
	return &UserHadler{}
}

func (uh *UserHadler) CreacUSer() {

}
