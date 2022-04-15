package main


import (
	"fmt"
	"gomud/character"
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

	for _, class := range []string{"fighter","rogue","mage","cleric"} {
		fmt.Printf("Stats for %10s: %s (range %d - %d)\n", class, gomud.RollStats(min, max, class).Text(), min, max)
	}
}
