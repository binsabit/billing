package db

import (
	"context"
	"database/sql"
	"github.com/volatiletech/null/v8"
	"gitlab.com/binsabit/billing/enums"
	"gitlab.com/binsabit/billing/sqlc/models"
)

type HalykPayTransaction struct {
	CompanyID  int64 `json:"company_id"`
	SenderType int32 `json:"sender_type"`
	SenderID   int64 `json:"sender_id"`
	Amount     int64 `json:"amount"`
}

type HalykConfirmTransaction struct {
	InvoiceID int64     `json:"invoice_id"`
	Code      int       `json:"code"`
	Amount    uint64    `json:"amount"`
	Response  null.JSON `json:"response"`
}

func (s *Store) CreateTransaction(ctx context.Context, transaction HalykPayTransaction) (int64, error) {
	var id int64
	err := s.execTx(ctx, func(q *models.Queries) error {
		result, err := q.CreateTransaction(ctx, models.CreateTransactionParams{
			CompanyID:  transaction.CompanyID,
			SenderType: transaction.SenderType,
			Amount:     transaction.Amount,
			SenderID: sql.NullInt64{
				Int64: transaction.SenderID,
				Valid: transaction.SenderType != enums.Kaspi || transaction.SenderType != enums.KaspiInvoice || transaction.SenderType != enums.Qiwi,
			},
			Status:     enums.TransactionNotPaid,
			ProviderID: enums.HalykBank,
			Type:       enums.Deposit,
		})

		if err != nil {
			return err
		}

		id, err = result.LastInsertId()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return -1, err
	}
	return id, err
}
