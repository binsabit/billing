package app

import "github.com/gofiber/fiber/v2"

func (s *Server) CompanyTransactions(app *fiber.App) {
	transactions := app.Group("api/v1/transactions")

	transactions.Post("/halyk/pay", s.TansactionDepositHalyk)
	//transactions.Post("/deposit/halyk/verify", s.TransactionDepositVerifyHalyk)
}
