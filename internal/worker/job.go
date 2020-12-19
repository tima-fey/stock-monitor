package worker

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/stock-monitor/internal/stockParser"
	"time"
)

func ComparePrice(stocks []stockParser.Stock, hourThreshold, dayThreshold, weekThreshold int) (answer string,  err error){
	answer = ""
	for _, st := range stocks{
		warn := true
		currentPrise, err := st.GetCurrentPrice()
		if err != nil {
			return "", err
		}
		hourPrise, err := st.GetPrevPrice("hour")
		if err != nil {
			return "", err
		}
		dayPrise, err := st.GetPrevPrice("day")
		if err != nil {
			return "", err
		}
		weekPrise, err := st.GetPrevPrice("week")
		if err != nil {
			return "", err
		}
		var hourLost int
		var dayLost int
		var weekLost int
		if hourPrise != 0 {
			hourLost = int((hourPrise - currentPrise) / hourPrise * 100)
			if  hourLost > hourThreshold{
				warn = true
			}
		}
		if dayPrise != 0 {
			dayLost = int((dayPrise - currentPrise) / dayPrise * 100)
			if dayLost > dayThreshold {
				warn = true
			}
		}
		if weekPrise != 0 {
			weekLost = int((weekPrise - currentPrise) / weekPrise * 100)

			if weekLost > weekThreshold {
				warn = true
			}
		}
		if warn == true {
			answer = fmt.Sprintf("%s \n Name %s Code %s hour %d%% day %d%% week %d%%", answer, st.Name, st.Code, hourLost, dayLost, weekLost)
		}
		time.Sleep(time.Second)
	}
	return answer, nil
}

func Scheduler(stocks []stockParser.Stock, hourThreshold, dayThreshold, weekThreshold int, interval time.Duration, botApi *tgbotapi.BotAPI, user int64)  {
	for {
		if checkTime() == false {
			time.Sleep(interval)
			continue
		}
		answer, err:= ComparePrice(stocks, hourThreshold, dayThreshold, weekThreshold)
		if err != nil {
			fmt.Println(err)
		}
		if answer != "" {
			msg := tgbotapi.NewMessage(user, answer)
			botApi.Send(msg)
		}
		time.Sleep(interval)
	}
}
func checkTime() bool {
	if time.Now().Hour() > 22{
		return false
	}
	if time.Now().Hour() < 10{
		return false
	}
	if int(time.Now().Weekday()) == 1 {
		return false
	}
	if int(time.Now().Weekday()) == 6 {
		return false
	}
	return true
}
