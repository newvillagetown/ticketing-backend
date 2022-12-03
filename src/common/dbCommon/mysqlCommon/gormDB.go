package mysqlCommon

import (
	"database/sql"
	"gorm.io/gorm"
)

type GormModel struct {
	ID        string `json:"id" gorm:"primaryKey;column:id""`
	CreatedAt int64  `json:"createdAt" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt int64  `json:"updatedAt" gorm:"autoUpdateTime;column:updated_at"`
	IsDeleted bool   `json:"isDeleted" gorm:"default:false; column:is_deleted"`
}

type GormUser struct {
	GormModel GormModel `gorm:"embedded"`
	Name      string    `json:"name"  gorm:"column:name"`
	Email     string    `json:"email"  gorm:"column:email"`
}

/*
 user 테이블 생성 기존꺼 일단 그대로 두고 임시로 작업
create table gorm_users (
	id varchar(200),
	name varchar(200),
	email varchar(200),
	is_deleted TINYINT(1) not null,
	created_at varchar(200),
	updated_at varchar(200),
    PRIMARY KEY (id)
	);
*/

type GormUserAuth struct {
	ID         string `json:"id" gorm:"primaryKey; column:id"`
	CreateAt   string `json:"createAt" gorm:"autoCreateTime; column:created_at"`
	UpdatedAt  string `json:"updatedAt" gorm:"autoUpdateTime; column:updated_at"`
	Provider   string `json:"provider" gorm:"column:provider"`
	UserID     uint   `json:"userID" gorm:"column:user_id"`
	LastSignIn int64  `json:"lastSignIn" gorm:"autoUpdateTime; column:last_sign_in"`
	IsDeleted  bool   `json:"isDeleted" gorm:"default:false; column:is_deleted"`
}

/*
 userAuth 테이블 생성 기존꺼 일단 그대로 두고 임시로 작업
create table gorm_users (
	id varchar(200),
	name varchar(200),
	email varchar(200),
	is_deleted TINYINT(1) not null,
	created_at varchar(200),
	updated_at varchar(200),
    PRIMARY KEY (id)
	);

	create table gorm_user_auth (
	id int,
	provider varchar(200),
	user_id varchar(200),
	last_sign_in int,
	createdAt varchar(200),
	updatedAt varchar(200)
	);
*/

type GormProduct struct {
	gorm.Model
	LastUpdated int64        `json:"lastUpdated" gorm:"autoUpdateTime; column:last_updated"`
	IsDeleted   sql.NullBool `json:"isDeleted" gorm:"default:false; column:is_deleted"`
	Name        string       `json:"name" gorm:"column:name"`
	Description string       `json:"description" gorm:"column:description"`
	Category    string       `json:"category" gorm:"column:category"`
	PerAmount   int64        `json:"perAmount" gorm:"column:per_amount"`
	ImgUrl      string       `json:"imgUrl" gorm:"column:img_url"`
	TotalCount  int64        `json:"totalCount" gorm:"column:total_count"`
	RestCount   int64        `json:"restCount" gorm:"column:rest_count"`
	StartDate   int64        `json:"startDate" gorm:"column:start_date"`
	EndDate     int64        `json:"endDate" gorm:"column:end_date"`
}

/*
 Product 테이블 생성 기존꺼 일단 그대로 두고 임시로 작업
*/
