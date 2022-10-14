package user

import (
	"log"
	"net/http"

	"assignment2/pkg/domain/message"
	"assignment2/pkg/domain/user"

	"github.com/gin-gonic/gin"
)

type UserHdlImpl struct {
	userUsecase user.UserUsecase
}

func NewUserHandler(userUsecase user.UserUsecase) user.UserHandler {
	return &UserHdlImpl{userUsecase: userUsecase}
}

func (u *UserHdlImpl) GetUserByEmailHdl(ctx *gin.Context) {
	log.Printf("%T - GetUserByEmailHdl is invoked]\n", u)
	defer log.Printf("%T - GetUserByEmailHdl executed\n", u)

	email := ctx.Params.ByName("email")
	println(email)
	if email == ""{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "param can't be empty string"})
	}

	user, err := u.userUsecase.GetUserByEmailSvc(ctx, email)

	if err != nil {
		switch err.Error() {
		case "BAD_REQUEST":
			ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
				Code:  80,
				Error: "invalid processing payload",
			})
			return
		case "INTERNAL_SERVER_ERROR":
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, message.Response{
				Code:  99,
				Error: "something went wrong",
			})
			return
		case "NOT_FOUND":
			ctx.AbortWithStatusJSON(http.StatusOK, message.Response{
				Code:  99,
				Error: "user not found",
			})
			return
		}
	}
	// response result for the user if success
	ctx.JSONP(http.StatusOK, message.Response{
		Code:    0,
		Message: "success get user",
		Data:    user,
	})
}

// Insert New User
// @Summary this api will insert user with unique email
// @Schemes
// @Description insert new user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} User
// @Router /v1/user [post]
func (u *UserHdlImpl) InsertUserHdl(ctx *gin.Context) {
	// JSON: struktur data yang bisa dibaca secara manusiawi
	// dan digunakan secara masive untuk mengirimkan payload
	// dari client -> server atau sebaliknya
	// {"first_name":"Tara", "last_name":"Calman", "email":"calman@email.com"}
	// first_name, last_name, email -> json property
	// Tara, Calman, calman@email.com -> json property value
	// yang ingin dipecahkan oleh json, standardize payload around world
	// selain json: protobuf, form, csv

	log.Printf("%T - InsertUserHdl is invoked]\n", u)
	defer log.Printf("%T - InsertUserHdl executed\n", u)

	// binding / mendapatkan body payload dari request
	log.Println("binding body payload from request")
	var user user.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Code:  80,
			Error: "failed to bind payload",
		})
		return
	}
	// check apakah email kosong atau tidak: kalau kosong lempar BAD_REQUEST
	log.Println("check email from request")
	if user.Email == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Code:  80,
			Error: "email should not be empty",
		})
		return
	}
	// call service/usecase untuk menginsert data
	log.Println("calling insert service usecase")
	result, err := u.userUsecase.InsertUserSvc(ctx, user)
	if err != nil {
		switch err.Error() {
		case "BAD_REQUEST":
			ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
				Code:  80,
				Error: "invalid processing payload",
			})
			return
		case "INTERNAL_SERVER_ERROR":
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, message.Response{
				Code:  99,
				Error: "something went wrong",
			})
			return
		}
	}
	// response result for the user if success
	ctx.JSONP(http.StatusOK, message.Response{
		Code:    0,
		Message: "success insert user",
		Data:    result,
	})
}
