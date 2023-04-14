package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

// Data yang akan dikirimkan dalam format JSON
type Data struct {
	ValueWater int `json:"water"`
	ValueWind  int `json:"wind"`
}

func main() {
	// URL untuk POST request
	url := "https://jsonplaceholder.typicode.com/posts"

	// Seed untuk random number generator
	rand.Seed(time.Now().UnixNano())

	for {
		// Membuat data dengan nilai acak
		data := Data{
			ValueWater: rand.Intn(100) + 1,
			ValueWind:  rand.Intn(100) + 1,
		}

		// Menentukan status air dan angin
		StatusWater := getStatus(data.ValueWater, 5, 8)
		StatusWind := getStatus(data.ValueWind, 6, 15)

		// Mengirimkan POST request dengan data JSON
		body, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
		if err != nil {
			panic(err)
		}

		// Menampilkan hasil POST request pada terminal
		fmt.Println("Status Code:", resp.StatusCode)
		fmt.Println("Response Body:")
		if resp.Body != nil {
			//defer resp.Body.Close()
			fmt.Println(readResponseBody(resp.Body))
			//fmt.Printf("Value Water: %d, Value Wind: %d\n\n", data.ValueWater, data.ValueWind)
			fmt.Printf("Value Water: %s\nValue Wind: %s\n\n", StatusWater, StatusWind)
		}
		fmt.Println()

		// Menunggu selama 15 detik sebelum mengirimkan POST request lagi
		time.Sleep(15 * time.Second)
	}
}

// Fungsi untuk menentukan status air dan angin
func getStatus(value int, thresholdSafe int, thresholdWarning int) string {
	if value < thresholdSafe {
		return "aman"
	} else if value < thresholdWarning {
		return "siaga"
	} else {
		return "bahaya"
	}
}

// Fungsi untuk membaca response body sebagai string
func readResponseBody(body io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	return buf.String()
}
