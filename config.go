package main

import (
//	"context"
//	"encoding/json"
//	"errors"
//	"fmt"
	"log"
//	"net/http"
//	"os"
	"sync"
	"sync/atomic"
)

func configure() (*AppConfig, error) {
	return &AppConfig{
		buffer:			make(chan struct{}, 2),
		pool:			make(chan func()),
		totalApiCalls: 		atomic.Int32{},
		wg: 			&sync.WaitGroup{},
	}, nil
}

type AppConfig struct {
	buffer			chan struct{}
	pool			chan func()
	totalApiCalls 		atomic.Int32
	wg			*sync.WaitGroup
}

func (cfg *AppConfig) Start() {
	go func() {
		for {
			select {
				case f := <-cfg.pool:
					log.Println("Large task started")
					cfg.buffer <- struct{}{}
					go f()
			}
		}
	}()
}





























/*
func (cfg *crawlerConfig) addPage(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, exists := cfg.pages[normalizedURL]; exists {
		cfg.pages[normalizedURL]++
		return false
	}
	cfg.pages[normalizedURL]++
	return true
} 
*/



/*
	type crawlerConfig struct {
	pages 					map[string]int
	baseURL					*url.URL
	concurrencyControl			chan struct{}
	mu 					*sync.Mutex
	wg 					*sync.WaitGroup
	}

	func configure(rawBaseURL string, maxConcurrency int) (*crawlerConfig, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error: '%v' occured while attempting to parse URL: '%s' from config", err, rawBaseURL)
	}

	return &crawlerConfig{
		pages: 			make(map[string]int),
		baseURL: 		baseURL,
		concurrencyControl: 	make(chan struct{}, maxConcurrency),
		mu: 			&sync.Mutex{},
		wg: 			&sync.WaitGroup{},
	}, nil
}
*/
