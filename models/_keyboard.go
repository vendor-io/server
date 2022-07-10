package models

import (
	"gorm.io/gorm"
)

type KeyboardType string

const (
	SixtyPercent       KeyboardType = "60%"
	SixtyFivePercent   KeyboardType = "65%"
	SeventyFivePercent KeyboardType = "75%"
	TenKeyLess         KeyboardType = "TKL"
	FullCompact        KeyboardType = "95%"
	Fullsize           KeyboardType = "100%"
)

type KeyboardStruct struct {
	gorm.Model
	Name          string
	Type          KeyboardType
	Switches      string
	AmountInStock uint32
	Price         uint64
}
