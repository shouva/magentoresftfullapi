package main

import "github.com/jinzhu/gorm"

var g *gorm.DB

func open(host, port, dbname, username, password string) (oGorm *gorm.DB, err error) {
	oGorm, err = gorm.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	return
}
