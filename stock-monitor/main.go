package main

import (
	"fmt"
	"github.com/stock-monitor/internal/configReader"
	"github.com/stock-monitor/internal/stocksNames"
	"github.com/stock-monitor/internal/tg"
	"github.com/stock-monitor/internal/worker"
	"sync"
	"time"
)

func main(){
	config := configReader.Config{}
	err := config.Parse("/Users/tima-fey/go/src/github.com/stock-monitor/config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	bot, err := tg.StartBot(config.Token)
	if err != nil {
		fmt.Println(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go worker.Scheduler(stocksNames.StocksLite, 1, 2, 3, time.Minute, bot,136012973 )
	wg.Wait()
}