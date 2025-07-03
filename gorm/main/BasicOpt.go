package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

var DB *gorm.DB

func init() {
	dsn := "sams_test:$1qaz2wsx$@tcp(rm-2ze9584p1u9q7tq5svo.mysql.rds.aliyuncs.com:3306)/vision_cloud_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "t_",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

type Student struct {
	gorm.Model

	Name  string `gorm:"type:varchar(255);not null" json:"name"`
	Age   int    `gorm:"type:int;not null" json:"age"`
	Grade string `gorm:"type:varchar(50)" json:"grade"`
}

type Employee struct {
	gorm.Model
	Name       string  `gorm:"type:varchar(128);not null" json:"name"`
	Department string  `gorm:"type:varchar(128);not null" json:"department"`
	Salary     float64 `gorm:"type:decimal(10,2)" json:"salary"`
}

func main() {
	DB.AutoMigrate(&Student{}, &Employee{}, &Book{})
	// EmployeeOpt()
	//BookOpt()
	BlogOpt()
}

func StudentOpt() {
	// 参数定义
	//student := Student{
	//	Name:  "李四",
	//	Age:   15,
	//	Grade: "三年级",
	//}
	//
	//// 插入记录
	//createResult := DB.Create(&student)
	//if createResult.Error != nil {
	//	log.Fatal(createResult.Error)
	//}

	// 查询记录
	var students []Student
	queryResult := DB.Where("age > ?", 18).Find(&students)
	if queryResult.Error != nil {
		log.Fatal(queryResult.Error)
	}
	fmt.Printf("SQL语句: %v\n", queryResult.Statement.SQL.String())
	fmt.Printf("查询结果: %v", students)

	// 更新记录
	updateResult := DB.Model(&Student{}).Where("grade = ?", "三年级").
		UpdateColumn("grade", "四年级")
	if updateResult.Error != nil {
		log.Fatal(updateResult.Error)
	}
	fmt.Printf("更新成功，影响行数: %d\n", updateResult.RowsAffected)

	// 删除记录
	deletResult := DB.Where("age < ?", 18).Delete(&Student{})
	if deletResult.Error != nil {
		log.Fatal(deletResult.Error)
	}
}

func EmployeeOpt() {
	//ems := []Employee{
	//	{Name: "张三", Department: "技术部1", Salary: 10000.00},
	//	{Name: "李四", Department: "技术部1", Salary: 12000.00},
	//	{Name: "王五", Department: "技术部2", Salary: 15000.00},
	//	{Name: "赵六", Department: "技术部3", Salary: 13000.00},
	//}
	//DB.Create(&ems)

	rows, err := DB.Model(&Employee{}).Where("department = ?", "技术部1").Rows()
	if err != nil {
		log.Fatal(err)
	}
	var employees []Employee
	for rows.Next() {
		var employee Employee
		if err := DB.ScanRows(rows, &employee); err != nil {
			log.Fatal(err)
		}
		employees = append(employees, employee)
	}
	fmt.Println(employees)
	rows.Close()

	var maxSalaryEmps []Employee
	DB.Where("salary = (?)", DB.Table("t_employee").Select("MAX(salary) AS max_salary")).Find(&maxSalaryEmps)
	fmt.Println(maxSalaryEmps)
}

type Book struct {
	gorm.Model
	Title  string  `gorm:"type:varchar(255);not null" json:"title"`
	Author string  `gorm:"type:varchar(32);not null" json:"author"`
	Price  float64 `gorm:"type:decimal(10,2);not null" json:"price"`
}

func BookOpt() {
	//books := []Book{
	//	{Title: "毛选第一版", Author: "张三", Price: 129.00},
	//	{Title: "毛选第二版", Author: "张三", Price: 129.00},
	//	{Title: "毛选第三版", Author: "张三", Price: 129.00},
	//	{Title: "毛选经典版", Author: "张三", Price: 229.00},
	//	{Title: "毛选典藏版", Author: "张三", Price: 329.00},
	//}
	//DB.Create(&books)

	var books []Book
	result := DB.Where("price > ?", 129).Find(&books)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println(books)
}

type User struct {
	gorm.Model
	Name    string `gorm:"type:varchar(128);not null" json:"name"`
	Age     int    `gorm:"type:int;not null" json:"age"`
	Posts   []Post `gorm:"foreignKey:UserID"`
	PostNum uint   `gorm:"not null;default:0" json:"post_num"`
}

type Post struct {
	gorm.Model
	Title      string `gorm:"type:varchar(128);not null" json:"title"`
	Content    string `gorm:"type:text;not null" json:"content"`
	UserID     uint
	User       User      `gorm:"foreignKey:UserID"`
	Comments   []Comment `gorm:"foreignKey:PostID"`
	CommentNum uint      `gorm:"not null;default:0" json:"comment_num"`
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("post_num", gorm.Expr("post_num + ?", 1))
	return
}

type Comment struct {
	gorm.Model
	PostID  uint
	Post    Post   `gorm:"foreignKey:PostID"`
	Content string `gorm:"type:varchar(128);not null" json:"content"`
	UserID  uint
	User    User `gorm:"foreignKey:UserID"`
}

func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(&Post{}).Where("id = ?", c.PostID).UpdateColumn("comment_num", gorm.Expr("comment_num + ?", 1))
	return
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	tx.Model(&Post{}).Where("id = ? and commnet_num > 0", c.PostID).UpdateColumn("comment_num", gorm.Expr("comment_num - ?", 1))
	return
}

func BlogOpt() {
	DB.Migrator().DropTable(&User{}, &Post{}, &Comment{})
	err := DB.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		return
	}

	users := []User{
		{Name: "张三", Age: 20},
		{Name: "李四", Age: 20},
		{Name: "王五", Age: 22},
		{Name: "赵六", Age: 24},
	}
	DB.Create(&users)
	posts := []Post{
		{Title: "Post 1", Content: "Content 1", UserID: 1},
		{Title: "Post 2", Content: "Content 2", UserID: 1},
		{Title: "Post 3", Content: "Content 3", UserID: 2},
		{Title: "Post 4", Content: "Content 4", UserID: 3},
	}
	DB.Create(&posts)
	comments := []Comment{
		{Content: "Comment 1", PostID: 1, UserID: 4},
		{Content: "Comment 2", PostID: 1, UserID: 4},
		{Content: "Comment 3", PostID: 2, UserID: 4},
		{Content: "Comment 4", PostID: 1, UserID: 4},
	}
	result := DB.Create(&comments)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	var myPosts []Post
	myPostsResult := DB.Where(&Post{UserID: 1}).Preload("Comments").Find(&myPosts)
	if myPostsResult.Error != nil {
		log.Fatal(myPostsResult.Error)
	}
	for _, post := range posts {
		fmt.Println(post)
	}

	var maxCommentPosts []Post
	DB.Model(&Post{}).Where("comment_num = (?)", DB.Table("t_post").Select("MAX(comment_num) maxCommentNum")).Find(&maxCommentPosts)
	fmt.Println(maxCommentPosts)
}
