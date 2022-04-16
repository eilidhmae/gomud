package main
// created by Eilidh Robey erobey@sonatype.com

import (
	"fmt"
	"log"
	mud "gomud/adventure"
	"os"
)

const welcomeBanner = `
 \ \        /      |                                    |                ___|                         \  |  |   |  __ \  
  \ \  \   /  _ \  |   __|   _ \   __ ` + "`" + `__ \    _ \      __|   _ \      \___ \    _ \   __ \    _` + "`" + ` |  |\/ |  |   |  |   | 
   \ \  \ /   __/  |  (     (   |  |   |   |   __/      |    (   |           |  (   |  |   |  (   |  |   |  |   |  |   | 
    \_/\_/  \___| _| \___| \___/  _|  _|  _| \___|     \__| \___/      _____/  \___/  _|  _| \__,_| _|  _| \___/  ____/  

`


func main() {
	fmt.Println(welcomeBanner)
	errorHandler := make(chan error)
	quit := make(chan bool)

	// logging
	f, err := os.OpenFile("mud.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("error opening mud.log: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	c, err := mud.LoginWithOS()
	if err != nil {
		log.Fatal(err)
	}
	go c.PromptWithOS(quit, errorHandler)
	for {
		select {
		case <-quit:
			return
		case err := <-errorHandler:
			log.Println(err)
		}
	}
}
