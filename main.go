package main
// created by Eilidh Robey erobey@sonatype.com

import (
	"fmt"
	mud "gomud/adventure"
)

const welcomeBanner = `
 \ \        /      |                                    |                ___|                         \  |  |   |  __ \  
  \ \  \   /  _ \  |   __|   _ \   __ ` + "`" + `__ \    _ \      __|   _ \      \___ \    _ \   __ \    _` + "`" + ` |  |\/ |  |   |  |   | 
   \ \  \ /   __/  |  (     (   |  |   |   |   __/      |    (   |           |  (   |  |   |  (   |  |   |  |   |  |   | 
    \_/\_/  \___| _| \___| \___/  _|  _|  _| \___|     \__| \___/      _____/  \___/  _|  _| \__,_| _|  _| \___/  ____/  

`


func main() {
	fmt.Println(welcomeBanner)
	errorHandler := make(chan string)
	quit := make(chan bool)

	c := mud.Login()
	go c.Prompt(quit, errorHandler)
	for {
		select {
		case <-quit:
			return
		case err := <-errorHandler:
			fmt.Printf("%s\n", err)
		}
	}
}
