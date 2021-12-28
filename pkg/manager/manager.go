// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/onosproject/chronos-exporter/pkg/collector"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var log = logging.GetLogger("manager")

// Manager single point of entry for the topology system.
type Manager struct {
	Config collector.AetherModel
	ImagePath string
}

func NewManager(configData []byte, imagePath string) *Manager {
	modelConfig, err := collector.LoadModel(configData)
	if err != nil {
		log.Fatal("Error unmarshalling configuration %v", err)
	}
	log.Infof("Config model loaded, with %d sites", len(modelConfig.Sites))
	return &Manager{
		Config: *modelConfig,
		ImagePath: imagePath,
	}
}

func (mgr *Manager) Run() {
	router := gin.New()
	mgr.Config.Collect()

	router.GET("/metrics", prometheusHandler())
	router.GET("/config", mgr.handleConfig)
	router.OPTIONS("/config", mgr.handleConfig)

	router.GET("/images/:image", mgr.handleImage)
	router.OPTIONS("/images/:image", mgr.handleImage)

	err := router.Run("0.0.0.0:" + "2112")
	if err != nil {
		log.Fatal(err)
	}
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (mgr *Manager) handleConfig(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:4200")
	switch c.Request.Method {
	case "OPTIONS":
		c.Header("Vary", "Origin")
		c.Header("Vary", "Access-Control-Request-Method")
		c.Header("Vary", "Access-Control-Request-Headers")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token, Authorization")
		c.Header("Access-Control-Allow-Methods", "GET, OPTIONS")
		c.Status(http.StatusOK)
	case "GET":
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, mgr.Config)
	default:
		c.Status(http.StatusNotImplemented)
	}
}

func (mgr *Manager) handleImage(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:4200")
	switch c.Request.Method {
	case "OPTIONS":
		c.Header("Vary", "Origin")
		c.Header("Vary", "Access-Control-Request-Method")
		c.Header("Vary", "Access-Control-Request-Headers")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token, Authorization")
		c.Header("Access-Control-Allow-Methods", "GET, OPTIONS")
		c.Status(http.StatusOK)
	case "GET":
		c.Header("Content-Type", "image/png")
		image, okay := c.Params.Get("image")
		if !okay {
			c.String(http.StatusNotFound, "error")
		}
		imagePath := mgr.ImagePath + image
		c.File(imagePath)
	default:
		c.Status(http.StatusNotImplemented)
	}
}
