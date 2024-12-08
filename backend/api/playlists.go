package api

import (
	"errors"
	"fmt"
	"net/http"
	db "ourplaylist/db/sqlc"
	"ourplaylist/token"

	"github.com/gin-gonic/gin"
)

type addPlaylistReq struct {
	Username   string `json:"username" binding:"required"`
	Link       string `json:"link" binding:"required"`
	Desription string `json:"description" binding:"required"`
}

type addPlaylistRes struct {
	Username string      `json:"username"`
	Playlist db.Playlist `json:"playlist"`
}

func (server *Server) addPlaylist(ctx *gin.Context) {
	var req addPlaylistReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.Username != req.Username {
		err := errors.New("account doesnt belong to the user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	var res addPlaylistRes

	playlist, err := server.store.CreatePlaylist(ctx, db.CreatePlaylistParams{
		UserID:      user.ID,
		Link:        req.Link,
		Description: req.Desription,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	res.Playlist = playlist
	res.Username = user.Username

	ctx.JSON(http.StatusOK, res)

}

// type getPlaylistReq struct {
// 	Username string `json:"username" binding:"requierd"`
// }

type getPlaylistRes struct {
	Username       string        `json:"username"`
	UserPlaylists  []db.Playlist `json:"user_playlists"`
	LikedPlaylists []db.Playlist `json:"liked_playlists"`
}

func (server *Server) getPlaylists(ctx *gin.Context) {
	// var req getPlaylistReq

	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	// 	return
	// }

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	// if authPayload.Username != req.Username {
	// 	err := errors.New("account doesnt belong to the user")
	// 	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	// 	return
	// }

	user, err := server.store.GetUser(ctx, authPayload.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	var res getPlaylistRes

	res.Username = user.Username

	res.UserPlaylists, err = server.store.GetPlaylistsByUserID(ctx, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	res.LikedPlaylists, err = server.store.GetLikedPlaylistsByUser(ctx, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, res)

}

type getProfileReq struct {
	Username string `json:"username" binding:"required"`
}

type getProfileRes struct {
	Username  string        `json:"username"`
	Playlists []db.Playlist `json:"user_playlists"`
}

func (server *Server) getProfile(ctx *gin.Context) {
	var req getProfileReq
	req.Username = ctx.Query("username")
	fmt.Println()
	if req.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username query parameter is required"})
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	var res getProfileRes

	res.Username = user.Username

	res.Playlists, err = server.store.GetPlaylistsByUserID(ctx, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, res)

}

type likePlaylistReq struct {
	Username string `json:"username" binding:"required"`
	Playlist int32  `json:"playlist_id" binding:"required"`
}

type likePlaylistRes struct {
	Username string      `json:"username"`
	Playlist db.Playlist `json:"playlist"`
}

func (server *Server) likePlaylist(ctx *gin.Context) {
	var req likePlaylistReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.Username != req.Username {
		err := errors.New("account doesnt belong to the user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	var res likePlaylistRes
	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	fmt.Println("tu dochodzi")
	res.Username = user.Username

	playlist_res, err := server.store.LikePlaylistTx(ctx, db.LikePlaylistTxParams{
		UserID:     user.ID,
		PlaylistID: req.Playlist,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	fmt.Println("tu dochodzi")

	res.Playlist = playlist_res.Playlist

	ctx.JSON(http.StatusOK, res)

}
