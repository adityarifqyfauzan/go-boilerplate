// handler or presentation layer implementation
package user

import "github.com/gin-gonic/gin"

type handler struct {
}

func NewHandler() handler {
	return handler{}
}

func (h handler) Me(r *gin.Context) {

}

func (h handler) Login(r *gin.Context) {

}

func (h handler) Register(r *gin.Context) {

}

func (h handler) Logout(r *gin.Context) {

}

func (h handler) RefreshToken(r *gin.Context) {

}

func (h handler) ChangePassword(r *gin.Context) {

}

func (h handler) ForgotPassword(r *gin.Context) {

}
