package dao

// import (
// 	"github.com/jinzhu/gorm"
// 	// for postgres
// 	"fmt"

// 	_ "github.com/jinzhu/gorm/dialects/postgres"
// )

// type DBClient struct {
// 	client *gorm.DB
// }

// func (m *DBClient) Connect() {
// 	config := getDbConfig()
// 	client, err := gorm.Open(
// 		"postgres",
// 		fmt.Sprintf(
// 			"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
// 			config.Url,
// 			config.Port,
// 			config.DbName,
// 			config.UserName,
// 			config.Password,
// 		),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	m.client = client
// }
// func (m *DBClient) Disconnect() {
// 	m.client.Close()
// }
