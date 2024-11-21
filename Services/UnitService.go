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
	IUnitService interface {
		Create(dto *Dto.CreateUnitDto) (res *Response.Unit, err error)
		Update(dto *Dto.UpdateUnitDto, id string) (res *Response.Unit, err error, statusCode int)
		Delete(id string) (err error, statusCode int)
	}

	UnitService struct {
		unitRepository Repositories.IUnitRepository
	}
)

func UnitProvider(unitRepository Repositories.IUnitRepository) *UnitService {
	return &UnitService{unitRepository: unitRepository}
}

func (s *UnitService) Create(dto *Dto.CreateUnitDto) (res *Response.Unit, err error) {
	unit := &Models.Unit{
		Name:      dto.Name,
		CompanyID: dto.CompanyID,
		Status:    dto.Status,
		Code:      dto.Code,
	}

	unit, err = s.unitRepository.Create(unit)
	if err != nil {
		return nil, err
	}

	res = Response.ToCompanyResponse(unit)

	return res, nil
}

func (s *UnitService) Update(dto *Dto.UpdateUnitDto, id string) (res *Response.Unit, err error, statusCode int) {
	_, err = s.unitRepository.FindById(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, http.StatusBadRequest
	}
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	unit := &Models.Unit{
		Name:      dto.Name,
		CompanyID: dto.CompanyID,
		Status:    dto.Status,
		Code:      dto.Code,
	}

	unit, err = s.unitRepository.Update(unit, id)

	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	res = Response.ToCompanyResponse(unit)

	return res, nil, http.StatusOK
}

func (s *UnitService) Delete(id string) (err error, statusCode int) {

	_, err = s.unitRepository.FindById(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, http.StatusBadRequest
	}
	if err != nil {
		return err, http.StatusInternalServerError
	}

	err = s.unitRepository.Delete(id)

	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}
