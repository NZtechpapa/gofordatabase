# gofordatabase
1.  command :

For postgres
docker run -d --name fredpostgres -v postgres:/home/fred/Documents/work/database/gofordatabase/postgres -p 54320:5432 postgres
docker logs -f my_postgres
docker exec -it fredpostgres psql -U postgres
CREATE USER dbuser WITH PASSWORD 'lei123456'
CREATE DATABASE barfoot OWNER dbuser

2. 
For sql server:

sudo docker run -e 'ACCEPT_EULA=Y' -e 'SA_PASSWORD=Lei123456' -p 14330:1433 --name sql1 -d microsoft/mssql-server-linux 
docker run -e 'ACCEPT_EULA=Y' -e 'SA_PASSWORD=Lei123456' -p 14330:1433 --name sql1 -d microsoft/mssql-server-linux 
docker run -e 'ACCEPT_EULA=Y' -e 'MSSQL_SA_PASSWORD=Lei123456' -p 14330:1433 -v /home/fred/Documents/work/database/gofordatabase/mssql:/var/opt/mssql -d microsoft/mssql-server-linux
sudo docker exec -it sql1 "bash"

/opt/mssql-tools/bin/sqlcmd -S localhost -U SA -P 'Lei123456'
CREATE DATABASE barfootDB
GO
Fredtest

3.  
sqlcmd -S 127.0.0.1 -p 14330 -U SA -P Lei123456 -d barfootDB -i ./CreateTestData.sql

sqlcmd -S 127.0.0.1  -U SA -P Lei123456 -d Fredtest 




type User struct {
    ID                       uint
    IndexID                  string //  2121212
    Heldbool                 bool   // false
    UnitID                   string //  1212
    OwnerSurname1            string //Jordan
    OwnerType1               int    // Individual
    OtherNames1              string //Angela Lynn
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







    // Insert Customers
    fmt.Println("------ InsertCustomer ------------")

    customerObj := models.Customer{
        Email:     "test@gmail.com",
        Nick:      "test1",
        Firstname: "FN_1",
        Lastname:  "LN_1",
        Age:       15}

    customerObj1 := models.Customer{
        Email:     "test2@gmail.com",
        Nick:      "test2",
        Firstname: "FN_2",
        Lastname:  "LN_2",
        Age:       35}

    insertRes := dbmanagers.InsertCustomer(customerObj)
    fmt.Println("Customer was inserted : ", insertRes)

    insertRes1 := dbmanagers.InsertCustomer(customerObj1)
    fmt.Println("Customer was inserted : ", insertRes1)

    resultDb = dbmanagers.GetCustomers()

    for _, item := range resultDb {
        fmt.Println(item.ToString())
    }

    // Update Customer
    fmt.Println("------ Update Customer ------------")

    customerUpd := resultDb[0]

    customerUpd.Nick = "test_upd"
    customerUpd.Email = "test_upd@gmail.com"

    updateRes := dbmanagers.UpdateCustomer(customerUpd)
    fmt.Println("Customer was updated : ", updateRes)

    // Delete Customer
    fmt.Println("------ Delete Customer ------------")
    updateRes = dbmanagers.DeleteCustomer(customerUpd)
    fmt.Println("Customer was deleted : ", updateRes)

    fmt.Println("------ GetCustomer ------------")
    resultDb = dbmanagers.GetCustomers()
    fmt.Printf("There are customers %d in DB \n\n", len(resultDb))