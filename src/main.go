package main

import (
	"database/sql"
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
	InitDB()
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
	l.Use(JWTAuth(SigningKey))
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

	log.Printf("Serving on port: 4242")
	e.Run(":4242")
}
