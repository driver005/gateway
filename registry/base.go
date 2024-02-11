package registry

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"time"

	"github.com/driver005/gateway/config"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/middlewares"
	"github.com/driver005/gateway/migrations"
	"github.com/driver005/gateway/repository"
	"github.com/driver005/gateway/routes"
	"github.com/driver005/gateway/routes/admin"
	"github.com/driver005/gateway/services"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/go-playground/validator/v10"
	"github.com/sarulabs/di"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	dbLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"

	// "github.com/gofiber/fiber/v3/middleware/csrf"
	"github.com/gofiber/fiber/v3/middleware/favicon"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Base struct {
	logger       *zap.SugaredLogger
	buildVersion string
	buildHash    string
	buildDate    string
	trc          trace.Tracer
	database     sql.Database
	container    di.Container
	config       *config.Config
	session      *session.Store
	middleware   *middlewares.Handler
	adminRouter  fiber.Router
	storeRouter  fiber.Router
	router       *fiber.App
	routes       *routes.Routes
	migrations   *migrations.Handler
	validator    *validator.Validate

	//Interfaces
	priceSelectionStrategy interfaces.IPriceSelectionStrategy
	taxCalculationStrategy interfaces.ITaxCalculationStrategy
	inventoryService       interfaces.IInventoryService
	stockLocationService   interfaces.IStockLocationService
	cacheService           interfaces.ICacheService
	pricingModuleService   interfaces.IPricingModuleService
	fileService            interfaces.IFileService
	batchJobStrategy       interfaces.IBatchJobStrategy

	//Repository
	addressRepo                       *repository.AddressRepo
	analyticsConfigRepo               *repository.AnalyticsConfigRepo
	batchJobRepo                      *repository.BatchJobRepo
	cartRepo                          *repository.CartRepo
	claimImageRepo                    *repository.ClaimImageRepo
	claimItemRepo                     *repository.ClaimItemRepo
	claimTagRepo                      *repository.ClaimTagRepo
	claimRepo                         *repository.ClaimRepo
	countryRepo                       *repository.CountryRepo
	currencyRepo                      *repository.CurrencyRepo
	customShippingOptionRepo          *repository.CustomShippingOptionRepo
	customerGroupRepo                 *repository.CustomerGroupRepo
	customerRepo                      *repository.CustomerRepo
	discountConditionRepo             *repository.DiscountConditionRepo
	discountRuleRepo                  *repository.DiscountRuleRepo
	discountRepo                      *repository.DiscountRepo
	draftOrderRepo                    *repository.DraftOrderRepo
	fulfillmentProviderRepo           *repository.FulfillmentProviderRepo
	fulfillmentRepo                   *repository.FulfillmentRepo
	giftCardTransactionRepo           *repository.GiftCardTransactionRepo
	giftCardRepo                      *repository.GiftCardRepo
	idempotencyKeyRepo                *repository.IdempotencyKeyRepo
	imageRepo                         *repository.ImageRepo
	inviteRepo                        *repository.InviteRepo
	ineItemAdjustmentRepo             *repository.LineItemAdjustmentRepo
	ineItemTaxLineRepo                *repository.LineItemTaxLineRepo
	ineItemRepo                       *repository.LineItemRepo
	moneyAmountRepo                   *repository.MoneyAmountRepo
	noteRepo                          *repository.NoteRepo
	notificationProviderRepo          *repository.NotificationProviderRepo
	notificationRepo                  *repository.NotificationRepo
	oAuthRepo                         *repository.OAuthRepo
	orderEditRepo                     *repository.OrderEditRepo
	orderItemChangeRepo               *repository.OrderItemChangeRepo
	orderRepo                         *repository.OrderRepo
	paymentCollectionRepo             *repository.PaymentCollectionRepo
	paymentProviderRepo               *repository.PaymentProviderRepo
	paymentSessionRepo                *repository.PaymentSessionRepo
	paymentRepo                       *repository.PaymentRepo
	priceListRepo                     *repository.PriceListRepo
	productCategoryRepo               *repository.ProductCategoryRepo
	productCollectionRepo             *repository.ProductCollectionRepo
	productOptionValueRepo            *repository.ProductOptionValueRepo
	productOptionRepo                 *repository.ProductOptionRepo
	productTagRepo                    *repository.ProductTagRepo
	productTaxRateRepo                *repository.ProductTaxRateRepo
	productTypeRepo                   *repository.ProductTypeRepo
	productVariantInventoryItem       *repository.ProductVariantInventoryItem
	productVariantRepo                *repository.ProductVariantRepo
	productRepo                       *repository.ProductRepo
	publishableApiKeySalesChannelRepo *repository.PublishableApiKeySalesChannelRepo
	publishableApiKeyRepo             *repository.PublishableApiKeyRepo
	refundRepo                        *repository.RefundRepo
	regionRepo                        *repository.RegionRepo
	returnItemRepo                    *repository.ReturnItemRepo
	returnReasonRepo                  *repository.ReturnReasonRepo
	returnRepo                        *repository.ReturnRepo
	salesChannelLocationRepo          *repository.SalesChannelLocationRepo
	salesChannelRepo                  *repository.SalesChannelRepo
	shippingMethodTaxLineRepo         *repository.ShippingMethodTaxLineRepo
	shippingMethodRepo                *repository.ShippingMethodRepo
	shippingOptionRequirementRepo     *repository.ShippingOptionRequirementRepo
	shippingOptionRepo                *repository.ShippingOptionRepo
	shippingProfileRepo               *repository.ShippingProfileRepo
	shippingTaxRateRepo               *repository.ShippingTaxRateRepo
	stagedJobRepo                     *repository.StagedJobRepo
	storeRepo                         *repository.StoreRepo
	swapRepo                          *repository.SwapRepo
	taxProviderRepo                   *repository.TaxProviderRepo
	taxRateRepo                       *repository.TaxRateRepo
	trackingLinkRepo                  *repository.TrackingLinkRepo
	userRepo                          *repository.UserRepo

	//Services
	analyticsConfigService         *services.AnalyticsConfigService
	authService                    *services.AuthService
	batchJobService                *services.BatchJobService
	cartService                    *services.CartService
	claimItemService               *services.ClaimItemService
	claimService                   *services.ClaimService
	csvParserService               *services.CsvParserService
	currencyService                *services.CurrencyService
	customShippingOptionService    *services.CustomShippingOptionService
	customerGroupService           *services.CustomerGroupService
	customerService                *services.CustomerService
	discountConditionService       *services.DiscountConditionService
	discountService                *services.DiscountService
	draftOrderService              *services.DraftOrderService
	eventBusService                *services.Bus
	defaultFileService             *services.DefaultFileService
	fulfillmentProviderService     *services.FulfillmentProviderService
	fulfillmentService             *services.FulfillmentService
	giftCardService                *services.GiftCardService
	idempotencyKeyService          *services.IdempotencyKeyService
	inviteService                  *services.InviteService
	lineItemAdjustmentService      *services.LineItemAdjustmentService
	lineItemService                *services.LineItemService
	newTotalsService               *services.NewTotalsService
	noteService                    *services.NoteService
	notificationService            *services.NotificationService
	oAuthService                   *services.OAuthService
	orderItemChangeService         *services.OrderItemChangeService
	orderEditService               *services.OrderEditService
	orderService                   *services.OrderService
	paymentCollectionService       *services.PaymentCollectionService
	paymentProviderService         *services.PaymentProviderService
	paymentService                 *services.PaymentService
	priceListService               *services.PriceListService
	pricingService                 *services.PricingService
	productCategoryService         *services.ProductCategoryService
	productCollectionService       *services.ProductCollectionService
	productTagService              *services.ProductTagService
	productTaxRateService          *services.ProductTaxRateService
	productTypeService             *services.ProductTypeService
	productVariantInventoryService *services.ProductVariantInventoryService
	productVariantService          *services.ProductVariantService
	productService                 *services.ProductService
	publishableApiKeyService       *services.PublishableApiKeyService
	regionService                  *services.RegionService
	returnReasonService            *services.ReturnReasonService
	returnedService                *services.ReturnService
	salesChannelInventoryService   *services.SalesChannelInventoryService
	salesChannelLocationService    *services.SalesChannelLocationService
	salesChannelService            *services.SalesChannelService
	defaultSearchService           *services.DefaultSearchService
	shippingOptionService          *services.ShippingOptionService
	shippingProfileService         *services.ShippingProfileService
	shippingTaxRateService         *services.ShippingTaxRateService
	stagedJobService               *services.StagedJobService
	storeService                   *services.StoreService
	strategyResolverService        *services.StrategyResolverService
	swapService                    *services.SwapService
	systemProviderService          *services.SystemProviderService
	systemTaxService               *services.SystemTaxService
	taxProviderService             *services.TaxProviderService
	taxRateService                 *services.TaxRateService
	tockenService                  *services.TockenService
	totalsService                  *services.TotalsService
	userService                    *services.UserService
	flagRouter                     *services.FlagRouter

	//Routes
	adminAnalyticsConfig   *admin.AnalyticsConfig
	adminApp               *admin.App
	adminAuth              *admin.Auth
	adminBatch             *admin.Batch
	adminCollection        *admin.Collection
	adminCurrencie         *admin.Currencie
	adminCustomerGroup     *admin.CustomerGroup
	adminCustomer          *admin.Customer
	adminDiscount          *admin.Discount
	adminDraftOrder        *admin.DraftOrder
	adminGiftCard          *admin.GiftCard
	adminInventoryItem     *admin.InventoryItem
	adminInvite            *admin.Invite
	adminNote              *admin.Note
	adminNotification      *admin.Notification
	adminOrderEdit         *admin.OrderEdit
	adminOrder             *admin.Order
	adminPaymentCollection *admin.PaymentCollection
	adminPayment           *admin.Payment
	adminPriceList         *admin.PriceList
	adminProductCategory   *admin.ProductCategory
	adminProductTag        *admin.ProductTag
	adminProductType       *admin.ProductType
	adminProduct           *admin.Product
	adminPublishableApiKey *admin.PublishableApiKey
	adminRegion            *admin.Region
	adminReservation       *admin.Reservation
	adminReturnReason      *admin.ReturnReason
	adminReturn            *admin.Return
	adminSalesChannel      *admin.SalesChannel
	adminShippingOption    *admin.ShippingOption
	adminShippingProfile   *admin.ShippingProfile
	adminStockLocation     *admin.StockLocation
	adminStore             *admin.Store
	adminSwap              *admin.Swap
	adminTaxRate           *admin.TaxRate
	adminUpload            *admin.Upload
	adminUser              *admin.User
	adminVariant           *admin.Variant
}

