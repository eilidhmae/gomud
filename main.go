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
	stats := gomud.RollStats(min, max)

	fmt.Printf("Stats: %v (range %d - %d)\n", stats.Raw, min, max)
}
