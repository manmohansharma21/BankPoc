package domain

import (
	"github.com/manmohansharma21/bankpoc/dto"
	"github.com/manmohansharma21/bankpoc/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

// Secondary Port interfaces means repository
type AccountRepostory interface {
	Save(Account) (*Account, *errs.AppError)
}
