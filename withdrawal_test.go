package kucoin

import (
	"testing"
)

func TestApiService_Withdrawals(t *testing.T) {
	s := NewApiServiceFromEnv()
	p := &PaginationParam{CurrentPage: 1, PageSize: 10}
	rsp, err := s.Withdrawals("", "", 0, 0, p)
	if err != nil {
		t.Fatal(err)
	}
	ws := WithdrawalsModel{}
	if _, err := rsp.ReadPaginationData(&ws); err != nil {
		t.Fatal(err)
	}

	for _, w := range ws {
		t.Log(ToJsonString(w))
		switch {
		case w.Id == "":
			t.Error("Empty key 'id'")
		case w.Address == "":
			t.Error("Empty key 'address'")
		case w.Currency == "":
			t.Error("Empty key 'currency'")
		case w.Amount == "":
			t.Error("Empty key 'amount'")
		case w.Fee == "":
			t.Error("Empty key 'fee'")
		case w.Status == "":
			t.Error("Empty key 'status'")
		case w.CreatedAt == 0:
			t.Error("Empty key 'createdAt'")
		case w.UpdatedAt == 0:
			t.Error("Empty key 'updatedAt'")
		}
	}
}

func TestApiService_WithdrawalQuotas(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.WithdrawalQuotas("BTC")
	if err != nil {
		t.Fatal(err)
	}
	wq := &WithdrawalQuotasModel{}
	if err := rsp.ReadData(wq); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(wq))
	switch {
	case wq.Currency == "":
		t.Error("Empty key 'currency'")
	case wq.AvailableAmount == "":
		t.Error("Empty key 'availableAmount'")
	case wq.RemainAmount == "":
		t.Error("Empty key 'remainAmount'")
	case wq.WithdrawMinSize == "":
		t.Error("Empty key 'withdrawMinSize'")
	case wq.LimitBTCAmount == "":
		t.Error("Empty key 'limitBTCAmount'")
	case wq.InnerWithdrawMinFee == "":
		t.Error("Empty key 'innerWithdrawMinFee'")
	case wq.WithdrawMinFee == "":
		t.Error("Empty key 'withdrawMinFee'")
	case wq.Precision == 0:
		t.Error("Empty key 'precision'")
	}
}

func TestApiService_ApplyWithdrawal(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.ApplyWithdrawal("BTC", "xx", "0.01", map[string]string{})
	if err != nil {
		t.Fatal(err)
	}
	w := &ApplyWithdrawalResultModel{}
	if err := rsp.ReadData(w); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(w))
	switch {
	case w.WithdrawalId == "":
		t.Error("Empty key 'withdrawalId'")
	}
}

func TestApiService_CancelWithdrawal(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.CancelWithdrawal("xxx")
	if err != nil {
		t.Fatal(err)
	}
	w := &CancelWithdrawalResultModel{}
	if err := rsp.ReadData(w); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(w))
	switch {
	case len(w.CancelledWithdrawIds) == 0:
		t.Error("Empty key 'cancelledWithdrawIds'")
	}
}
