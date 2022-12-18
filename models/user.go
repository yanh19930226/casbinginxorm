package models

import (
	"errors"
	"fmt"
)

// User example
type User struct {
	ID       int    `json:"id" format:"int64"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
}

// AddUser
type AddUser struct {
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

// UpdateUser
type UpdateUser struct {
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

var userMaxID = 3
var users = []User{
	{ID: 1, Name: "admin", Password: "admin", Mobile: "18650482503"},
	{ID: 2, Name: "yanh", Password: "yanh", Mobile: "18650482503"},
	{ID: 3, Name: "root", Password: "root", Mobile: "18650482503"},
}

// UsersAll example
func UsersAll(q string) ([]User, error) {
	if q == "" {
		return users, nil
	}
	as := []User{}
	for k, v := range users {
		if q == v.Name {
			as = append(as, users[k])
		}
	}
	return as, nil
}

// UserOne example
func UserOne(id int) (User, error) {
	for _, v := range users {
		if id == v.ID {
			return v, nil
		}
	}
	return User{}, errors.New("no rows in result set")
}

// Insert example
func (a User) Insert() (int, error) {
	userMaxID++
	a.ID = userMaxID
	a.Name = fmt.Sprintf("user_%s", a.Name)
	a.Mobile = fmt.Sprintf("user_%s", a.Mobile)
	a.Password = fmt.Sprintf("user_%s", a.Password)
	users = append(users, a)
	return userMaxID, nil
}

// Delete example
func DeleteUser(id int) error {
	for k, v := range users {
		if id == v.ID {
			users = append(users[:k], users[k+1:]...)
			return nil
		}
	}
	return errors.New("delete error")
}

// Update example
func (a User) Update() error {
	for k, v := range users {
		if a.ID == v.ID {
			users[k].Name = a.Name
			users[k].Mobile = a.Mobile
			users[k].Password = a.Password
			return nil
		}
	}
	return errors.New("update error")
}
