// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package collector

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// ServerImpl -
type ServerImpl struct {
	config *AetherModel
}

func NewServer(config *AetherModel) ServerInterface {
	return &ServerImpl{
		config: config,
	}
}

// GetAetherConfiguration (GET /config)
func (s *ServerImpl) GetAetherConfiguration(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, s.config)
}

// DeleteApplicationsApplication (DELETE /config/applications/{application-id})
func (s *ServerImpl) DeleteApplicationsApplication(ctx echo.Context, applicationId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// GetApplicationsApplication (GET /config/applications/{application-id})
func (s *ServerImpl) GetApplicationsApplication(ctx echo.Context, applicationId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// PostApplicationsApplication (POST /config/applications/{application-id})
func (s *ServerImpl) PostApplicationsApplication(ctx echo.Context, applicationId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// DeleteSitesSite (DELETE /config/sites/{site-id})
func (s *ServerImpl) DeleteSitesSite(ctx echo.Context, siteId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// GetSitesSite (GET /config/sites/{site-id})
func (s *ServerImpl) GetSitesSite(ctx echo.Context, siteId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// PostSitesSite (POST /config/sites/{site-id})
func (s *ServerImpl) PostSitesSite(ctx echo.Context, siteId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// DeleteSitesSiteDeviceGroupsDeviceGroup (DELETE /config/sites/{site-id}/device-groups/{device-group-id})
func (s *ServerImpl) DeleteSitesSiteDeviceGroupsDeviceGroup(ctx echo.Context, siteId string, deviceGroupId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// GetSitesSiteDeviceGroupsDeviceGroup (GET /config/sites/{site-id}/device-groups/{device-group-id})
func (s *ServerImpl) GetSitesSiteDeviceGroupsDeviceGroup(ctx echo.Context, siteId string, deviceGroupId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// PostSitesSiteDeviceGroupsDeviceGroup (POST /config/sites/{site-id}/device-groups/{device-group-id})
func (s *ServerImpl) PostSitesSiteDeviceGroupsDeviceGroup(ctx echo.Context, siteId string, deviceGroupId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// DeleteSitesSiteDeviceGroupsDeviceGroupDevicesDevice (DELETE /config/sites/{site-id}/device-groups/{device-group-id}/devices/{dev-id})
func (s *ServerImpl) DeleteSitesSiteDeviceGroupsDeviceGroupDevicesDevice(ctx echo.Context, siteId string, deviceGroupId string, devId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// GetSitesSiteDeviceGroupsDeviceGroupDevicesDevice (GET /config/sites/{site-id}/device-groups/{device-group-id}/devices/{dev-id})
func (s *ServerImpl) GetSitesSiteDeviceGroupsDeviceGroupDevicesDevice(ctx echo.Context, siteId string, deviceGroupId string, devId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// PostSitesSiteDeviceGroupsDeviceGroupDevicesDevice (POST /config/sites/{site-id}/device-groups/{device-group-id}/devices/{dev-id})
func (s *ServerImpl) PostSitesSiteDeviceGroupsDeviceGroupDevicesDevice(ctx echo.Context, siteId string, deviceGroupId string, devId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// DeleteSitesSiteDevicesDevice (DELETE /config/sites/{site-id}/devices/{serial-number})
func (s *ServerImpl) DeleteSitesSiteDevicesDevice(ctx echo.Context, siteId string, serialNumber string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// GetSitesSiteDevicesDevice (POST /config/sites/{site-id}/devices/{serial-number})
func (s *ServerImpl) GetSitesSiteDevicesDevice(ctx echo.Context, siteId string, serialNumber string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// PostSitesSiteDevicesDevice (POST /config/sites/{site-id}/devices/{serial-number})
func (s *ServerImpl) PostSitesSiteDevicesDevice(ctx echo.Context, siteId string, serialNumber string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// DeleteSitesSiteSimsSim (DELETE /config/sites/{site-id}/sims/{iccid})
func (s *ServerImpl) DeleteSitesSiteSimsSim(ctx echo.Context, siteId string, iccid string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// GetSitesSiteSimsSim (GET /config/sites/{site-id}/sims/{iccid})
func (s *ServerImpl) GetSitesSiteSimsSim(ctx echo.Context, siteId string, iccid string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// PostSitesSiteSimsSim (POST /config/sites/{site-id}/sims/{iccid})
func (s *ServerImpl) PostSitesSiteSimsSim(ctx echo.Context, siteId string, iccid string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// DeleteSitesSiteSlicesSlice (DELETE /config/sites/{site-id}/slices/{slice-id})
func (s *ServerImpl) DeleteSitesSiteSlicesSlice(ctx echo.Context, siteId string, sliceId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// GetSitesSiteSlicesSlice (GET /config/sites/{site-id}/slices/{slice-id})
func (s *ServerImpl) GetSitesSiteSlicesSlice(ctx echo.Context, siteId string, sliceId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// PostSitesSiteSlicesSlice (POST /config/sites/{site-id}/slices/{slice-id})
func (s *ServerImpl) PostSitesSiteSlicesSlice(ctx echo.Context, siteId string, sliceId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// DeleteSitesSiteSlicesSliceApplicationsApplication (DELETE /config/sites/{site-id}/slices/{slice-id}/applications/{app-id})
func (s *ServerImpl) DeleteSitesSiteSlicesSliceApplicationsApplication(ctx echo.Context, siteId string, sliceId string, appId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// GetSitesSiteSlicesSliceApplicationsApplication (GET /config/sites/{site-id}/slices/{slice-id}/applications/{app-id})
func (s *ServerImpl) GetSitesSiteSlicesSliceApplicationsApplication(ctx echo.Context, siteId string, sliceId string, appId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// PostSitesSiteSlicesSliceApplicationsApplication (POST /config/sites/{site-id}/slices/{slice-id}/applications/{app-id})
func (s *ServerImpl) PostSitesSiteSlicesSliceApplicationsApplication(ctx echo.Context, siteId string, sliceId string, appId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// DeleteSitesSiteSlicesSliceDeviceGroupsDeviceGroup (DELETE /config/sites/{site-id}/slices/{slice-id}/device-groups/{dg-id})
func (s *ServerImpl) DeleteSitesSiteSlicesSliceDeviceGroupsDeviceGroup(ctx echo.Context, siteId string, sliceId string, dgId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// GetSitesSiteSlicesSliceDeviceGroupsDeviceGroup (GET /config/sites/{site-id}/slices/{slice-id}/device-groups/{dg-id})
func (s *ServerImpl) GetSitesSiteSlicesSliceDeviceGroupsDeviceGroup(ctx echo.Context, siteId string, sliceId string, dgId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// PostSitesSiteSlicesSliceDeviceGroupsDeviceGroup (POST /config/sites/{site-id}/slices/{slice-id}/device-groups/{dg-id})
func (s *ServerImpl) PostSitesSiteSlicesSliceDeviceGroupsDeviceGroup(ctx echo.Context, siteId string, sliceId string, dgId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// DeleteSitesSiteSmallCellsSmallCell (DELETE /config/sites/{site-id}/small-cells/{small-cell-id})
func (s *ServerImpl) DeleteSitesSiteSmallCellsSmallCell(ctx echo.Context, siteId string, smallCellId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// GetSitesSiteSmallCellsSmallCell (GET /config/sites/{site-id}/small-cells/{small-cell-id})
func (s *ServerImpl) GetSitesSiteSmallCellsSmallCell(ctx echo.Context, siteId string, smallCellId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// PostSitesSiteSmallCellsSmallCell (POST /config/sites/{site-id}/small-cells/{small-cell-id})
func (s *ServerImpl) PostSitesSiteSmallCellsSmallCell(ctx echo.Context, siteId string, smallCellId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}
