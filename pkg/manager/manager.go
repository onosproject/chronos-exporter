// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package manager

import (
	"fmt"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/onosproject/chronos-exporter/pkg/alerts"
	"github.com/onosproject/chronos-exporter/pkg/collector"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("manager")

// Manager single point of entry for the topology system.
type Manager struct {
	ImagePath    string
	SitePlanPath string
	echoRouter   *echo.Echo
}

func NewManager(configData []byte, imagePath, sitePlanPath string, allowCorsOrigins []string) *Manager {
	modelConfig, err := collector.LoadModel(configData)
	if err != nil {
		log.Fatal("Error unmarshalling configuration", err)
	}
	log.Infof("Config model loaded, with %d sites", len(modelConfig.Sites))
	mgr := Manager{
		ImagePath:    imagePath,
		SitePlanPath: sitePlanPath,
	}

	mgr.echoRouter = echo.New()
	if len(allowCorsOrigins) > 0 {
		mgr.echoRouter.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: allowCorsOrigins,
			AllowHeaders: []string{echo.HeaderAccessControlAllowOrigin, echo.HeaderContentType, echo.HeaderAuthorization},
		}))
	}
	collector.RegisterHandlers(mgr.echoRouter, collector.NewServer(modelConfig))
	alerts.RegisterHandlers(mgr.echoRouter, alerts.NewServer())
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(mgr.echoRouter)

	mgr.echoRouter.Use(middleware.Static(imagePath))
	mgr.echoRouter.Use(middleware.Static(sitePlanPath))

	// TODO - write a custome middleware for serving the rasa model with cache
	//  as described here https://rasa.com/docs/rasa/model-storage/#how-to-configure-your-server
	mgr.echoRouter.Use(middleware.Static("/opt"))

	return &mgr
}

func (mgr *Manager) Run(port uint) {
	log.Infof("Starting Manager on port %d", port)
	mgr.echoRouter.Logger.Fatal(mgr.echoRouter.Start(fmt.Sprintf(":%d", port)))
	log.Warn("Manager Stopping")
}
