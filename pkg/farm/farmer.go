package farm

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type Farmer struct {
}

// Фермер приходит и забирает яйца
func FarmerComes(FarmerCheckMinDelay int, FarmerCheckMaxDelay int, FarmerMaxNeededQuantity int,
	FarmerMinNeededQuantity int, mutex *sync.Mutex, eggsInFridge *int64) {
	rand.Seed(time.Now().Unix())
	for {
		farmerCheckDelay := random(FarmerCheckMinDelay, FarmerCheckMaxDelay)
		// Calling Sleep method
		time.Sleep(time.Duration(farmerCheckDelay) * time.Second)
		eggsQuantityNeeded := random(FarmerMinNeededQuantity, FarmerMaxNeededQuantity)
		mutex.Lock()
		if *eggsInFridge <= eggsQuantityNeeded {
			log.Print("Фермер взял ", eggsInFridge, " яиц ")
			*eggsInFridge = 0
			log.Print("Количество яиц в холодильнике: ", *eggsInFridge)
		}
		if *eggsInFridge > eggsQuantityNeeded {
			*eggsInFridge -= eggsQuantityNeeded
			log.Print("Фермер взял ", eggsQuantityNeeded, " яиц ")
			log.Print("Количество яиц в холодильнике: ", *eggsInFridge)
		}
		mutex.Unlock()
	}
}
