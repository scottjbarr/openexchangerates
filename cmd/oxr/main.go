package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	oxr "github.com/scottjbarr/openexchangerates"
)

func main() {
	cmd := ""
	appID := ""
	base := ""
	symbols := ""
	date := ""

	flag.StringVar(&cmd, "cmd", "", "Command latest|historical (required)")
	flag.StringVar(&appID, "appid", "", "App ID (required)")
	flag.StringVar(&base, "base", "", "Base currency")
	flag.StringVar(&symbols, "symbols", "", "Symbols comma separated")
	flag.StringVar(&date, "date", "", "Date YYYY-MM-DD")
	flag.Parse()

	if len(appID) == 0 || !isValid(cmd) {
		flag.Usage()
		os.Exit(1)
	}

	c := oxr.New(appID)

	switch cmd {
	case "latest":
		resp, err := getLatest(c, base, symbols)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n", dump(resp.LatestData))

		return

	case "historical":
		resp, err := getHistorical(c, date, base, symbols)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n", dump(resp.HistoricalData))

		return

	default:

	}
}

func isValid(cmd string) bool {
	switch cmd {
	case "latest", "historical":
		return true
	}

	return false
}

func dump(o interface{}) []byte {
	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}

	return b
}

func getLatest(c oxr.Client, base, symbols string) (*oxr.LatestResponse, error) {
	currencies := strings.Split(symbols, ",")

	params := oxr.LatestParams{
		Base:    "USD",
		Symbols: currencies,
	}

	rates, err := c.Latest(&params)
	if err != nil {
		return nil, err
	}

	return rates, nil
}

func getHistorical(c oxr.Client, date, base, symbols string) (*oxr.HistoricalResponse, error) {
	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}

	currencies := strings.Split(symbols, ",")

	params := oxr.HistoricalParams{
		Base:    "USD",
		Symbols: currencies,
	}

	rates, err := c.Historical(d, &params)
	if err != nil {
		return nil, err
	}

	return rates, nil
}
