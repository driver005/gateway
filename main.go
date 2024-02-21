package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./query",
		FieldNullable:     false,
		FieldCoverable:    false,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})
	gormdb, _ := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/medusa-test"))
	g.UseDB(gormdb)

	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
	// cmd.Execute()
}

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

// gormdb, _ := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/medusa"))
// g.UseDB(gormdb)

// g.ApplyBasic(
// 	// Generate structs from all tables of current database
// 	g.GenerateAllTable()...,
// )
// // Generate the code
// g.Execute()
