//+build mage

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func createFile(name string) (file *os.File) {
	file, err := os.Create("dashboard.yaml")

	if err != nil {
		fmt.Println("Error creating file")
		panic(err)
	}

	return file
}

func httpCall(url string) (output *http.Response) {
	output, err := http.Get(url)

	if err != nil {
		fmt.Println("Error in get url")
		panic(err)
	}
	return output
}

func GetDashboard() {
	url := "https://url"

	output := httpCall(url)

	file := createFile("dashboard.json")

	_, err := io.Copy(file, output.Body)

	if err != nil {
		fmt.Println("Error copying file")
		panic(err)
	}
}
