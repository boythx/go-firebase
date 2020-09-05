package main

import (
	"context"
	"firebase/models"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type App struct {
	client *firestore.Client
	ctx    context.Context
}

func main() {
	route := App{}
	route.Init()
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/", func(c echo.Context) error {
		return route.Home(c)
	})
	e.Logger.Fatal(e.Start(":9000"))
}

func (route *App) Init() {
	route.ctx = context.Background()
	sa := option.WithCredentialsFile("income.json")
	app, err := firebase.NewApp(route.ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	route.client, err = app.Firestore(route.ctx)
	if err != nil {
		log.Fatalln(err)
	}

}

func (route *App) Home(c echo.Context) error {
	IncomesData := []models.Income{}
	iter := route.client.Collection("incomes").Documents(route.ctx)
	for {
		IncomeData := models.Income{}
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		mapstructure.Decode(doc.Data(), &IncomeData)
		IncomesData = append(IncomesData, IncomeData)
	}
	return c.JSON(http.StatusOK, IncomesData)
}
