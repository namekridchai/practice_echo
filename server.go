package main

import (
	"fmt"      // formatting and printing values to the console.
	"log"      // logging messages to the console.
	"net/http" // Used for build HTTP servers and clients.

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Port we listen on.
const portNum string = ":8080"

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

// Handler functions.
func Home(c echo.Context) error {
	// fmt.Fprintf(w, "Homepage")
	return c.String(http.StatusAccepted, "hello home")
}

func CreateUser(c echo.Context) error {

	var user User
	err := c.Bind(&user)

	if err != nil {
		return err
	}

	users = append(users, user)
	fmt.Println(users)
	return c.JSON(http.StatusCreated, user)

}

func GetUser(c echo.Context) error {
	return c.JSON(http.StatusAccepted, users)

}

func DisplayHello(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		fmt.Println("hello console")
		return next(c)

	})
}

func HandleBasicAuth(username string, password string, c echo.Context) (bool, error) {
	if username == "joe" && password == "secret" {
		return true, nil
	}
	return false, nil
}

func main() {
	log.Println("Starting our simple http server.")

	e := echo.New()
	e.GET("/", Home)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(DisplayHello)

	g := e.Group("/api")

	g.Use(middleware.BasicAuth(HandleBasicAuth))

	g.GET("/users", GetUser)
	g.POST("/users", CreateUser)

	e.Logger.Fatal(e.Start(portNum))
	fmt.Println("To close connection CTRL+C :-)")

}
