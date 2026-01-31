package configs

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var onceDb sync.Once

var customLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags),
	logger.Config{
		LogLevel: logger.Info,
	},
)

func InitDB() {
	onceDb.Do(func() {
		dbhost := os.Getenv("dbhost")
		dbname := os.Getenv("dbname")
		dbpassword := os.Getenv("dbpassword")
		dbusername := os.Getenv("dbusername")
		dbport := os.Getenv("dbport")
		// dsn := "host=" + dbhost + " user=" + dbusername + " password=" + dbpassword + " dbname=" + dbname + " port=" + dbport
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
			dbhost,
			dbusername,
			dbpassword,
			dbname,
			dbport,
		)

		var db *gorm.DB
		var err error

		connectingToDb := func() error {
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
				PrepareStmt: false, // fix error karena pooler supabase
				Logger:      customLogger,
				NamingStrategy: schema.NamingStrategy{
					SingularTable: true,
				},
			})

			if err != nil {
				log.Fatal("Error when try connecting to database: ", err.Error())
				return err
			}

			// Active debug mode
			// db.Debug()

			DB = db
			return nil
		}

		b := backoff.NewExponentialBackOff()
		b.MaxElapsedTime = 15 * time.Second

		err = backoff.Retry(connectingToDb, b)
		if err != nil {
			log.Fatalln("Max retry reached, cannot connect to the database: ", err.Error())
			panic(err)
		}
	})
}

func GetDB() *gorm.DB {
	return DB
}
