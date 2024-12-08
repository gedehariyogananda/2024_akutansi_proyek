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
	ISubUserService interface {
		Create(subUserDto *Dto.CreateSubUserDto) (res *Response.SubUser, err error)
		Update(subUserDto *Dto.UpdateSubUserDto, id string) (res *Response.SubUser, statusCode int, err error)
		Delete(id string) (statusCode int, err error)
		FindByID(id string) (res *Response.SubUser, statusCode int, err error)
		FindAll(companyId string, dto *Common.Query) (res []*Response.SubUser, meta Common.Meta, err error)
	}
	SubUserService struct {
		SubUserRepository Repositories.ISubUserRepository
	}
)

func SubUserProvider(subUserRepository Repositories.ISubUserRepository) *SubUserService {
	return &SubUserService{SubUserRepository: subUserRepository}
}

func (s *SubUserService) Create(subUserDto *Dto.CreateSubUserDto) (res *Response.SubUser, err error) {
	subUser := &Models.SubUser{
		Name:        subUserDto.Name,
		EmployeeKey: subUserDto.EmployeeKey,
		Status:      subUserDto.Status,
		CompanyID:   subUserDto.CompanyID,
	}

	subUser, err = s.SubUserRepository.Create(subUser)
	if err != nil {
		return nil, err
	}

	return Response.ToSubUser(subUser), nil
}

func (s *SubUserService) FindByID(id string) (res *Response.SubUser, statusCode int, err error) {
	subUser, err := s.SubUserRepository.FindByID(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, http.StatusNotFound, err
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	res = Response.ToSubUser(subUser)

	return res, http.StatusOK, nil
}

func (s *SubUserService) Update(subUserDto *Dto.UpdateSubUserDto, id string) (res *Response.SubUser, statusCode int, err error) {
	subUser, err := s.SubUserRepository.FindByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, http.StatusNotFound, err
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	subUser = &Models.SubUser{
		Name:        subUserDto.Name,
		EmployeeKey: subUserDto.EmployeeKey,
		Status:      subUserDto.Status,
		CompanyID:   subUserDto.CompanyID,
		Password:    subUserDto.Password,
	}

	subUser, err = s.SubUserRepository.Update(subUser, id)

	subUser.ID = id
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	res = Response.ToSubUser(subUser)

	return res, http.StatusOK, nil

}

func (s *SubUserService) Delete(id string) (statusCode int, err error) {
	_, err = s.SubUserRepository.FindByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound, err
	}

	err = s.SubUserRepository.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *SubUserService) FindAll(companyId string, dto *Common.Query) (res []*Response.SubUser, meta Common.Meta, err error) {
	subUsers, totalData, err := s.SubUserRepository.FindAll(companyId, dto)
	if err != nil {
		return nil, Common.Meta{}, err
	}

	for _, subUser := range subUsers {
		res = append(res, Response.ToSubUser(subUser))
	}

	meta = Common.Meta{
		Limit:     dto.Limit,
		Page:      dto.Page,
		TotalData: totalData,
	}
	return res, meta, nil
}
