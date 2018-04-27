package db

import (
  "os"
  "strconv"

  _ "github.com/go-sql-driver/mysql" // for sqlx
  "github.com/jmoiron/sqlx"
  "github.com/subosito/gotenv"
)

func InitMysql() (*sqlx.DB, error) {
  gotenv.Load()
  dbPort := os.Getenv("DATABASE_PORT")
  if dbPort == "" {
    dbPort = "3306"
  }

  dbPoolSize, _ := strconv.Atoi(os.Getenv("DATABASE_POOL"))
  if dbPoolSize == 0 {
    dbPoolSize = 10
  }

  dataSourceName := os.Getenv("DATABASE_USERNAME") + ":" + os.Getenv("DATABASE_PASSWORD") + "@(" + os.Getenv("DATABASE_HOST") + ":" + (string(dbPort)) + ")/" + os.Getenv("DATABASE_NAME") + "?parseTime=true"
  db, err := sqlx.Open("mysql", dataSourceName)

  if err != nil {
    return nil, err
  }

  db.SetMaxIdleConns(dbPoolSize)

  err = db.Ping()
  if err != nil {
    return nil, err
  }

  return db, err
}
