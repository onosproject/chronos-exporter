// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package alerts

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// ServerImpl -
type ServerImpl struct {
}

func NewServer() ServerInterface {
	return new(ServerImpl)
}

// GetAetherAlerts (GET /alerts)
func (s *ServerImpl) GetAetherAlerts(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// ClearAlert (POST /alerts/{alert-id}/clear)
func (s *ServerImpl) ClearAlert(ctx echo.Context, alertId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}
