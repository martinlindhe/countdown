package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func format(d time.Duration) string {
	//mil := d.Milliseconds() % 1000
	sec := int(d.Seconds()) % 60
	min := int(d.Minutes())
	//return fmt.Sprintf("%v:%02v.%03v", min, sec, mil)
	return fmt.Sprintf("%v:%02v", min, sec)
}

func main() {

	in := ""

	if len(os.Args) >= 2 {
		in = os.Args[1]
	} else {
		in = "25m"
	}

	duration, err := time.ParseDuration(in)
	if err != nil {
		log.Fatal(err)
	}

	finished := time.Now().Add(duration)

	fmt.Printf("Counting down for %s\n", format(duration))
	fmt.Print(format(time.Until(finished)))

	for {
		time.Sleep(250 * time.Millisecond)

		if time.Now().After(finished) {
			fmt.Print("\rFinished!        \n")
			break
		}

		fmt.Print("\r", format(time.Until(finished)))
	}
}
