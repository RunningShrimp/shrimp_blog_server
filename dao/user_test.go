package dao

import (
	"os"
	"reflect"
	"shrimp_blog_sever_v2/config"
	"shrimp_blog_sever_v2/model"
	"testing"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	config.InitGoble("../")

	os.Exit(m.Run())
}

func TestUserDaoImpl_SelectById(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want *model.User
	}{
		{
			name: "test",
			args: args{},
			want: &model.User{
				ID:       1,
				Username: "",
				Password: "",
				Email:    "",
				Avatar:   "",
				Status:   0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDaoImpl{}
			t.Log(u.SelectById(1))
		})
	}
}

func TestUserDaoImpl_SelectByPage(t *testing.T) {
	type args struct {
		page     int
		pageSize int
	}
	tests := []struct {
		name string
		args args
		want []*model.User
	}{
		{
			name: "test",
			args: args{},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDaoImpl{}

			users := u.SelectByPage(1, 10)

			for _, user := range users {

				t.Log(*user)
			}
		})
	}
}

func TestUserDaoImpl_SelectByIds(t *testing.T) {
	type args struct {
		ids []int
	}
	tests := []struct {
		name string
		args args
		want []*model.User
	}{
		{
			name: "test",
			args: args{ids: []int{1, 2}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDaoImpl{}
			users := u.SelectByIds(tt.args.ids...)

			for _, user := range users {

				t.Log(*user)
			}
		})
	}
}

func TestUserDaoImpl_SelectStatusById(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				id: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDaoImpl{}
			if got := u.SelectStatusById(tt.args.id); got != tt.want {
				t.Errorf("SelectStatusById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDaoImpl_DeleteById(t *testing.T) {
	type args struct {
		t *model.User
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDaoImpl{}
			u.DeleteById(tt.args.t)
		})
	}
}

func TestUserDaoImpl_Insert(t *testing.T) {
	type args struct {
		t *model.User
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDaoImpl{}
			u.Insert(tt.args.t)
		})
	}
}

func TestUserDaoImpl_SelectById1(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want *model.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDaoImpl{}
			if got := u.SelectById(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDaoImpl_SelectByIds1(t *testing.T) {
	type args struct {
		ids []int
	}
	tests := []struct {
		name string
		args args
		want []*model.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDaoImpl{}
			if got := u.SelectByIds(tt.args.ids...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectByIds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDaoImpl_SelectByPage1(t *testing.T) {
	type args struct {
		page     int
		pageSize int
	}
	tests := []struct {
		name string
		args args
		want []*model.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDaoImpl{}
			if got := u.SelectByPage(tt.args.page, tt.args.pageSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectByPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDaoImpl_SelectStatusById1(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDaoImpl{}
			if got := u.SelectStatusById(tt.args.id); got != tt.want {
				t.Errorf("SelectStatusById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDaoImpl_Update(t *testing.T) {
	type args struct {
		t *model.User
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDaoImpl{}
			u.Update(tt.args.t)
		})
	}
}
