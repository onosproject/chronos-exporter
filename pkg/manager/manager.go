// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package manager

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/onosproject/chronos-exporter/pkg/aether"
	"github.com/onosproject/chronos-exporter/pkg/alerts"
	"github.com/onosproject/chronos-exporter/pkg/collector"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var log = logging.GetLogger("manager")

// Manager single point of entry for the topology system.
type Manager struct {
	echoRouter  *echo.Echo
	configModel *collector.AetherModel
	aetherModel *aether.Enterprises
	alertModel  *alerts.AetherAlerts
}

func NewManager(configData []byte, aetherData []byte, alertData []byte, imagePath,
	sitePlanPath string, allowCorsOrigins []string) *Manager {

	modelConfig, err := collector.LoadModel(configData)
	if err != nil {
		log.Fatal("Error unmarshalling configuration", err)
	}

	aetherConfig, err := aether.LoadModel(aetherData)
	if err != nil {
		log.Fatal("Error unmarshalling configuration", err)
	}

	modelAlert, err := alerts.LoadModel(alertData)
	if err != nil {
		log.Fatal("Error unmarshalling alerts", err)
	}

	log.Infof("Config model loaded, with %d sites", len(modelConfig.Sites))
	mgr := Manager{
		configModel: modelConfig,
		alertModel:  modelAlert,
		aetherModel: aetherConfig,
	}

	mgr.echoRouter = echo.New()
	if len(allowCorsOrigins) > 0 {
		mgr.echoRouter.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: allowCorsOrigins,
			AllowHeaders: []string{echo.HeaderAccessControlAllowOrigin, echo.HeaderContentType, echo.HeaderAuthorization},
		}))
	}

	collector.RegisterHandlers(mgr.echoRouter, collector.NewServer(modelConfig))
	aether.RegisterHandlers(mgr.echoRouter, aether.NewServer(aetherConfig))
	alerts.RegisterHandlers(mgr.echoRouter, alerts.NewServer(modelAlert))

	mgr.echoRouter.Use(middleware.Static(imagePath))
	mgr.echoRouter.Use(middleware.Static(sitePlanPath))

	mgr.echoRouter.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	return &mgr
}

func (mgr *Manager) Run(port uint) {
	log.Infof("Starting Manager on port %d", port)
	mgr.configModel.Collect()
	mgr.aetherModel.Collect()
	mgr.echoRouter.Logger.Fatal(mgr.echoRouter.Start(fmt.Sprintf(":%d", port)))
	log.Warn("Manager Stopping")
}
