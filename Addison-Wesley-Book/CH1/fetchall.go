package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

// fetch multiple urls concurrently using go routines and channels

var wg sync.WaitGroup
type urlInfo struct {
	URLName string
	time int64
	size int
	err string
}


func fetch(url string) (*urlInfo, error){
	data := urlInfo{URLName:url}

	s := time.Now().Unix()

	response, err := http.Get(url)
	if err!= nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err!= nil {
		return nil, err
	}

	data.size = len(body)
	response.Body.Close()
	data.time = time.Now().Unix() - s
	return &data, nil
}


func main(){
	for _, url := range os.Args[1:]{
		wg.Add(1)
		go func() {
			info, err := fetch(url)
			if err != nil{
				panic(err)
			}

			fmt.Println(info)
			wg.Done()
		}()
	}
	wg.Wait()
}
