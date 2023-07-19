package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"gitlab.com/binsabit/billing/enums"
	"gitlab.com/binsabit/billing/models"
	"time"
)

type HalykPayTransaction struct {
	CompanyID  int64           `json:"company_id"`
	SenderType int16           `json:"sender_type"`
	SenderID   int64           `json:"sender_id"`
	Amount     decimal.Decimal `json:"amount"`
}

type HalykConfirmTransaction struct {
	InvoiceID int64           `json:"invoice_id"`
	Code      int             `json:"code"`
	Amount    decimal.Decimal `json:"amount"`
	Response  string          `json:"response"`
}

func (s Storage) tx(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = fn(tx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (s Storage) CreateTransaction(ctx context.Context, data *HalykPayTransaction) (int64, error) {

	//init, err := s.IsInit(ctx, data.CompanyID)
	//if err != nil {
	//	return -1, err
	//}

	tx := &models.CompanyBalanceTransaction{
		CompanyID:  data.CompanyID,
		SenderType: data.SenderType,
		SenderID: null.Int64{
			Int64: data.SenderID,
			//check if the proveded method payment has user id
			Valid: data.SenderType != 2 || data.SenderType != 6 || data.SenderType != 3,
		},
		Amount: data.Amount,
		Status: enums.TransactionNotPaid,
		IsInit: false,
		Type:   enums.Deposit,
	}
	err := tx.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return -1, err
	}
	return tx.ID, nil
}

func (s Storage) ConfirmTransaction(ctx context.Context, data HalykConfirmTransaction) error {
	err := s.tx(ctx, func(tx *sql.Tx) error {
		transaction, err := models.FindCompanyBalanceTransaction(ctx, s.db, data.InvoiceID)
		if err != nil {
			return err
		}
		if transaction == nil {
			return sql.ErrNoRows
		}

		if data.Amount != transaction.Amount {
			return fmt.Errorf("amount of transaction is diffrent")
		}

		transaction.Status = enums.TransactionPaid
		transaction.PaidAt = time.Now()

		_, err = transaction.Update(ctx, tx, boil.Infer())

		return err
	})

	return err
}

//func (s Storage) IsInit(ctx context.Context, companyID int64) (bool, error) {
//	row, err := models.CompanyBalanceTransactions(qm.Where("comapany_id=?", companyID)).One(ctx, s.db)
//	if err != nil {
//		return true, err
//	}
//	if row != nil {
//		return false, nil
//	}
//	return true, nil
//}
