package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realwebdev/blog/internal/modules/user/requests/auth"
	userService "github.com/realwebdev/blog/internal/modules/user/services"
	"github.com/realwebdev/blog/pkg/converters"
	"github.com/realwebdev/blog/pkg/errors"
	"github.com/realwebdev/blog/pkg/html"
	"github.com/realwebdev/blog/pkg/sessions"
)

type Controller struct {
	userServiceInterface userService.UserServiceInterface
}

func New(uSI userService.UserServiceInterface) *Controller {
	return &Controller{
		userServiceInterface: uSI,
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
		validationErrors := errors.FromValidation(err)

		sessions.Set(c, "errors", converters.MapToString(validationErrors))
		c.Redirect(http.StatusFound, "/register")
		return
	}

	// create the user
	user, err := controller.userServiceInterface.RegisterUser(request)
	if err != nil {
		c.Redirect(http.StatusFound, "/register")
	}

	log.Printf("The user created successfully with a name %s \n", user.Name)
	c.Redirect(http.StatusFound, "/")

	// (service layer to create the user inside the database)

	// Check if the there is any error on the user creation

	// after creating the user redirect to the homepage
}
