# gofordatabase

For database server
1. config  sql server :
For sql server:
method 1:  using docker from mscrosoft.  pull the lastest version from docker hub
and run the following command: (modify the your path -v)
step 1:
docker run -e 'ACCEPT_EULA=Y' -e 'MSSQL_SA_PASSWORD=Lei123456' -p 14330:1433 -v /home/fred/Documents/work/database/gofordatabase/mssql:/var/opt/mssql -d microsoft/mssql-server-linux
step 2: test

sudo docker exec -it sql1 "bash"
/opt/mssql-tools/bin/sqlcmd -S localhost -U SA -P 'Lei123456'
CREATE DATABASE test_table
GO
show the table !  Done!

2.  Env config:
2.1  install the last1 golang  version  from the offical website
2.2  run :  go run  maim.go  under the pwd of src/ 
2.3  use  go get ***  to install the relative lib  folllowing the error info


3.  modify the sql serer info  in   usermanager.go

4.  add the new table  at  usermodel.go











