package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const (
	height    = 6
	width     = 25
	layerSize = width * height
)

func readInput(path string) [][]byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Could not read file :%s\n", err)
	}

	var imgData [][]byte
	for layer := 0; layer < len(data); layer += layerSize {
		layerData := data[layer : layer+layerSize]
		imgData = append(imgData, layerData)
	}
	return imgData
}

func main() {
	imgData := readInput("input.txt")

	var decodedImg []byte
	for i := 0; i < layerSize; i++ {
		var pixel byte = 2
		for _, layer := range imgData {
			if layer[i] != '2' {
				pixel = layer[i]
				break
			}
		}

		// to read easily
		if pixel == '1' {
			pixel = '@'
		} else {
			pixel = ' '
		}

		decodedImg = append(decodedImg, pixel)
	}

	// print each row
	for y := 0; y < height; y++ {
		fmt.Println(string(decodedImg[y*width : y*width+width]))
	}
}
