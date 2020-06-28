# oAuth2 api with GO and Postgres

This project is developed to explore Golang and compare it to NodeJS oauth2 approach.

Recently I was playing with Python/Flask api, NodeJS and now with Go. I use all these api with PostgreSQL in Docker containers. During thinkering with Python for example I discovered alpine versions which make these containers siginificantly smaller. I recently had a look at dotnet core for creating api's with C#. I noticed that C# and MS SQL docker containers are large. So I tried to reduce the size running dotnet core on alpine. Next Postgres in alpine version seem to be significantly smaller than MSSQL. So my decision for primary database became Postgres.

Go has basic module for communication with [sql databases](https://golang.org/pkg/database/sql/).

Additional light library on the top of standard is sqlx. In this example we will [use sqlx](https://github.com/jmoiron/sqlx).

## Installation

```bash
# install sqlx module
go get github.com/jmoiron/sqlx

```

## Sqlx modules

sqlx is a library which provides a set of extensions on go's standard database/sql library. The sqlx versions of sql.DB, sql.TX, sql.Stmt, et al. all leave the underlying interfaces untouched, so that their interfaces are a superset on the standard ones. This makes it relatively painless to integrate existing codebases using database/sql with sqlx.

## Postgres database driver (pq)

This is [basic postgres driver for go](https://pkg.go.dev/mod/github.com/lib/pq).
After lookin around I decided to use this driver only.

At this moment there are lot of different packages. The basic use of go database package and [pq driver is demonstrated in this youtube video](https://youtube.com/watch?v=tOBosuiJIHM)
