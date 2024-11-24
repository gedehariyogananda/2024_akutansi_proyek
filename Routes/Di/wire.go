//go:build wireinject
// +build wireinject

package Di

import (
	"2024_akutansi_project/Controllers"
	"2024_akutansi_project/Middleware"
	"2024_akutansi_project/Repositories"
	"2024_akutansi_project/Services"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func DIAuth(db *gorm.DB, redis *redis.Client) *Controllers.AuthController {
	panic(wire.Build(wire.NewSet(
		Repositories.UserRepositoryProvider,
		Services.AuthServiceProvider,
		Controllers.AuthControllerProvider,
		Services.JwtServiceProvider,
		Repositories.SubUserRepositoryProvider,
		Repositories.CompanyRepositoryProvider,

		wire.Bind(new(Controllers.IAuthController), new(*Controllers.AuthController)),
		wire.Bind(new(Repositories.ISubUserRepository), new(*Repositories.SubUserRepository)),
		wire.Bind(new(Services.IAuthService), new(*Services.AuthService)),
		wire.Bind(new(Repositories.IUserRepository), new(*Repositories.UserRepository)),
		wire.Bind(new(Services.IJwtService), new(*Services.JwtService)),
		wire.Bind(new(Repositories.ICompanyRepository), new(*Repositories.CompanyRepository)),
	),
	))

	return &Controllers.AuthController{}
}

func DICommonMiddleware(db *gorm.DB, redis *redis.Client) *Middleware.CommondMiddleware {
	panic(wire.Build(wire.NewSet(
		Middleware.CommonMiddlewareProvider,
		Services.JwtServiceProvider,
		Repositories.UserRepositoryProvider,

		wire.Bind(new(Services.IJwtService), new(*Services.JwtService)),
		wire.Bind(new(Repositories.IUserRepository), new(*Repositories.UserRepository)),
		wire.Bind(new(Middleware.ICommonMiddleware), new(*Middleware.CommondMiddleware)),
	),
	))

	return &Middleware.CommondMiddleware{}
}

func DIInvoice(db *gorm.DB) *Controllers.InvoiceController {
	panic(wire.Build(wire.NewSet(
		Repositories.InvoiceRepositoryProvider,
		Services.InvoiceServiceProvider,
		Controllers.InvoiceControllerProvider,
		Repositories.InvoiceItemRepositoryProvider,
		Repositories.SellableProductRepositoryProvider,
		Repositories.ReceiptRepositoryProvider,
		Repositories.MaterialProductRepositoryProvider,

		wire.Bind(new(Controllers.IInvoiceController), new(*Controllers.InvoiceController)),
		wire.Bind(new(Services.IInvoiceService), new(*Services.InvoiceService)),
		wire.Bind(new(Repositories.IInvoiceRepository), new(*Repositories.InvoiceRepository)),
		wire.Bind(new(Repositories.IInvoiceItemRepository), new(*Repositories.InvoiceItemRepository)),
		wire.Bind(new(Repositories.ISellableProductRepository), new(*Repositories.SellableProductRepository)),
		wire.Bind(new(Repositories.IReceiptRepository), new(*Repositories.ReceiptRepository)),
		wire.Bind(new(Repositories.IMaterialProductRepository), new(*Repositories.MaterialProductRepository)),
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
func DIWaitingList(db *gorm.DB) *Controllers.WaitingListController {
	panic(wire.Build(wire.NewSet(
		Repositories.WaitingListRepositoryProvider,
		Services.WaitingListServiceProvider,
		Controllers.WaitingListControllerProvider,

		wire.Bind(new(Controllers.IWaitingListController), new(*Controllers.WaitingListController)),
		wire.Bind(new(Services.IWaitingListService), new(*Services.WaitingListService)),
		wire.Bind(new(Repositories.IWaitingListRepository), new(*Repositories.WaitingListRepository)),
	),
	))

	return &Controllers.WaitingListController{}
}

func DIUnit(db *gorm.DB) *Controllers.UnitController {
	panic(wire.Build(wire.NewSet(
		Repositories.UnitProvider,
		Services.UnitProvider,
		Controllers.UnitProvider,

		wire.Bind(new(Controllers.IUnitController), new(*Controllers.UnitController)),
		wire.Bind(new(Services.IUnitService), new(*Services.UnitService)),
		wire.Bind(new(Repositories.IUnitRepository), new(*Repositories.UnitRepository)),
	),
	))

	return &Controllers.UnitController{}
}
