package handlers

import (
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/Emy/prom-opendata-zh-parking/internal/types"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/robfig/cron"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

var promReg = prometheus.NewRegistry()
var promHandler = promhttp.HandlerFor(
	promReg,
	promhttp.HandlerOpts{
		EnableOpenMetrics: true,
	},
)

// The frame counter gets increased whenever an update has been polled from the open data platform.
// Its intended purpose is to indicate if there had been an update the last time data was requested
// from the prometheus instances.
//
// The counter gets reset to zero when the server is restarted.
var frameCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name:      "frames_total",
		Namespace: "zurich_parking",
		Help:      "Indicates if there had been a change in the values. (Increments when new data is pulled from the open data platform. Zeroes the counter on server restart)",
	},
)

var freeSpaces = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: "zurich_parking",
		Name:      "free_spaces",
		Help:      "Number of free parking spaces per parking lot in Zurich",
	},
	[]string{"lot"}, // Labels for grouping
)

// This function is being called to initialize all needed handlers and event schedulers that are required for an automatic
// retrieval of parking data from the open data platform.
func InitializePrometheusHandling() {
	updatePrometheusData()
	promReg.MustRegister(frameCounter)
	promReg.MustRegister(freeSpaces)
	http.Handle("/metrics", promHandler)
	logger.Info("promHandler.go: Prometheus Handler initialized.")
	enableAutoScheduledFetch()
}

func updatePrometheusData() {
	logger.Debug("promHandler.go: Updating API data...")

	fetchedData := fetchData()
	if fetchedData == nil {
		return
	}

	for _, item := range []types.ParkingData(*fetchedData) {
		freeSpaces.With(prometheus.Labels{"lot": item.Name}).Set(float64(item.SpacesLeft))
	}
}

func fetchData() *[]types.ParkingData {
	base, err := url.Parse("https://www.pls-zh.ch")
	if err != nil {
		logger.Error("promHandler.go: Could not parse baseURL.")
		return nil
	}
	base.Path += "/plsFeed/rss"

	client := &http.Client{
		Transport: &http.Transport{},
	}

	req, err := http.NewRequest("GET", base.String(), nil)
	if err != nil {
		// Handle error
		return nil
	}

	// We do a tiny amount of tomfoolery in this line :3
	req.Header.Set("User-Agent", randomUA())

	response, err := client.Do(req)
	if err != nil {
		logger.Error("promHandler.go: WebRequest > Could not read response body.", "Error", err)
		return nil
	}
	logger.Debug("promHandler.go: WebRequest > Got Response with", "Statuscode", response.StatusCode)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	var result types.RSS
	xml.Unmarshal([]byte(body), &result)

	var parkingData []types.ParkingData
	for _, item := range result.Channel.Items {
		// Split description to get status and available spaces
		descParts := strings.Split(item.Description, " / ")
		if len(descParts) < 2 {
			continue
		}

		status := descParts[0]
		spaces := 0
		fmt.Sscanf(descParts[1], "%d", &spaces)

		parkingData = append(parkingData, types.ParkingData{
			Name:       item.Title,
			URL:        item.Link,
			Status:     status,
			SpacesLeft: spaces,
		})
	}
	logger.Debug("promHandler.go: Fetched and unmarshalled the data from the open data platform sucessfully.")
	frameCounter.Inc()
	return &parkingData
}

// This function registers and calls the scheduled update job to fetch new data from the open data platform.
//
// Currently the interval is set that the job gets executed every 5 minutes on the clock with a 15 second delay to
// account for delays on the open data platform data base.
func enableAutoScheduledFetch() {
	c := cron.New()
	c.AddFunc("15 */5 * * *", func() { updatePrometheusData() }) // every 5 minutes with a 15 second delay.
	c.Start()
	logger.Info("promHandler.go: Enabled scheduled fetching of API responses from the open data platform.")
}
