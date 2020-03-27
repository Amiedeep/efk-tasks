//+build mage

package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	kibanaURL        = "http://kibana_url:port"
	dashoard         = "dashoboard_id"
	dashoardFileName = "filename.json"
)

func createFile(name string) (file *os.File) {
	file, err := os.Create(name)

	if err != nil {
		fmt.Println("Error creating file")
		panic(err)
	}

	return file
}

func httpGetCall(url string) (output *http.Response) {
	output, err := http.Get(url)

	if err != nil {
		fmt.Println("Error in get url")
		panic(err)
	}
	return output
}

func ExportDashboard() {
	url := kibanaURL + "/api/kibana/dashboards/export?dashboard=" + dashoard

	fmt.Println("URL: ", url)
	output := httpGetCall(url)

	file := createFile(dashoardFileName)

	_, err := io.Copy(file, output.Body)

	if err != nil {
		fmt.Println("Error copying file")
		panic(err)
	}
}

func ImportDashboard() {
	url := kibanaURL + "/api/kibana/dashboards/import?force=true"
	fmt.Println("URL: ", url)
	dashboardFile, err := ioutil.ReadFile(dashoardFileName)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(dashboardFile))
	req.Header.Set("kbn-xsrf", "true")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	fmt.Println(resp.Status)
	// out, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(out))

	if err != nil {
		fmt.Println("Error exporting dashboard")
		fmt.Println(err.Error())

		panic(err)
	}
}

func DeleteSavedObject() {
	objectType := "dashboard"
	id := "dashoboard id"
	// space := "default"
	url := kibanaURL + "/api/saved_objects/" + objectType + "/" + id
	// url := kibanaURL + "/s/" + space + "/api/saved_objects/" + objectType + "/" + id
	fmt.Println("URL: ", url)

	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Set("kbn-xsrf", "true")

	client := &http.Client{}
	resp, err := client.Do(req)

	fmt.Println(resp.Status)
	// out, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(out))

	if err != nil {
		fmt.Println("Error deleting object")
		fmt.Println(err.Error())

		panic(err)
	}
}
