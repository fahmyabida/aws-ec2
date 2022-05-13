package model

import "gorm.io/datatypes"

type User struct {
	Id            int    `gorm:"primaryKey" json:"id"`
	Username      string `json:"username"`
	Password      string `json:"-"`
	Name          string `json:"name"`
	Role          string `json:"role"`
}

type UserContoh struct {
	Id            int    `gorm:"primaryKey" json:"id"`
	Username      string `json:"username"`
	Password      string `json:"-"`
	Name          string `json:"name"`
	Role          string `json:"role"`
	IdJenisMuatan datatypes.JSON `gorm:column="id_jenis_muatan"`
}

func CobaInsert() {
	user := UserContoh{}
	user.IdJenisMuatan = []byte("[1,2,3,4,5]")
}

type UserCostum struct { // table user_custom
	Id       int    `json:"id"`
	Username string `json:"username"`
}

// TableName overrides the table name used by User to `profiles`
func (UserCostum) TableName() string {
	return "users"
}
