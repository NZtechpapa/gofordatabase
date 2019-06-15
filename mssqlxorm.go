//package mssqlxorm
package main

// Replace with your own connection parameters
var (
    server   = "localhost"
    port     = 14330
    user     = "sa"
    password = "Lei123456"
    database = "barfootDB"
)

type User struct {
    ID                       uint
    IndexID                  string //  2121212
    Heldbool                 bool   // false
    UnitID                   string //  1212
    OwnerSurname1            string //Jordan
    OwnerType1               int    // Individual
    OtherNames1              string //Angela Lynn
    CorporateName1           string //  ""
    OwnerSurname2            string //Jordan
    OwnerType2               int    // Individual
    OtherNames2              string //Angela Lynn
    CorporateName2           string //  ""
    RafID                    bigint //  1864011
    TlaId                    bigint // 5
    Images                   string // http://media-archive.realestate.co.nz/listings/1285127/9f33292e37a00e4072f956ac961db3c4
    SaleID                   bigint // 5
    Street                   string // Stredwick Drive
    Suffix                   string
    Tenure                   string
    ValRef                   string
    WardId                   bigint
    Garages                  int
    TlaName                  string
    Bedrooms                 int
    Category                 string
    LandArea                 int
    Latitude                 float64
    ListDate                 time
    Postcode                 int
    RegionId                 int
    SaleAsIs                 bool
    SaleDate                 time
    SellDays                 int
    SuburbId                 int
    TownCity                 string
    WardName                 string
    Bathrooms                int
    FloorArea                float32
    LandValue                float64
    ListPrice                float64
    ListingID                bigint  // 1285127
    Longitude                float64 // 174.7392
    SalePrice                float64 //382000
    Valuation                float64 //0
    ExternalId               string  //LEG-123607
    StreetNumber             string  //  120
    RegionName               string  // Auckland Region
    SuburbName               string  // torbay
    Suppressed               bool    // false
    SaleMethod               int     //     P - Private Treaty(Neg.)
    CreatedDate              date    // 2004-09-13
    Garages_DVR              int     // 1
    DVR_LandValue            float32 //690000.0
    DVR_Valuation            float32 //690000.0
    IsNewDwelling            bool
    ValuationDate            date    //  null
    SettlementDate           date    //  2004-09-13
    LastUpdatedDate          date    // 2004-08-06
    LastChangeDate           date    // 2019-02-19
    DVR_ValuationDate        date    //7/2017
    ValuationImprovement     float32 //0
    DVR_ValuationImprovement float64 // 410000.0
}

func main() {

    db, err := gorm.Open("mssql", "sqlserver://sa:Lei123456@localhost:14330?database=Master")
    defer db.Close()

    if err != nil {
        panic("failed to connect database")
    }
    db.DropTable(&User{})   // Delete old table if it exists.
    db.CreateTable(&User{}) // Create a new table.

    fmt.Printf("Connected!\n")

    // Add(Create) data
    db.Create(&user)

    fmt.Println(user)
}
