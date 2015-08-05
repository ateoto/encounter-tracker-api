package main

import (
	"database/sql"
	"encoding/json"
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

type Token struct {
	AccessToken string `json:"access_token"`
}

type Simple struct {
	Message string `json:"message"`
}

func CreateUser(c *echo.Context) error {
	var u User

	if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
		log.Fatal(err)
	}

	email := u.Email
	password := u.Password
	username := u.Username

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
		return c.JSON(http.StatusConflict, Simple{Message: "Email or username already exists"})
	}

	return c.JSON(http.StatusCreated, Simple{Message: "User successfully created"})
}

func LoginUser(c *echo.Context) error {
	// Post email and/or username with password.
	var u User

	if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
		log.Fatal(err)
	}

	email := u.Email
	password := u.Password
	username := u.Username

	if password == "" {
		return c.JSON(http.StatusBadRequest, Simple{Message: "Must specify a password"})
	}

	// TODO: This looks like it could be simplified when it's not one in the morning.

	if email != "" && username != "" {
		if err := db.QueryRow("SELECT * FROM users WHERE email = $1 AND username = $2", email, username).Scan(&u.Id, &u.Email, &u.Username, &u.Password); err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusForbidden, Simple{Message: "Forbidden"})
			}
			log.Fatal(err)
		}
	} else if email != "" && username == "" {
		if err := db.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&u.Id, &u.Email, &u.Username, &u.Password); err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusForbidden, Simple{Message: "Forbidden"})
			}
			log.Fatal(err)
		}
	} else if email == "" && username != "" {
		if err := db.QueryRow("SELECT * FROM users WHERE username = $1", username).Scan(&u.Id, &u.Email, &u.Username, &u.Password); err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusForbidden, Simple{Message: "Forbidden"})
			}
			log.Fatal(err)
		}
	} else {
		return c.JSON(http.StatusBadRequest, Simple{Message: "Must provide a username/email and password"})
	}

	// We now have the user object from the database. Time to match passwords.
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return c.JSON(http.StatusForbidden, Simple{Message: "Forbidden"})
	} else {
		// JWT Stuff

		token := jwt.New(jwt.SigningMethodHS256)

		token.Claims["user_id"] = u.Id
		token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		tokenString, err := token.SignedString([]byte(SigningKey))
		if err != nil {
			log.Fatal(err)
		}

		var jsonToken Token

		jsonToken.AccessToken = tokenString

		return c.JSON(http.StatusOK, jsonToken)
	}
}

func LogoutUser(c *echo.Context) error {
	// Invalidate token
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
