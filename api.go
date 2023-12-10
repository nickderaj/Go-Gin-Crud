package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Run() {
	r := gin.Default()

	r.GET("/", HealthCheck)
	r.GET("/users", GetUsers)
	r.GET("/users/:id", GetUser)
	r.POST("/users", CreateUser)
	r.PUT("/users/:id", UpdateUser)
	r.DELETE("/users/:id", DeleteUser)

	err := r.Run()
	if err != nil {
		log.Fatal("Failed to run Server")
	}
}

func HandleError(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err,
	})
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		HandleError(c, "invalid id")
		return
	}

	var user User
	res := DB.First(&user, id) // find product with integer primary key

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		HandleError(c, "user not found")
		return
	}

	if res.Error != nil {
		HandleError(c, "error fetching user")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "retrieved user",
		"user":    user,
	})
}

func GetUsers(c *gin.Context) {
	var users []User
	res := DB.Find(&users)

	if res.Error != nil {
		HandleError(c, "error fetching users")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "retrieved users",
		"users":   users,
	})
}

func CreateUser(c *gin.Context) {
	var body struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	err := c.Bind(&body)
	if err != nil {
		return
	}

	if body.Name == "" || body.Email == "" {
		HandleError(c, "missing required parameters")
		return
	}

	user := User{Name: body.Name, Email: body.Email}
	res := DB.Create(&user)

	if res.Error != nil {
		HandleError(c, "failed to create user")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "created user",
		"user":    user,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		HandleError(c, "invalid id")
		return
	}

	var body struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	err := c.Bind(&body)
	if err != nil {
		return
	}

	var user User
	res := DB.First(&user, id) // find product with integer primary key

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		HandleError(c, "user not found")
		return
	}

	if res.Error != nil {
		HandleError(c, "error fetching user")
		return
	}

	DB.Model(&user).Updates(User{
		Name:  body.Name,
		Email: body.Email,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "updated user",
		"user":    user,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		HandleError(c, "invalid id")
		return
	}

	res := DB.Delete(&User{}, id)

	if res.Error != nil {
		HandleError(c, "error deleting user")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted",
	})
}
