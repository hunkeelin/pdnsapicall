package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type editrecord struct{}

func loader(name, recordtype, ctype, content, zone string, ttl int, ch chan string) {
	f := editrecord{}
	host, port, apikey, rzone, autoptr := readconfig()
	fqdn := name + "."
	url := fmt.Sprintf("http://%s:%s/api/v1/servers/localhost/zones/%s", host, port, zone)
	body, status_code := f.worker(fqdn, recordtype, ctype, content, zone, url, apikey, ttl)
	ch <- body
	ch <- status_code
	if autoptr && recordtype == "A" {
		g := strings.Split(content, ".")
		content = fqdn
		fqdn = g[3] + "." + g[2] + "." + rzone
		recordtype = "PTR"
		url := fmt.Sprintf("http://%s:%s/api/v1/servers/localhost/zones/%s", host, port, rzone)
		body, status_code = f.worker(fqdn, recordtype, ctype, content, rzone, url, apikey, ttl)
		ch <- body
		ch <- status_code
	}
	close(ch)
}
func (f *editrecord) worker(fqdn, recordtype, ctype, content, zone, url, apikey string, ttl int) (returnbody, returnstatus string) {
	payload := makePayload(fqdn, recordtype, ctype, content, ttl, false)
	encodepayload, _ := json.Marshal(payload)
	ebody := bytes.NewReader(encodepayload)
	req, err := http.NewRequest("PATCH", url, ebody)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("X-API-Key", apikey)

	client := http.Client{
		Timeout: 500 * time.Millisecond,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body), string(resp.Status)
}
