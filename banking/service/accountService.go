package service

import (
	"time"

	"github.com/manmohansharma21/bankpoc/banking-lib/errs"
	"github.com/manmohansharma21/bankpoc/banking/domain"
	"github.com/manmohansharma21/bankpoc/banking/dto"
)

const dbTSLayout = "2006-01-02 15:04:05"

// Primary port interface
type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
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

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	// incoming request validation
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	// server side validation for checking the available balance in the account
	if req.IsTransactionTypeWithdrawal() {
		account, err := s.repo.FindById(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}
	// if all is well, build the domain object & save the transaction
	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}
	transaction, appError := s.repo.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}
	response := transaction.ToDto()
	return &response, nil
}

// Helper function to create NewAccountDefaultService
func NewAccountService(repo domain.AccountRepostory) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
