package view

import (
	"errors"
	"log"
	"time"

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
	input.Created = time.Now().String()

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

// UpdateUser - update user details
func (p Provider) UpdateUser(c *gin.Context) {
	userDetails := model.User{}
	_id := c.Param("id")
	c.BindJSON(&userDetails)

	userDetails.ID = _id

	err := p.db.UpdateUser(userDetails)
	if err != nil {
		util.HTTPError(c, constants.BadRequestCode, err)
		return
	}
	c.JSON(200, map[string]string{"Status": "OK"})
}

// Login - user login
func (p Provider) Login(c *gin.Context) {
	input := model.User{}

	c.BindJSON(&input)
	log.Printf("%#v", input)
	user, err := p.db.GetUserByEmail(input.Email)
	log.Printf("%#v", user)
	if err != nil {
		util.HTTPError(c, constants.BadRequestCode, err)
		return
	}

	hashCheck := util.GeneratePasswordHash(input.Password, user.Salt)
	log.Print(*hashCheck)
	if user.Password == *hashCheck {
		token := model.Token{
			UserID:  user.ID,
			Created: time.Now().String(),
			ID:      uuid.New().String(),
		}
		_, err := p.db.CreateToken(token)
		if err != nil {
			util.HTTPError(c, constants.BadRequestCode, err)
			return
		}
		c.JSON(200, token)
		return
	}

	util.HTTPError(c, constants.BadRequestCode, errors.New("Please provide proper credentials"))
	return
}

// Logout - logs the user out
func (p Provider) Logout(c *gin.Context) {
	token := c.Param("id")

	_, err := p.db.DeleteToken(token)
	if err != nil {
		util.HTTPError(c, constants.BadRequestCode, err)
		return
	}
	c.JSON(200, map[string]string{"Status": "OK"})
}
