package migrations

import (
	"github.com/driver005/gateway/interfaces"
)

type Handler struct {
	r         Registry
	migrators []interfaces.IMigrator
}

func New(r Registry) *Handler {
	return &Handler{
		r,
		[]interfaces.IMigrator{
			&InitialSchema1611063162649{r: r},
			&CountriesCurrencies1611063174563{r: r},
			&Claims1612284947120{r: r},
			&Indexes1612353094577{r: r},
			&Notifications1613146953072{r: r},
			&ProductTypeCategoryTags1613146953073{r: r},
			&DraftOrders1613384784316{r: r},
			&TrackingLinks1613656135167{r: r},
			&CartContext1614684597235{r: r},
			&ReturnReason1615891636559{r: r},
			&DiscountUsageCount1615970124120{r: r},
			&DiscountUsage1617002207608{r: r},
			&NullablePassword1619108646647{r: r},
			&NoNotification1623231564533{r: r},
			&GcRemoveUniqueOrder1624287602631{r: r},
			&SoftDeletingUniqueConstraints1624610325746{r: r},
			&EnsureCancellationFieldsExist1625560513367{r: r},
			&AddDiscountableToProduct1627995307200{r: r},
			&AllowBackorderSwaps1630505790603{r: r},
			&RankColumnWithDefaultValue1631104895519{r: r},
			&EnforceUniqueness1631261634964{r: r},
			&ValidDurationForDiscount1631696624528{r: r},
			&NestedReturnReasons1631800727788{r: r},
			&StatusOnProduct1631864388026{r: r},
			&AddNotes1632220294687{r: r},
			&DeleteDateOnShippingOptionRequirements1632828114899{r: r},
			&ExtendedUserApi1633512755401{r: r},
			&AddCustomShippingOptions1633614437919{r: r},
			&OrderTaxRateToRealType1638543550000{r: r},
			&ExternalIdOrder1638952072999{r: r},
			&NewTaxSystem1641636508055{r: r},
			&CustomerGroups1644943746861{r: r},
			&DiscountConditions1646324713514{r: r},
			&UpdateMoneyAmountAddPriceList1646915480108{r: r},
			&AddLineItemAdjustments1648600574750{r: r},
			&TaxLineConstraints1648641130007{r: r},
			&AddBatchJobModel1649775522087{r: r},
			&SalesChannel1656949291839{r: r},
			&TaxedGiftCardTransactions1657098186554{r: r},
			&ExtendedBatchJob1657267320181{r: r},
			&TaxInclusivePricing1659501357661{r: r},
			&PaymentSessionUniqCartIdProviderId1660040729000{r: r},
			&MultiPaymentCart1661345741249{r: r},
			&SwapFulfillmentStatusRequiresAction1661863940645{r: r},
			&OrderEditing1663059812399{r: r},
			&LinteItemOriginalItemRelation1663059812400{r: r},
			&PaymentCollection1664880666982{r: r},
			&AddAnalyticsConfig1666173221888{r: r},
			&PublishableApiKey1667815005070{r: r},
			&UpdateCustomerEmailConstraint1669032280562{r: r},
			&AddTaxRateToGiftCards1670855241304{r: r},
			&MultiLocation1671711415179{r: r},
			&ProductCategory1672906846559{r: r},
			&PaymentSessionIsInitiated1672906846560{r: r},
			&StagedJobOptions1673003729870{r: r},
			&UniquePaySessCartId1673550502785{r: r},
			&ProductCategoryProduct1674455083104{r: r},
			&MultiLocationSoftDelete1675689306130{r: r},
			&ProductCategoryRank1677234878504{r: r},
			&EnsureRequiredQuantity1678093365811{r: r},
			&LineitemAdjustmentsAmount1678093365812{r: r},
			&CategoryRemoveSoftDelete1679950221063{r: r},
			&CategoryCreateIndexes1679950645253{r: r},
			&ProductDomainImpovedIndexes1679950645254{r: r},
			&ProductSearchGinIndexes1679950645254{r: r},
			&AddSalesChannelMetadata1680714052628{r: r},
			&AddDescriptionToProductCategory1680857773272{r: r},
			&LineItemTaxAdjustmentOnCascadeDelete1680857773272{r: r},
			&AddTableProductShippingProfile1680857773273{r: r},
			&DropProductIdFkSalesChannels1680857773273{r: r},
			&DropVariantIdFkMoneyAmount1680857773273{r: r},
			&UpdateReturnReasonIndex1692870898423{r: r},
			&LineitemProductId1692870898424{r: r},
			&AddTimestempsToProductShippingProfiles1692870898425{r: r},
			&DropMoneyAmountConstraintsForPricingModule1692953518123{r: r},
			&DropFksIsolatedProduct1694602553610{r: r},
			&ProductSalesChannelsLink1698056997411{r: r},
			&CartSalesChannelsLink1698160215000{r: r},
			&DropNonNullConstraintPriceList1699371074198{r: r},
			&AddMetadataToProductCategory1699564794649{r: r},
			&OrderSalesChannelsLink1701860329931{r: r},
			&PublishableKeySalesChannelsLink1701894188811{r: r},
		},
	}
}

func (h *Handler) Add(migrators []interfaces.IMigrator) {
	h.migrators = append(h.migrators, migrators...)
}

func (h *Handler) Up() error {
	for _, migrator := range h.migrators {
		h.r.Logger().Infof("Running migration: %s \n", migrator.GetName())
		if err := migrator.Up(); err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) Down() error {
	for i := len(h.migrators) - 1; i >= 0; i-- {
		h.r.Logger().Infof("Droping migration: %s \n", h.migrators[i].GetName())
		if err := h.migrators[i].Down(); err != nil {
			return err
		}
	}
	return nil
}
