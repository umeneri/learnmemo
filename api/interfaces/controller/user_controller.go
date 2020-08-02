package controller

import (
	"api/interfaces/auth"
	"api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lunny/log"
)

type UserController interface {
	Index(c *gin.Context)
	Login(c *gin.Context)
	GoogleCallback(c *gin.Context)
}

type userController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(useCase usecase.UserUseCase) UserController {
	return &userController{
		userUseCase: useCase,
	}
}

func (t *userController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "")
}

func (t *userController) Login(c *gin.Context) {
	url := auth.GetGoogleOAuthUrl()
	c.Redirect(http.StatusTemporaryRedirect, url)
}


func (t *userController) GoogleCallback(c *gin.Context) {
	content, err := auth.GetUserInfo(c.Query("state"), c.Query("code"))

	log.Printf("Content: %s\n", content)

	t.userUseCase.LoginUser(content)

	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

}