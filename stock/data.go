package stock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Stocks struct {
	Price string
	Date  string
}

type StockResponse struct {
	MetaData        map[string]string            `json:"Meta Data"`
	TimeSeriesDaily map[string]map[string]string `json:"Time Series (Daily)"`
}

func GetDaily() Stocks {
	api := os.Getenv("STOCK_API_KEY")
	if api == "" {
		log.Print("Err no STOCK_API_KEY.")
		return Stocks{}
	}
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=VTSMX&apikey=%s", api)
	resp, err := http.Get(url)
	if err != nil {
		log.Print("http.Get err: ", err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("ioutil.ReadAll err: ", err)
	}

	s := StockResponse{}
	err = json.Unmarshal(bodyBytes, &s)
	if err != nil {
		log.Print("unmarshal error: ", err)
	}
	date := s.MetaData["3. Last Refreshed"]
	prce := s.TimeSeriesDaily[date]["4. close"]
	p := Stocks{
		Price: prce,
		Date:  date,
	}
	return p
}
