package main

import (
	"fmt"
	"github.com/xEtarusx/vote-analyser/models"
	"sort"
)

func WorstVotedLayersWithMinimumVotes(layerData []models.LayerData, amount int, minimumVotesTotal int) {
	sort.Slice(layerData, func(i, j int) bool {
		return layerData[i].VotePercentagePositive() < layerData[j].VotePercentagePositive()
	})

	counter := 1
	for _, layer := range layerData {
		if layer.TotalVotes() < minimumVotesTotal {
			continue
		}

		fmt.Printf("%s (+%d : -%d) with a negative of %.2f%%\n", layer.Name, layer.Upvotes, layer.Downvotes, layer.VotePercentageNegative())

		if counter == amount {
			break
		}
		counter++
	}
}
