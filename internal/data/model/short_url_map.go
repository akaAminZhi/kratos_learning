package model

const TableNameShortUrlMap = "short_url_map"

type ShortUrlMap struct {
	ID   int64   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id" redis:"id"`
	Lurl *string `gorm:"column:lurl;" json:"lurl" redis:"lurl"`
	Md5  string  `gorm:"column:md5;" json:"md5" redis:"md5"`
	Surl string  `gorm:"column:surl;" json:"surl" redis:"surl"`
}

func (s *ShortUrlMap) TableName() string {
	return TableNameShortUrlMap
}
