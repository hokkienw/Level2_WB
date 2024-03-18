package main

import (
	"fmt"
	"log"
	"time"
	"github.com/beevik/ntp"
)



func main() {
	time, err := GetTime("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time)
}

func GetTime(server string) (time.Time, error){
	time, err := ntp.Time(server)
	if err != nil {
		log.Fatal(err)
	}
	return time, err
}
