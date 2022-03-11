package main

import (
	"log"
	"os/exec"
)

func main(){
	cmd := exec.Command("/bin/bash","-c","./write_word.sh")
	err := cmd.Run()
	if err != nil{
		log.Println(err)
	}
	log.Println("ok")
}