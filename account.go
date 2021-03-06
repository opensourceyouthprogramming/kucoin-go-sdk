package kucoin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// An AccountModel represents an account.
type AccountModel struct {
	Id        string `json:"id"`
	Currency  string `json:"currency"`
	Type      string `json:"type"`
	Balance   string `json:"balance"`
	Available string `json:"available"`
	Holds     string `json:"holds"`
}

// An AccountsModel is the set of *AccountModel.
type AccountsModel []*AccountModel

// Accounts returns a list of accounts.
// See the Deposits section for documentation on how to deposit funds to begin trading.
func (as *ApiService) Accounts(currency, typo string) (*ApiResponse, error) {
	p := map[string]string{}
	if currency != "" {
		p["currency"] = currency
	}
	if typo != "" {
		p["type"] = typo
	}
	req := NewRequest(http.MethodGet, "/api/v1/accounts", p)
	return as.Call(req)
}

// Account returns an account when you know the accountId.
func (as *ApiService) Account(accountId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/accounts/"+accountId, nil)
	return as.Call(req)
}

// CreateAccount creates an account according to type(main|trade) and currency
func (as *ApiService) CreateAccount(typo, currency string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/accounts", map[string]string{"currency": currency, "type": typo})
	return as.Call(req)
}

// An AccountHistoryModel represents account activity either increases or decreases your account balance.
type AccountHistoryModel struct {
	Currency  string          `json:"currency"`
	Amount    string          `json:"amount"`
	Fee       string          `json:"fee"`
	Balance   string          `json:"balance"`
	BizType   string          `json:"bizType"`
	Direction string          `json:"direction"`
	CreatedAt int64           `json:"createdAt"`
	Context   json.RawMessage `json:"context"`
}

// An AccountHistoriesModel the set of *AccountHistoryModel.
type AccountHistoriesModel []*AccountHistoryModel

// AccountHistories returns a list about account activity.
// Account activity either increases or decreases your account balance.
// Items are paginated and sorted latest first.
func (as *ApiService) AccountHistories(accountId string, startAt, endAt int64, pagination *PaginationParam) (*ApiResponse, error) {
	p := map[string]string{}
	if startAt > 0 {
		p["startAt"] = IntToString(startAt)
	}
	if endAt > 0 {
		p["endAt"] = IntToString(endAt)
	}
	pagination.ReadParam(p)
	req := NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/accounts/%s/ledgers", accountId), p)
	return as.Call(req)
}

// An AccountHoldModel represents the holds on an account for any active orders or pending withdraw requests.
// As an order is filled, the hold amount is updated.
// If an order is canceled, any remaining hold is removed.
// For a withdraw, once it is completed, the hold is removed.
type AccountHoldModel struct {
	Currency   string `json:"currency"`
	HoldAmount string `json:"holdAmount"`
	BizType    string `json:"bizType"`
	OrderId    string `json:"orderId"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}

// An AccountHoldsModel is the set of *AccountHoldModel.
type AccountHoldsModel []*AccountHoldModel

// AccountHolds returns a list of currency hold.
// Holds are placed on an account for any active orders or pending withdraw requests.
// As an order is filled, the hold amount is updated.
// If an order is canceled, any remaining hold is removed.
// For a withdraw, once it is completed, the hold is removed.
func (as *ApiService) AccountHolds(accountId string, pagination *PaginationParam) (*ApiResponse, error) {
	p := map[string]string{}
	pagination.ReadParam(p)
	req := NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/accounts/%s/holds", accountId), p)
	return as.Call(req)
}

// An InterTransferResultModel represents the result of a inner-transfer operation.
type InterTransferResultModel struct {
	OrderId string `json:"orderId"`
}

// InnerTransfer makes a currency transfer internally.
// The inner transfer interface is used for assets transfer among the accounts of a user and is free of charges on the platform.
// For example, a user could transfer assets for free form the main account to the trading account on the platform.
func (as *ApiService) InnerTransfer(clientOid, payAccountId, recAccountId, amount string) (*ApiResponse, error) {
	p := map[string]string{
		"clientOid":    clientOid,
		"payAccountId": payAccountId,
		"recAccountId": recAccountId,
		"amount":       amount,
	}
	req := NewRequest(http.MethodPost, "/api/v1/accounts/inner-transfer", p)
	return as.Call(req)
}
