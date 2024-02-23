package models

import "github.com/google/uuid"

type OrderGiftCard struct {
	OrderId    uuid.NullUUID `gorm:"column:order_id;type:character varying;primaryKey;index:IDX_e62ff11e4730bb3adfead979ee,priority:1" json:"order_id"`
	GiftCardId uuid.NullUUID `gorm:"column:gift_card_id;type:character varying;primaryKey;index:IDX_f2bb9f71e95b315eb24b2b84cb,priority:1" json:"gift_card_id"`
}
