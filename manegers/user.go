package manegers

import "github.come/salmanfaris22/Skill-map/models"

type UserMangege struct {
}

func NewUSerManager() *UserMangege {
	return &UserMangege{}
}
func (u *UserMangege) Creat(user *models.User) (*models.User, error) {
	return nil, nil
}
