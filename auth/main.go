package main

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/dgrijalva/jwt-go.v3"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	MSISDN   string `json:"msisdn"`
}

var db *sql.DB

const (
	tokenSecretKey = "secret_key"
)

func main() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/orderfaz1")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/register", registerHandler)
	router.POST("/login", loginHandler)

	router.Run(":8080")
}

func registerHandler(c *gin.Context) {
	msisdn := c.PostForm("msisdn")
	name := c.PostForm("name")
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Check if username already exists
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// Check if MSISDN already exists
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE msisdn = ?", msisdn).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MSISDN already exists"})
		return
	}

	// Ensure MSISDN has a prefix of "62"
	if !strings.HasPrefix(msisdn, "62") {
		msisdn = "62" + msisdn
	}

	// Hash the password
	hashedPassword, err := hashPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	// Generate UUID
	id := generateUUID()

	// Insert user data into the database
	insertSQL := "INSERT INTO users (id, name, username, password, msisdn) VALUES (?, ?, ?, ?, ?)"
	_, err = db.Exec(insertSQL, id, name, username, hashedPassword, msisdn)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func loginHandler(c *gin.Context) {
	msisdn := c.PostForm("msisdn")
	password := c.PostForm("password")

	// Retrieve user from the database
	var user User
	err := db.QueryRow("SELECT id, password FROM users WHERE msisdn = ?", msisdn).Scan(&user.ID, &user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid MSISDN or password"})
		return
	}

	// Check if the password matches
	if verifyPassword(password, user.Password) {
		token, err := generateJWT(user.ID)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid MSISDN or password"})
	}
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func verifyPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func generateUUID() string {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

func generateJWT(userID string) (string, error) {
	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Generate encoded token
	tokenString, err := token.SignedString([]byte(tokenSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
