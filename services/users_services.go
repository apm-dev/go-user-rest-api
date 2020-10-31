package services

import (
	"github.com/apm-dev/go-user-rest-api/domain/users"
	"github.com/apm-dev/go-user-rest-api/requests"
	"github.com/apm-dev/go-user-rest-api/utils/crypto_utils"
	"github.com/apm-dev/go-user-rest-api/utils/date_utils"
	"github.com/apm-dev/go-user-rest-api/utils/errors"
)

var UserService userServiceInterface = &userService{}

type userService struct {
}

type userServiceInterface interface {
	GetUsers() (users.Users, *errors.RestError)
	FindUser(id int64) (*users.User, *errors.RestError)
	CreateUser(u users.User) (*users.User, *errors.RestError)
	StoreUser(r *requests.UserStoreRequest) (*users.User, *errors.RestError)
	UpdateUser(r *requests.UserUpdateRequest, id int64, isPartial bool) (*users.User, *errors.RestError)
	DeleteUser(id int64) *errors.RestError
}

func (*userService) GetUsers() (users.Users, *errors.RestError) {
	return users.User{}.All()
}

func (*userService) FindUser(id int64) (*users.User, *errors.RestError) {
	u := users.User{ID: id}
	err := u.Find()
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (*userService) CreateUser(u users.User) (*users.User, *errors.RestError) {
	if err := u.Insert(); err != nil {
		return nil, err
	}
	return &u, nil
}

func (*userService) StoreUser(r *requests.UserStoreRequest) (*users.User, *errors.RestError) {
	user := users.User{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Password:  crypto_utils.GetSHA256(r.Password),
		CreatedAt: date_utils.GetNow(),
		UpdatedAt: date_utils.GetNow(),
	}
	if err := user.Insert(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (*userService) UpdateUser(r *requests.UserUpdateRequest, id int64, isPartial bool) (*users.User,
	*errors.RestError) {
	var user *users.User
	var err *errors.RestError
	if isPartial {
		user, err = UserService.FindUser(id)
		if err != nil {
			return nil, err
		}
		if r.FirstName != "" {
			user.FirstName = r.FirstName
		}
		if r.LastName != "" {
			user.LastName = r.LastName
		}
		if r.Email != "" {
			user.Email = r.Email
		}
	} else {
		user = &users.User{
			ID:        id,
			FirstName: r.FirstName,
			LastName:  r.LastName,
			Email:     r.Email,
		}
	}
	user.UpdatedAt = date_utils.GetNow()
	if err := user.Update(); err != nil {
		return nil, err
	}
	return user, nil
}

func (*userService) DeleteUser(id int64) *errors.RestError {
	err := (&users.User{ID: id}).Delete()
	if err != nil {
		return err
	}
	return nil
}
