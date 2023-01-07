package dao

import (
	"shrimp_blog_sever/model"
	"testing"
)

func TestGenSelectFiled(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "getString",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(genSelectFiledString[model.User]())
		})
	}
}

func Test_genInsertValueString(t *testing.T) {
	user := model.User{
		ID:         0,
		Username:   "sadad",
		Password:   "dasas",
		Email:      "sasda",
		Avatar:     "sada",
		CreateTime: nil,
		UpdateTime: nil,
		DeleteTime: nil,
		Status:     0,
	}

	t.Log(genInsertValueAndPlaceholder[model.User](user))
}

func Test_genUpdateFieldString(t *testing.T) {
	user := model.User{
		ID:         0,
		Username:   "sadad",
		Password:   "dasas",
		Email:      "sasda",
		Avatar:     "sada",
		CreateTime: nil,
		UpdateTime: nil,
		DeleteTime: nil,
		Status:     0,
	}

	t.Log(genUpdateFieldString[model.User](user))
}
