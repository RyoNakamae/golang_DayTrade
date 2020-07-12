package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	_url := "https://www.nikkei.com/markets/worldidx/chart/nk225/"

	doc, err := goquery.NewDocument(_url)
	if err != nil {
		panic(err)
	}

	u := url.URL{}
	u.Scheme = doc.Url.Scheme
	u.Host = doc.Url.Host

	// ページtitleの取得
	// title := doc.Find("title").Text()
	// fmt.Println(title)

	//economic_value_now のものが日経平均っぽい
	economic_value_now := doc.Find(".economic_value_now").Text()
	fmt.Println(economic_value_now)
	economic_value_now = strings.Replace(economic_value_now, ",", "", -1)
	fmt.Println(economic_value_now)

	economic_value_time := doc.Find(".economic_value_time").Text()
	fmt.Println(economic_value_time)

	sampleText := strconv.Itoa(2)
	sampleText2 := fmt.Sprint(2.123)
	sampleText3 := "abc,123"

	now := time.Now().Format("2006/01/02")

	// file2, err := os.Create("test_o2.csv")
	file2, err := os.OpenFile("test_o2.csv", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer file2.Close()

	writer := csv.NewWriter(file2)
	writer.UseCRLF = true //デフォルトはLFのみ

	writer.Write([]string{now, "01", "1", sampleText, sampleText2, sampleText3, economic_value_time, economic_value_now})

	writer.Flush()
}
