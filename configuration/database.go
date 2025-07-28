package configuration

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/exception"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(config Config) *gorm.DB {
	username := config.Get("DATASOURCE_USERNAME")
	password := config.Get("DATASOURCE_PASSWORD")
	host := config.Get("DATASOURCE_HOST")
	port := config.Get("DATASOURCE_PORT")
	dbName := config.Get("DATASOURCE_DB_NAME")
	maxPoolOpen, err := strconv.Atoi(config.Get("DATASOURCE_POOL_MAX_CONN"))
	maxPoolIdle, err := strconv.Atoi(config.Get("DATASOURCE_POOL_IDLE_CONN"))
	maxPollLifeTime, err := strconv.Atoi(config.Get("DATASOURCE_POOL_LIFE_TIME"))
	exception.PanicLogging(err)

	loggerDb := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   loggerDb,
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	exception.PanicLogging(err)

	sqlDB, err := db.DB()
	exception.PanicLogging(err)

	sqlDB.SetMaxOpenConns(maxPoolOpen)
	sqlDB.SetMaxIdleConns(maxPoolIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPollLifeTime))) * time.Millisecond)

	//autoMigrate
	if err := autoMigrateInOrder(db); err != nil {
		log.Fatal("Auto migration failed:", err)
	}

	return db
}

func autoMigrateInOrder(db *gorm.DB) error {
	log.Println("Starting auto migration...")

	parentTables := []interface{}{
		&entity.User{},
		&entity.Kategori{},
	}

	for _, table := range parentTables {
		log.Printf("Migrating %T...", table)
		if err := db.AutoMigrate(table); err != nil {
			return err
		}
	}

	dependentTables := []interface{}{
		&entity.Toko{},
		&entity.Alamat{},
		&entity.Produk{},
		&entity.FotoProduk{},
		&entity.LogProduk{},
		&entity.Trx{},
		&entity.DetailTrx{},
	}

	for _, table := range dependentTables {
		log.Printf("Migrating %T...", table)
		if err := db.AutoMigrate(table); err != nil {
			return err
		}
	}

	log.Println("Auto migration completed successfully")
	return nil
}
