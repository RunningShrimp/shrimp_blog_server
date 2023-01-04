package dao

import (
	"shrimp_blog_sever_v2/model"
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
