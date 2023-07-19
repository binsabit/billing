package app

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gitlab.com/binsabit/billing/config"
	"gitlab.com/binsabit/billing/database"
)

type Server struct {
	config  config.Config
	router  *fiber.App
	storage *database.Storage
}

func NewServer(cfg config.Config, storage *database.Storage) (*Server, error) {
	server := &Server{
		config:  cfg,
		storage: storage,
	}

	server.setupRoutes()

	return server, nil
}

func (s *Server) setupRoutes() {
	router := fiber.New(fiber.Config{
		JSONDecoder: sonic.Unmarshal,
		JSONEncoder: sonic.Marshal,
	})
	router.Use(cors.New())
	s.CompanyTransactions(router)

	s.router = router
}

func (s *Server) Start(addr string) error {
	return s.router.Listen(addr)
}
