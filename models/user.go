package models

import (
	"crud/db"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Roles    int    `json:"roles"`
}

func FetchAllUser() (Response, error) {
	var obj User
	var arrobj []User
	var res Response

	con := db.CreateConf()

	sqlStatement := "SELECT * FROM user"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.ID, &obj.Username, &obj.Email, &obj.Roles)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func StoreUser(username string, email string, roles string) (Response, error) {
	var res Response

	con := db.CreateConf()

	sqlStatement := "INSERT user (username, email, roles) VALUES (?, ?, ?)"

	stmnt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmnt.Exec(username, email, roles)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdateUser(id int, username string, email string, roles string) (Response, error) {
	var res Response

	con := db.CreateConf()

	sqlStatement := "UPDATE user SET username = ?, email = ?, roles = ? WHERE id = ?"

	stmnt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmnt.Exec(username, email, roles, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteUser(id int) (Response, error) {
	var res Response

	con := db.CreateConf()

	sqlStatement := "DELETE FROM user WHERE id = ?"

	stmnt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmnt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
