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
	min := 5
	max := 18
	quit := make(chan bool)

	class := mud.ClassPrompt()
	fmt.Printf("Stats for %10s: %s\n\n", class, mud.RollStats(min, max, class).Text())

	// start a thread handling commands from user
	go mud.Prompt(quit)
	
	// help: displays choices with brief explanation
	// inventory: display contents of inventory
	// look: describe current room
	// north/south/east/west: enter another room in given direction
	// fight: attack mob in current room (create fight thread that returns to action thread)
	// wait for quit signal from Prompt()
	<-quit
}
