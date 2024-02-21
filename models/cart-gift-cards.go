package models

import "github.com/google/uuid"

type CartGiftCard struct {
	CartId     uuid.NullUUID `gorm:"column:cart_id;type:character varying;primaryKey;index:IDX_d38047a90f3d42f0be7909e8ae,priority:1" json:"cart_id"`
	GiftCardId uuid.NullUUID `gorm:"column:gift_card_id;type:character varying;primaryKey;index:IDX_0fb38b6d167793192bc126d835,priority:1" json:"gift_card_id"`
}
