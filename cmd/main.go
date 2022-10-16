package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/StepanAnisin/chickenfarm/pkg/farm"
)

func main() {
	ranch := new(farm.Ranch)
	go farm.InitRanch(ranch)

	server := &http.Server{Addr: ":8080", Handler: nil}
	go func() {
		//Запуск Хэндлера
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			eggsCount := farm.GetEggsCount(ranch)
			fmt.Fprintf(w, "Количество яиц в холодильнике: %d", eggsCount)
		})
		log.Fatal(server.ListenAndServe())
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(ctx)

}
