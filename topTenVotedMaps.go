package main

import (
	"fmt"
	"github.com/xEtarusx/vote-counter/models"
	"sort"
)

func TopVotedLayersWithMinimumVotes(layerData []models.LayerData, amount int, minimumVotesTotal int) {
	sort.Slice(layerData, func(i, j int) bool {
		return layerData[i].VotePercentagePositive() > layerData[j].VotePercentagePositive()
	})

	counter := 1
	for _, layer := range layerData {
		if layer.TotalVotes() < minimumVotesTotal {
			continue
		}

		fmt.Printf("%s (+%d : -%d) with a positive of %.2f%%\n", layer.Name, layer.Upvotes, layer.Downvotes, layer.VotePercentagePositive())

		if counter == amount {
			break
		}
		counter++
	}
}
