package service

import (
	"time"

	"github.com/manmohansharma21/bankpoc/domain"
	"github.com/manmohansharma21/bankpoc/dto"
	"github.com/manmohansharma21/bankpoc/errs"
)

// Primary port interface
type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

// Implementation for this primary port
type DefaultAccountService struct {
	repo domain.AccountRepostory //Reference to secodary port
}

// newAccount implementation
func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate() // Responsibility of validating incoming request lies with the service layer.
	if err != nil {
		return nil, err
	}

	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"), //time.RFC3339, modified Format to save in database.
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponseDto()

	return &response, nil
}

// Helper function to create NewAccountDefaultService
func NewAccountService(repo domain.AccountRepostory) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
