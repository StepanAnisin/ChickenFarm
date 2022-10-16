package farm

import (
	"log"
	"math"
	"math/rand"
	"sync"
	"time"
)

// Генерация рандомного числа в диапазоне
func random(min, max int) int64 {
	return int64(rand.Intn(max-min) + min)
}

// Функция спауна яиц курицей
// 1. Генерируем время через сколько заспаунилось яйцо (eggSpawnDelay)
// 2. Генерируем сколько яиц заспаунилось (eggsMin, eggsMax)
// 3. Складываем в любой свободный ресурс
// 4. Приходит фермер, забирает яйца.
func CarryEggs(number int, mutex *sync.Mutex, eggsMinSpawnCount int, eggsMaxSpawnCount int,
	eggsSpawnMinDelay int, eggsSpawnMaxDelay int, eggsInFridge *int64, wg *sync.WaitGroup) {
	rand.Seed(time.Now().Unix())
	for {
		eggsSpawnDelay := random(eggsSpawnMinDelay, eggsSpawnMaxDelay)
		// Calling Sleep method
		time.Sleep(time.Duration(eggsSpawnDelay) * time.Second)
		eggsSpawnCount := random(eggsMinSpawnCount, eggsMaxSpawnCount)
		mutex.Lock()
		if math.MaxInt64-*eggsInFridge > eggsSpawnCount {
			*eggsInFridge += eggsSpawnCount
			log.Print("Курица ", number, " снесла ", eggsSpawnCount, " яиц с задержкой ", eggsSpawnDelay)
			log.Print("Количество яиц в холодильнике: ", *eggsInFridge)
		} else {
			wg.Done()
			mutex.Unlock()
			break
		}
		mutex.Unlock()
	}
}
