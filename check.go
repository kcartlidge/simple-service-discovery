package main

import (
	"net/http"
	"sync"

	config "github.com/kcartlidge/simples-config"
)

var result EndpointResults

// EndpointResults ... The outcome of all endpoint checks.
type EndpointResults struct {
	mtx       sync.Mutex
	Endpoints []EndpointResult `json:"endpoints"`
}

// EndpointResult ... The outcome of a single endpoint's check.
type EndpointResult struct {
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
	Status   int    `json:"status"`
}

func performChecks(endpoints map[int]config.Entry) {
	r := EndpointResults{}
	r.Endpoints = []EndpointResult{}

	for _, ep := range endpoints {
		epr := EndpointResult{
			Name:     ep.Key,
			Endpoint: ep.Value,
			Status:   http.StatusInternalServerError,
		}
		res, err := http.Get(epr.Endpoint)
		if err == nil {
			res.Body.Close()
			epr.Status = res.StatusCode
		}
		r.Endpoints = append(r.Endpoints, epr)
	}

	r.mtx.Lock()
	result.Endpoints = r.Endpoints
	r.mtx.Unlock()
}
