package api

import (
	db "ourplaylist/db/sqlc"
	"ourplaylist/token"
	"ourplaylist/util"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
	config     util.Config
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.SecretKey)
	if err != nil {
		return nil, err
	}
	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Replace with your frontend's URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(corsConfig))
	authRoute := router.Group("/").Use(authMiddleware(server.tokenMaker))
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	authRoute.POST("/playlist/add", server.addPlaylist)
	authRoute.GET("/playlists", server.getPlaylists)
	authRoute.GET("/profile", server.getProfile)
	authRoute.POST("playlist/like", server.likePlaylist)

	server.router = router
	return server, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
