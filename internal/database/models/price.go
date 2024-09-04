package models

type Price struct {
	FullPrice int   `bson:"fullPrice"`
	Taxes     []Tax `bson:"taxes"`
}
