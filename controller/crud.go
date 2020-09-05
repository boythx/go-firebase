package controller

import (
	"context"
	initial "firebase/initFirebase"
	"firebase/models"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
)

var ctx = context.Background()
var client = initial.Init(ctx)

func Home(c echo.Context) error {
	IncomesData := []models.Income{}
	// client := initial.Init(ctx)
	iter := client.Collection("income-v2").Documents(ctx)
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

func AddData(c echo.Context) error {
	price, error := strconv.ParseFloat(c.FormValue("price"), 10)
	if error != nil {
		return nil
	}
	IncomesData := models.Income{
		Date:  c.FormValue("date"),
		List:  c.FormValue("list"),
		Price: price,
	}
	// client := initial.Init(ctx)
	_, _, err := client.Collection("income-v2").Add(ctx, IncomesData)

	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	return c.JSON(http.StatusCreated, nil)

}
