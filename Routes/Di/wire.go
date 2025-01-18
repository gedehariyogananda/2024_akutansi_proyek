//go:build wireinject
// +build wireinject

package Di

import (
	"2024_akutansi_project/Connector"
	"2024_akutansi_project/Controllers"
	"2024_akutansi_project/Middleware"
	"2024_akutansi_project/Repositories"
	"2024_akutansi_project/Services"

	"firebase.google.com/go/messaging"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/google/wire"
	"github.com/minio/minio-go/v7"
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
		Repositories.AccountProvider,

		wire.Bind(new(Controllers.IAuthController), new(*Controllers.AuthController)),
		wire.Bind(new(Repositories.ISubUserRepository), new(*Repositories.SubUserRepository)),
		wire.Bind(new(Services.IAuthService), new(*Services.AuthService)),
		wire.Bind(new(Repositories.IUserRepository), new(*Repositories.UserRepository)),
		wire.Bind(new(Services.IJwtService), new(*Services.JwtService)),
		wire.Bind(new(Repositories.ICompanyRepository), new(*Repositories.CompanyRepository)),
		wire.Bind(new(Repositories.IAccountRepository), new(*Repositories.AccountRepository)),
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
		Repositories.SellableStockRepositoryProvider,
		Repositories.AccountProvider,
		Repositories.MaterialStockRepositoryProvider,
		Repositories.JournalEntriesProvider,
		Repositories.TransactionRepositoryProvider,
		Services.JournalEntriesProvider,
		Repositories.PurchaseRepositoryProvider,

		wire.Bind(new(Controllers.IInvoiceController), new(*Controllers.InvoiceController)),
		wire.Bind(new(Services.IInvoiceService), new(*Services.InvoiceService)),
		wire.Bind(new(Repositories.IInvoiceRepository), new(*Repositories.InvoiceRepository)),
		wire.Bind(new(Repositories.IInvoiceItemRepository), new(*Repositories.InvoiceItemRepository)),
		wire.Bind(new(Repositories.ISellableProductRepository), new(*Repositories.SellableProductRepository)),
		wire.Bind(new(Repositories.IReceiptRepository), new(*Repositories.ReceiptRepository)),
		wire.Bind(new(Repositories.ISellableStockRepository), new(*Repositories.SellableStockRepository)),
		wire.Bind(new(Repositories.IMaterialStockRepository), new(*Repositories.MaterialStockRepository)),
		wire.Bind(new(Repositories.IAccountRepository), new(*Repositories.AccountRepository)),
		wire.Bind(new(Repositories.IJournalEntriesRepository), new(*Repositories.JournalEntriesRepository)),
		wire.Bind(new(Repositories.IMaterialProductRepository), new(*Repositories.MaterialProductRepository)),
		wire.Bind(new(Services.IJournalEntriesService), new(*Services.JournalEntriesService)),
		wire.Bind(new(Repositories.ITransactionRepository), new(*Repositories.TransactionRepository)),
		wire.Bind(new(Repositories.IPurchaseRepository), new(*Repositories.PurchaseRepository)),
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

func DIProfile(db *gorm.DB, mongo *mongo.Client) *Controllers.ProfileController {
	panic(wire.Build(wire.NewSet(
		Repositories.ProfileRepositoryProvider,
		Services.ProfileServiceProvider,
		Controllers.ProfileControllerProvider,
		Connector.ShopeeConnectorProvider,

		wire.Bind(new(Controllers.IProfileController), new(*Controllers.ProfileController)),
		wire.Bind(new(Services.IProfileService), new(*Services.ProfileService)),
		wire.Bind(new(Repositories.IProfileRepository), new(*Repositories.ProfileRepository)),
		wire.Bind(new(Connector.IShopeeConnector), new(*Connector.ShopeeConnector)),
	),
	))

	return &Controllers.ProfileController{}
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

func DIWorker(db *gorm.DB, mongo *mongo.Client, messaging *messaging.Client) *Services.WorkerService {
	panic(wire.Build(wire.NewSet(
		Services.WorkerServiceProvider,
		Repositories.DeviceTokenRepositoryProvider,
		Repositories.NotificationRepositoryProvider,
		Repositories.SellableStockRepositoryProvider,
		Repositories.SellableProductRepositoryProvider,
		Repositories.MaterialStockRepositoryProvider,
		Repositories.MaterialProductRepositoryProvider,

		wire.Bind(new(Services.IWorkerService), new(*Services.WorkerService)),
		wire.Bind(new(Repositories.IDeviceTokenRepository), new(*Repositories.DeviceTokenRepository)),
		wire.Bind(new(Repositories.INotificationRepository), new(*Repositories.NotificationRepository)),
		wire.Bind(new(Repositories.ISellableStockRepository), new(*Repositories.SellableStockRepository)),
		wire.Bind(new(Repositories.ISellableProductRepository), new(*Repositories.SellableProductRepository)),
		wire.Bind(new(Repositories.IMaterialStockRepository), new(*Repositories.MaterialStockRepository)),
		wire.Bind(new(Repositories.IMaterialProductRepository), new(*Repositories.MaterialProductRepository)),
	),
	))

	return &Services.WorkerService{}
}

func DITax(db *gorm.DB) *Controllers.TaxController {
	panic(wire.Build(wire.NewSet(
		Repositories.TaxRepositoryProvider,
		Services.TaxServiceProvider,
		Controllers.TaxControllerProvider,

		wire.Bind(new(Controllers.ITaxController), new(*Controllers.TaxController)),
		wire.Bind(new(Services.ITaxService), new(*Services.TaxService)),
		wire.Bind(new(Repositories.ITaxRepository), new(*Repositories.TaxRepository)),
	),
	))

	return &Controllers.TaxController{}
}
func DISubUser(db *gorm.DB) *Controllers.SubUserController {
	panic(wire.Build(wire.NewSet(
		Repositories.SubUserRepositoryProvider,
		Services.SubUserProvider,
		Controllers.SubUserProvider,
		wire.Bind(new(Controllers.ISubUserController), new(*Controllers.SubUserController)),
		wire.Bind(new(Services.ISubUserService), new(*Services.SubUserService)),
		wire.Bind(new(Repositories.ISubUserRepository), new(*Repositories.SubUserRepository)),
	)))
	return &Controllers.SubUserController{}
}

func DIAccount(db *gorm.DB) *Controllers.AccountController {
	panic(wire.Build(wire.NewSet(
		Repositories.AccountProvider,
		Services.AccountProvider,
		Controllers.AccountProvider,
		wire.Bind(new(Controllers.IAccountController), new(*Controllers.AccountController)),
		wire.Bind(new(Services.IAccountService), new(*Services.AccountService)),
		wire.Bind(new(Repositories.IAccountRepository), new(*Repositories.AccountRepository)),
	)))
	return &Controllers.AccountController{}
}

func DISellableProduct(db *gorm.DB, minio *minio.Client) *Controllers.SellableProductController {
	panic(wire.Build(wire.NewSet(
		Repositories.SellableProductRepositoryProvider,
		Repositories.PromoItemRepositoryProvider,
		Repositories.ReceiptRepositoryProvider,
		Services.SellableProductServiceProvider,
		Services.StorageServiceProvider,
		Controllers.SellableProductControllerProvider,

		// Bind interfaces ke implementasinya
		wire.Bind(new(Repositories.ISellableProductRepository), new(*Repositories.SellableProductRepository)),
		wire.Bind(new(Repositories.IPromoItemRepository), new(*Repositories.PromoItemRepository)),
		wire.Bind(new(Repositories.IReceiptRepository), new(*Repositories.ReceiptRepository)),
		wire.Bind(new(Services.ISellableProductService), new(*Services.SellableProductService)),
		wire.Bind(new(Services.IStorageService), new(*Services.StorageService)),
		wire.Bind(new(Controllers.ISellableProductController), new(*Controllers.SellableProductController)),
	),
	))

	return &Controllers.SellableProductController{}
}

func DIMaterialProduct(db *gorm.DB) *Controllers.MaterialProductController {
	panic(wire.Build(wire.NewSet(
		Repositories.MaterialProductRepositoryProvider,
		Services.MaterialProductServiceProvider,
		Controllers.MaterialProductControllerProvider,
		Repositories.MaterialConversionRepositoryProvider,

		wire.Bind(new(Controllers.IMaterialProductController), new(*Controllers.MaterialProductController)),
		wire.Bind(new(Services.IMaterialProductService), new(*Services.MaterialProductService)),
		wire.Bind(new(Repositories.IMaterialProductRepository), new(*Repositories.MaterialProductRepository)),
		wire.Bind(new(Repositories.IMaterialConversionRepository), new(*Repositories.MaterialConversionRepository)),
	),
	))

	return &Controllers.MaterialProductController{}
}

func DIStockOpname(db *gorm.DB) *Controllers.StockOpnameController {
	panic(wire.Build(wire.NewSet(
		Repositories.StockOpnameRepositoryProvider,
		Repositories.SellableStockRepositoryProvider,
		Repositories.MaterialStockRepositoryProvider,
		Services.StockOpnameServiceProvider,
		Controllers.StockOpnameControllerProvider,

		wire.Bind(new(Controllers.IStockOpnameController), new(*Controllers.StockOpnameController)),
		wire.Bind(new(Services.IStockOpnameService), new(*Services.StockOpnameService)),
		wire.Bind(new(Repositories.IStockOpnameRepository), new(*Repositories.StockOpnameRepository)),
		wire.Bind(new(Repositories.ISellableStockRepository), new(*Repositories.SellableStockRepository)),
		wire.Bind(new(Repositories.IMaterialStockRepository), new(*Repositories.MaterialStockRepository)),
	),
	))

	return &Controllers.StockOpnameController{}
}

func DITransaction(db *gorm.DB) *Controllers.TransactionController {
	panic(wire.Build(wire.NewSet(
		Repositories.TransactionRepositoryProvider,
		Services.TransactionServiceProvider,
		Controllers.TransactionControllerProvider,

		wire.Bind(new(Controllers.ITransactionController), new(*Controllers.TransactionController)),
		wire.Bind(new(Services.ITransactionService), new(*Services.TransactionService)),
		wire.Bind(new(Repositories.ITransactionRepository), new(*Repositories.TransactionRepository)),
	),
	))

	return &Controllers.TransactionController{}
}

func DIJournalEntries(db *gorm.DB) *Controllers.JournalEntriesController {
	panic(wire.Build(wire.NewSet(
		Repositories.JournalEntriesProvider,
		Services.JournalEntriesProvider,
		Controllers.JournalEntriesProvider,
		Repositories.AccountProvider,
		Repositories.TransactionRepositoryProvider,
		Repositories.PurchaseRepositoryProvider,

		wire.Bind(new(Controllers.IJournalEntriesController), new(*Controllers.JournalEntriesController)),
		wire.Bind(new(Services.IJournalEntriesService), new(*Services.JournalEntriesService)),
		wire.Bind(new(Repositories.IJournalEntriesRepository), new(*Repositories.JournalEntriesRepository)),
		wire.Bind(new(Repositories.IAccountRepository), new(*Repositories.AccountRepository)),
		wire.Bind(new(Repositories.ITransactionRepository), new(*Repositories.TransactionRepository)),
		wire.Bind(new(Repositories.IPurchaseRepository), new(*Repositories.PurchaseRepository)),
	),
	))

	return &Controllers.JournalEntriesController{}
}
func DiPurchase(db *gorm.DB) *Controllers.PurchaseController {
	panic(wire.Build(wire.NewSet(
		Repositories.SellableProductRepositoryProvider,
		Repositories.MaterialProductRepositoryProvider,
		Repositories.PurchaseRepositoryProvider,
		Repositories.PurchaseSellableProductRepositoryProvider,
		Repositories.PurchaseMaterialProductRepositoryProvider,
		Repositories.MaterialStockRepositoryProvider, // Menambahkan provider untuk MaterialStockRepository
		Repositories.SellableStockRepositoryProvider, // Menambahkan provider untuk SellableStockRepository
		Services.PurchaseServiceProvider,
		Controllers.PurchaseControllerProvider,

		wire.Bind(new(Controllers.IPurchaseController), new(*Controllers.PurchaseController)),
		wire.Bind(new(Services.IPurchaseService), new(*Services.PurchaseService)),
		wire.Bind(new(Repositories.ISellableProductRepository), new(*Repositories.SellableProductRepository)),
		wire.Bind(new(Repositories.IMaterialProductRepository), new(*Repositories.MaterialProductRepository)),
		wire.Bind(new(Repositories.IPurchaseRepository), new(*Repositories.PurchaseRepository)),
		wire.Bind(new(Repositories.IPurchaseSellableProductRepository), new(*Repositories.PurchaseSellableProductRepository)),
		wire.Bind(new(Repositories.IPurchaseMaterialProductRepository), new(*Repositories.PurchaseMaterialProductRepository)),
		wire.Bind(new(Repositories.IMaterialStockRepository), new(*Repositories.MaterialStockRepository)),
		wire.Bind(new(Repositories.ISellableStockRepository), new(*Repositories.SellableStockRepository)),
	)))
}
