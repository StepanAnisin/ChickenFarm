package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/StepanAnisin/chickenfarm/pkg/farm"
)

func main() {
	ranch := new(farm.Ranch)
	go farm.InitRanch(ranch)

	//Запуск Хэндлера
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		eggsCount := farm.GetEggsCount(ranch)
		fmt.Fprintf(w, "Количество яиц в холодильнике: %d", eggsCount)
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
