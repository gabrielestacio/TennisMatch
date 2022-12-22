package main

import (
	"fmt"
	"time"
)

func player11(quadra chan string) {
	event := <-quadra
	if event == "Começa o jogo" {
		fmt.Println("vai esperar por mim 1")
	} else {
		fmt.Println("Obrigado 1")
	}

	quadra <- "Devolveu"
}

func player22(quadra chan string) {
	event := <-quadra
	if event == "Começa o jogo" {
		fmt.Println("vai esperar por mim 1")
	} else {
		fmt.Println("Obrigado 1")
	}

	quadra <- "Devolveu"
}

func main() {
	quadra := make(chan string, 1)
	quadra <- "Começa o jogo"

	go player11(quadra)
	go player22(quadra)

	time.Sleep(time.Second * 5)

	fmt.Print("...")
}
