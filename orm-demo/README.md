# ORM demo with GO and SQLite(?)

This demo project is based on [youtube video](https://youtube.cin/watch?v=VAGodyl84OY)

It uses [this ORM library](https://gorm.io/docs/) to implement ORM and connect to SQLite.

## Installation

```bash
# get orm package
go get -u github.com/jinzhu/gorm
# get sqlite3 driver
# did not worked properly
go get github.com/mattn/go-sqlite3

# needed to install g++ to build
sudo apt-get install g++

# for postgres installing 
go get github.com/jinzhu/gorm/dialects/postgres



```
