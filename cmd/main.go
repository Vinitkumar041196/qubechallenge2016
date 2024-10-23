package main

import "log"

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate)
	log.Println("Initializing server...")
}
