package model

import (
	"defaultproject/util"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database(connString string) {
	// 初始化GORM日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level(这里记得根据需求改一下)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: newLogger,
	})
	// Error
	if connString == "" || err != nil {
		util.Log().Error("mysql lost: %v", err)
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		util.Log().Error("mysql lost: %v", err)
		panic(err)
	}

	//设置连接池
	//空闲
	sqlDB.SetMaxIdleConns(1)
	//打开
	sqlDB.SetMaxOpenConns(1)
	DB = db

	fmt.Println("DB initialization complete")
	// migration()
}

func PostgresDB(connString string) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level(这里记得根据需求改一下)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		Logger: newLogger,
	})
	// Error
	if connString == "" || err != nil {
		util.Log().Error("postgres lost: %v", err)
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		util.Log().Error("postgres lost: %v", err)
		panic(err)
	}

	// 设置连接池
	// 空闲
	sqlDB.SetMaxIdleConns(1)
	// 打开
	sqlDB.SetMaxOpenConns(1)
	DB = db

	fmt.Println("DB initialization complete")
	migration()
	insertSampleData(db)
}

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func insertSampleData(db *gorm.DB) error {
	// 检查表中是否已有数据
	var count int64
	if err := db.Model(&Book{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		books := []Book{
			{Title: "The Go Programming Language", AvailableCopies: 5},
			{Title: "Clean Code: A Handbook of Agile Software Craftsmanship", AvailableCopies: 3},
			{Title: "Design Patterns: Elements of Reusable Object-Oriented Software", AvailableCopies: 2},
			{Title: "The Pragmatic Programmer: Your Journey to Mastery", AvailableCopies: 4},
			{Title: "Introduction to Algorithms", AvailableCopies: 1},
			{Title: "You Don't Know JS: Up & Going", AvailableCopies: 6},
			{Title: "Effective Java", AvailableCopies: 3},
			{Title: "Refactoring: Improving the Design of Existing Code", AvailableCopies: 2},
			{Title: "Domain-Driven Design: Tackling Complexity in the Heart of Software", AvailableCopies: 4},
			{Title: "The Mythical Man-Month: Essays on Software Engineering", AvailableCopies: 1},
		}
		if err := db.Create(&books).Error; err != nil {
			return err
		}
		fmt.Println("Inserted 10 sample books.")
	} else {
		fmt.Println("Table is not empty, skipping sample data insertion.")
	}
	return nil
}
