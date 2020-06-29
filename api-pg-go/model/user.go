package model

import (
	"time"

	"../pgdb"
)

// User structure
type User struct {
	ID         string    `pg:"id, type:uuid, default:public.uuid_generate_v4()" json:"id"`
	Roles      string    `pg:"roles, type:VARCHAR(100), NOT NULL" json:"roles"`
	FirstName  string    `pg:"first_name, notnull" json:"first_name"`
	LastName   string    `pg:"last_name, notnull" json:"last_name"`
	Email      string    `pg:"email, pk, unique" json:"email"`
	Password   string    `pg:"password, notnull" json:"-"`
	BirthDate  string    `pg:"birth_date, null" json:"birth_date"`
	CreateDate time.Time `pg:"createdate, notnull, default: CURRENT_DATE" json:"createdate"`
}

// GetAllUsers will extract all users from users table
func GetAllUsers() ([]User, error) {
	rows, err := pgdb.DB.Query("SELECT * FROM users;")
	// check for errors
	if err != nil {
		return nil, err
	}
	//close at the end
	defer rows.Close()
	//create users slice
	users := make([]User, 0)
	// extract rows
	for rows.Next() {
		//create new user instance
		user := User{}
		//load from db
		err := rows.Scan(&user.ID, &user.Roles,
			&user.FirstName, &user.LastName, &user.Email,
			&user.Password, &user.BirthDate, &user.CreateDate)
		//check for error
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	//check for rows error
	if err := rows.Err(); err != nil {
		// log.Println(err)
		return nil, err
	}
	return users, nil
}

// AddNewUser to postgres database. On success returns
// user unique id generate by postgres
func AddNewUser(user User) (uid string, e error) {
	var id string

	err := pgdb.DB.QueryRow(`INSERT INTO users (roles,first_name, last_name,email,password,birth_date)
	VALUES ($1,$2,$3,$4,$5,$6)
	RETURNING id createdate;`, user.Roles, user.FirstName, user.LastName, user.Email, user.Password, user.BirthDate).Scan(&id)

	if err != nil {
		return "", err
	}
	return id, nil
}
