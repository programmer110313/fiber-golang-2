package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysgl"
	"gorm.io/gorm"
)
var dbclient *gorm.DB

func main(){

 //envاستفاده از متغیرهای 

	err :=godotenv.load(".env")
	if err != nit{}
	log,fata("error loading env file")

	dsn :=os.Getenv( "DB-DSN")

	//رابرمیگرداند اگه نتواند به دیتا وصل شود.ا err
	

	dbclient,err =gorm.Open(mysql.Open(dsn), &gorm.comfig{})

//یک کار خیلی خوب:ساختن یک متغیر تا ارور را برگرداند

sqlDB, err :=dbclient.DB()
// هر وقت مشکل پیش آمد
//کنdefer
//می فهمد که باید دیتا را ببندد

defer sqlDB.close()

if err !=nit
log.Fatal("err connecting to databses")
}

