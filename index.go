package main

import (
	"fmt"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"strings"

	"github.com/google/uuid"
	"github.com/kbinani/screenshot"
	"github.com/pbnjay/memory"
)

func getPcConfig() {

	m := memory.TotalMemory() / 1000 * 100
	fmt.Print(strconv.Itoa(int(m)))

	response, error := http.Get("https://api.ipify.org?format=json")
	if error != nil {
		fmt.Println(error)
	}

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Println(error)
	}
	response.Body.Close()
	ip := strings.Replace(strings.Replace(strings.Split(string(body), ":")[1], "}", "", -1), `"`, "", -1)
	fmt.Printf("\n" + ip)
}

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

func addAutoStart() {

}

func getUUID() {
	id := uuid.New()
	fmt.Println(id.String())

}

func main() {
	getPcConfig()
	createScreenShots()
	getUUID()
}
