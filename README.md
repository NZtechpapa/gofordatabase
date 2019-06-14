# gofordatabase
1.  command :

For postgres
docker run -d --name fredpostgres -v my_dbdata:/var/lib/postgresql/data -p 54320:5432 postgres
docker logs -f my_postgres
docker exec -it fredpostgres psql -U postgres



For sql server:

sudo docker run -e 'ACCEPT_EULA=Y' -e 'SA_PASSWORD=Lei123456' -p 14330:1433 --name sql1 -d microsoft/mssql-server-linux 

sudo docker exec -it sql1 "bash"
