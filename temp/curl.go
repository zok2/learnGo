package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	baseUrl := "http://pic3.sfacg.com/Pic/OnlineComic4/ZLSZ/ZP/"
	appendUrl := "?p=http://www.baidu.com/img/bdlogo.gif"
	num := 506
	for i := 23; i <= 110; i++ {
		for j := 1; j <= 110; j++ {
			formattedNumberi := fmt.Sprintf("%03d", i)
			formattedNumberj := fmt.Sprintf("%03d", j)
			imageURL := fmt.Sprintf("%s%s/%s.jpg%s", baseUrl, formattedNumberi, formattedNumberj, appendUrl)
			response, err := http.Get(imageURL)

			if err != nil {
				fmt.Println("Error while fetching the image:", err)
				continue
			}
			if response.StatusCode != http.StatusOK {
				fmt.Println("Error while fetching the image:", err)
				continue
			}
			defer response.Body.Close()

			// Create or open the destination file
			num++
			k := num

			destinationPath := fmt.Sprintf("/Users/zok2/Downloads/A%d.jpg", k)
			file, err := os.Create(destinationPath)
			if err != nil {
				fmt.Println("Error while creating the file:", err)
				continue
			}
			defer file.Close()

			// Copy the response body to the destination file
			_, err = io.Copy(file, response.Body)
			time.Sleep(1)
		}
	}
}
