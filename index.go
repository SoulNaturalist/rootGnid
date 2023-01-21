package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
	"github.com/pbnjay/memory"
	utils "https://github.com/SoulNaturalist/rootGnid/utils"
)

func createScreenShots() {
	n := screenshot.NumActiveDisplays()
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)

		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	}
}

func getPcConfig() {
	m := memory.TotalMemory() / 1000 * 100
	fmt.Printf("Total system memory:%d", m)
}

func main() {
	createScreenShots()
	getPcConfig()

}
