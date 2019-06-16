package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	//	"strconv"
)

type Barfoottable1 struct {
	ID    int    `xorm:"ID PK int autoincr"`
	Held  bool   `json:"Held"`
	Unit  string `json:"Unit"`
	Owner []struct {
		Surname       string `json:"Surname"`
		OwnerType     string `json:"OwnerType"`
		OtherNames    string `json:"OtherNames"`
		CorporateName string `json:"CorporateName"`
	} `json:"Owner"`
	RafID                   int     `json:"RafID"`
	TlaID                   int     `json:"TlaId"`
	Images                  string  `json:"Images"`
	SaleID                  int     `json:"SaleID"  xorm:"SaleID int default 0"`
	Street                  string  `json:"Street"`
	Suffix                  string  `json:"Suffix"`
	Tenure                  string  `json:"Tenure"`
	ValRef                  string  `json:"ValRef"`
	WardID                  int     `json:"WardId"`
	Garages                 int     `json:"Garages"`
	TlaName                 string  `json:"TlaName"`
	Bedrooms                int     `json:"Bedrooms"`
	Category                string  `json:"Category"`
	LandArea                int     `json:"LandArea"`
	Latitude                string  `json:"Latitude"`
	ListDate                string  `json:"ListDate"`
	Postcode                string  `json:"Postcode"`
	RegionID                int     `json:"RegionId"`
	SaleAsIs                bool    `json:"SaleAsIs"`
	SaleDate                string  `json:"SaleDate"`
	SellDays                int     `json:"SellDays"`
	SuburbID                int     `json:"SuburbId"`
	TownCity                string  `json:"TownCity"`
	WardName                string  `json:"WardName"`
	Bathrooms               int     `json:"Bathrooms"`
	FloorArea               int     `json:"FloorArea"`
	LandValue               int     `json:"LandValue"`
	ListPrice               int     `json:"ListPrice"`
	ListingID               int     `json:"ListingID"`
	Longitude               string  `json:"Longitude"`
	SalePrice               int     `json:"SalePrice"`
	Valuation               int     `json:"Valuation"`
	ExternalID              string  `json:"ExternalId"`
	RegionName              string  `json:"RegionName"`
	SaleMethod              string  `json:"SaleMethod"`
	SuburbName              string  `json:"SuburbName"`
	Suppressed              bool    `json:"Suppressed"`
	CreatedDate             string  `json:"CreatedDate"`
	GaragesDVR              int     `json:"Garages-DVR"`
	StreetNumber            string  `json:"StreetNumber"`
	DVRLandValue            float64 `json:"DVR_LandValue"`
	DVRValuation            float64 `json:"DVR_Valuation"`
	IsNewDwelling           bool    `json:"IsNewDwelling"`
	ValuationDate           string  `json:"ValuationDate"`
	LastChangeDate          string  `json:"LastChangeDate"`
	SettlementDate          string  `json:"SettlementDate"`
	LastUpdatedDate         string  `json:"LastUpdatedDate"`
	DVRValuationDate        string  `json:"DVR_ValuationDate"`
	ValuationImprovement    int     `json:"ValuationImprovement"`
	DVRValuationImprovement float64 `json:"DVR_ValuationImprovement"`
}

func (barfoottable *Barfoottable1) ToString() string {

	var buffer bytes.Buffer
	buffer.WriteString("------- User --------")

	return buffer.String()
}

func (barfoottable *Barfoottable1) ToJson() string {
	result, err := json.Marshal(barfoottable)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	return string(result)
}
