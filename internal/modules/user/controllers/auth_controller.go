package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realwebdev/blog/internal/modules/user/requests/auth"
	userService "github.com/realwebdev/blog/internal/modules/user/services"
	"github.com/realwebdev/blog/pkg/html"
)

type Controller struct {
	userServiceInterface userService.UserServiceInterface
}

func New(ctrl Controller) *Controller {
	return &Controller{
		userServiceInterface: ctrl.userServiceInterface,
	}
}

func (controller *Controller) Register(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{
		"title": "Register",
	})
}

func (controller *Controller) HandleRegister(c *gin.Context) {
	// validate the request
	var request auth.RegisterRequest

	if err := c.ShouldBind(&request); err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	// create the user
	user, err := controller.userServiceInterface.RegisterUser(request)
	if err != nil {
		c.Redirect(http.StatusFound, "/register")
	}

	log.Printf("The user created successfully with a name %s \n", user.Name)

	// (service layer to create the user inside the database)

	// Check if the there is any error on the user creation

	// after creating the user redirect to the homepage

	c.JSON(http.StatusOK, gin.H{
		"message": "Register done",
	})
}
