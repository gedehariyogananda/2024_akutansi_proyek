package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
	"2024_akutansi_project/Utils"
	"net/http"
)

type (
	ITransactionService interface {
		GetAll(query *Common.Query) (data []*Models.Transaction, meta Common.Meta, statusCode int, err error)
		Store(request *Dto.CreateTransactionDto, companyCode string) (recordCode string, statusCode int, err error)
	}

	TransactionService struct {
		transactionRepository Repositories.ITransactionRepository
	}
)

func TransactionServiceProvider(transactionRepository Repositories.ITransactionRepository) *TransactionService {
	return &TransactionService{transactionRepository: transactionRepository}
}

func (s *TransactionService) GetAll(query *Common.Query) (data []*Models.Transaction, meta Common.Meta, statusCode int, err error) {
	data, totalData, err := s.transactionRepository.GetAllByCompany(query)
	if err != nil {
		return nil, Common.Meta{}, http.StatusInternalServerError, err
	}

	meta = Common.PaginateMetadata(nil, totalData, query.Limit, query.Page)

	return data, meta, http.StatusOK, nil
}

func (s *TransactionService) Store(request *Dto.CreateTransactionDto, companyCode string) (recordCode string, statusCode int, err error) {
	latestPrefFormat, err := s.transactionRepository.CountByDateCompanyID(request.CompanyID, request.Date)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	data, err := s.transactionRepository.Store(&Models.Transaction{
		Title:                 request.Title,
		Name:                  request.Name,
		AdditionalData:        request.AdditionalData,
		Date:                  Utils.ParseDateStringToDate(request.Date, nil),
		DueDate:               Utils.ParseDateStringToDate(*request.DueDate, nil),
		Note:                  request.Note,
		TransactionID:         request.TransactionID,
		PaymentMethod:         request.PaymentMethod,
		PaymentType:           request.PaymentType,
		TransactionRecordCode: Utils.GenerateTransactionRecord(latestPrefFormat, companyCode),
		Amount:                request.Amount,
		CompanyID:             request.CompanyID,
	})

	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return data.TransactionRecordCode, http.StatusOK, nil
}
