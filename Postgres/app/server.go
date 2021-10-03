package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}
type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Welcome to " + appConfig.AppName)

	server.InitializeDB(dbConfig)
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Pritnf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (server *Server) InitializeDB(dbConfig DBConfig) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort, "shella", "password", "gits.id", "5432")
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nill {
		panic("Failed on Connecting to the database server")
	}
	for _, model := range RegisterModels() {
		err = server.DB.Debug()AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println('"Database migrated successfully')
}
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}
	var DBConfig = DBConfig{}

	err := godotenv.load()
	if err != nil {
		log.Fatal("Error on loading .env file")
	}

	appConfig.AppName = os.Getenv("App_Name", "GoTugas")
	appConfig.AppEnv = os.Getenv("App_Env", "Development")
	appConfig.AppPort = os.Getenv("App_Port", "9000")

	dbConfig.DBHost = getEnv("DB_Host", "localhost")
	dbConfig.DBUser = getEnv("DB_User", "User")
	dbConfig.DBPassword = getEnv("DB_Password", "passsword")

	dbCOnfig.DBName = getEnv("DB_Name", "dbname")
	dbConfig.DBPort = getEnv("DB_Port", "5432")

	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)
}
