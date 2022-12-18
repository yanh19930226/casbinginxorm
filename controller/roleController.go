package controller

import (
	"casbinginxorm/models"
	"casbinginxorm/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct{}

// ListRoles godoc
//
//	@Summary		List roles
//	@Description	get roles
//	@Tags			roles
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string	false	"name search by name"
//	@Success		200	{array}		models.Role
//	@Failure		400	{object}	utils.HTTPError
//	@Failure		404	{object}	utils.HTTPError
//	@Failure		500	{object}	utils.HTTPError
//	@Router			/api/v1/roles [get]
func (c *RoleController) ListRoles(ctx *gin.Context) {
	q := ctx.Request.URL.Query().Get("name")
	users, err := models.RolesAll(q)
	if err != nil {
		utils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// AddRole godoc
//
//	@Summary		Add an role
//	@Description	add by json role
//	@Tags			roles
//	@Accept			json
//	@Produce		json
//	@Param			role	body		models.AddRole	true	"Add role"
//	@Success		200		{object}	models.Role
//	@Failure		400		{object}	utils.HTTPError
//	@Failure		404		{object}	utils.HTTPError
//	@Failure		500		{object}	utils.HTTPError
//	@Router			/api/v1/roles [post]
func (c *RoleController) AddRole(ctx *gin.Context) {
	var addRole models.AddRole
	if err := ctx.ShouldBindJSON(&addRole); err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	// //Validation
	// if err := addUser.Validation(); err != nil {
	// 	utils.NewError(ctx, http.StatusBadRequest, err)
	// 	return
	// }
	role := models.Role{
		Name:   addRole.Name,
		Status: addRole.Status,
	}
	lastID, err := role.Insert()
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	role.ID = lastID
	ctx.JSON(http.StatusOK, role)
}

// UpdateRole godoc
//
//	@Summary		Update an role
//	@Description	Update by json role
//	@Tags			roles
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Role ID"
//	@Param			role	body		models.UpdateRole	true	"Update role"
//	@Success		200		{object}	models.Role
//	@Failure		400		{object}	utils.HTTPError
//	@Failure		404		{object}	utils.HTTPError
//	@Failure		500		{object}	utils.HTTPError
//	@Router			/api/v1/roles/{id} [patch]
func (c *RoleController) UpdateRole(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	var updateRole models.UpdateRole
	if err := ctx.ShouldBindJSON(&updateRole); err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	role := models.Role{
		ID:     aid,
		Name:   updateRole.Name,
		Status: updateRole.Status,
	}
	err = role.Update()
	if err != nil {
		utils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, role)
}

// DeleteRole godoc
//
//	@Summary		Delete an role
//	@Description	Delete by role ID
//	@Tags			roles
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Role ID"	Format(int64)
//	@Success		204	{object}	models.Role
//	@Failure		400	{object}	utils.HTTPError
//	@Failure		404	{object}	utils.HTTPError
//	@Failure		500	{object}	utils.HTTPError
//	@Router			/api/v1/roles/{id} [delete]
func (c *RoleController) DeleteRole(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	err = models.DeleteRole(aid)
	if err != nil {
		utils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
