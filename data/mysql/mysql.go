package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitSqlite() {
	var err error
	db, err = sql.Open("sqlite3", "../../test.db")
	if err != nil {
		fmt.Println(err)
	}
	stmt, err := db.Prepare("INSERT INTO USER_INFO(username,uid,sex,age,phone) values(?,?,?,?,?)")
	if err != nil {
		fmt.Println(err)
	}
	_, err = stmt.Exec("baskwsos", "222222", "male", 18, "13333333333")
	if err != nil {
		fmt.Println(err)
	}
	// id, err := res.LastInsertId()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(id)
	// //更新数据
	// stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// res, err = stmt.Exec("astaxieupdate", id)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// affect, err := res.RowsAffected()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(affect)
}

func Select() {
	rows, err := db.Query("SELECT * FROM USER_INFO")
	if err != nil {
		fmt.Println("no data")
	}
	for rows.Next() {
		var uid int
		var username string
		var phone string
		var sex string
		var age int
		var id int

		err := rows.Scan(&id, &username, &uid, &sex, &age, &phone)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("uid", uid)
		fmt.Println("username", username)
		fmt.Println("sex", sex)
		fmt.Println("phone", phone)
	}
}
