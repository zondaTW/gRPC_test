package weather

import (
	"fmt"
	"strings"
	"github.com/anaskhan96/soup"
	"os"

	crawlerPB "../../../pb/crawler"
)

type weatherInfo struct {
	date string
	weather string
	temperature string
}

func zip(lists ...[]soup.Root) func() []soup.Root {
	zip := make([]soup.Root, len(lists))
	i := 0
	return func() []soup.Root {
		for j := range lists {
			if i >= len(lists[j]) {
				return nil
			}
			zip[j] = lists[j][i]
		}
		i++
		return zip
	}
}

func GetWeatherInfo(url string) []*crawlerPB.WeatherReply_Info {
	resp, err := soup.Get(url)
	if err != nil {
		fmt.Println("http transport error is:", err)
		os.Exit(1)
	}

	doc := soup.HTMLParse(resp)
	table := doc.Find("table", "class", "FcstBoxTable01")
	thead := table.Find("thead")
	tr := thead.Find("tr")
	ths := tr.FindAll("th")

	tbody := table.Find("tbody")
	tr = tbody.Find("tr")
	weather_imgs := tr.FindAll("img")

	tbody = table.Find("tbody")
	tr = tbody.Find("tr")
	tds := tr.FindAll("td")

	weatherInfoArray := make([]*crawlerPB.WeatherReply_Info, 0)
	iter := zip(ths[1:], weather_imgs, tds)
	for tuple := iter(); tuple != nil; tuple = iter() {
		temp := strings.Replace(tuple[2].Text(), "\t", "", -1)
		temp = strings.Replace(temp, "\n", "", -1)
		temp = strings.Replace(temp, " ", "", -1)

		weatherInfoTemp := &crawlerPB.WeatherReply_Info{
			Date: tuple[0].Text(),
			Weather: tuple[1].Attrs()["title"],
			Temperature: temp}
		weatherInfoArray = append(weatherInfoArray, weatherInfoTemp)
	}
	return weatherInfoArray
}