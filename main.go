package main

import "log"

func main() {
	log.Println("d") // nocheck:thislog
	log.Println("d") // nocheck:thislog
	log.Println("d") // nocheck:thislog
	log.Println("d")
}
