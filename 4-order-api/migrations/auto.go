package main

import (
	"os"
	"purple/4-order-api/internal/auth"
	"purple/4-order-api/internal/order"
	"purple/4-order-api/internal/product"
	"purple/4-order-api/internal/user"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&product.Product{},
		&auth.Session{},
		&user.User{},
		&order.Order{},
	)
	DropUnusedColumns(
		db,
		&user.User{},
    )
}

func DropUnusedColumns(DB *gorm.DB, dsts ...any) {
	for _, dst := range dsts {
		stmt := &gorm.Statement{DB: DB}
		stmt.Parse(dst)
		fields := stmt.Schema.Fields
		columns, _ := DB.Debug().Migrator().ColumnTypes(dst)

		for i := range columns {
			found := false
			for j := range fields {
				if columns[i].Name() == fields[j].DBName {
					found = true
					break
				}
			}
			if !found {
				DB.Migrator().DropColumn(dst, columns[i].Name())
			}
		}
	} 
}