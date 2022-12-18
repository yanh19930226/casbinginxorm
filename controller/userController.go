package controller

import (
	"casbinginxorm/models"
	"casbinginxorm/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// ListUsers godoc
//
//	@Summary		List users
//	@Description	get users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string	false	"name search by name"
//	@Success		200	{array}		models.User
//	@Failure		400	{object}	utils.HTTPError
//	@Failure		404	{object}	utils.HTTPError
//	@Failure		500	{object}	utils.HTTPError
//	@Router			/api/v1/users [get]
func (c *UserController) ListUsers(ctx *gin.Context) {
	q := ctx.Request.URL.Query().Get("name")
	users, err := models.UsersAll(q)
	if err != nil {
		utils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// AddUser godoc
//
//	@Summary		Add an user
//	@Description	add by json user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.AddUser	true	"Add user"
//	@Success		200		{object}	models.User
//	@Failure		400		{object}	utils.HTTPError
//	@Failure		404		{object}	utils.HTTPError
//	@Failure		500		{object}	utils.HTTPError
//	@Router			/api/v1/users [post]
func (c *UserController) AddUser(ctx *gin.Context) {
	var addUser models.AddUser
	if err := ctx.ShouldBindJSON(&addUser); err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	// //Validation
	// if err := addUser.Validation(); err != nil {
	// 	utils.NewError(ctx, http.StatusBadRequest, err)
	// 	return
	// }
	account := models.User{
		Name: addUser.Name,
	}
	lastID, err := account.Insert()
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	account.ID = lastID
	ctx.JSON(http.StatusOK, account)
}

// UpdateUser godoc
//
//	@Summary		Update an user
//	@Description	Update by json user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"User ID"
//	@Param			user	body		models.UpdateUser	true	"Update user"
//	@Success		200		{object}	models.User
//	@Failure		400		{object}	utils.HTTPError
//	@Failure		404		{object}	utils.HTTPError
//	@Failure		500		{object}	utils.HTTPError
//	@Router			/api/v1/users/{id} [patch]
func (c *UserController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	var updateUser models.UpdateUser
	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	user := models.User{
		ID:       aid,
		Name:     updateUser.Name,
		Mobile:   updateUser.Mobile,
		Password: updateUser.Password,
	}
	err = user.Update()
	if err != nil {
		utils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// DeleteUser godoc
//
//	@Summary		Delete an user
//	@Description	Delete by user ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"	Format(int64)
//	@Success		204	{object}	models.User
//	@Failure		400	{object}	utils.HTTPError
//	@Failure		404	{object}	utils.HTTPError
//	@Failure		500	{object}	utils.HTTPError
//	@Router			/api/v1/users/{id} [delete]
func (c *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	err = models.DeleteUser(aid)
	if err != nil {
		utils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
