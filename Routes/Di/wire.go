//go:build wireinject
// +build wireinject

package Di

import (
	"2024_akutansi_project/Controllers"
	"2024_akutansi_project/Middleware"
	"2024_akutansi_project/Repositories"
	"2024_akutansi_project/Services"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func DIAuth(db *gorm.DB) *Controllers.AuthController {
	panic(wire.Build(wire.NewSet(
		Repositories.AuthRepositoryProvider,
		Services.AuthServiceProvider,
		Controllers.AuthControllerProvider,
		Services.JwtServiceProvider,

		wire.Bind(new(Controllers.IAuthController), new(*Controllers.AuthController)),
		wire.Bind(new(Services.IAuthService), new(*Services.AuthService)),
		wire.Bind(new(Repositories.IAuthRepository), new(*Repositories.AuthRepository)),
		wire.Bind(new(Services.IJwtService), new(*Services.JwtService)),
	),
	))

	return &Controllers.AuthController{}
}

func DICommonMiddleware(db *gorm.DB) *Middleware.CommondMiddleware {
	panic(wire.Build(wire.NewSet(
		Middleware.CommonMiddlewareProvider,
		Services.JwtServiceProvider,
		Repositories.AuthRepositoryProvider,

		wire.Bind(new(Services.IJwtService), new(*Services.JwtService)),
		wire.Bind(new(Repositories.IAuthRepository), new(*Repositories.AuthRepository)),
		wire.Bind(new(Middleware.ICommonMiddleware), new(*Middleware.CommondMiddleware)),
	),
	))

	return &Middleware.CommondMiddleware{}
}

func DICompany(db *gorm.DB) *Controllers.CompanyController {
	panic(wire.Build(wire.NewSet(
		Repositories.CompanyRepositoryProvider,
		Services.CompanyServiceProvider,
		Controllers.CompanyControllerProvider,
		Repositories.UserCompanyRepositoryProvider,

		wire.Bind(new(Controllers.ICompanyController), new(*Controllers.CompanyController)),
		wire.Bind(new(Services.ICompanyService), new(*Services.CompanyService)),
		wire.Bind(new(Repositories.ICompanyRepository), new(*Repositories.CompanyRepository)),
		wire.Bind(new(Repositories.IUserCompanyRepository), new(*Repositories.UserCompanyRepository)),
	),
	))

	return &Controllers.CompanyController{}
}
