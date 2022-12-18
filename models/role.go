package models

import (
	"errors"
)

// Role example
type Role struct {
	ID     int    `json:"id" format:"int64"`
	Name   string `json:"name"`
	Status int    `json:"status" format:"int64"`
}

// AddRole
type AddRole struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
}

// UpdateRole
type UpdateRole struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
}

var roleMaxID = 3
var roles = []Role{
	{ID: 1, Name: "admin", Status: 1},
	{ID: 2, Name: "customer", Status: 0},
	{ID: 3, Name: "root", Status: 1},
}

// RolesAll example
func RolesAll(q string) ([]Role, error) {
	if q == "" {
		return roles, nil
	}
	as := []Role{}
	for k, v := range roles {
		if q == v.Name {
			as = append(as, roles[k])
		}
	}
	return as, nil
}

// RoleOne example
func RoleOne(id int) (Role, error) {
	for _, v := range roles {
		if id == v.ID {
			return v, nil
		}
	}
	return Role{}, errors.New("no rows in result set")
}

// Insert example
func (a Role) Insert() (int, error) {
	roleMaxID++
	a.ID = roleMaxID
	roles = append(roles, a)
	return roleMaxID, nil
}

// Delete example
func DeleteRole(id int) error {
	for k, v := range roles {
		if id == v.ID {
			roles = append(roles[:k], roles[k+1:]...)
			return nil
		}
	}
	return errors.New("delete error")
}

// Update example
func (a Role) Update() error {
	for k, v := range roles {
		if a.ID == v.ID {
			roles[k].Name = a.Name
			roles[k].Status = a.Status
			return nil
		}
	}
	return errors.New("update error")
}
