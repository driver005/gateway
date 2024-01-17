package migrations

import "reflect"

type TaxedGiftCardTransactions1657098186554 struct {
	r Registry
}

func (m *TaxedGiftCardTransactions1657098186554) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *TaxedGiftCardTransactions1657098186554) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "gift_card_transaction" ADD "is_taxable" boolean`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "gift_card_transaction" ADD "tax_rate" real`).Error; err != nil {
		return err
	}
	return nil
}
func (m *TaxedGiftCardTransactions1657098186554) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "gift_card_transaction" DROP COLUMN "is_taxable"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "gift_card_transaction" DROP COLUMN "tax_rate"`).Error; err != nil {
		return err
	}
	return nil
}
