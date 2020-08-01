package view

import (
	"github.com/Blaze2305/url_short/internal/pkg/constants"
	"github.com/Blaze2305/url_short/internal/pkg/model"
	"github.com/Blaze2305/url_short/internal/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateUser - create a user
func (p Provider) CreateUser(c *gin.Context) {
	input := model.User{}

	c.BindJSON(&input)

	_id := uuid.New().String()
	input.ID = _id

	passhash, salt := util.GenerateSHA256Hash(input.Password)

	input.Password = *passhash
	input.Salt = *salt

	user, err := p.db.CreateUser(input)
	if err != nil {
		util.HTTPError(c, constants.BadRequestCode, err)
		return
	}

	c.JSON(200, *user)
}

// GetUser - get the details of the user given id
func (p Provider) GetUser(c *gin.Context) {
	userID := c.Param("id")

	user, err := p.db.GetUser(userID)
	if err != nil {
		util.HTTPError(c, constants.BadRequestCode, err)
		return
	}

	c.JSON(200, *user)
}

// DeleteUser - delete a user with the given id
func (p Provider) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	user, err := p.db.DeleteUser(userID)
	if err != nil {
		util.HTTPError(c, constants.BadRequestCode, err)
		return
	}
	c.JSON(200, map[string]string{"Status": "OK", "ID": *user})
}
