package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"database/sql"
)

var DB *gorm.DB

//初始化数据库
func InitDB() (*gorm.DB, error) {
	if DB != nil {
		return DB, nil
	}
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=LifeLong sslmode=disable password=postgres")
	if err == nil {
		DB = db
		db.SingularTable(true)
		db.AutoMigrate(&Userinfo{}, &Smslog{}, &Wish{})
		return db, err
	}
	return nil, err
}


//创建
func Insert(model interface{}) error {
	InitDB()
	return DB.Create(model).Error
}

//删除
func Delete(model interface{}) error {
	InitDB()
	return DB.Delete(model).Error
}

//更新
func Update(model interface{}) error {
	InitDB()
	return DB.Update(model).Error
}

//通过sql语句获取数据
func QueryBySQL(sql string) (*sql.Rows, error) {
	InitDB()
	rows, err := DB.Raw(sql).Rows() // (*sql.Rows, error)
	defer rows.Close()
	if err == nil {
		return rows, err
	}
	return nil, err
}

//执行sql
func ExecBySQL(sql string) (error) {
	InitDB()
	return DB.Exec(sql).Error
}

