package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type (
	IAccountService interface {
		FindByID(id string) (res *Response.Account, statusCode int, err error)
		Create(dto *Dto.CreateAccountDto) (res *Response.Account, err error)
	}

	AccountService struct {
		AccountRepository Repositories.IAccountRepository
	}
)

func AccountProvider(accountRepository Repositories.IAccountRepository) *AccountService {
	return &AccountService{AccountRepository: accountRepository}
}

func (s *AccountService) FindByID(id string) (res *Response.Account, statusCode int, err error) {
	account, err := s.AccountRepository.FindByID(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, http.StatusNotFound, err
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	res = Response.ToAccount(account)

	return res, http.StatusOK, nil
}

func (s *AccountService) Create(dto *Dto.CreateAccountDto) (res *Response.Account, err error) {
	account := &Models.Account{
		Name:      dto.Name,
		Code:      dto.Code,
		Type:      dto.Type,
		CompanyID: dto.CompanyID,
		IsLocked:  dto.IsLocked,
		Status:    dto.Status,
	}

	account, err = s.AccountRepository.Create(account)
	if err != nil {
		return nil, err
	}

	return Response.ToAccount(account), nil
}
