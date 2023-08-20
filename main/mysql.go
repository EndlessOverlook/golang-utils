package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_learn")
	if err != nil {
		panic(err)
	}
	// 最大空闲连接数，默认不配置，是2个最大空闲连接
	db.SetMaxIdleConns(5)
	// 最大连接数，默认不配置，是不限制最大连接数
	db.SetMaxOpenConns(100)
	// 连接最大存活时间
	db.SetConnMaxLifetime(time.Minute * 3)
	// 空闲连接最大存活时间
	db.SetConnMaxIdleTime(time.Minute * 1)
	err = db.Ping()
	if err != nil {
		log.Println("数据库连接失败")
		db.Close()
		panic(err)
	}
	DB = db

}

func save() {
	r, err := DB.Exec("insert into user (username,sex,email) values(?,?,?)", "test001", "man", "001@test.com")
	if err != nil {
		log.Println("执行sql语句出错")
		panic(err)
	}
	id, err := r.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("插入成功:", id)
}
func main() {
	defer DB.Close()
	save()
}

type User struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

func query(id int) (*User, error) {
	rows, err := DB.Query("select * from user where user_id=? limit 1", id)
	if err != nil {
		log.Println("查询出现错误:", err)
		return nil, errors.New(err.Error())
	}
	user := new(User)
	for rows.Next() {
		if err := rows.Scan(&user.UserId, &user.Username, &user.Sex, &user.Email); err != nil {
			log.Println("scan error:", err)
			return nil, errors.New(err.Error())
		}
	}
	return user, nil
}

func update(username string, id int) {
	ret, err := DB.Exec("update user set username=? where user_id=?", username, id)
	if err != nil {
		log.Println("更新出现问题:", err)
		return
	}
	affected, _ := ret.RowsAffected()
	fmt.Println("更新成功的行数:", affected)
}

func delete(id int) {
	ret, err := DB.Exec("delete from user where user_id=?", id)
	if err != nil {
		log.Println("删除出现问题:", err)
		return
	}
	affected, _ := ret.RowsAffected()
	fmt.Println("删除成功的行数:", affected)
}

func insertTx(username string) {
	tx, err := DB.Begin()
	if err != nil {
		log.Println("开启事务错误:", err)
		return
	}
	ret, err := tx.Exec("insert into user (username,sex,email) values (?,?,?)", username, "man", "test@test.com")
	if err != nil {
		log.Println("事务sql执行出错:", err)
		return
	}
	id, _ := ret.LastInsertId()
	fmt.Println("插入成功:", id)
	if username == "lisi" {
		fmt.Println("回滚...")
		_ = tx.Rollback()
	} else {
		_ = tx.Commit()
	}

}
