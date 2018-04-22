package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getzone() error {
	host, port, apikey, _, _ := readconfig()
	url := fmt.Sprintf("http://%s:%s/api/v1/servers/localhost/zones", host, port)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-API-Key", apikey)

	client := http.Client{
		Timeout: 500 * time.Millisecond,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	//newStr := buf.String()
	var g []Getstruct
	jerr := json.Unmarshal(body, &g)
	if jerr != nil {
		return jerr
	}
	d, err := json.MarshalIndent(g, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(d))
	return nil
}
