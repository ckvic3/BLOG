// 负责处理数据库相关的事务

package model

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

type Blog struct {
	Model
	Title   string `json:"title"`
	Content string `json:"content"`
}

func init() {
	var err error
	var dbType, dbName, user, password, host, tablePrefix string

	dbType = "mysql"
	dbName = "blog"
	user = "root"
	password = "root"
	host = "127.0.0.1"
	tablePrefix = "blog_"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, dbName)
	fmt.Println(dsn)

	db, err = gorm.Open(dbType, dsn)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("初始化数据库成功......")
	}

	fmt.Printf("%s\n", db.Dialect().GetName())

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

}

// 获取博文的总数量
func GetBlogsNum() (count int, err error) {

	err = db.Model(&Blog{}).Count(&count).Error
	return
}

// pageNum 表示当前页数
// pageSize 表示一页可展示的博文数量
func GetBlogs(pageNum int, pageSize int) (blogs []Blog, count int) {
	var err error

	err = db.Model(&Blog{}).Count(&count).Error
	if err != nil {
		log.Println("models fun GetBlogs Failed!")
		panic(err)
	}
	offset := pageNum * pageSize

	if offset > count {
		return
	}
	err = db.Offset(offset).Limit(pageSize).Find(&blogs).Error
	if err != nil {
		log.Println("models fun GetBlogs Failed!")
		panic(err)

	}

	return
}

// 根据id返回对应的博文内容
func GetBlog(id int) (blog Blog, err error) {

	err = db.Where("id = ?", id).First(&blog).Error
	//if err != nil {
	//	log.Println("models fun GetBlog Failed!")
	//}
	return
}

//
func CreateBlog(data map[string]interface{}) bool {

	var err error
	if data["id"] == 0 {
		err = db.Model(&Blog{}).Create(&Blog{Title: data["title"].(string),
			Content: data["content"].(string)}).Error
	} else {
		err = db.Model(&Blog{}).Where("id=?", data["id"]).Update(data).Error
	}
	if err != nil {
		panic(err)
	}
	return true
}

// 删除id对应的博文
func DeleteBlog(id int) bool {
	err := db.Where("id = ?", id).Delete(&Blog{}).Error
	if err != nil {
		panic(err)
	}
	return true
}

func UpdataBlog(id int, data map[string]interface{}) bool {

	err := db.Model(&Blog{}).Where("id = ?", id).Update(data).Error

	if err != nil {
		panic(err.Error())
	}
	return true
}

func (blog *Blog) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedOn", time.Now().Unix())
	return err
}

func (blog *Blog) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("ModifiedOn", time.Now().Unix())
	return err
}
