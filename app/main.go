package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"sync"
	"time"

	Config "github.com/StepanAnisin/Config"
)

var eggsInFridge int = 0 //  общий ресурс

func main() {

	config, err := Config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	rand.Seed(time.Now().Unix())
	var wg sync.WaitGroup
	wg.Add(6)

	//Запуск процесса спауна яиц
	// канал
	ch := make(chan bool)
	// определяем мьютекс
	var mutex sync.Mutex
	for i := 0; i < config.ChikensCount; i++ {
		go carryEggs(i, ch, &mutex, config.EggsMinSpawnCount, config.EggsMaxSpawnCount,
			config.EggsSpawnMinDelay, config.EggsSpawnMaxDelay)
	}
	go farmerComes(ch, config.FarmerCheckMinDelay, config.FarmerCheckMaxDelay, config.FarmerMaxNeededQuantity,
		config.FarmerMinNeededQuantity, &mutex)

	//Запуск Хэндлера
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		fmt.Print(eggsInFridge)
		fmt.Fprintf(w, "Количество яиц в холодильнике: ", eggsInFridge)
		mutex.Unlock()
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
	wg.Wait()
}

// Генерация рандомного числа в диапазоне
func random(min, max int) int {
	return rand.Intn(max-min) + min
}

// Функция спауна яиц курицей
// 1. Генерируем время через сколько заспаунилось яйцо (eggSpawnDelay)
// 2. Генерируем сколько яиц заспаунилось (eggsMin, eggsMax)
// 3. Складываем в любой свободный ресурс
// 4. Приходит фермер, всё забирает. Видимо, обнуляем этот некторый счетчик
func carryEggs(number int, ch chan bool, mutex *sync.Mutex, eggsMinSpawnCount int, eggsMaxSpawnCount int,
	eggsSpawnMinDelay int, eggsSpawnMaxDelay int) {
	for {
		eggsSpawnDelay := random(eggsSpawnMinDelay, eggsSpawnMaxDelay)
		// Calling Sleep method
		time.Sleep(time.Duration(eggsSpawnDelay) * time.Second)
		eggsSpawnCount := random(eggsMinSpawnCount, eggsMaxSpawnCount)
		mutex.Lock()
		if math.MaxInt64-eggsInFridge > eggsSpawnCount {
			eggsInFridge += eggsSpawnCount
		} else {
			eggsInFridge = 0
		}
		log.Print("Курица ", number, " снесла ", eggsSpawnCount, " яиц с задержкой ", eggsSpawnDelay)
		log.Print("Количество яиц в холодильнике: ", eggsInFridge)
		mutex.Unlock()
	}
}

// Фермер приходит и забирает яйца
func farmerComes(ch chan bool, FarmerCheckMinDelay int, FarmerCheckMaxDelay int, FarmerMaxNeededQuantity int,
	FarmerMinNeededQuantity int, mutex *sync.Mutex) {
	for {
		farmerCheckDelay := random(FarmerCheckMinDelay, FarmerCheckMaxDelay)
		// Calling Sleep method
		time.Sleep(time.Duration(farmerCheckDelay) * time.Second)
		eggsQuantityNeeded := random(FarmerMinNeededQuantity, FarmerMaxNeededQuantity)
		mutex.Lock()
		if eggsInFridge <= eggsQuantityNeeded {
			log.Print("Фермер взял ", eggsInFridge, " яиц ")
			eggsInFridge = 0
			log.Print("Количество яиц в холодильнике: ", eggsInFridge)
		}
		if eggsInFridge > eggsQuantityNeeded {
			eggsInFridge -= eggsQuantityNeeded
			log.Print("Фермер взял ", eggsQuantityNeeded, " яиц ")
			log.Print("Количество яиц в холодильнике: ", eggsInFridge)
		}
		mutex.Unlock()
	}
}
