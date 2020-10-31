package users

import (
	"time"
)

type Resource struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Mail      string    `json:"mail"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) Marshall() interface{} {
	return Resource{
		ID:        u.ID,
		Name:      u.FirstName + " " + u.LastName,
		Mail:      u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (users Users) Marshall() []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall()
	}
	return result
}
