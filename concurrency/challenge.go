package concurrency

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
)

var dSize uint64

func downloadSize(url string) (int, error) {
	//return 2000000, nil
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("bad status: %d %s", resp.StatusCode, resp.Status)
	}

	return strconv.Atoi(resp.Header.Get("Content-Length"))
}

func poolOfData(ch <-chan string, wg *sync.WaitGroup) {
	for stringValue := range ch {
		ds, err := downloadSize(stringValue)
		if err != nil {
			fmt.Println(err)
			wg.Done()
			continue
		}
		atomic.AddUint64(&dSize, uint64(ds))
		wg.Done()
	}
}

func downloadsSize(urls []string) (int, error) {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(len(urls))

	for i := 0; i < runtime.NumCPU(); i++ {
		go poolOfData(ch, &wg)
	}

	for _, url := range urls {
		ch <- url
	}

	wg.Wait()
	close(ch)

	return int(dSize), nil
}

func gen2020URLs() []string {
	var urls []string
	urlTemplate := "https://d37ci6vzurychx.cloudfront.net/trip-data/%s_tripdata_2020-%02d.parquet"
	// https://d37ci6vzurychx.cloudfront.net/trip-data/yellow_tripdata_2022-03.parquet
	// https://s3.amazonaws.com/nyc-tlc/trip+data/%s_tripdata_2020-%02d.csv
	for _, vendor := range []string{"yellow", "green"} {
		for month := 1; month <= 12; month++ {
			url := fmt.Sprintf(urlTemplate, vendor, month)
			urls = append(urls, url)
		}
	}
	return urls
}

func testChallenge() {
	urls := gen2020URLs()
	size, err := downloadsSize(urls)
	if err != nil {
		log.Fatal(err)
	}

	sizeGB := float64(size) / (1 << 30)
	fmt.Printf("size = %.2fGB\n", sizeGB)
}
