package main

import (
	"github.com/StepanAnisin/chickenfarm/pkg/farm"
	"log"
	"net/http"
)

func main() {
	ranch := new(farm.Ranch)
	farm.InitRanch(ranch)

	//Запуск Хэндлера
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		farm.GetEggsCount(ranch)
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
