package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
)

type (
	ISubUserService interface {
		Create(subUserDto *Dto.CreateSubUserDto) (res *Response.SubUser, err error)
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
