package infraestructure

import (
	"fmt"
	"log"
	"time"

	config "github.com/JuhethAriza/inventory/src/common/config"
	productEntities "github.com/JuhethAriza/inventory/src/modules/Producto/domain/entities"
	userEntities "github.com/JuhethAriza/inventory/src/modules/User/domain/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConnection struct {
	*gorm.DB
}

func NewDBConnection(cfg *config.Config) *DBConnection {
	fmt.Println(cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Dbname)

	// DSN con parámetros mejorados para conexiones más robustas
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=30s",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Dbname,
	)

	fmt.Println(dsn)

	// Intentar conectar con retry
	var db *gorm.DB
	var err error
	maxRetries := 5
	retryDelay := 2 * time.Second

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			// Verificar que la conexión funciona
			sqlDB, err := db.DB()
			if err == nil {
				err = sqlDB.Ping()
				if err == nil {
					// Configurar conexión pool
					sqlDB.SetMaxIdleConns(10)
					sqlDB.SetMaxOpenConns(100)
					sqlDB.SetConnMaxLifetime(time.Hour)
					break
				}
			}
		}
		
		if i < maxRetries-1 {
			log.Printf("Error connecting to database (attempt %d/%d): %s. Retrying in %v...", i+1, maxRetries, err, retryDelay)
			time.Sleep(retryDelay)
		}
	}

	if err != nil {
		log.Fatalf("Error to connect database after %d attempts: %s", maxRetries, err)
	}

	log.Println("Successfully connected to database")

	// AutoMigrate crea las tablas automáticamente basándose en las entidades
	err = db.AutoMigrate(
		&userEntities.User{},
		&productEntities.Product{},
	)
	if err != nil {
		log.Fatalf("Error migrating database : %s", err)
	}
	
	log.Println("Database migration completed successfully")

	return &DBConnection{db}
}
