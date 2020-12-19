package stockParser

import (
	"fmt"
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"github.com/piquette/finance-go/quote"
	"time"
)

type Stock struct {
	Name string
	Code string
	Downtime time.Time
}

func (s Stock)GetCurrentPrice() (price float64, err error) {
	fmt.Println(s)
	data, err := quote.Get(s.Code)
	if err != nil {
		return 0, err
	}
	return data.RegularMarketPrice, nil
}

func (s Stock)GetPrevPrice(interval string) (price float64, err error) {
	timeInterval := time.Hour
	datetimeInterval := datetime.OneHour

	switch{
	case interval == "hour":
		timeInterval = time.Hour
		datetimeInterval = datetime.OneHour
	case interval == "day":
		timeInterval = time.Hour * 24
		datetimeInterval = datetime.OneDay
	case interval == "week":
		timeInterval = time.Hour * 24 * 5
		datetimeInterval = datetime.FiveDay
	}

	startTime := time.Now().Add(timeInterval * -2)
	endTime := time.Now().Add(timeInterval * -1)
	data := &chart.Params{
		Symbol:   s.Code,
		Start:    datetime.New(&startTime),
		End:      datetime.New(&endTime),
		Interval: datetimeInterval,
	}
	iter := chart.Get(data)
	for iter.Next(){
		b := iter.Bar()
		result, _ := b.Close.Float64()
		return result , nil
	}
	return 0,nil
}
