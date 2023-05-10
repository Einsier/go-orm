package main

import (
	"fmt"

	goorm "github.com/einsier/go-orm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e, _ := goorm.NewEngine("sqlite3", "sample.db")
	defer e.Close()

	s := e.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	res, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := res.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
