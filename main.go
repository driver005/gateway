package main

import (
	"github.com/driver005/gateway/plugins/apikey"
	"github.com/driver005/gateway/plugins/auth"
	"github.com/driver005/gateway/plugins/cart"
	"github.com/driver005/gateway/plugins/currency"
	"github.com/driver005/gateway/plugins/customer"
	"github.com/driver005/gateway/plugins/fulfillment"
	"github.com/driver005/gateway/plugins/inventory"
	"github.com/driver005/gateway/plugins/notification"
	"github.com/driver005/gateway/plugins/order"
	"github.com/driver005/gateway/plugins/payment"
	"github.com/driver005/gateway/plugins/pricing"
	"github.com/driver005/gateway/plugins/product"
	"github.com/driver005/gateway/plugins/promotion"
	"github.com/driver005/gateway/plugins/region"
	"github.com/driver005/gateway/plugins/saleschannel"
	"github.com/driver005/gateway/plugins/stocklocation"
	"github.com/driver005/gateway/plugins/store"
	"github.com/driver005/gateway/plugins/tax"
	"github.com/driver005/gateway/plugins/user"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

//go:generate go run github.com/ogen-go/ogen/cmd/ogen -v --clean --config ./oas/ogen.yaml --target medusa ./openapi.full.yaml

// BackgroundCheck is your custom Workflow Definition.
// func BackgroundCheck(ctx workflow.Context, param string) (string, error) {
// 	Define the Activity Execution options
// 	StartToCloseTimeout or ScheduleToCloseTimeout must be set
// 	activityOptions := workflow.ActivityOptions{
// 		StartToCloseTimeout: 10 * time.Second,
// 	}
// 	ctx = workflow.WithActivityOptions(ctx, activityOptions)
// 	Execute the Activity synchronously (wait for the result before proceeding)
// 	var ssnTraceResult string
// 	err := workflow.ExecuteActivity(ctx, SSNTraceActivity, param).Get(ctx, &ssnTraceResult)
// 	if err != nil {
// 		return "", err
// 	}
// 	Make the results of the Workflow available
// 	return ssnTraceResult, nil
// }

// SSNTraceActivity is your custom Activity Definition.
// func SSNTraceActivity(ctx context.Context, param string) (*string, error) {
// 	This is where a call to another service is made
// 	Here we are pretending that the service that does SSNTrace returned "pass"
// 	result := "pass"
// 	fmt.Println("pass")
// 	return &result, nil
// }

// func main() {
// 	Initialize a Temporal Client
// 	Specify the Namespace in the Client options
// 	clientOptions := client.Options{
// 		Namespace: "backgroundcheck_namespace",
// 	}
// 	temporalClient, err := client.Dial(clientOptions)
// 	if err != nil {
// 		log.Fatalln("Unable to create a Temporal Client", err)
// 	}
// 	defer temporalClient.Close()
// 	Create a new Worker
// 	yourWorker := worker.New(temporalClient, "backgroundcheck-boilerplate-task-queue-local", worker.Options{})
// 	Register Workflows
// 	yourWorker.RegisterWorkflow(BackgroundCheck)
// 	Register Activities
// 	yourWorker.RegisterActivity(SSNTraceActivity)
// 	Start the Worker Process
// 	err = yourWorker.Run(worker.InterruptCh())
// 	if err != nil {
// 		log.Fatalln("Unable to start the Worker Process", err)
// 	}

// 	workflowOptions := client.StartWorkflowOptions{
// 		...
// 	}

// 	temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, "test")

// 	cmd.Execute()
// 	g := gen.NewGenerator(gen.Config{
// 		OutPath:           "./query",
// 		FieldNullable:     false,
// 		FieldCoverable:    false,
// 		FieldWithIndexTag: true,
// 		FieldWithTypeTag:  true,
// 	})
// 	gormdb, _ := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/medusa-user"))
// 	g.UseDB(gormdb)

// 	g.ApplyBasic(
// 		Generate structs from all tables of current database
// 		g.GenerateAllTable()...,
// 	)
// 	Generate the code
// 	g.Execute()
// }

// gormdb, _ := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/medusa-test2"))
// gormdb.AutoMigrate(
// 	models.Address{},
// 	models.AnalyticsConfig{},
// 	models.BatchJob{},
// 	models.Cart{},
// 	models.ClaimImage{},
// 	models.ClaimItem{},
// 	models.ClaimOrder{},
// 	models.ClaimTag{},
// 	models.Country{},
// 	models.Currency{},
// 	models.CustomShippingOption{},
// 	models.Customer{},
// 	models.CustomerGroup{},
// 	models.Discount{},
// 	models.DiscountCondition{},
// 	models.DiscountConditionCustomerGroup{},
// 	models.DiscountConditionProduct{},
// 	models.DiscountConditionProductCollection{},
// 	models.DiscountConditionProductTag{},
// 	models.DiscountConditionProductType{},
// 	models.DiscountRule{},
// 	models.DraftOrder{},
// 	models.Fulfillment{},
// 	models.FulfillmentItem{},
// 	models.FulfillmentProvider{},
// 	models.GiftCard{},
// 	models.GiftCardTransaction{},
// 	models.IdempotencyKey{},
// 	models.Image{},
// 	models.Invite{},
// 	models.LineItem{},
// 	models.LineItemAdjustment{},
// 	models.LineItemTaxLine{},
// 	models.MoneyAmount{},
// 	models.Note{},
// 	models.Notification{},
// 	models.NotificationProvider{},
// 	models.OAuth{},
// 	models.Order{},
// 	models.OrderEdit{},
// 	models.OrderItemChange{},
// 	models.Payment{},
// 	models.PaymentCollection{},
// 	models.PaymentProvider{},
// 	models.PaymentSession{},
// 	models.PriceList{},
// 	models.Product{},
// 	models.ProductCategory{},
// 	models.ProductCollection{},
// 	models.ProductOption{},
// 	models.ProductOptionValue{},
// 	models.ProductTag{},
// 	models.ProductTaxRate{},
// 	models.ProductType{},
// 	models.ProductTypeTaxRate{},
// 	models.ProductVariant{},
// 	models.ProductVariantInventoryItem{},
// 	models.ProductVariantMoneyAmount{},
// 	models.PublishableApiKey{},
// 	models.PublishableApiKeySalesChannel{},
// 	models.Refund{},
// 	models.Region{},
// 	models.Return{},
// 	models.ReturnItem{},
// 	models.ReturnReason{},
// 	models.SalesChannel{},
// 	models.SalesChannelLocation{},
// 	models.ShippingMethod{},
// 	models.ShippingMethodTaxLine{},
// 	models.ShippingOption{},
// 	models.ShippingOptionRequirement{},
// 	models.ShippingProfile{},
// 	models.ShippingTaxRate{},
// 	models.StagedJob{},
// 	models.Store{},
// 	models.Swap{},
// 	models.TaxLine{},
// 	models.TaxProvider{},
// 	models.TaxRate{},
// 	models.TrackingLink{},
// 	models.User{},
// )

// spec := oas.NewOpenAPI()
// spec.Parse(".", []string{}, "", false, "admin")
// d, err := yaml.Marshal(&spec)
// if err != nil {
// 	log.Fatalf("error: %v", err)
// }
// _ = os.WriteFile("./admin.base.yaml", d, 0644)

// g := gen.NewGenerator(gen.Config{
// 	OutPath:           "./query",
// 	FieldNullable:     false,
// 	FieldCoverable:    false,
// 	FieldWithIndexTag: true,
// 	FieldWithTypeTag:  true,
// })
// gormdb, _ := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/medusa-test"))
// g.UseDB(gormdb)

// g.ApplyBasic(
// 	Generate structs from all tables of current database
// 	g.GenerateAllTable()...,
// )
// Generate the code
// g.Execute()

// var dbList = []string{
// 	"api-key",
// 	"auth",
// 	"cart",
// 	"customer",
// 	"fulfillment",
// 	"order",
// 	"payment",
// 	"pricing",
// 	"product",
// 	"region",
// 	"sales-channel",
// 	"tax",
// 	"user",
// }

// func main() {
// 	for _, name := range dbList {
// 		g := gen.NewGenerator(gen.Config{
// 			FieldWithIndexTag: true,
// 			FieldWithTypeTag:  true,
// 			OutPath:           "./query",
// 			WithUnitTest:      true,
// 		})
// 		gormdb, _ := gorm.Open(postgres.Open(fmt.Sprintf("postgres://postgres:postgres@localhost:5432/medusa-%s", name)))
// 		g.UseDB(gormdb)

// 		g.ApplyBasic(
// 			// Generate structs from all tables of current database
// 			g.GenerateAllTable()...,
// 		)
// 		// Generate the code
// 		g.Execute()
// 	}
// }

func main() {

	gormdb, _ := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/medusa-test"))

	apikey.Init(gormdb)
	auth.Init(gormdb)
	cart.Init(gormdb)
	currency.Init(gormdb)
	customer.Init(gormdb)
	fulfillment.Init(gormdb)
	inventory.Init(gormdb)
	notification.Init(gormdb)
	order.Init(gormdb)
	payment.Init()
	pricing.Init(gormdb)
	product.Init(gormdb)
	promotion.Init(gormdb)
	region.Init(gormdb)
	saleschannel.Init(gormdb)
	stocklocation.Init(gormdb)
	store.Init(gormdb)
	tax.Init(gormdb)
	user.Init(gormdb)

	// db, err := gormdb.DB()
	// if err != nil {
	// 	panic(err)
	// }

	// if err := goose.Up(db, "test"); err != nil {
	// 	panic(err)
	// }

	g := gen.NewGenerator(gen.Config{
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		OutPath:           "./query",
		WithUnitTest:      true,
		Mode:              gen.WithQueryInterface,
	})

	g.UseDB(gormdb)

	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}
