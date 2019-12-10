package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const (
	height = 6
	width  = 25
)

func readInput(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Could not read file :%s\n", err)
	}

	return data
}

func main() {
	imgData := readInput("input.txt")
	layerSize := width * height

	zeroInSelectedLayer := layerSize + 1
	selectedLayerFactor := -1
	for layer := 0; layer < len(imgData); layer += layerSize {
		layerData := imgData[layer : layer+layerSize]
		zeroCount := 0
		oneCount := 0
		twoCount := 0

		for _, pixel := range layerData {
			switch pixel {
			case '0':
				zeroCount++
			case '1':
				oneCount++
			case '2':
				twoCount++
			default:
				log.Fatalf("Unknown pixel=%s in layer=%d\n", string(pixel), layer/layerSize)
			}
		}

		if zeroCount < zeroInSelectedLayer {
			zeroInSelectedLayer = zeroCount
			selectedLayerFactor = oneCount * twoCount
		}
	}

	fmt.Println(selectedLayerFactor)
}
