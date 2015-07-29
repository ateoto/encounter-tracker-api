package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(c *echo.Context) error {
	email := c.Form("email")
	password := c.Form("password")
	username := c.Form("username")

	// TODO: Need to enforce these fields or return malformed request error.

	// Lookup to see if user or email are in the database already.
	var existingUsers int
	if err := db.QueryRow("SELECT COUNT(id) FROM users WHERE email = $1 OR username = $2", email, username).Scan(&existingUsers); err != nil {
		log.Fatal(err)
	}

	if existingUsers <= 0 {
		hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		if err != nil {
			log.Fatal(err)
		}

		// User/email does not exist in database. Create new user and return user id.
		var userId int
		if err := db.QueryRow("INSERT INTO users(email, username, password) VALUES ($1, $2, $3) RETURNING id", email, username, hashed).Scan(&userId); err != nil {
			log.Fatal(err)
		}

	} else {
		return c.JSON(http.StatusConflict, "Email or Username already exists")
	}

	return c.JSON(http.StatusCreated, "User successfully created")
}

func LoginUser(c *echo.Context) error {
	// Post email and/or username with password.
	var user User

	email := c.Form("email")
	password := c.Form("password")
	username := c.Form("username")

	if password == "" {
		return c.JSON(http.StatusBadRequest, "Must specify a password")
	}

	// TODO: This looks like it could be simplified when it's not one in the morning.

	if email != "" && username != "" {
		if err := db.QueryRow("SELECT * FROM users WHERE email = $1 AND username = $2", email, username).Scan(&user.Id, &user.Email, &user.Username, &user.Password); err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusForbidden, "No thanks")
			}
			log.Fatal(err)
		}
	} else if email != "" && username == "" {
		if err := db.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user.Id, &user.Email, &user.Username, &user.Password); err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusForbidden, "No thanks")
			}
			log.Fatal(err)
		}
	} else if email == "" && username != "" {
		if err := db.QueryRow("SELECT * FROM users WHERE username = $1", username).Scan(&user.Id, &user.Email, &user.Username, &user.Password); err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusForbidden, "No thanks")
			}
			log.Fatal(err)
		}
	} else {
		return c.JSON(http.StatusBadRequest, "Must specify email or username.")
	}

	// We now have the user object from the database. Time to match passwords.
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return c.JSON(http.StatusForbidden, "No thanks")
	} else {
		// JWT Stuff

		token := jwt.New(jwt.SigningMethodHS256)

		token.Claims["user_id"] = user.Id
		token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		tokenString, err := token.SignedString([]byte(SigningKey))
		if err != nil {
			log.Fatal(err)
		}

		return c.JSON(http.StatusOK, tokenString)
	}
}

func LogoutUser(c *echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "Not Yet")
}

func UserDetail(c *echo.Context) error {
	// Must be authorized as the user
	// Return id, email, username
	return c.JSON(http.StatusNotImplemented, "Not Yet")
}

func DeleteUser(c *echo.Context) error {
	// Must be authorized as the user
	// Logout and delete from database
	return c.JSON(http.StatusNotImplemented, "Not Yet")
}
