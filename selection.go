package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/xEtarusx/vote-analyser/models"
)

var TextTopTenVotedMaps = "Top 10 voted Layers (minimum 20 total Votes)"
var TextWorstTenVotedMaps = "Worst 10 voted Layers (minimum 20 total Votes)"
var TextOutOfRotationMaps = "Maps out of current rotation"
var ExitApp = "Exit"

func ShowPrompt(layerData []models.LayerData) bool {
	fmt.Println("")

	prompt := promptui.Select{
		Label: "What do you want to see?",
		Items: []string{TextTopTenVotedMaps, TextWorstTenVotedMaps, TextOutOfRotationMaps, ExitApp},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return true
	}

	switch result {
	case TextTopTenVotedMaps:
		TopVotedLayersWithMinimumVotes(layerData, 10, 20)
		break
	case TextWorstTenVotedMaps:
		WorstVotedLayersWithMinimumVotes(layerData, 10, 20)
		break
	case TextOutOfRotationMaps:
		OutOfRotationMaps(layerData)
		break
	case ExitApp:
		return true
	}

	return false
}
