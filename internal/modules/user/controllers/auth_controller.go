package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/realwebdev/blog/internal/modules/user/requests/auth"
	userService "github.com/realwebdev/blog/internal/modules/user/services"
	"github.com/realwebdev/blog/pkg/converters"
	"github.com/realwebdev/blog/pkg/errors"
	"github.com/realwebdev/blog/pkg/html"
	"github.com/realwebdev/blog/pkg/old"
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

	var request auth.RegisterRequest

	if err := c.ShouldBind(&request); err != nil {
		validationErrors := errors.FromValidation(err)

		sessions.Set(c, "errors", converters.MapToString(validationErrors))

		sessions.Set(c, "old", converters.UrlValuesToString(old.FromContext(c)))

		c.Redirect(http.StatusFound, "/register")
		return
	}

	if exists, _ := controller.userServiceInterface.CheckUserExist(request.Email); exists {
		sessions.Set(c, "errors", converters.MapToString(map[string]string{
			"email": "Email already exists",
		}))
		sessions.Set(c, "old", converters.UrlValuesToString(old.FromContext(c)))
		c.Redirect(http.StatusFound, "/register")
		return
	}

	// create the user
	user, err := controller.userServiceInterface.RegisterUser(request)
	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	// login the user
	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	log.Printf("The user created successfully with a name %s \n", user.Name)
	c.Redirect(http.StatusFound, "/")
}

func (controller *Controller) Login(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/login", gin.H{
		"title": "Login",
	})
}

func (controller *Controller) HandleLogin(c *gin.Context) {
	var request auth.LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		validationErrors := errors.FromValidation(err)

		sessions.Set(c, "errors", converters.MapToString(validationErrors))

		sessions.Set(c, "old", converters.UrlValuesToString(old.FromContext(c)))

		c.Redirect(http.StatusFound, "/login")
		return
	}

	user, err := controller.userServiceInterface.HandleUserLogin(request)
	if err != nil {
		sessions.Set(c, "errors", converters.MapToString(map[string]string{
			"email": "Invalid credentials",
		}))
		sessions.Set(c, "old", converters.UrlValuesToString(old.FromContext(c)))

		c.Redirect(http.StatusFound, "/login")
		return
	}

	// login the user
	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	log.Printf("The user logged in successfully with a name %s \n", user.Name)
	c.Redirect(http.StatusFound, "/")
}

func (controller *Controller) HandleLogout(c *gin.Context) {
	sessions.Remove(c, "auth")
	c.Redirect(http.StatusFound, "/")
}
