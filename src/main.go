package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

var db *sql.DB

var migrateUrl string

var SigningKey = os.Getenv("ET_SIGNING_KEY")

const (
	Bearer = "Bearer"
)

func main() {
	var err error

	db_host := os.Getenv("ET_DB_HOST")
	db_user := os.Getenv("ET_DB_USER")
	db_pass := os.Getenv("ET_DB_PASS")
	db_name := os.Getenv("ET_DB_NAME")

	connection := fmt.Sprintf("host=%v sslmode=disable user=%v password=%v dbname=%v", db_host, db_user, db_pass, db_name)
	migrateUrl = fmt.Sprintf("postgres://%v@%v/%v?sslmode=disable&password=%v", db_user, db_host, db_name, db_pass)

	log.Printf(connection)

	db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(100)

	RunMigrations()

	e := echo.New()

	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.StripTrailingSlash())

	//e.Get("/", index)

	m := e.Group("/monsters")
	m.Use(JWTAuth(SigningKey))

	m.Get("", IndexMonster)
	m.Post("", CreateMonster)
	m.Get("/:id", ReadMonster)
	m.Put("/:id", UpdateMonster)
	m.Delete("/:id", DeleteMonster)

	l := e.Group("/languages")
	l.Get("", LanguageIndex)
	l.Post("", CreateLanguage)
	l.Get("/:id", LanguageDetail)
	l.Delete("/:id", DeleteLanguage)

	u := e.Group("/users")
	u.Post("", CreateUser)
	u.Post("/login", LoginUser)
	u.Get("/:id", UserDetail)
	u.Delete("/:id", DeleteUser)

	c := e.Group("/campaigns")
	c.Post("", CreateCampaign)
	c.Get("", IndexCampaign)
	c.Get("/:id", ReadCampaign)
	c.Put("/:id", UpdateCampaign)
	c.Delete("/:id", DeleteCampaign)

	/*
		en := e.Group("/encounters")
		en.Post("", CreateEncounter)
		en.Get("", IndexEncounter)
		en.Get("/:id", ReadEncounter)
		en.Put("/:id", UpdateEncounter)
		en.Delete("/:id", DeleteEncounter)
	*/

	e.Run(":4242")
}
