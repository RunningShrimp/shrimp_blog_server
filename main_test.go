package main

import (
	"os"
	"shrimp_blog_sever/app"
	"shrimp_blog_sever/config"
	_ "shrimp_blog_sever/config"
	"shrimp_blog_sever/model"
	"testing"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	config.Init(".")

	os.Exit(m.Run())
}

func TestDatabase(t *testing.T) {
	rows, err := app.DBClient.Queryx("select * from user")
	if err != nil {
		t.Error(err)
	}

	for rows.Next() {
		result := make(map[string]interface{})
		if err := rows.MapScan(result); err != nil {
			t.Error(err)
		}
		t.Log(result)
	}
}

func TestGetUser(t *testing.T) {
	rows, err := app.DBClient.Queryx("select * from user")

	if err != nil {
		t.Error(err)
	}

	for rows.Next() {
		result := make(map[string]interface{})
		if err := rows.MapScan(result); err != nil {
			t.Error(err)
		}
		user := model.User{}
		rows.StructScan(&user)

		//t.Log(result)
		//
		//user.MapToStruct(result)

		t.Log(user.DeleteTime.IsZero())
	}
}
