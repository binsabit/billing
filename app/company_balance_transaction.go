package app

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/binsabit/billing/db"
	"log"
)

func (s Server) TransactionDepositHalyk(ctx *fiber.Ctx) error {

	reqData := db.HalykPayTransaction{}

	if err := ctx.BodyParser(&reqData); err != nil {
		log.Printf("wrong request data: %v", err)

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	id, err := s.storage.CreateTransaction(ctx.Context(), reqData)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"invoice_id": id,
	})
}

func (s Server) TransactionDepositVerifyHalyk(ctx *fiber.Ctx) error {
	return nil
}
