package main

import (
	"fmt"
	"time"
)

func speak(person, text string, qty int) {
	for i := 0; i < qty; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s) %s (iteration %d)\n", person, text, i+1)
	}
}

func main() {
	//speak("Breno", "How are you?", 3)
	//speak("Vanessa", "I can't only speak after you! ", 1)

	//Command go creates a go routine (Something like a thread)
	// go speak("Breno", "Hello!", 10)
	// go speak("Vanessa", "Hi!", 10)

	// time.Sleep(time.Second * 5)
	// fmt.Println("FIM!")

	go speak("Maria", "I get it", 10)
	speak("John", "Hello", 5)
}