func NewRegistry() *Base {
	r := new(Base)
	r.Config()
	return r
}

func (m *Base) WithBuildInfo(version, hash, date string) *Base {
	m.buildVersion = version
	m.buildHash = hash
	m.buildDate = date
	return m
}

func (m *Base) BuildVersion() string {
	return m.buildVersion
}

func (m *Base) BuildDate() string {
	return m.buildDate
}

func (m *Base) BuildHash() string {
	return m.buildHash
}

func (m *Base) Logger() *zap.SugaredLogger {
	if m.logger == nil {
		level, err := zap.ParseAtomicLevel(m.Config().Logger.Level)
		if err != nil {
			panic(err)
		}
		var encoderConfig zapcore.EncoderConfig
		if m.Config().Logger.Development {
			encoderConfig = zap.NewDevelopmentEncoderConfig()
		} else {
			encoderConfig = zap.NewProductionEncoderConfig()
		}
		config := zap.Config{
			Level:         level,
			Development:   m.Config().Logger.Development,
			Encoding:      m.Config().Logger.Encoding,
			EncoderConfig: encoderConfig,
			Sampling: &zap.SamplingConfig{
				Initial:    100,
				Thereafter: 100,
			},
			OutputPaths: []string{
				"stderr",
			},
			ErrorOutputPaths: []string{
				"stderr",
			},
			InitialFields: map[string]interface{}{
				"pid": os.Getpid(),
			},
		}
		m.logger = zap.Must(config.Build()).Sugar()
		defer m.logger.Sync()
	}
	return m.logger
}

