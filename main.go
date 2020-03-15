package main

import (
	"log"
	"os/exec"
)

func main() {

	c := exec.Command(`D:\Games\Steam\Steam.exe`)
	err := c.Start()
	if err != nil {
		log.Println(err)
	}

	log.Println("all good")

}
