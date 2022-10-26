package app

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/fahmyabdul/golibs"
	"github.com/fahmyabdul/golibs/databases"

	"github.com/gomodule/redigo/redis"
	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/rintik-io/rintik-auth/configs"
)

type Databases struct {
	RedisPool   *redis.Pool
	MongoConn   *mongo.Database
	PostgreGORM *databases.PostgreCoreGORM
	SQLiteConn  *sql.DB
}

// DatabaseConnection :
func (a *Application) DatabaseConnection() error {
	a.Databases = Databases{}

	if configs.Properties.Databases.Redis.Status {
		err := a.RedisConnection()
		if err != nil {
			return err
		}
	}

	if configs.Properties.Databases.Mongo.Status {
		err := a.MongoConnection()
		if err != nil {
			return err
		}
	}

	if configs.Properties.Databases.Postgre.Status {
		err := a.PostgreConnection()
		if err != nil {
			return err
		}
	}

	if configs.Properties.Databases.Sqlite.Status {
		err := a.SQLiteConnection()
		if err != nil {
			return err
		}
	}

	return nil
}

// DatabaseConnectionClose :
func (a *Application) DatabaseConnectionClose() {
	if a.Databases.RedisPool != nil {
		log.Println("| Closing Redis Connection")
		a.Databases.RedisPool.Close()
	}

	if a.Databases.MongoConn != nil {
		log.Println("| Closing MongoDB Connection")
		a.Databases.MongoConn.Client().Disconnect(context.TODO())
	}

	if a.Databases.PostgreGORM != nil {
		log.Println("| Closing PostgreSQL GORM Connection")
		a.Databases.PostgreGORM.Close()
	}
}

// MongoConnection : connection to mongo
func (a *Application) MongoConnection() error {
	confMongo := configs.Properties.Databases.Mongo
	mongoConn, err := databases.NewMongo(
		confMongo.Host,
		confMongo.User,
		confMongo.Pass,
		confMongo.DB,
		confMongo.Srv,
		confMongo.Cluster,
		confMongo.RsName,
		golibs.Log,
	)
	if err != nil {
		golibs.Log.Println("| Mongo | Connection | Failed |", err.Error())
		return err
	}

	a.Databases.MongoConn = mongoConn

	return nil
}

// RedisConnection : Redis
func (a *Application) RedisConnection() error {
	confRedis := configs.Properties.Databases.Redis
	redisPool, err := databases.NewRedis(
		confRedis.Host,
		confRedis.Auth,
		confRedis.DB,
		confRedis.MaxIdle,
		confRedis.MaxActive,
		golibs.Log,
	)
	if err != nil {
		golibs.Log.Println("| Redis | Connection | Error |", err.Error())
		return err
	}

	a.Databases.RedisPool = redisPool

	golibs.Log.Println("| Redis | Connection | Success")

	return nil
}

// PostgreConnection :
func (a *Application) PostgreConnection() error {
	confPostgre := configs.Properties.Databases.Postgre
	gormPostgre, err := databases.NewPostgreGORM(
		confPostgre.Host,
		confPostgre.Port,
		confPostgre.User,
		confPostgre.Pass,
		confPostgre.DB,
		golibs.Log,
	)
	if err != nil {
		golibs.Log.Println("| PostgreSQL GORM | Connection | Error |", err.Error())
		return err
	}

	a.Databases.PostgreGORM = gormPostgre

	golibs.Log.Println("| PostgreSQL GORM | Connection | Success")

	return nil
}

// SQLiteConnection:
func (a *Application) SQLiteConnection() error {
	confSqlite := configs.Properties.Databases.Sqlite

	err := os.MkdirAll(confSqlite.Path, os.ModePerm)
	if err != nil {
		golibs.Log.Println("| SQLite | Connection | Error |", err.Error())
		return err
	}

	file_path := filepath.Join(confSqlite.Path, confSqlite.File)
	db, err := sql.Open("sqlite3", file_path)
	if err != nil {
		golibs.Log.Println("| SQLite | Connection | Error |", err.Error())
		return err
	}

	a.Databases.SQLiteConn = db

	golibs.Log.Printf("| SQLite | Connection | Success, db : %s\n", file_path)

	return nil
}
