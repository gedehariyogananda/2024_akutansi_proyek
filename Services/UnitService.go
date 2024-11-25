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
	IUnitService interface {
		Create(dto *Dto.CreateUnitDto) (res *Response.Unit, err error)
		Update(dto *Dto.UpdateUnitDto, id string) (res *Response.Unit, statusCode int, err error)
		Delete(id string) (err error, statusCode int)
		FindAll(companyID string, query *Common.Query) (res []*Response.Unit, meta Common.Meta, err error)
		FindByID(id string) (res *Response.Unit, statusCode int, err error)
	}

	UnitService struct {
		unitRepository Repositories.IUnitRepository
	}
)

func UnitProvider(unitRepository Repositories.IUnitRepository) *UnitService {
	return &UnitService{unitRepository: unitRepository}
}

func (c *UnitService) FindByID(id string) (res *Response.Unit, statusCode int, err error) {
	unit, err := c.unitRepository.FindById(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, http.StatusBadRequest, err
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	res = Response.ToUnit(unit)

	return res, http.StatusOK, nil
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

	res = Response.ToUnit(unit)

	return res, nil
}

func (s *UnitService) Update(dto *Dto.UpdateUnitDto, id string) (res *Response.Unit, statusCode int, err error) {
	_, err = s.unitRepository.FindById(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, http.StatusBadRequest, err
	}
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	unit := &Models.Unit{
		Name:      dto.Name,
		CompanyID: dto.CompanyID,
		Status:    dto.Status,
		Code:      dto.Code,
	}

	unit, err = s.unitRepository.Update(unit, id)

	unit.ID = id

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	res = Response.ToUnit(unit)

	return res, http.StatusOK, nil
}

func (s *UnitService) Delete(id string) (err error, statusCode int) {

	_, err = s.unitRepository.FindById(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err, http.StatusBadRequest
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

func (s *UnitService) FindAll(companyID string, query *Common.Query) (res []*Response.Unit, meta Common.Meta, err error) {
	var units []*Models.Unit

	units, total, err := s.unitRepository.FindAll(companyID, query)

	if err != nil {
		return nil, Common.Meta{}, err
	}

	res = Response.ToUnitSlice(units)

	meta = Common.Meta{
		TotalData: int64(total),
		Limit:     query.Limit,
		Page:      query.Page,
	}

	return res, meta, nil
}
