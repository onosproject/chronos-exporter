// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package alerts

import (
	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v2"
	"net/http"
	"time"
)

// ServerImpl -
type ServerImpl struct {
	alert *AetherAlerts
}

type response struct {
	Message string `json:"message"`
	AlertId string `json:"alert-id"`
}

func NewServer(alert *AetherAlerts) ServerInterface {
	return &ServerImpl{
		alert: alert,
	}
}

func LoadModel(modelData []byte) (*AetherAlerts, error) {

	model := new(AetherAlerts)

	err := yaml.UnmarshalStrict(modelData, model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

// GetAetherAlerts (GET /alerts)
func (s *ServerImpl) GetAetherAlerts(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, s.alert)
}

// ClearAlert (POST /alerts/{alert-id}/clear)
func (s *ServerImpl) ClearAlert(ctx echo.Context, alertId string) error {

	for index, al := range *s.alert {
		if al.AlertId == alertId {
			var timeNow time.Time
			timeNow = time.Now()
			s1 := *s.alert
			s1[index].ClearedAt = &timeNow
			resp := &response{
				Message: "Alert cleared",
				AlertId: alertId,
			}
			return ctx.JSON(http.StatusOK, &resp)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound)
}
