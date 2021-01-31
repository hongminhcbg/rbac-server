package repo

import (
	"fmt"
	"rbac/models"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const mysqlURL = "external:hongminh@tcp(192.168.100.12:3306)/rbac?charset=utf8mb4&parseTime=True&loc=Local"

func Test_userImpl_Create(t *testing.T) {

	db, err := gorm.Open(mysql.Open(mysqlURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("connect mysql success")

	type args struct {
		user models.User
	}
	tests := []struct {
		name    string
		repo    *userImpl
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "normal case",
			repo: &userImpl{
				db: db,
			},
			args: args{
				user: models.User{
					Name: "minhnh21",
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.repo.Create(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userImpl.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}
