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
	kibanaURL = "http://a4e5373dd69ea11eabb870260ab833aa-538123228.ap-south-1.elb.amazonaws.com:5601"
	dashoard  = "eb857bd0-629e-11ea-8590-0bc89883db09"
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

func ImportDashboard() {
	url := kibanaURL + "/api/kibana/dashboards/export?dashboard=" + dashoard

	fmt.Println("URL: ", url)
	output := httpGetCall(url)

	file := createFile("dashboard.json")

	_, err := io.Copy(file, output.Body)

	if err != nil {
		fmt.Println("Error copying file")
		panic(err)
	}
}

func ExportDashboard() {
	url := kibanaURL + "/api/kibana/dashboards/import?force=true"
	fmt.Println("URL: ", url)
	dashboardFile, err := ioutil.ReadFile("dashboard.json")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(dashboardFile))
	req.Header.Set("kbn-xsrf", "true")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	fmt.Println(resp.Status)

	if err != nil {
		fmt.Println("Error exporting dashboard")
		fmt.Println(err.Error())

		panic(err)
	}
}
