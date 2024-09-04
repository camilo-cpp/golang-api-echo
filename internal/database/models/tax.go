package models

type Tax struct {
	TaxId   string  `bson:"taxId"`
	TaxType string  `bson:"taxType"`
	Rate    float64 `bson:"rate"`
}
