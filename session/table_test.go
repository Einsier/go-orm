package session

import "testing"

type User struct {
	Name string `goorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	s := NewSession().Model(&User{})
	s.DropTable()
	s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}
