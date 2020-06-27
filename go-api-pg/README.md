# oAuth2 api with GO and Postgres

This project is developed to explore Golang and compare it to NodeJS oauth2 approach. It uses go-pg modules.

Recently I was playing with Python/Flask api, NodeJS and now with Go. I use all these api with PostgreSQL in Docker containers. During thinkering with Python for example I discovered alpine versions which make these containers siginificantly smaller. I recently had a look at dotnet core for creating api's with C#. I noticed that C# and MS SQL docker containers are large. So I tried to reduce the size running dotnet core on alpine. Next Postgres in alpine version seem to be significantly smaller than MSSQL. So my decision for primary database became Postgres.

## Installation

```bash
# install pg module
go get github.com/go-pg/pg


```

## Pg go module

Pg go module is used to connect to Postgres. It supports inital creation of tables. The tables should be defined as stuct types (data models). See examples in db folder.
