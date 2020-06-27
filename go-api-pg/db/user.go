package db

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

/*
User type for users table. By default pg module will convert PascalCase type names to snake_case.
If you want to have custom tablename provide tableName struct as below.
Additional instructions can be provided using backticks and sql namespace.
*/
type User struct {
	tableName  struct{}  `pg:"users"`
	ID         string    `pg:"id, type:uuid, default:public.uuid_generate_v4()" json:"id"`
	Roles      string    `pg:"roles, type:VARCHAR(100), NOT NULL" json:"roles"`
	FirstName  string    `pg:"first_name, notnull" json:"first_name"`
	LastName   string    `pg:"last_name, notnull" json:"last_name"`
	Email      string    `pg:"email, pk, unique" json:"email"`
	password   string    `pg:"password, notnull"`
	BirthDate  time.Time `pg:"birth_date, null" json:"birth_date"`
	CreateDate time.Time `pg:"createdate, notnull, default: CURRENT_DATE"`
}

// Insert the user to database
func (u *User) Insert(db *pg.DB) (*User, error) {
	// err := db.Insert(u)
	user, err := db.Model(u).Returning("*").Insert()
	if err != nil {
		log.Printf("Failed to insert user...%v\n", err)
		return nil, err
	}
	log.Printf("User inserted...%v", u.Email)
	return user, nil
}

// CreateUsersTable will create table is db passed to this method.
func CreateUsersTable(db *pg.DB) error {
	opt := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.CreateTable(&User{}, opt)
	if err != nil {
		log.Printf("Failed to create users table: %v", err)
		return err
	}
	return nil
}
