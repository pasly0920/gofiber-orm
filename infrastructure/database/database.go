package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDB() (*gorm.DB, error) {
	//viper.SetConfigName("database")     // 파일명 설정
	//viper.AddConfigPath("../../config") // 파일 경로 설정
	//viper.SetConfigType("yaml")         // 설정 파일 형식 설정
	//
	//if err := viper.ReadInConfig(); err != nil {
	//	return nil, err
	//}
	//
	//host := viper.GetString("host")
	//port := viper.GetInt("port")
	//username := viper.GetString("username")
	//password := viper.GetString("password")
	//database := viper.GetString("database")
	//
	//// Connect to the database
	//dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, database)

	dsn := "host=localhost user=postgres password=password dbname=identity port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %s", err.Error())
		return nil, err
	}

	return db, nil
}
