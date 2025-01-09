package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // Import driver MySQL
	"log"
)

var DB *sql.DB

// InitDB khởi tạo kết nối cơ sở dữ liệu
func InitDB() {
	var err error

	// Thay đổi thông tin kết nối phù hợp với MySQL của bạn
	dsn := "root:123456@tcp(localhost:3306)/todo_app"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Không thể kết nối đến MySQL: %v", err)
	}

	// Kiểm tra kết nối
	if err = DB.Ping(); err != nil {
		log.Fatalf("Không thể ping MySQL: %v", err)
	}

	fmt.Println("Kết nối MySQL thành công!")
}
