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
		Repositories.CompanyRepositoryProvider,

		wire.Bind(new(Controllers.IAuthController), new(*Controllers.AuthController)),
		wire.Bind(new(Services.IAuthService), new(*Services.AuthService)),
		wire.Bind(new(Repositories.IAuthRepository), new(*Repositories.AuthRepository)),
		wire.Bind(new(Services.IJwtService), new(*Services.JwtService)),
		wire.Bind(new(Repositories.ICompanyRepository), new(*Repositories.CompanyRepository)),
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
		Repositories.PaymentMethodRepositoryProvider,

		wire.Bind(new(Controllers.ICompanyController), new(*Controllers.CompanyController)),
		wire.Bind(new(Services.ICompanyService), new(*Services.CompanyService)),
		wire.Bind(new(Repositories.ICompanyRepository), new(*Repositories.CompanyRepository)),
		wire.Bind(new(Repositories.IUserCompanyRepository), new(*Repositories.UserCompanyRepository)),
		wire.Bind(new(Repositories.IPaymentMethodRepository), new(*Repositories.PaymentMethodRepository)),
	),
	))

	return &Controllers.CompanyController{}
}

func DISaleableProduct(db *gorm.DB) *Controllers.SaleableProductController {
	panic(wire.Build(wire.NewSet(
		Repositories.SaleableProductRepositoryProvider,
		Services.SaleableProductServiceProvider,
		Controllers.SaleableProductControllerProvider,
		Repositories.MaterialProductRepositoryProvider,
		Repositories.CategoryRepositoryProvider,

		wire.Bind(new(Controllers.ISaleableProductController), new(*Controllers.SaleableProductController)),
		wire.Bind(new(Services.ISaleableProductService), new(*Services.SaleableProductService)),
		wire.Bind(new(Repositories.IMaterialProductRepository), new(*Repositories.MaterialProductRepository)),
		wire.Bind(new(Repositories.ISaleableProductRepository), new(*Repositories.SaleableProductRepository)),
		wire.Bind(new(Repositories.ICategoryRepository), new(*Repositories.CategoryRepository)),
	),
	))

	return &Controllers.SaleableProductController{}
}

func DIInvoice(db *gorm.DB) *Controllers.InvoiceController {
	panic(wire.Build(wire.NewSet(
		Repositories.InvoiceRepositoryProvider,
		Services.InvoiceServiceProvider,
		Controllers.InvoiceControllerProvider,
		Repositories.InvoiceMaterialRepositoryProvider,
		Repositories.InvoiceSaleableRepositoryProvider,
		Repositories.SaleableProductRepositoryProvider,
		Repositories.PaymentMethodRepositoryProvider,
		Repositories.CompanyRepositoryProvider,

		wire.Bind(new(Controllers.IInvoiceController), new(*Controllers.InvoiceController)),
		wire.Bind(new(Services.IInvoiceService), new(*Services.InvoiceService)),
		wire.Bind(new(Repositories.IInvoiceRepository), new(*Repositories.InvoiceRepository)),
		wire.Bind(new(Repositories.IInvoiceMaterialRepository), new(*Repositories.InvoiceMaterialRepository)),
		wire.Bind(new(Repositories.IInvoiceSaleableRepository), new(*Repositories.InvoiceSaleableRepository)),
		wire.Bind(new(Repositories.ISaleableProductRepository), new(*Repositories.SaleableProductRepository)),
		wire.Bind(new(Repositories.IPaymentMethodRepository), new(*Repositories.PaymentMethodRepository)),
		wire.Bind(new(Repositories.ICompanyRepository), new(*Repositories.CompanyRepository)),
	),
	))

	return &Controllers.InvoiceController{}
}

func DICategory(db *gorm.DB) *Controllers.CategoryController {
	panic(wire.Build(wire.NewSet(
		Repositories.CategoryRepositoryProvider,
		Services.CategoryServiceProvider,
		Controllers.CategoryControllerProvider,

		wire.Bind(new(Controllers.ICategoryController), new(*Controllers.CategoryController)),
		wire.Bind(new(Services.ICategoryService), new(*Services.CategoryService)),
		wire.Bind(new(Repositories.ICategoryRepository), new(*Repositories.CategoryRepository)),
	),
	))

	return &Controllers.CategoryController{}
}

func DIPaymentMethod(db *gorm.DB) *Controllers.PaymentMethodController {
	panic(wire.Build(wire.NewSet(
		Repositories.PaymentMethodRepositoryProvider,
		Services.PaymentMethodServiceProvider,
		Controllers.PaymentMethodControllerProvider,

		wire.Bind(new(Controllers.IPaymentMethodController), new(*Controllers.PaymentMethodController)),
		wire.Bind(new(Services.IPaymentMethodService), new(*Services.PaymentMethodService)),
		wire.Bind(new(Repositories.IPaymentMethodRepository), new(*Repositories.PaymentMethodRepository)),
	),
	))

	return &Controllers.PaymentMethodController{}
}
