package main

import (
  "log"
  "os"

  _ "github.com/go-sql-driver/mysql"
  "github.com/jmoiron/sqlx"
  "github.com/subosito/gotenv"
)

var create_table_virtual_product_agents = "CREATE TABLE virtual_product_agents (" +
  "`id` INT(11) NOT NULL AUTO_INCREMENT," +
  "`deleted` TINYINT(1) DEFAULT 0," +
  "`user_id` INT(11) DEFAULT NULL," +
  "`ktp` VARCHAR(255) DEFAULT NULL," +
  "`status` TINYINT(4) DEFAULT 0," +
  "`referrer_id` INT(11) DEFAULT NULL," +
  "`cashback_id` INT(11) DEFAULT NULL," +
  "`reason` VARCHAR(255) DEFAULT NULL," +
  "`referrer_commission_paid_at` DATETIME DEFAULT NULL," +
  "`kyc_verified_by` VARCHAR(255) DEFAULT NULL," +
  "`kyc_verified_at` DATETIME DEFAULT NULL," +
  "`kyc_rejected_by` INT(11) DEFAULT NULL," +
  "`kyc_rejected_at` DATETIME DEFAULT NULL," +
  "`kyc_rejection_reason` VARCHAR(255) DEFAULT NULL," +
  "`agent_type` VARCHAR(50) DEFAULT NULL," +
  "`register_from` VARCHAR(20) DEFAULT NULL," +
  "`kyc_from` VARCHAR(20) DEFAULT NULL," +
  "`joined_at` DATETIME DEFAULT NULL," +
  "`created_at` DATETIME DEFAULT NULL," +
  "`updated_at` DATETIME DEFAULT NULL," +
  "PRIMARY KEY (`id`)" +
  ") ENGINE = InnoDB DEFAULT CHARSET = utf8"

func main() {
  gotenv.Load()
  port := os.Getenv("DATABASE_PORT")
  if port == "" {
    port = "3306"
  }

  dataSourceName := os.Getenv("DATABASE_USERNAME") + ":" + os.Getenv("DATABASE_PASSWORD") + "@(" + os.Getenv("DATABASE_HOST") + ":" + (string(port)) + ")/" + os.Getenv("DATABASE_NAME")
  log.Println("CONNECTING TO:", os.Getenv("DATABASE_HOST")+":"+(string(port))+"/"+os.Getenv("DATABASE_NAME"))
  db, err := sqlx.Open("mysql", dataSourceName)
  defer db.Close()

  log.Println("MIGRATING...")
  if err != nil {
    panic(err.Error())
  }

  db.MustExec(create_table_virtual_product_agents)
  log.Println("DONE")
}
