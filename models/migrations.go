package models

type Migration struct {
	Id        int32  `gorm:"column:id;type:integer;primaryKey;autoIncrement:true" json:"id"`
	Timestamp int64  `gorm:"column:timestamp;type:bigint;not null" json:"timestamp"`
	Name      string `gorm:"column:name;type:character varying;not null" json:"name"`
}

func (*Migration) TableName() string {
	return "migrations"
}
