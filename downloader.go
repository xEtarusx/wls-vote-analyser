package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/xEtarusx/vote-analyser/models"
	"strconv"
	"time"
)

func DownloadMapData() ([]models.LayerData, error) {
	c := colly.NewCollector()

	// Increase http timeout because the website is loading super slow (~12s)
	c.SetRequestTimeout(30 * time.Second)

	var mapData []models.LayerData
	c.OnHTML(".table.table-sm tbody tr", func(e *colly.HTMLElement) {
		// first column: id
		id, err := strconv.Atoi(e.ChildText("td:first-child"))
		if err != nil {
			fmt.Println("Could not convert id " + e.ChildText("td:first-child") + " to int")
			return
		}

		// third column: upvotes
		upvotes, err := strconv.Atoi(e.ChildText("td:nth-child(3)"))
		if err != nil {
			fmt.Println("Could not convert id " + e.ChildText("td:nth-child(3)") + " to int")
			return
		}

		// fourth column: downvotes
		downvotes, err := strconv.Atoi(e.ChildText("td:nth-child(4)"))
		if err != nil {
			fmt.Println("Could not convert id " + e.ChildText("td:nth-child(4)") + " to int")
			return
		}

		mapStats := models.LayerData{
			Id:        id,
			Name:      e.ChildText("td:nth-child(2)"),
			Upvotes:   upvotes,
			Downvotes: downvotes,
		}

		mapData = append(mapData, mapStats)
	})

	err := c.Visit("https://mapvote.welovesquad.de/")

	return mapData, err
}
