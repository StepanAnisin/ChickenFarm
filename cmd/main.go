package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	chicken "github.com/StepanAnisin/chickenfarm/pkg/chicken"
	config "github.com/StepanAnisin/chickenfarm/pkg/configreader"
	farmer "github.com/StepanAnisin/chickenfarm/pkg/farmer"
)

var eggsInFridge int = 0 //  общий ресурс

func main() {

	config, err := config.LoadConfig("../config/app.env")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	var wg sync.WaitGroup
	wg.Add(config.ChikensCount + 1)
	//Запуск процесса спауна яиц
	// канал
	ch := make(chan bool)
	// определяем мьютекс
	var mutex sync.Mutex
	for i := 0; i < config.ChikensCount; i++ {
		go chicken.CarryEggs(i, ch, &mutex, config.EggsMinSpawnCount, config.EggsMaxSpawnCount,
			config.EggsSpawnMinDelay, config.EggsSpawnMaxDelay, &eggsInFridge)
	}
	go farmer.FarmerComes(ch, config.FarmerCheckMinDelay, config.FarmerCheckMaxDelay, config.FarmerMaxNeededQuantity,
		config.FarmerMinNeededQuantity, &mutex, &eggsInFridge)

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
