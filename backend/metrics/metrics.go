package metrics

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const durationTimer = "durationTimer"

var cMetrics = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "app_requests_duration_seconds",
		Help:       "The total requests duration in seconds",
		Objectives: map[float64]float64{0.5: 0.05, 0.95: 0.005, 0.99: 0.001},
	},
	[]string{"route"},
)

func parseURL(c *gin.Context) string {
	method := strings.ToLower(c.Request.Method)

	url := c.Request.URL.String()
	url = strings.Replace(url, "/", "_", -1)

	query := fmt.Sprintf("?%s", c.Request.URL.RawQuery)
	url = strings.Replace(url, query, "", -1)

	for _, p := range c.Params {
		url = strings.Replace(url, p.Value, fmt.Sprintf(":%s", p.Key), -1)
	}

	return fmt.Sprintf("%s%s", method, url)
}

// Middleware ...
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := parseURL(c)
		duration, err := cMetrics.GetMetricWithLabelValues(url)
		if err != nil {
			panic(err.Error())
		}
		c.Set(durationTimer, prometheus.NewTimer(duration))
	}
}

// Get ...
func Get(c *gin.Context) {
	h := promhttp.Handler()
	h.ServeHTTP(c.Writer, c.Request)
}

// RequestEnd ...
func RequestEnd(c *gin.Context) {
	durationTimer, _ := c.Get(durationTimer)
	durationTimer.(*prometheus.Timer).ObserveDuration()
}
