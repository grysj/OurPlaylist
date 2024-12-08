package api

import (
	"database/sql"
	"net/http"
	db "ourplaylist/db/sqlc"
	"ourplaylist/util"
	"time"

	"github.com/gin-gonic/gin"
)

type createUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Mail     string `json:"mail" binding:"required"`
}

type userRes struct {
	Username  string    `json:"username"`
	Mail      string    `json:"mail"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userRes {
	return userRes{
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		Mail:      user.Mail,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashed, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	arg := db.CreateUserParams{
		Username: req.Username,
		Password: hashed,
		Mail:     req.Mail,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	res := newUserResponse(user)

	ctx.JSON(http.StatusOK, res)

}

type loginUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginUserRes struct {
	AccessToken string  `json:"Token"`
	User        userRes `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	accessToken, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := loginUserRes{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, res)

}
