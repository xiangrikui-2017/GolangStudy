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
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

type Account struct {
	gorm.Model
	Name    string  `gorm:"type:varchar(128);not null" json:"name"`
	Code    string  `gorm:"type:varchar(128);not null" json:"code"`
	Balance float64 `gorm:"type:decimal(10,2);not null" json:"balance"`
}

type Transaction struct {
	gorm.Model
	FromAccountId string  `gorm:"type:varchar(32);not null" json:"from_account_id"`
	ToAccountId   string  `gorm:"type:varchar(32);not null" json:"to_account_id"`
	Amount        float64 `gorm:"type:decimal(10,2);not null" json:"amount"`
}

func main() {
	// 创建表
	if err := DB.AutoMigrate(&Account{}, &Transaction{}); err != nil {
		return
	}

	/**
	初始化Account表

	initAccounts := []Account{
		{Name: "张三", Code: "A1", Balance: 100},
		{Name: "王五", Code: "B1", Balance: 100},
	}
	DB.Create(&initAccounts)
	*/
	/**
	编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，
	需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
	并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
	*/
	trade := Transaction{
		FromAccountId: "A1",
		ToAccountId:   "B1",
		Amount:        100.0,
	}
	tx := DB.Begin()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			tx.Rollback()
		}
	}()
	var fromAccount Account
	if result := DB.First(&fromAccount, "code = ?", trade.FromAccountId); result.Error != nil {
		panic("付款账户查询失败")
	}
	if fromAccount.Balance < trade.Amount {
		panic("余额不足")
	}
	var toAccount Account
	if result := DB.First(&toAccount, "code = ?", trade.ToAccountId); result.Error != nil {
		panic("收账户查询失败")
	}
	aResult := DB.Model(&fromAccount).Where("code = ?", trade.FromAccountId).
		Where("balance >= ?", trade.Amount).
		Update("balance", gorm.Expr("balance - ?", trade.Amount))
	if aResult.Error != nil {
		panic(aResult.Error)
	}
	bResult := DB.Model(&toAccount).Where("code = ?", trade.ToAccountId).
		Update("balance", gorm.Expr("balance + ?", trade.Amount))
	if bResult.Error != nil {
		panic(bResult.Error)
	}
	txSave := DB.Create(&trade)
	if txSave.Error != nil {
		panic(txSave.Error)
	}

	tx.Commit()

	var accounts []Account
	query := DB.Where("code", []string{trade.FromAccountId, trade.ToAccountId}).Find(&accounts)
	if query.Error != nil {
		log.Fatal(query.Error)
	}
	fmt.Println(accounts)
}
