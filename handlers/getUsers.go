package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"users-paging-app/models"
)

func (s Service) GetUsers(c *gin.Context) {
	count := c.Query("count")
	offset := c.Query("offset")
	if count == "" {
		c.String(http.StatusBadRequest, "missing count parameter")
		return
	}
	if offset == "" {
		offset = "0"
	}

	var users = make([]models.User, 0)
	rows, err := s.Db.Query("SELECT * FROM users LIMIT ?,?;", offset, count)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	for rows.Next() {
		var us models.User
		err = rows.Scan(&us.Id, &us.Name)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		users = append(users, us)
	}

	usersResponse := models.UsersResponse{
		Count: len(users),
		Users: users,
	}

	c.JSON(http.StatusOK, usersResponse)
}
