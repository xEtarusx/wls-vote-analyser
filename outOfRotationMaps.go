package main

import (
	"fmt"
	"github.com/xEtarusx/vote-counter/models"
)

func OutOfRotationMaps(layerData []models.LayerData) {
	for _, layer := range layerData {
		// Only layers with a total of 50 or more votes can be voted out of rotation
		// Quote of Harakiri: "50 Votes insgesamt pro Layer. Negativ-Votes müssen positive um 20% übersteigen, damit das Layer fliegt."
		if layer.TotalVotes() < 50 {
			continue
		}

		if layer.VotePercentageDifference() <= -20 {
			fmt.Printf("%s (+%d : -%d) with a negative of %.2f%%\n", layer.Name, layer.Upvotes, layer.Downvotes, layer.VotePercentageNegative())
		}
	}
}
