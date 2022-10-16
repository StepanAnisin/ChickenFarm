package farm

import (
	"fmt"
	"github.com/StepanAnisin/chickenfarm/pkg/config"
	"log"
	"sync"
)

type Ranch struct {
	eggsInFridge int
	mutex        sync.Mutex
}

func InitRanch(ranch *Ranch) {
	cfg, err := config.LoadConfig("config/app.env")
	if err != nil {
		log.Fatal("can not load config:", err)
	}
	var wg sync.WaitGroup
	wg.Add(cfg.ChikensCount + 1)
	defer wg.Wait()
	for i := 0; i < cfg.ChikensCount; i++ {
		go CarryEggs(i, &ranch.mutex, cfg.EggsMinSpawnCount, cfg.EggsMaxSpawnCount,
			cfg.EggsSpawnMinDelay, cfg.EggsSpawnMaxDelay, &ranch.eggsInFridge)
	}
	go FarmerComes(cfg.FarmerCheckMinDelay, cfg.FarmerCheckMaxDelay, cfg.FarmerMaxNeededQuantity,
		cfg.FarmerMinNeededQuantity, &ranch.mutex, &ranch.eggsInFridge)
}

func GetEggsCount(ranch *Ranch) {
	ranch.mutex.Lock()
	fmt.Print(ranch.eggsInFridge)
	fmt.Printf("Количество яиц в холодильнике: %d", ranch.eggsInFridge)
	ranch.mutex.Unlock()
}
