package migrations

import (
	"reflect"
	"strings"

	"github.com/driver005/gateway/utils"
)

type CountriesCurrencies1611063174563 struct {
	r Registry
}

func (m *CountriesCurrencies1611063174563) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *CountriesCurrencies1611063174563) Up() error {
	countries := utils.Countries
	for _, c := range countries {
		query := `INSERT INTO "country" ("iso_2", "iso_3", "num_code", "name", "display_name") VALUES ($1, $2, $3, $4, $5)`
		iso2 := strings.ToLower(c.Alpha2)
		iso3 := strings.ToLower(c.Alpha3)
		numeric := c.Numeric
		name := strings.ToUpper(c.Name)
		display := c.Name
		if err := m.r.Context().Exec(query, iso2, iso3, numeric, name, display).Error; err != nil {
			return err
		}
	}

	currencies := utils.Currencies
	for _, c := range currencies {
		query := `INSERT INTO "currency" ("code", "symbol", "symbol_native", "name") VALUES ($1, $2, $3, $4)`
		code := strings.ToLower(c.Code)
		sym := c.Symbol
		nat := c.SymbolNative
		name := c.Name
		if err := m.r.Context().Exec(query, code, sym, nat, name).Error; err != nil {
			return err
		}
	}

	return nil
}

func (m *CountriesCurrencies1611063174563) Down() error {
	countries := utils.Countries
	for _, c := range countries {
		query := `DELETE FROM "country" WHERE iso_2 = $1`
		if err := m.r.Context().Exec(query, c.Alpha2).Error; err != nil {
			return err
		}
	}

	currencies := utils.Currencies
	for _, c := range currencies {
		query := `DELETE FROM "currency" WHERE code = $1`
		if err := m.r.Context().Exec(query, c.Code).Error; err != nil {
			return err
		}
	}

	return nil
}
