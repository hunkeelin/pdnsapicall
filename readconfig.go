package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func readconfig() (rhost, rport, rapikey, rrzone string, rautoptr bool) {
	var varcheck []bool
	err := godotenv.Load("dnsapi.conf")
	if err != nil {
		log.Fatal("Error loading dnsapi.conf file; is it in the same directory as the software?")
	}
	port, isempty := os.LookupEnv("port")
	varcheck = append(varcheck, isempty)
	autoptr, isempty := os.LookupEnv("autoptr")
	varcheck = append(varcheck, isempty)
	fautoptr, err := strconv.ParseBool(autoptr)
	if err != nil {
		log.Fatal("check autoptr whether it's true/false")
	}
	host, isempty := os.LookupEnv("host")
	varcheck = append(varcheck, isempty)
	apikey, isempty := os.LookupEnv("apikey")
	varcheck = append(varcheck, isempty)
	rzone, isempty := os.LookupEnv("reversezone")
	varcheck = append(varcheck, isempty)
	return host, port, apikey, rzone, fautoptr
}
func checkvar(listvar []bool) {
	for _, check := range listvar {
		if check == false {
			fmt.Println("configuration error")
			log.Fatal("check your configuration")
		}
	}
}
