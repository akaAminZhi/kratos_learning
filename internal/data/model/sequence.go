package model

const TableNameSequence = "sequence"

type Sequence struct {
	ID   int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id" redis:"id"`
	Stub string `gorm:"column:stub;not null" json:"stub" redis:"stub"`
}

func (s *Sequence) TableName() string {
	return TableNameSequence
}
