package chicken

import (
	"log"
	"math"
	"math/rand"
	"sync"
	"time"
)

// Генерация рандомного числа в диапазоне
func random(min, max int) int {
	return rand.Intn(max-min) + min
}

// Функция спауна яиц курицей
// 1. Генерируем время через сколько заспаунилось яйцо (eggSpawnDelay)
// 2. Генерируем сколько яиц заспаунилось (eggsMin, eggsMax)
// 3. Складываем в любой свободный ресурс
// 4. Приходит фермер, всё забирает. Видимо, обнуляем этот некторый счетчик
func CarryEggs(number int, ch chan bool, mutex *sync.Mutex, eggsMinSpawnCount int, eggsMaxSpawnCount int,
	eggsSpawnMinDelay int, eggsSpawnMaxDelay int, eggsInFridge *int) {
	rand.Seed(time.Now().Unix())
	for {
		eggsSpawnDelay := random(eggsSpawnMinDelay, eggsSpawnMaxDelay)
		// Calling Sleep method
		time.Sleep(time.Duration(eggsSpawnDelay) * time.Second)
		eggsSpawnCount := random(eggsMinSpawnCount, eggsMaxSpawnCount)
		mutex.Lock()
		if math.MaxInt64-*eggsInFridge > eggsSpawnCount {
			*eggsInFridge += eggsSpawnCount
		} else {
			*eggsInFridge = 0
		}
		log.Print("Курица ", number, " снесла ", eggsSpawnCount, " яиц с задержкой ", eggsSpawnDelay)
		log.Print("Количество яиц в холодильнике: ", *eggsInFridge)
		mutex.Unlock()
	}
}
