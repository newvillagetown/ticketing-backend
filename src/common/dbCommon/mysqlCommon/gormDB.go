package mysqlCommon

//TODO 추후 테이블 생성하는 실행파일 만들어야됨.

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
	created_at int,
	updated_at int,
    PRIMARY KEY (id)
	);
*/

type GormUserAuth struct {
	GormModel  GormModel `gorm:"embedded"`
	Provider   string    `json:"provider" gorm:"column:provider"`
	UserID     string    `json:"userID" gorm:"column:user_id"`
	LastSignIn int64     `json:"lastSignIn" gorm:"column:last_sign_in"`
}

/*
 userAuth 테이블 생성 기존꺼 일단 그대로 두고 임시로 작업

create table gorm_users (
	id varchar(200),
	provider varchar(200) not null,
	user_id varchar(200) not null,
	last_sign_in int,
	is_deleted TINYINT(1) not null,
	created_at int,
	updated_at int,
    PRIMARY KEY (id),
	FOREIGN KEY (user_id)
	REFERENCES gorm_user(id) on update cascade
	);
*/

type GormProduct struct {
	GormModel   GormModel `gorm:"embedded"`
	Name        string    `json:"name" gorm:"column:name"`
	Description string    `json:"description" gorm:"column:description"`
	Category    string    `json:"category" gorm:"column:category"`
	PerAmount   int64     `json:"perAmount" gorm:"column:per_amount"`
	ImgUrl      string    `json:"imgUrl" gorm:"column:img_url"`
	TotalCount  int64     `json:"totalCount" gorm:"column:total_count"`
	RestCount   int64     `json:"restCount" gorm:"column:rest_count"`
	StartDate   int64     `json:"startDate" gorm:"column:start_date"`
	EndDate     int64     `json:"endDate" gorm:"column:end_date"`
}

/*
 Product 테이블 생성 기존꺼 일단 그대로 두고 임시로 작업
	create table gorm_products (
	id varchar(200),
	created_at int,
	updated_at int,
	is_deleted TINYINT(1) not null,
	name varchar(200) not null,
	description text,
	category varchar(200) not null,
	per_amount int not null,
	img_url varchar(1000),
	total_count int not null,
	rest_count int not null,
	start_date int not null,
	end_date int not null,
    PRIMARY KEY (id),
	);

*/
