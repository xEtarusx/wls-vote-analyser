package main

import (
	"fmt"
)

func main() {
	fmt.Print("Downloading data (Takes about 10s) ...")
	layerData, err := DownloadMapData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(" done")

	for {
		exit := ShowPrompt(layerData)

		if exit {
			break
		}
	}
}
