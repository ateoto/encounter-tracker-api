package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

type Size struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	SpaceSquares string `json:"space_squares"`
	SpaceHexes   string `json:"space_hexes"`
}

type MonsterModel struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	HitPoints       int    `json:"hit_points"`
	ChallengeRating string `json:"challenge_rating"`
	XpReward        int    `json:"xp_reward"`
	ArmorClass      int    `json:"armor_class"`
	ArmorType       string `json:"armor_type"`
	SizeId          int    `json:"size_id"`
}

func IndexMonster(c *echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "Not Yet")
}

func CreateMonster(c *echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "Not Yet")
}

func ReadMonster(c *echo.Context) error {
	monsterId := c.Param("id")
	log.Printf(monsterId)
	// Get Monster

	// Get MonsterSkills
	return c.JSON(http.StatusNotImplemented, "Not Yet")
}

func UpdateMonster(c *echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "Not Yet")
}

func DeleteMonster(c *echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "Not Yet")
}
