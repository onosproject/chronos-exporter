// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package manager

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/onosproject/chronos-exporter/pkg/collector"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

var log = logging.GetLogger("manager")

// Manager single point of entry for the topology system.
type Manager struct {
	Config       collector.AetherModel
	ImagePath    string
	SitePlanPath string
}

func NewManager(configData []byte, imagePath, sitePlanPath string) *Manager {
	modelConfig, err := collector.LoadModel(configData)
	if err != nil {
		log.Fatal("Error unmarshalling configuration", err)
	}
	log.Infof("Config model loaded, with %d sites", len(modelConfig.Sites))
	return &Manager{
		Config:       *modelConfig,
		ImagePath:    imagePath,
		SitePlanPath: sitePlanPath,
	}
}

func (mgr *Manager) Run() {
	router := gin.New()
	mgr.Config.Collect()

	router.GET("/metrics", prometheusHandler())
	router.GET("/config", mgr.handleConfig)
	router.GET("/images/:image", mgr.handleImage)
	router.GET("/site-plans/:site/:site-plan", mgr.handleSitePlan)

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: false,
		AllowOrigins:    []string{"http://localhost:4200"},
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:4200"
		},
		MaxAge:           12 * time.Hour,
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
	}))

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
	switch c.Request.Method {
	case "GET":
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, mgr.Config)
	default:
		c.Status(http.StatusNotImplemented)
	}
}

func (mgr *Manager) handleImage(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		c.Header("Content-Type", "image/png")
		image, okay := c.Params.Get("image")
		if !okay {
			c.String(http.StatusNotFound, "error")
		}
		imagePath := mgr.ImagePath + image
		c.File(imagePath)
		log.Infof("Serving file %s", imagePath)
	default:
		c.Status(http.StatusNotImplemented)
	}
}

func (mgr *Manager) handleSitePlan(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		c.Header("Content-Type", "image/svg+xml")
		sitePath, okPath := c.Params.Get("site")
		sitePlan, okPlan := c.Params.Get("site-plan")
		if !okPath || !okPlan {
			c.String(http.StatusNotFound, "error")
		}
		sitePlanPath := mgr.SitePlanPath + sitePath + "/" + sitePlan
		c.File(sitePlanPath)
		log.Infof("Serving file %s", sitePlanPath)
	default:
		c.Status(http.StatusNotImplemented)
	}
}
