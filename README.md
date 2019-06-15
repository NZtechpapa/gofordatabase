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


sudo docker exec -it sql1 "bash"

/opt/mssql-tools/bin/sqlcmd -S localhost -U SA -P 'Lei123456'
CREATE DATABASE TestDB

Fredtest

3.  
sqlcmd -S 127.0.0.1 -p 14330 -U SA -P Lei123456 -d barfootDB -i ./CreateTestData.sql

sqlcmd -S 127.0.0.1  -U SA -P Lei123456 -d Fredtest 