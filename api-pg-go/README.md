# Basic api demo with GO and Postgres

This demo project uses basis Go database/sql module and official pg driver for postgres.

The intention is to create simple api with less as possible dependencies, so we use:

- basic go database module and postgres driver
- basic go http router (defaul mux)
- encode/json for json marshaling

## Module stucture

There are some ideas about MVC structure but I am inclined using a custom structure:

- pgdb: module responsible for postgres database connection
- model: for data types/models. The user model contains direct connection to pgdb currently. Im am considering interface implementation for loose coupling?
- routes module: container routes and calls appropriate method in models
