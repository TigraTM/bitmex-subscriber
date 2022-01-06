package http

import (
	usersApp "bitmex-subscriber/internal/users/app"
	usersWeb "bitmex-subscriber/internal/users/web"
	"bitmex-subscriber/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Service struct {
	log  *zap.SugaredLogger
	conn *sqlx.DB
}

func New(log *zap.SugaredLogger, conn *sqlx.DB) *Service {
	return &Service{
		log:  log,
		conn: conn,
	}
}

func (h *Service) Init() *gin.Engine {
	router := gin.New()

	repo := db.DB{Conn: h.conn}

	usersApp := usersApp.New(repo)

	usersHandler := usersWeb.New(usersApp)

	apiV1 := router.Group("/api/v1")

	{
		usersHandler.Init(apiV1)
	}

	return router
}
