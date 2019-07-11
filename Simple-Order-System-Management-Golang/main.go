package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/middleware"
	_customerHttpDeliver "gitlab.warungpintar.co/enrinal/intern-diary/simple-order/customer/delivery/http"
	_customerRepo "gitlab.warungpintar.co/enrinal/intern-diary/simple-order/customer/repository/mysql"
	_cusotmerUsecase "gitlab.warungpintar.co/enrinal/intern-diary/simple-order/customer/usecase"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbName := viper.GetString(`database.dbname`)
	connection := fmt.Sprintf("%s:@tcp(%s:%s)/%s", dbUser, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	customerrepo := _customerRepo.NewMysqlCustomerRepository(dbConn)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	customerusecase :=  _cusotmerUsecase.NewCustomerUsecase(customerrepo, timeoutContext)
	_customerHttpDeliver.NewCustomerHandler(e,customerusecase)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
