package infraestructure

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "os"
)

func ConnectDatabase() *gorm.DB {
    dsn := os.Getenv("DATABASE_URL")
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Error al conectar con la base de datos:", err)
    }

    return db
}
