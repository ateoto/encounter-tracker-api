package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

type Campaign struct {
	Id     int64  `json:"id"`
	UserId int64  `json:"userId"`
	Name   string `json:"name"`
}

type Campaigns []Campaign

func CreateCampaign(c *echo.Context) error {
	// POST /campaign
	// User is decoded from JWT Token
	name := c.Form("name")

	log.Printf(name)
	// Grab from JWT instead
	userid := 5

	var campaign Campaign

	if err := db.QueryRow("INSERT INTO campaigns(userid, name) VALUES ($1, $2) RETURNING id", userid, name).Scan(&campaign.Id); err != nil {
		log.Printf("I can't even put it in.")
		log.Fatal(err)
	}

	if err := db.QueryRow("SELECT userid, name FROM campaigns WHERE id=$1", campaign.Id).Scan(&campaign.UserId, &campaign.Name); err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusCreated, campaign)
}

func IndexCampaign(c *echo.Context) error {
	// GET /campaign
	// returns JSON Array of Campaigns

	// Get User id from JWT Token

	userid := 5

	rows, err := db.Query("SELECT * FROM campaigns WHERE userid = $1", userid)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var campaigns Campaigns

	for rows.Next() {
		var campaign Campaign
		err := rows.Scan(&campaign.Id, &campaign.UserId, &campaign.Name)
		if err != nil {
			log.Fatal(err)
		}

		campaigns = append(campaigns, campaign)
	}

	return c.JSON(http.StatusOK, campaigns)
}

func ReadCampaign(c *echo.Context) error {
	// GET /campaign/:id
	campaignId := c.Param("id")
	// Get User id from JWT Token

	userid := 5

	var campaign Campaign

	if err := db.QueryRow("SELECT * FROM campaigns WHERE id = $1 AND userid = $2", campaignId, userid).Scan(&campaign.Id, &campaign.UserId, &campaign.Name); err != nil {
		if err == sql.ErrNoRows {
			return c.NoContent(http.StatusNoContent)
		} else {
			log.Fatal(err)
		}
	}

	return c.JSON(http.StatusOK, campaign)
}

func UpdateCampaign(c *echo.Context) error {
	// PUT /campaign/:id
	return c.JSON(http.StatusNotImplemented, "Not Yet")
}

func DeleteCampaign(c *echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "Not Yet")
}
