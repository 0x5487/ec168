package order

import (
	"time"

	"github.com/rotisserie/eris"
)

var (
	ErrOrderNotFound = eris.NewGlobal("order not found")
)

type Status int32

const (
	Created          Status = 1
	Paid             Status = 2
	Seattled         Status = 3
	PartitalSeattled Status = 4 // Parlay
	Failed           Status = 5
	Voided           Status = 6
)

type Type int32

const (
	Normal Type = 1
	Active Type = 2
)

// Order 代表訂單資訊
type Order struct {
	ID              string
	Type            Type   // 1: 正常訂單 2:活動
	OrderNO         string `json:"order_no" db:"order_no"`
	MerchantID      int32
	AgentID         int32
	MemberID        int32 `json:"member_id" db:"member_id"`
	Member          string
	GameID          int64
	GameCategoryID  int32
	Currency        int32 `json:"currency"`
	BeforeAmount    int64 `json:"before_amount" db:"before_amount"`
	NetAmount       int64 `json:"net_amount" db:"net_amount"`
	BetAmount       int64 `json:"bet_amount" db:"bet_amount"`
	JackpotAmount   int64
	ValidBetAmount  int64     `json:"valid_bet_amount" db:"valid_bet_amount"`
	PayoutAmount    int64     `json:"payout_amount" db:"payout_amount"`
	AfterAmount     int64     `json:"after_amount" db:"after_amount"`
	BetAt           time.Time `json:"bet_at" db:"bet_at"`
	Status          Status
	IsParley        bool
	ClientIP        string
	DeviceType      string
	Result          string
	Hash            string
	Note            string
	VendorUpdatedAt time.Time // 串接第三方使用
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	ForceUpdate     bool
}

type LineItemStatus int32

var (
	LineItemUnseattled LineItemStatus = 1
	LineItemSeattled   LineItemStatus = 2
	LineItemCancelled  LineItemStatus = 3
)

type LineItem struct {
	ID              string
	OrderID         string
	GameID          int64
	Round           string `json:"round_id" db:"round_id"`
	Market          string
	Outcome         string
	Odds            int32
	Status          LineItemStatus
	BetAmount       int64 `json:"bet_amount" db:"bet_amount"`
	JackpotAmount   int64
	ValidBetAmount  int64     `json:"valid_bet_amount" db:"valid_bet_amount"`
	PayoutAmount    int64     `json:"payout_amount" db:"payout_amount"`
	SeattledAt      time.Time `json:"seattle_at" db:"seattle_at"`
	SeattledVersion string
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

type FindOrderOptions struct {
	ID string
}

type Settlement struct {
	GameID          int64
	Round           string
	Market          string
	Result          string
	SeattledVersion string
}

// Servicer 代表訂單相關的業務邏輯
type Servicer interface {
	Order(orderUUID string) (Order, error)
	Orders(opts FindOrderOptions) ([]Order, error)
	CreateOrder(order *Order) error
	SeattledOrder(settlement Settlement) error
}

type Repository interface {
	Insert(order *Order) error
}
