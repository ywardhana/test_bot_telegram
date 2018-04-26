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
  "PRIMARY KEY (`id`)," +
  "KEY `index agent_on_created_at` (`created_at`)," +
  "KEY `index agent_on_deleted_and_joined_at` (`deleted`, `joined_at`)," +
  "KEY `index agent_on_ktp` (`ktp`)," +
  "KEY `index agent_on_referrer_id` (`referrer_id`)," +
  "KEY `index agent_on_updated_at` (`updated_at`)," +
  "KEY `index agent_on_user_id` (`user_id`)," +
  "KEY `index agent_on_kyc_verified_by_and_kyc_verified_at` (`kyc_verified_by`, `kyc_verified_at`)," +
  "KEY `index agent_on_kyc_rejected_by_and_kyc_rejected_at` (`kyc_rejected_by`, `kyc_rejected_at`)," +
  "KEY `index agent_on_register_from` (`register_from`)," +
  "KEY `index agent_on_kyc_from` (`kyc_from`)" +
  ") ENGINE = InnoDB DEFAULT CHARSET = utf8"

var create_table_agent_coordinators = "CREATE TABLE agent_coordinators (" +
  "`id` INT(11) NOT NULL AUTO_INCREMENT," +
  "`agent_id` INT(11) DEFAULT NULL," +
  "`created_at` DATETIME DEFAULT NULL," +
  "`updated_at` DATETIME DEFAULT NULL," +
  "PRIMARY KEY (`id`)," +
  "KEY `index agent_coordinator_on_agent_id` (`agent_id`)," +
  "KEY `index agent_coordinator_on_updated_at` (`updated_at`)," +
  "KEY `index agent_coordinator_on_created_at` (`created_at`)" +
  ") ENGINE = InnoDB DEFAULT CHARSET = utf8"

var create_table_agent_images = "CREATE TABLE agent_images (" +
  "`id` INT(11) NOT NULL AUTO_INCREMENT," +
  "`agent_id` INT(11) DEFAULT NULL," +
  "`type` VARCHAR(255) DEFAULT NULL," +
  "`replaceable` TINYINT(1) DEFAULT 0," +
  "`data_file_name` VARCHAR(255) DEFAULT NULL," +
  "`data_content_type` VARCHAR(255) DEFAULT NULL," +
  "`data_file_size` INT(11) DEFAULT NULL," +
  "`data_updated_at` DATETIME DEFAULT NULL," +
  "`created_at` DATETIME DEFAULT NULL," +
  "`updated_at` DATETIME DEFAULT NULL," +
  "PRIMARY KEY (`id`)," +
  "KEY `index agent_images_on_agent_id` (`agent_id`)," +
  "KEY `index agent_images_on_agent_id_and_type` (`agent_id`, `type`)," +
  "KEY `index agent_images_on_agent_id_and_type_and_replaceable` (`agent_id`, `type`, `replaceable`)," +
  "KEY `index agent_images_on_updated_at` (`updated_at`)," +
  "KEY `index agent_images_on_created_at` (`created_at`)" +
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
  db.MustExec(create_table_agent_coordinators)
  db.MustExec(create_table_agent_images)
  log.Println("DONE")
}
