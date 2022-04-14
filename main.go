package main


import (
	"fmt"
	"gomud/character"
)

const welcomeBanner = `
 __          __ ______  _        _____   ____   __  __  ______ 
 \ \        / /|  ____|| |      / ____| / __ \ |  \/  ||  ____|
  \ \  /\  / / | |__   | |     | |     | |  | || \  / || |__   
   \ \/  \/ /  |  __|  | |     | |     | |  | || |\/| ||  __|  
    \  /\  /   | |____ | |____ | |____ | |__| || |  | || |____ 
     \/  \/    |______||______| \_____| \____/ |_|  |_||______|

`


func main() {
	fmt.Println(welcomeBanner)
	min := 5
	max := 18

	for _, class := range []string{"fighter","rogue","mage","cleric"} {
		fmt.Printf("Stats for %10s: %s (range %d - %d)\n", class, gomud.RollStats(min, max, class).Text(), min, max)
	}
}