func (m *Base) Config() *config.Config {
	if m.config == nil {
		con, err := config.LoadConfig("./config")
		if err != nil {
			m.Logger().Fatal("Unable to create config service")
		}
		m.config = con
	}
	return m.config
}

func (m *Base) Validator() *validator.Validate {
	if m.validator == nil {
		validate := validator.New()

		// register function to get tag name from json tags.
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		m.validator = validate
	}

	return m.validator
}

func (m *Base) Tracer(ctx context.Context) trace.Tracer {
	if m.trc == nil {
		tp := otel.GetTracerProvider()
		m.trc = tp.Tracer("github.com/driver005/gateway", trace.WithInstrumentationVersion(m.BuildVersion()))
	}

	return m.trc
}

func (m *Base) Database() sql.Database {
	return m.database
}

func (m *Base) Container() di.Container {
	return m.container
}

func (m *Base) GetDialector(database config.Database) gorm.Dialector {
	var credentials string
	if len(database.Username) > 0 {
		credentials = database.Username
		if len(database.Username) > 0 {
			credentials += fmt.Sprintf(":%s", database.Password)
		}
		credentials += "@"
	}

	switch database.Type {
	case "mysql":
		return mysql.New(mysql.Config{
			DriverName: database.Type,
			DSN: fmt.Sprintf(
				"%stcp(%s:%d)/%s",
				credentials,
				database.Host,
				database.Port,
				database.DBname,
			),
		})
	case "pgx", "pq", "postgres", "cockroach":
		return postgres.New(postgres.Config{
			DriverName: database.Type,
			DSN: fmt.Sprintf(
				"postgres://%s%s:%d/%s",
				credentials,
				database.Host,
				database.Port,
				database.DBname,
			),
		})
	case "sqlite":
		return sqlite.Open(
			database.DBname,
		)
	case "sqlserver":
		return sqlserver.New(sqlserver.Config{
			DriverName: database.Type,
			DSN: fmt.Sprintf(
				"sqlserver://%s%s:%d?database=%s",
				credentials,
				database.Host,
				database.Port,
				database.DBname,
			),
		})
	default:
		m.Logger().Fatal("Unknown database driver, please change the config file")
		return nil
	}
}

