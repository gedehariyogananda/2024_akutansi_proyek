package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
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
		Update(dto *Dto.UpdateAccountDto, id string) (res *Response.Account, statusCode int, err error)
		Delete(id string) (statusCode int, err error)
		FindAll(companyID string, dto *Common.Query) (res []*Response.Account, meta Common.Meta, err error)
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
		Status:    dto.Status,
	}

	account, err = s.AccountRepository.Create(account)
	if err != nil {
		return nil, err
	}

	return Response.ToAccount(account), nil
}

func (s *AccountService) Update(dto *Dto.UpdateAccountDto, id string) (res *Response.Account, statusCode int, err error) {
	account, err := s.AccountRepository.FindByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, http.StatusNotFound, err
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	account = &Models.Account{
		Name:      dto.Name,
		Code:      dto.Code,
		Type:      dto.Type,
		CompanyID: dto.CompanyID,
	}

	account, err = s.AccountRepository.Update(account, id)
	account.ID = id
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return Response.ToAccount(account), http.StatusOK, nil
}

func (s *AccountService) Delete(id string) (statusCode int, err error) {
	_, err = s.AccountRepository.FindByID(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound, err
	}

	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = s.AccountRepository.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *AccountService) FindAll(companyID string, dto *Common.Query) (res []*Response.Account, meta Common.Meta, err error) {
	accounts, totalData, err := s.AccountRepository.FindAll(companyID, dto)

	if err != nil {
		return nil, meta, err
	}

	meta = Common.Meta{
		TotalData: totalData,
		Page:      dto.Page,
		Limit:     dto.Limit,
	}

	return Response.ToAccountSlice(accounts), meta, nil
}
