package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

type Language struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Standard bool   `json:"standard"`
}

type Languages []Language

func LanguageIndex(c *echo.Context) error {
	rows, err := db.Query("SELECT * FROM languages")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	languages := Languages{}

	for rows.Next() {
		var language Language
		err := rows.Scan(&language.Id, &language.Name, &language.Standard)
		if err != nil {
			log.Fatal(err)
		}

		languages = append(languages, language)
	}

	return c.JSON(http.StatusOK, languages)
}

func CreateLanguage(c *echo.Context) error {
	var language Language

	jwt_claims, _ := c.Get("claims").(map[string]interface{})
	user_id := jwt_claims["user_id"]

	if err := json.NewDecoder(c.Request().Body).Decode(&language); err != nil {
		log.Fatal(err)
	}

	err := db.QueryRow("INSERT INTO languages(user, name, standard) VALUES ($1, $2, $3) RETURNING *", language.Name, language.Standard).Scan(user_id, &language.Id, &language.Name, &language.Standard)
	if err != nil {
		log.Fatal(err)
	}

	/*
		err = db.QueryRow("SELECT name, standard FROM languages WHERE id=$1", language.Id).Scan(&language.Name, &language.Standard)
		if err != nil {
			log.Printf("Fail on requery")
			log.Fatal(err)
		}
	*/
	return c.JSON(http.StatusCreated, language)
}

func LanguageDetail(c *echo.Context) error {
	languageId := c.Param("id")

	var language Language

	err := db.QueryRow("SELECT * FROM languages WHERE id =$1", languageId).Scan(&language.Id, &language.Name, &language.Standard)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.NoContent(http.StatusNoContent)
		} else {
			log.Fatal(err)
		}
	}

	return c.JSON(http.StatusOK, language)
}

func DeleteLanguage(c *echo.Context) error {
	languageId := c.Param("id")

	_, err := db.Exec("DELETE FROM languages WHERE id=$1", languageId)

	if err != nil {
		log.Fatal(err)
	}

	return c.NoContent(http.StatusNoContent)
}

func UpdateLanguage(c *echo.Context) error {
	return c.String(http.StatusOK, "Update Language")
}
