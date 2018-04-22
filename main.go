package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 7 {
		fmt.Println("wrong number of argument")
		fmt.Println("ARECORD: utest1 A REPLACE 1.1.1.1 domain 60")
		fmt.Println("DELETE: utest1 A DELETE 1.1.1.1 domain 60")
		fmt.Println("CNAME: utest1 CNAME REPLACE cname. klin-pro.com 60")
		fmt.Println("CNAME Note: the dot at the end of the cname host is important")
		os.Exit(0)
	} else {
		ch := make(chan string)
		ttl, err := strconv.Atoi(os.Args[6])
		if err != nil {
			log.Fatal(err)
		}
		fqdn := os.Args[1] + "." + os.Args[5]
		go loader(fqdn, os.Args[2], os.Args[3], os.Args[4], os.Args[5], ttl, ch)
		for elem := range ch {
			fmt.Println(elem)
		}
	}
}
