package enums

// providers
const (
	Kaspi          = 2 //webhook
	Qiwi           = 3 //webhook
	CompanyBill    = 4
	KaspiInvoice   = 5 //chet na oplatu
	HalykBank      = 6
	CompanyBalance = 7
	ProskladAdmin  = 10
)

// services type
const (
	Prosklad = iota + 1
	ExtraPoint
	Integration
	Teaching
)

// /transaction status
const (
	TransactionNotPaid = iota
	TransactionPaid
	TransactionRefund
	TransactionPaymentError
)

// company_carts status
const (
	CartCreated = iota + 1
	CartPaid
	CartCanceled
)

// transaction type
const (
	Deposit = iota + 1
	Withdrawal
	Transfer
	Reset
	Migrate
)
