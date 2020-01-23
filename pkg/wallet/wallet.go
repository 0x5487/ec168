package wallet

import (
	"time"
)

type Wallet struct {
	ID         int64
	MerchantID int32
	MemberID   int32 `json:"member_id" db:"member_id"`
	Member     string
	Currency   int32 `json:"currency"`
	Amount     int64 `json:"amount"`
	Note       string
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type TransactionType int32

var (
	WAGER       TransactionType = 1
	PAYOUT      TransactionType = 2
	REFUND      TransactionType = 3
	DEDUCT      TransactionType = 4
	WITHDRAW    TransactionType = 5
	DEPOSIT     TransactionType = 6
	TRANSFERIN  TransactionType = 7
	TRANSFEROUT TransactionType = 8
)

type Transaction struct {
	ID           int64
	WalletID     int64
	TransUUID    string
	MerchantID   int32
	MemberID     int32
	Member       string
	Type         TransactionType
	Currency     int32
	BeforeAmount int64
	Withdraw     int64
	Deposit      int64
	AfterAmount  int64
	Delta        int64
	Summary      string
	Note         string
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type Servicer interface {
	CreateWallet(memberID int32, currency int32) error
	Withdraw(memberID int32, currency int32, amount int64, allowNegative bool, note string) error
	Deposit(memberID int32, currency int32, amount int64, note string) error
}

type TransactionRepository interface {
	InsertTX() error
}