func (m *Base) Init(ctx context.Context) error {
	if m.database == nil {
		// new db connection
		database, err := gorm.Open(m.GetDialector(m.Config().MasterDatabase), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger:                                   dbLogger.Default.LogMode(dbLogger.Silent),
		})

		if err != nil {
			return err
		}

		m.database, err = sql.NewManager(database, m.logger)

		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Base) Context() *gorm.DB {
	return m.database.DB(context.Background())
}

func (m *Base) Manager(ctx context.Context) *gorm.DB {
	return m.database.DB(ctx)
}

func (m *Base) Migration() *migrations.Handler {
	if m.migrations == nil {
		m.migrations = migrations.New(m)
	}
	return m.migrations
}

func (m *Base) Session() *session.Store {
	if m.session == nil {
		m.session = session.New()
	}
	return m.session
}

func (m *Base) Middleware() *middlewares.Handler {
	if m.middleware == nil {
		m.middleware = middlewares.New(m)
	}
	return m.middleware
}

func (m *Base) NewRouter(app *fiber.App) *fiber.App {
	if m.router == nil {
		m.router = app
	}
	return m.router
}

func (m *Base) AdminRouter() fiber.Router {
	if m.adminRouter == nil {
		m.adminRouter = m.router.Group("/admin")
	}
	return m.adminRouter
}

func (m *Base) StoreRouter() fiber.Router {
	if m.storeRouter == nil {
		m.storeRouter = m.router.Group("/store")
	}
	return m.storeRouter
}

func (m *Base) Routes() *routes.Routes {
	if m.routes == nil {
		m.routes = routes.New(m)
	}
	return m.routes
}

func (m *Base) AddRoutes(router *fiber.App) {
	m.NewRouter(router)
	m.Routes().SetRoutes()
}

func (m *Base) Setup() {
	app := fiber.New(fiber.Config{
		ServerHeader:   m.Config().Server.ServerName,
		AppName:        m.Config().Server.ServerName,
		WriteTimeout:   time.Duration(m.Config().Server.Timeout) * time.Second,
		ReadTimeout:    time.Duration(m.Config().Server.Timeout) * time.Second,
		StrictRouting:  true,
		ReadBufferSize: m.Config().Server.RequestBodyLimit,
		ErrorHandler:   utils.ErrorHandler,
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		//lint:ignore S1005 unnecessary assignment to the blank identifier
		_ = <-c
		fmt.Println("\nGracefully shutting down...")
		_ = app.Shutdown()
	}()

	app.Use(favicon.New())

	m.AddRoutes(app)

	if err := app.Listen(fmt.Sprintf("%s:%d", m.Config().Server.Host, m.Config().Server.Port)); err != nil {
		m.Logger().Fatal(err)
	}
}
