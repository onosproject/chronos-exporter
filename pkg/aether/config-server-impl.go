// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package aether

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/onosproject/chronos-exporter/pkg/utils"
	"net/http"
)

// ServerImpl -
type ServerImpl struct {
	config *Enterprises
}

func NewServer(config *Enterprises) ServerInterface {
	return &ServerImpl{
		config: config,
	}
}

// DeleteConnectivityServices DELETE /connectivity-services
// (DELETE /aether/v2.0.0/{target}/connectivity-services)
func (w *ServerImpl) DeleteConnectivityServices(ctx echo.Context, target Target) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// GetConnectivityServices GET /connectivity-services
// (GET /aether/v2.0.0/{target}/connectivity-services)
func (w *ServerImpl) GetConnectivityServices(ctx echo.Context, target Target) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// PostConnectivityServices POST /connectivity-services
// (POST /aether/v2.0.0/{target}/connectivity-services)
func (w *ServerImpl) PostConnectivityServices(ctx echo.Context, target Target) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// DeleteConnectivityServicesConnectivityService DELETE /connectivity-services/connectivity-service
// (DELETE /aether/v2.0.0/{target}/connectivity-services/connectivity-service/{connectivity-service-id})
func (w *ServerImpl) DeleteConnectivityServicesConnectivityService(ctx echo.Context, target Target, connectivityServiceId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// GET /connectivity-services/connectivity-service
// (GET /aether/v2.0.0/{target}/connectivity-services/connectivity-service/{connectivity-service-id})
func (w *ServerImpl) GetConnectivityServicesConnectivityService(ctx echo.Context, target Target, connectivityServiceId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// POST /connectivity-services/connectivity-service
// (POST /aether/v2.0.0/{target}/connectivity-services/connectivity-service/{connectivity-service-id})
func (w *ServerImpl) PostConnectivityServicesConnectivityService(ctx echo.Context, target Target, connectivityServiceId string) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// DELETE /enterprises
// (DELETE /aether/v2.0.0/{target}/enterprises)
func (w *ServerImpl) DeleteEnterprises(ctx echo.Context, target Target) error {
	if w.config == nil {
		return ctx.String(http.StatusBadRequest, "No enterprises are present")
	}
	w.config = nil
	return ctx.String(http.StatusOK, "Enterprises ate delete successfully")
}

// GET /enterprises
// (GET /aether/v2.0.0/{target}/enterprises)
func (w *ServerImpl) GetEnterprises(ctx echo.Context, target Target) error {
	return ctx.JSON(http.StatusOK, w.config)
}

// POST /enterprises
// (POST /aether/v2.0.0/{target}/enterprises)
func (w *ServerImpl) PostEnterprises(ctx echo.Context, target Target) error {
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	var ents Enterprises
	err = json.Unmarshal(body, &ents)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	w.config = &ents
	return ctx.String(http.StatusOK, "Enterprises are added successfully")
}

// DELETE /enterprises/enterprise
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id})
func (w *ServerImpl) DeleteEnterprisesEnterprise(ctx echo.Context, target Target, enterpriseId string) error {
	e1 := *w.config.Enterprise
	for index, e := range *w.config.Enterprise {
		if e.EnterpriseId == enterpriseId {
			*w.config.Enterprise = append(e1[:index], e1[index+1:]...)
			return ctx.String(http.StatusOK, "Enterprise with id:"+enterpriseId+" deleted successfully")
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id})
func (w *ServerImpl) GetEnterprisesEnterprise(ctx echo.Context, target Target, enterpriseId string) error {
	for _, e := range *w.config.Enterprise {
		if e.EnterpriseId == enterpriseId {
			return ctx.JSON(http.StatusOK, e)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id})
func (w *ServerImpl) PostEnterprisesEnterprise(ctx echo.Context, target Target, enterpriseId string) error {
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for _, e := range *w.config.Enterprise {
		if e.EnterpriseId == enterpriseId {
			return ctx.String(http.StatusConflict, "Enterprise "+*e.DisplayName+"already exists with enterprise id "+
				enterpriseId)
		}
	}

	var ent EnterprisesEnterprise
	err = json.Unmarshal(body, &ent)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	*w.config.Enterprise = append(*w.config.Enterprise, ent)
	return ctx.String(http.StatusOK, "Enterprise "+*ent.DisplayName+"added successfully")
}

// DELETE /enterprises/enterprise/{enterprise-id}/application
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/application/{application-id})
func (w *ServerImpl) DeleteEnterprisesEnterpriseApplication(ctx echo.Context, target Target, enterpriseId string, applicationId string) error {
	var flag bool
	var entIndex, appIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for aIndex, app := range *ent.Application {
				if app.ApplicationId == applicationId {
					flag = true
					entIndex = eIndex
					appIndex = aIndex
					break
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		app1 := *e1[entIndex].Application

		app1 = append(app1[:appIndex], app1[appIndex+1:]...)
		e1[entIndex].Application = &app1
		return ctx.String(http.StatusOK, "Application with Id:"+applicationId+" deleted from Enteprise Id: "+enterpriseId)
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/application
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/application/{application-id})
func (w *ServerImpl) GetEnterprisesEnterpriseApplication(ctx echo.Context, target Target, enterpriseId string, applicationId string) error {
	for _, e := range *w.config.Enterprise {
		if e.EnterpriseId == enterpriseId {
			for _, app := range *e.Application {
				if app.ApplicationId == applicationId {
					return ctx.JSON(http.StatusOK, app)
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/application
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/application/{application-id})
func (w *ServerImpl) PostEnterprisesEnterpriseApplication(ctx echo.Context, target Target, enterpriseId string, applicationId string) error {
	var flag bool
	var entIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, app := range *ent.Application {
				if app.ApplicationId == applicationId {
					return ctx.String(http.StatusConflict, "Application "+*app.DisplayName+"already exists with enterprise id "+
						enterpriseId)
				}
			}
			flag = true
			entIndex = eIndex
			break
		}
	}

	var app EnterprisesEnterpriseApplication
	err = json.Unmarshal(body, &app)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise

		*e1[entIndex].Application = append(*e1[entIndex].Application, app)
		return ctx.String(http.StatusOK, "Application with Id:"+app.ApplicationId+" added successfully to Enterprise Id: "+
			enterpriseId)
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/application/{application-id}/endpoint
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/application/{application-id}/endpoint/{endpoint-id})
func (w *ServerImpl) DeleteEnterprisesEnterpriseApplicationEndpoint(ctx echo.Context, target Target, enterpriseId string, applicationId string, endpointId string) error {
	var flag bool
	var entIndex, appIndex, eptIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for aIndex, app := range *ent.Application {
				if app.ApplicationId == applicationId {
					for epIndex, ep := range *app.Endpoint {
						if ep.EndpointId == endpointId {
							flag = true
							entIndex = eIndex
							appIndex = aIndex
							eptIndex = epIndex
							break
						}
					}
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		app1 := *e1[entIndex].Application
		ept1 := *app1[appIndex].Endpoint

		ept1 = append(ept1[:eptIndex], ept1[eptIndex+1:]...)
		app1[appIndex].Endpoint = &ept1
		return ctx.String(http.StatusOK, "Endpoint with Id: "+endpointId+" deleted successfully")
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/application/{application-id}/endpoint
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/application/{application-id}/endpoint/{endpoint-id})
func (w *ServerImpl) GetEnterprisesEnterpriseApplicationEndpoint(ctx echo.Context, target Target, enterpriseId string, applicationId string, endpointId string) error {
	for _, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, app := range *ent.Application {
				if app.ApplicationId == applicationId {
					for _, ep := range *app.Endpoint {
						if ep.EndpointId == endpointId {
							return ctx.JSON(http.StatusOK, ep)
						}
					}
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/application/{application-id}/endpoint
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/application/{application-id}/endpoint/{endpoint-id})
func (w *ServerImpl) PostEnterprisesEnterpriseApplicationEndpoint(ctx echo.Context, target Target, enterpriseId string, applicationId string, endpointId string) error {
	var flag bool
	var entIndex, appIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for aIndex, app := range *ent.Application {
				if app.ApplicationId == applicationId {
					for _, ept := range *app.Endpoint {
						if ept.EndpointId == endpointId {
							return ctx.String(http.StatusConflict, "Endpoint with Id: "+ept.EndpointId+" already exists")
						}
					}
					flag = true
					entIndex = eIndex
					appIndex = aIndex
					break
				}
			}
		}
	}

	var ept EnterprisesEnterpriseApplicationEndpoint
	err = json.Unmarshal(body, &ept)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise
		app1 := *e1[entIndex].Application

		*app1[appIndex].Endpoint = append(*app1[entIndex].Endpoint, ept)
		return ctx.String(http.StatusOK, "Endpoint with Id: "+endpointId+" added successfully")
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/application/{application-id}/endpoint/{endpoint-id}/mbr
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/application/{application-id}/endpoint/{endpoint-id}/mbr)
func (w *ServerImpl) DeleteEnterprisesEnterpriseApplicationEndpointMbr(ctx echo.Context, target Target, enterpriseId string, applicationId string, endpointId string) error {
	var flag bool
	var entIndex, appIndex, eptIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for aIndex, app := range *ent.Application {
				if app.ApplicationId == applicationId {
					for epIndex, ep := range *app.Endpoint {
						if ep.EndpointId == endpointId {
							flag = true
							entIndex = eIndex
							appIndex = aIndex
							eptIndex = epIndex
							break
						}
					}
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		app1 := *e1[entIndex].Application
		ept1 := *app1[appIndex].Endpoint

		if ept1[eptIndex].Mbr == nil {
			return ctx.String(http.StatusBadRequest, "Mbr is not present")
		}
		ept1[eptIndex].Mbr = nil
		return ctx.String(http.StatusOK, "Mbr delete from Endpoint with Id: "+endpointId)
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/application/{application-id}/endpoint/{endpoint-id}/mbr
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/application/{application-id}/endpoint/{endpoint-id}/mbr)
func (w *ServerImpl) GetEnterprisesEnterpriseApplicationEndpointMbr(ctx echo.Context, target Target, enterpriseId string, applicationId string, endpointId string) error {
	for _, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, app := range *ent.Application {
				if app.ApplicationId == applicationId {
					for _, ept := range *app.Endpoint {
						if ept.EndpointId == endpointId {
							if ept.Mbr != nil {
								return ctx.JSON(http.StatusOK, ept.Mbr)
							}
						}
					}
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/application/{application-id}/endpoint/{endpoint-id}/mbr
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/application/{application-id}/endpoint/{endpoint-id}/mbr)
func (w *ServerImpl) PostEnterprisesEnterpriseApplicationEndpointMbr(ctx echo.Context, target Target, enterpriseId string, applicationId string, endpointId string) error {
	var flag bool
	var entIndex, appIndex, eptIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for aIndex, app := range *ent.Application {
				if app.ApplicationId == applicationId {
					for epIndex, ept := range *app.Endpoint {
						if ept.EndpointId == endpointId {
							flag = true
							eptIndex = epIndex
							entIndex = eIndex
							appIndex = aIndex
							break
						}
					}
				}
			}
		}
	}

	var eptMbr EnterprisesEnterpriseApplicationEndpointMbr
	err = json.Unmarshal(body, &eptMbr)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise
		app1 := *e1[entIndex].Application
		ep1 := *app1[appIndex].Endpoint

		//TO DO: confirm which update method is fine
		/*ep1[eptIndex].Mbr.Uplink = eptMbr.Uplink
		ep1[eptIndex].Mbr.Downlink = eptMbr.Downlink*/

		ep1[eptIndex].Mbr = &eptMbr
		return ctx.String(http.StatusOK, "Mbr added to Endpoint with Id: "+endpointId+" successfully")
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/connectivity-service
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/connectivity-service/{connectivity-service})
func (w *ServerImpl) DeleteEnterprisesEnterpriseConnectivityService(ctx echo.Context, target Target, enterpriseId string, connectivityService string) error {
	var flag bool
	var entIndex, csIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for cIndex, app := range *ent.ConnectivityService {
				if app.ConnectivityService == connectivityService {
					flag = true
					entIndex = eIndex
					csIndex = cIndex
					break
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		cs1 := *e1[entIndex].ConnectivityService

		cs1 = append(cs1[:csIndex], cs1[csIndex+1:]...)
		e1[entIndex].ConnectivityService = &cs1
		return ctx.String(http.StatusOK, "Connectivity-Service with Id: "+connectivityService+" deleted successfully")
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/connectivity-service
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/connectivity-service/{connectivity-service})
func (w *ServerImpl) GetEnterprisesEnterpriseConnectivityService(ctx echo.Context, target Target, enterpriseId string, connectivityService string) error {
	for _, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, cs := range *ent.ConnectivityService {
				if cs.ConnectivityService == connectivityService {
					return ctx.JSON(http.StatusOK, cs)
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/connectivity-service
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/connectivity-service/{connectivity-service})
func (w *ServerImpl) PostEnterprisesEnterpriseConnectivityService(ctx echo.Context, target Target, enterpriseId string, connectivityService string) error {
	var flag bool
	var entIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, cs := range *ent.ConnectivityService {
				if cs.ConnectivityService == connectivityService {
					return ctx.String(http.StatusConflict, "Connectivity- Service with Id: "+connectivityService+
						"already exists with enterprise id "+
						enterpriseId)
				}
			}
			flag = true
			entIndex = eIndex
			break
		}
	}

	var cs EnterprisesEnterpriseConnectivityService
	err = json.Unmarshal(body, &cs)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise

		*e1[entIndex].ConnectivityService = append(*e1[entIndex].ConnectivityService, cs)
		return ctx.String(http.StatusOK, "Connectivity-Service with Id:"+cs.ConnectivityService+
			" added successfully to Enterprise Id: "+enterpriseId)
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/site
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id})
func (w *ServerImpl) DeleteEnterprisesEnterpriseSite(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
	var flag bool
	var entIndex, siteIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					flag = true
					entIndex = eIndex
					siteIndex = sIndex
					break
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site

		st1 = append(st1[:siteIndex], st1[siteIndex+1:]...)
		e1[entIndex].Site = &st1
		return ctx.String(http.StatusOK, "Site with Id: "+siteId+" deleted successfully")
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/site
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id})
func (w *ServerImpl) GetEnterprisesEnterpriseSite(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
	for _, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, st := range *ent.Site {
				if st.SiteId == siteId {
					return ctx.JSON(http.StatusOK, st)
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/site
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id})
func (w *ServerImpl) PostEnterprisesEnterpriseSite(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
	var flag bool
	var entIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, st := range *ent.Site {
				if st.SiteId == siteId {
					return ctx.String(http.StatusConflict, "Site with Id: "+siteId+
						"already exists with enterprise id "+enterpriseId)
				}
			}
			flag = true
			entIndex = eIndex
			break
		}
	}

	var st EnterprisesEnterpriseSite
	err = json.Unmarshal(body, &st)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise

		*e1[entIndex].Site = append(*e1[entIndex].Site, st)
		return ctx.String(http.StatusOK, "Site with Id:"+st.SiteId+" added successfully to Enterprise Id: "+enterpriseId)
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/site/{site-id}/device-group
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/device-group/{device-group-id})
func (w *ServerImpl) DeleteEnterprisesEnterpriseSiteDeviceGroup(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceGroupId string) error {
	var flag bool
	var entIndex, siteIndex, dgIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for dIndex, dg := range *st.DeviceGroup {
						if dg.DeviceGroupId == deviceGroupId {
							flag = true
							entIndex = eIndex
							siteIndex = sIndex
							dgIndex = dIndex
							break
						}
					}
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site
		dg1 := *st1[siteIndex].DeviceGroup

		dg1 = append(dg1[:dgIndex], dg1[dgIndex+1:]...)
		st1[siteIndex].DeviceGroup = &dg1
		return ctx.String(http.StatusOK, "Device Group with Id: "+deviceGroupId+" deleted successfully")
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/site/{site-id}/device-group
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/device-group/{device-group-id})
func (w *ServerImpl) GetEnterprisesEnterpriseSiteDeviceGroup(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceGroupId string) error {
	for _, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, dg := range *st.DeviceGroup {
						if dg.DeviceGroupId == deviceGroupId {
							return ctx.JSON(http.StatusOK, dg)
						}
					}
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/site/{site-id}/device-group
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/device-group/{device-group-id})
func (w *ServerImpl) PostEnterprisesEnterpriseSiteDeviceGroup(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceGroupId string) error {
	var flag bool
	var entIndex, siteIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, dg := range *st.DeviceGroup {
						if dg.DeviceGroupId == deviceGroupId {
							return ctx.String(http.StatusConflict, "Device Group with Id: "+dg.DeviceGroupId+" already exists")
						}
					}
					flag = true
					entIndex = eIndex
					siteIndex = sIndex
					break
				}
			}
		}
	}

	var dg EnterprisesEnterpriseSiteDeviceGroup
	err = json.Unmarshal(body, &dg)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site

		*st1[siteIndex].DeviceGroup = append(*st1[siteIndex].DeviceGroup, dg)
		return ctx.String(http.StatusOK, "Device Group with Id: "+deviceGroupId+" added successfully")
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/site/{site-id}/device-group/{device-group-id}/device
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/device-group/{device-group-id}/device/{device-id})
func (w *ServerImpl) DeleteEnterprisesEnterpriseSiteDeviceGroupDevice(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceGroupId string, deviceId string) error {
	var flag bool
	var entIndex, siteIndex, dgIndex, devIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for dIndex, dg := range *st.DeviceGroup {
						if dg.DeviceGroupId == deviceGroupId {
							for deIndex, dev := range *dg.Device {
								if dev.DeviceId == deviceId {
									flag = true
									entIndex = eIndex
									siteIndex = sIndex
									dgIndex = dIndex
									devIndex = deIndex
									break
								}
							}
						}
					}
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site
		dg1 := *st1[siteIndex].DeviceGroup
		dev1 := *dg1[dgIndex].Device

		dev1 = append(dev1[:devIndex], dev1[devIndex+1:]...)
		dg1[dgIndex].Device = &dev1
		return ctx.String(http.StatusOK, "Device "+deviceId+" delete from Device Group with Id: "+deviceGroupId)
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/site/{site-id}/device-group/{device-group-id}/device
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/device-group/{device-group-id}/device/{device-id})
func (w *ServerImpl) GetEnterprisesEnterpriseSiteDeviceGroupDevice(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceGroupId string, deviceId string) error {
	for _, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, dg := range *st.DeviceGroup {
						if dg.DeviceGroupId == deviceGroupId {
							for _, dev := range *dg.Device {
								if dev.DeviceId == deviceId {
									return ctx.JSON(http.StatusOK, dev)
								}
							}
						}
					}
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/site/{site-id}/device-group/{device-group-id}/device
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/device-group/{device-group-id}/device/{device-id})
func (w *ServerImpl) PostEnterprisesEnterpriseSiteDeviceGroupDevice(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceGroupId string, deviceId string) error {
	var flag bool
	var entIndex, siteIndex, dgIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for dIndex, dg := range *st.DeviceGroup {
						if dg.DeviceGroupId == deviceGroupId {
							for _, dev := range *dg.Device {
								if dev.DeviceId == deviceId {
									return ctx.String(http.StatusConflict, "Device with Id: "+dev.DeviceId+" already exists")
								}
							}
							flag = true
							entIndex = eIndex
							siteIndex = sIndex
							dgIndex = dIndex
							break
						}
					}

				}
			}
		}
	}

	var dev EnterprisesEnterpriseSiteDeviceGroupDevice
	err = json.Unmarshal(body, &dev)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site
		dg1 := *st1[siteIndex].DeviceGroup

		*dg1[dgIndex].Device = append(*dg1[dgIndex].Device, dev)
		return ctx.String(http.StatusOK, "Device with Id: "+deviceId+" added successfully")
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/site/{site-id}/device-group/{device-group-id}/mbr
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/device-group/{device-group-id}/mbr)
func (w *ServerImpl) DeleteEnterprisesEnterpriseSiteDeviceGroupMbr(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceGroupId string) error {
	var flag bool
	var entIndex, siteIndex, dgIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for dIndex, dg := range *st.DeviceGroup {
						if dg.DeviceGroupId == deviceGroupId {
							flag = true
							entIndex = eIndex
							siteIndex = sIndex
							dgIndex = dIndex
							break
						}
					}
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site
		dg1 := *st1[siteIndex].DeviceGroup

		if dg1[dgIndex].Mbr == nil {
			return ctx.String(http.StatusBadRequest, "Mbr is not present")
		}

		dg1[dgIndex].Mbr = nil
		return ctx.String(http.StatusOK, "Mbr delete from Device Group with Id: "+deviceGroupId)
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/site/{site-id}/device-group/{device-group-id}/mbr
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/device-group/{device-group-id}/mbr)
func (w *ServerImpl) GetEnterprisesEnterpriseSiteDeviceGroupMbr(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceGroupId string) error {
	for _, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, dg := range *st.DeviceGroup {
						if dg.DeviceGroupId == deviceGroupId {
							if dg.Mbr != nil {
								return ctx.JSON(http.StatusOK, dg.Mbr)
							}
						}
					}
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/site/{site-id}/device-group/{device-group-id}/mbr
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/device-group/{device-group-id}/mbr)
func (w *ServerImpl) PostEnterprisesEnterpriseSiteDeviceGroupMbr(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceGroupId string) error {
	var flag bool
	var entIndex, siteIndex, dgIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for dIndex, dg := range *st.DeviceGroup {
						if dg.DeviceGroupId == deviceGroupId {
							flag = true
							entIndex = eIndex
							siteIndex = sIndex
							dgIndex = dIndex
							break
						}
					}
				}
			}
		}
	}

	var dgMbr EnterprisesEnterpriseSiteDeviceGroupMbr
	err = json.Unmarshal(body, &dgMbr)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site
		dg1 := *st1[siteIndex].DeviceGroup

		//TO DO: confirm which update method is fine
		/*ep1[eptIndex].Mbr.Uplink = eptMbr.Uplink
		ep1[eptIndex].Mbr.Downlink = eptMbr.Downlink*/

		dg1[dgIndex].Mbr = &dgMbr
		return ctx.String(http.StatusOK, "Mbr added successfully to Device group with Id: "+deviceGroupId)
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/site/{site-id}/device
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/device/{device-id})
func (w *ServerImpl) DeleteEnterprisesEnterpriseSiteDevice(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceId string) error {
	var flag bool
	var entIndex, siteIndex, devIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for dIndex, dev := range *st.Device {
						if dev.DeviceId == deviceId {
							flag = true
							entIndex = eIndex
							siteIndex = sIndex
							devIndex = dIndex
							break
						}
					}
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site
		dev1 := *st1[siteIndex].Device

		dev1 = append(dev1[:devIndex], dev1[devIndex+1:]...)
		st1[siteIndex].Device = &dev1
		return ctx.String(http.StatusOK, "Device from Site with Id: "+siteId+" deleted successfully")
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/site/{site-id}/device
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/device/{device-id})
func (w *ServerImpl) GetEnterprisesEnterpriseSiteDevice(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceId string) error {
	for _, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, dev := range *st.Device {
						if dev.DeviceId == deviceId {
							return ctx.JSON(http.StatusOK, dev)
						}
					}
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/site/{site-id}/device
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/device/{device-id})
func (w *ServerImpl) PostEnterprisesEnterpriseSiteDevice(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceId string) error {
	var flag bool
	var entIndex, siteIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, dev := range *st.Device {
						if dev.DeviceId == deviceId {
							return ctx.String(http.StatusConflict, "Device with Id: "+deviceId+" already exists")
						}
					}
					flag = true
					entIndex = eIndex
					siteIndex = sIndex
					break
				}
			}
		}
	}

	var dev EnterprisesEnterpriseSiteDevice
	err = json.Unmarshal(body, &dev)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site

		*st1[siteIndex].Device = append(*st1[siteIndex].Device, dev)
		return ctx.String(http.StatusOK, "Device with Id: "+deviceId+" added successfully")
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/site/{site-id}/imsi-definition
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/imsi-definition)
func (w *ServerImpl) DeleteEnterprisesEnterpriseSiteImsiDefinition(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
	var flag bool
	var entIndex, siteIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					flag = true
					entIndex = eIndex
					siteIndex = sIndex
					break
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site

		if st1[siteIndex].ImsiDefinition == nil {
			return ctx.String(http.StatusBadRequest, "Imsi-definition in Site with Id: "+siteId+" is not present")
		}

		st1[siteIndex].ImsiDefinition = nil
		return ctx.String(http.StatusOK, "Imsi-definition from Site with Id: "+siteId+" deleted successfully")
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/site/{site-id}/imsi-definition
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/imsi-definition)
func (w *ServerImpl) GetEnterprisesEnterpriseSiteImsiDefinition(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
	for _, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, st := range *ent.Site {
				if st.SiteId == siteId {
					return ctx.JSON(http.StatusOK, st.ImsiDefinition)
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/site/{site-id}/imsi-definition
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/imsi-definition)
func (w *ServerImpl) PostEnterprisesEnterpriseSiteImsiDefinition(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
	var flag bool
	var entIndex, siteIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					flag = true
					entIndex = eIndex
					siteIndex = sIndex
					break
				}
			}
		}
	}

	var imsiDef EnterprisesEnterpriseSiteImsiDefinition
	err = json.Unmarshal(body, &imsiDef)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site

		st1[siteIndex].ImsiDefinition = &imsiDef
		return ctx.String(http.StatusOK, "Imsi-Definition added successfully to Site with Id: "+siteId)
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/site/{site-id}/ip-domain
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/ip-domain/{ip-domain-id})
func (w *ServerImpl) DeleteEnterprisesEnterpriseSiteIpDomain(ctx echo.Context, target Target, enterpriseId string, siteId string, ipDomainId string) error {
	var flag bool
	var entIndex, siteIndex, ipdIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for iIndex, ipd := range *st.IpDomain {
						if ipd.IpDomainId == ipDomainId {
							flag = true
							entIndex = eIndex
							siteIndex = sIndex
							ipdIndex = iIndex
							break
						}
					}
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site
		ipd1 := *st1[siteIndex].IpDomain

		ipd1 = append(ipd1[:ipdIndex], ipd1[ipdIndex+1:]...)
		st1[siteIndex].IpDomain = &ipd1
		return ctx.String(http.StatusOK, "Ip-Domain with Id: "+ipDomainId+" deleted successfully")
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/site/{site-id}/ip-domain
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/ip-domain/{ip-domain-id})
func (w *ServerImpl) GetEnterprisesEnterpriseSiteIpDomain(ctx echo.Context, target Target, enterpriseId string, siteId string, ipDomainId string) error {
	for _, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, ipd := range *st.IpDomain {
						if ipd.IpDomainId == ipDomainId {
							return ctx.JSON(http.StatusOK, ipd)
						}
					}
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/site/{site-id}/ip-domain
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/ip-domain/{ip-domain-id})
func (w *ServerImpl) PostEnterprisesEnterpriseSiteIpDomain(ctx echo.Context, target Target, enterpriseId string, siteId string, ipDomainId string) error {
	var flag bool
	var entIndex, siteIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, ipd := range *st.IpDomain {
						if ipd.IpDomainId == ipDomainId {
							return ctx.String(http.StatusConflict, "Ip-Domain with Id: "+ipDomainId+" already exists")
						}
					}
					flag = true
					entIndex = eIndex
					siteIndex = sIndex
					break
				}
			}
		}
	}

	var ipd EnterprisesEnterpriseSiteIpDomain
	err = json.Unmarshal(body, &ipd)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site

		*st1[siteIndex].IpDomain = append(*st1[siteIndex].IpDomain, ipd)
		return ctx.String(http.StatusOK, "Ip Domain with Id: "+ipDomainId+" added successfully")
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/site/{site-id}/monitoring
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/monitoring)
func (w *ServerImpl) DeleteEnterprisesEnterpriseSiteMonitoring(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
	var flag bool
	var entIndex, siteIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					flag = true
					entIndex = eIndex
					siteIndex = sIndex
					break
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site

		if st1[siteIndex].Monitoring == nil {
			return ctx.String(http.StatusBadRequest, "Monitoring in Site with Id: "+siteId+" is not present")
		}

		st1[siteIndex].Monitoring = nil
		return ctx.String(http.StatusOK, "Monitoring from Site with Id: "+siteId+" deleted successfully")
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/site/{site-id}/monitoring
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/monitoring)
func (w *ServerImpl) GetEnterprisesEnterpriseSiteMonitoring(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
	for _, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, st := range *ent.Site {
				if st.SiteId == siteId {
					return ctx.JSON(http.StatusOK, st.Monitoring)
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/site/{site-id}/monitoring
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/monitoring)
func (w *ServerImpl) PostEnterprisesEnterpriseSiteMonitoring(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
	var flag bool
	var entIndex, siteIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					flag = true
					entIndex = eIndex
					siteIndex = sIndex
					break
				}
			}
		}
	}

	var mo EnterprisesEnterpriseSiteMonitoring
	err = json.Unmarshal(body, &mo)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site

		st1[siteIndex].Monitoring = &mo
		return ctx.String(http.StatusOK, "Monitoring added successfully to Site with Id: "+siteId)
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/site/{site-id}/monitoring/edge-device
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/monitoring/edge-device/{edge-device-id})
func (w *ServerImpl) DeleteEnterprisesEnterpriseSiteMonitoringEdgeDevice(ctx echo.Context, target Target, enterpriseId string, siteId string, edgeDeviceId string) error {
	var flag bool
	var entIndex, siteIndex, edgIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for edIndex, ed := range *st.Monitoring.EdgeDevice {
						if ed.EdgeDeviceId == edgeDeviceId {
							flag = true
							entIndex = eIndex
							siteIndex = sIndex
							edgIndex = edIndex
							break
						}
					}
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site
		edg1 := *st1[siteIndex].Monitoring.EdgeDevice

		edg1 = append(edg1[:edgIndex], edg1[edgIndex+1:]...)
		st1[siteIndex].Monitoring.EdgeDevice = &edg1
		return ctx.String(http.StatusOK, "Monitoring from Site with Id: "+siteId+" deleted successfully")
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/site/{site-id}/monitoring/edge-device
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/monitoring/edge-device/{edge-device-id})
func (w *ServerImpl) GetEnterprisesEnterpriseSiteMonitoringEdgeDevice(ctx echo.Context, target Target, enterpriseId string, siteId string, edgeDeviceId string) error {
	for _, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, ed := range *st.Monitoring.EdgeDevice {
						if ed.EdgeDeviceId == edgeDeviceId {
							return ctx.JSON(http.StatusOK, ed)
						}
					}
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/site/{site-id}/monitoring/edge-device
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/monitoring/edge-device/{edge-device-id})
func (w *ServerImpl) PostEnterprisesEnterpriseSiteMonitoringEdgeDevice(ctx echo.Context, target Target, enterpriseId string, siteId string, edgeDeviceId string) error {
	var flag bool
	var entIndex, siteIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, edg := range *st.Monitoring.EdgeDevice {
						if edg.EdgeDeviceId == edgeDeviceId {
							return ctx.String(http.StatusConflict, "EdgeDevice with Id: "+edgeDeviceId+" already exists")
						}
					}
					flag = true
					entIndex = eIndex
					siteIndex = sIndex
					break
				}
			}
		}
	}

	var edg EnterprisesEnterpriseSiteMonitoringEdgeDevice
	err = json.Unmarshal(body, &edg)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site

		*st1[siteIndex].Monitoring.EdgeDevice = append(*st1[siteIndex].Monitoring.EdgeDevice, edg)
		return ctx.String(http.StatusOK, "EdgeDevice added successfully to Site with Id: "+siteId)
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/site/{site-id}/sim-card
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/sim-card/{sim-id})
func (w *ServerImpl) DeleteEnterprisesEnterpriseSiteSimCard(ctx echo.Context, target Target, enterpriseId string, siteId string, simId string) error {
	var flag bool
	var entIndex, siteIndex, simIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for siIndex, sim := range *st.SimCard {
						if sim.SimId == simId {
							flag = true
							entIndex = eIndex
							siteIndex = sIndex
							simIndex = siIndex
							break
						}
					}
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site
		sim1 := *st1[siteIndex].SimCard

		sim1 = append(sim1[:simIndex], sim1[simIndex+1:]...)
		st1[siteIndex].SimCard = &sim1
		return ctx.String(http.StatusOK, "Sim Card with Id: "+simId+" deleted successfully")
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/site/{site-id}/sim-card
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/sim-card/{sim-id})
func (w *ServerImpl) GetEnterprisesEnterpriseSiteSimCard(ctx echo.Context, target Target, enterpriseId string, siteId string, simId string) error {
	for _, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, sim := range *st.SimCard {
						if sim.SimId == simId {
							return ctx.JSON(http.StatusOK, sim)
						}
					}
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/site/{site-id}/sim-card
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/sim-card/{sim-id})
func (w *ServerImpl) PostEnterprisesEnterpriseSiteSimCard(ctx echo.Context, target Target, enterpriseId string, siteId string, simId string) error {
	var flag bool
	var entIndex, siteIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, sim := range *st.SimCard {
						if sim.SimId == simId {
							return ctx.String(http.StatusConflict, "Sim Card with Id: "+simId+" already exists")
						}
					}
					flag = true
					entIndex = eIndex
					siteIndex = sIndex
					break
				}
			}
		}
	}

	var sc EnterprisesEnterpriseSiteSimCard
	err = json.Unmarshal(body, &sc)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site

		*st1[siteIndex].SimCard = append(*st1[siteIndex].SimCard, sc)
		return ctx.String(http.StatusOK, "Sim Card with Id: "+simId+" added successfully")
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/site/{site-id}/slice
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id})
func (w *ServerImpl) DeleteEnterprisesEnterpriseSiteSlice(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string) error {
	var flag bool
	var entIndex, siteIndex, sliceIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for slIndex, slice := range *st.Slice {
						if slice.SliceId == sliceId {
							flag = true
							entIndex = eIndex
							siteIndex = sIndex
							sliceIndex = slIndex
							break
						}
					}
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site
		slice1 := *st1[siteIndex].Slice

		slice1 = append(slice1[:sliceIndex], slice1[sliceIndex+1:]...)
		st1[siteIndex].Slice = &slice1
		return ctx.String(http.StatusOK, "Slice with Id: "+sliceId+" deleted successfully")
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/site/{site-id}/slice
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id})
func (w *ServerImpl) GetEnterprisesEnterpriseSiteSlice(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string) error {
	for _, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, slice := range *st.Slice {
						if slice.SliceId == sliceId {
							return ctx.JSON(http.StatusOK, slice)
						}
					}
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/site/{site-id}/slice
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id})
func (w *ServerImpl) PostEnterprisesEnterpriseSiteSlice(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string) error {
	var flag bool
	var entIndex, siteIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, slice := range *st.Slice {
						if slice.SliceId == sliceId {
							return ctx.String(http.StatusConflict, "Slice with Id: "+sliceId+" already exists")
						}
					}
					flag = true
					entIndex = eIndex
					siteIndex = sIndex
					break
				}
			}
		}
	}

	var slice EnterprisesEnterpriseSiteSlice
	err = json.Unmarshal(body, &slice)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site

		*st1[siteIndex].Slice = append(*st1[siteIndex].Slice, slice)
		return ctx.String(http.StatusOK, "Sim Card with Id: "+sliceId+" added successfully")
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/device-group
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/device-group/{device-group})
func (w *ServerImpl) DeleteEnterprisesEnterpriseSiteSliceDeviceGroup(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string, deviceGroup string) error {
	var flag bool
	var entIndex, siteIndex, sliceIndex, dgIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for slIndex, slice := range *st.Slice {
						if slice.SliceId == sliceId {
							for dIndex, dg := range *slice.DeviceGroup {
								if dg.DeviceGroup == deviceGroup {
									flag = true
									entIndex = eIndex
									siteIndex = sIndex
									sliceIndex = slIndex
									dgIndex = dIndex
									break
								}
							}
						}
					}
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site
		slice1 := *st1[siteIndex].Slice
		dg1 := *slice1[sliceIndex].DeviceGroup

		dg1 = append(dg1[:dgIndex], dg1[dgIndex+1:]...)
		slice1[sliceIndex].DeviceGroup = &dg1
		return ctx.String(http.StatusOK, "Device Group "+deviceGroup+" delete from Slice with Id: "+sliceId)
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/device-group
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/device-group/{device-group})
func (w *ServerImpl) GetEnterprisesEnterpriseSiteSliceDeviceGroup(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string, deviceGroup string) error {
	for _, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, slice := range *st.Slice {
						if slice.SliceId == sliceId {
							for _, dg := range *slice.DeviceGroup {
								if dg.DeviceGroup == deviceGroup {
									return ctx.JSON(http.StatusOK, dg)
								}
							}
						}
					}
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/device-group
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/device-group/{device-group})
func (w *ServerImpl) PostEnterprisesEnterpriseSiteSliceDeviceGroup(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string, deviceGroup string) error {
	var flag bool
	var entIndex, siteIndex, sliceIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for slIndex, slice := range *st.Slice {
						if slice.SliceId == sliceId {
							for _, dg := range *slice.DeviceGroup {
								if dg.DeviceGroup == deviceGroup {
									return ctx.String(http.StatusConflict, "Device Group with Id: "+deviceGroup+" already exists")
								}
							}
							flag = true
							entIndex = eIndex
							siteIndex = sIndex
							sliceIndex = slIndex
							break
						}
					}

				}
			}
		}
	}

	var dg EnterprisesEnterpriseSiteSliceDeviceGroup
	err = json.Unmarshal(body, &dg)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site
		slice1 := *st1[siteIndex].Slice

		*slice1[sliceIndex].DeviceGroup = append(*slice1[sliceIndex].DeviceGroup, dg)
		return ctx.String(http.StatusOK, "Device Group with Id: "+deviceGroup+" added successfully")
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/filter
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/filter/{application})
func (w *ServerImpl) DeleteEnterprisesEnterpriseSiteSliceFilter(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string, application string) error {
	var flag bool
	var entIndex, siteIndex, sliceIndex, fiIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for slIndex, slice := range *st.Slice {
						if slice.SliceId == sliceId {
							for fIndex, fi := range *slice.Filter {
								if fi.Application == application {
									flag = true
									entIndex = eIndex
									siteIndex = sIndex
									sliceIndex = slIndex
									fiIndex = fIndex
									break
								}
							}
						}
					}
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site
		slice1 := *st1[siteIndex].Slice
		fi1 := *slice1[sliceIndex].Filter

		fi1 = append(fi1[:fiIndex], fi1[fiIndex+1:]...)
		slice1[sliceIndex].Filter = &fi1
		return ctx.String(http.StatusOK, "Filter "+application+" delete from Slice with Id: "+sliceId)
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/filter
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/filter/{application})
func (w *ServerImpl) GetEnterprisesEnterpriseSiteSliceFilter(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string, application string) error {
	for _, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, slice := range *st.Slice {
						if slice.SliceId == sliceId {
							for _, fi := range *slice.Filter {
								if fi.Application == application {
									return ctx.JSON(http.StatusOK, fi)
								}
							}
						}
					}
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/filter
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/filter/{application})
func (w *ServerImpl) PostEnterprisesEnterpriseSiteSliceFilter(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string, application string) error {
	var flag bool
	var entIndex, siteIndex, sliceIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for slIndex, slice := range *st.Slice {
						if slice.SliceId == sliceId {
							for _, fi := range *slice.Filter {
								if fi.Application == application {
									return ctx.String(http.StatusConflict, "Filter with Id: "+application+" already exists")
								}
							}
							flag = true
							entIndex = eIndex
							siteIndex = sIndex
							sliceIndex = slIndex
							break
						}
					}

				}
			}
		}
	}

	var fi EnterprisesEnterpriseSiteSliceFilter
	err = json.Unmarshal(body, &fi)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site
		slice1 := *st1[siteIndex].Slice

		*slice1[sliceIndex].Filter = append(*slice1[sliceIndex].Filter, fi)
		return ctx.String(http.StatusOK, "Filter with Id: "+application+" added successfully")
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/mbr
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/mbr)
func (w *ServerImpl) DeleteEnterprisesEnterpriseSiteSliceMbr(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string) error {
	var flag bool
	var entIndex, siteIndex, sliceIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for slIndex, slice := range *st.Slice {
						if slice.SliceId == sliceId {
							flag = true
							entIndex = eIndex
							siteIndex = sIndex
							sliceIndex = slIndex
							break
						}
					}
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site
		slice1 := *st1[siteIndex].Slice

		if slice1[sliceIndex].Mbr == nil {
			return ctx.String(http.StatusBadRequest, "Mbr in Slice with Id: "+sliceId+" is not present")
		}

		slice1[sliceIndex].Mbr = nil
		return ctx.String(http.StatusOK, "Mbr delete from Slice with Id: "+sliceId)
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/mbr
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/mbr)
func (w *ServerImpl) GetEnterprisesEnterpriseSiteSliceMbr(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string) error {
	for _, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, slice := range *st.Slice {
						if slice.SliceId == sliceId {
							if slice.Mbr != nil {
								return ctx.JSON(http.StatusOK, slice.Mbr)
							}
						}
					}
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/mbr
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/mbr)
func (w *ServerImpl) PostEnterprisesEnterpriseSiteSliceMbr(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string) error {
	var flag bool
	var entIndex, siteIndex, sliceIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for slIndex, slice := range *st.Slice {
						if slice.SliceId == sliceId {
							flag = true
							entIndex = eIndex
							siteIndex = sIndex
							sliceIndex = slIndex
							break
						}
					}

				}
			}
		}
	}

	var slMbr EnterprisesEnterpriseSiteSliceMbr
	err = json.Unmarshal(body, &slMbr)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site
		slice1 := *st1[siteIndex].Slice

		slice1[sliceIndex].Mbr = &slMbr
		return ctx.String(http.StatusOK, "Mbr added successfully to Slice with Id "+sliceId)
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/priority-traffic-rule
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/priority-traffic-rule/{priority-traffic-rule-id})
func (w *ServerImpl) DeleteEnterprisesEnterpriseSiteSlicePriorityTrafficRule(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string, priorityTrafficRuleId string) error {
	var flag bool
	var entIndex, siteIndex, sliceIndex, ptrIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for slIndex, slice := range *st.Slice {
						if slice.SliceId == sliceId {
							for pIndex, ptr := range *slice.PriorityTrafficRule {
								if ptr.PriorityTrafficRuleId == priorityTrafficRuleId {
									flag = true
									entIndex = eIndex
									siteIndex = sIndex
									sliceIndex = slIndex
									ptrIndex = pIndex
									break
								}
							}
						}
					}
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site
		slice1 := *st1[siteIndex].Slice
		ptr1 := *slice1[sliceIndex].PriorityTrafficRule

		ptr1 = append(ptr1[:ptrIndex], ptr1[ptrIndex+1:]...)
		slice1[sliceIndex].PriorityTrafficRule = &ptr1
		return ctx.String(http.StatusOK, "Priority Traffice Rule "+priorityTrafficRuleId+" delete from Slice with Id: "+sliceId)
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/priority-traffic-rule
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/priority-traffic-rule/{priority-traffic-rule-id})
func (w *ServerImpl) GetEnterprisesEnterpriseSiteSlicePriorityTrafficRule(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string, priorityTrafficRuleId string) error {
	for _, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, slice := range *st.Slice {
						if slice.SliceId == sliceId {
							for _, ptr := range *slice.PriorityTrafficRule {
								if ptr.PriorityTrafficRuleId == priorityTrafficRuleId {
									return ctx.JSON(http.StatusOK, ptr)
								}
							}
						}
					}
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/priority-traffic-rule
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/slice/{slice-id}/priority-traffic-rule/{priority-traffic-rule-id})
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/small-cell/{small-cell-id})
func (w *ServerImpl) PostEnterprisesEnterpriseSiteSlicePriorityTrafficRule(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string, priorityTrafficRuleId string) error {
	var flag bool
	var entIndex, siteIndex, sliceIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for slIndex, slice := range *st.Slice {
						if slice.SliceId == sliceId {
							for _, ptr := range *slice.PriorityTrafficRule {
								if ptr.PriorityTrafficRuleId == priorityTrafficRuleId {
									return ctx.String(http.StatusConflict, "Priority Traffic Rule with Id: "+priorityTrafficRuleId+" already exists")
								}
							}
							flag = true
							entIndex = eIndex
							siteIndex = sIndex
							sliceIndex = slIndex
							break
						}
					}

				}
			}
		}
	}

	var ptr EnterprisesEnterpriseSiteSlicePriorityTrafficRule
	err = json.Unmarshal(body, &ptr)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site
		slice1 := *st1[siteIndex].Slice

		*slice1[sliceIndex].PriorityTrafficRule = append(*slice1[sliceIndex].PriorityTrafficRule, ptr)
		return ctx.String(http.StatusOK, "Priority Traffic Rule with Id: "+priorityTrafficRuleId+" added successfully")
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/site/{site-id}/small-cell
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/small-cell/{small-cell-id})
func (w *ServerImpl) DeleteEnterprisesEnterpriseSiteSmallCell(ctx echo.Context, target Target, enterpriseId string, siteId string, smallCellId string) error {
	var flag bool
	var entIndex, siteIndex, scellIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for scIndex, sc := range *st.SmallCell {
						if sc.SmallCellId == smallCellId {
							flag = true
							entIndex = eIndex
							siteIndex = sIndex
							scellIndex = scIndex
							break
						}
					}
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site
		scell1 := *st1[siteIndex].SmallCell

		scell1 = append(scell1[:scellIndex], scell1[scellIndex+1:]...)
		st1[siteIndex].SmallCell = &scell1
		return ctx.String(http.StatusOK, "Small Cell  with Id: "+smallCellId+" deleted successfully")
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/site/{site-id}/small-cell
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/small-cell/{small-cell-id})
func (w *ServerImpl) GetEnterprisesEnterpriseSiteSmallCell(ctx echo.Context, target Target, enterpriseId string, siteId string, smallCellId string) error {
	for _, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, sc := range *st.SmallCell {
						if sc.SmallCellId == smallCellId {
							return ctx.JSON(http.StatusOK, sc)
						}
					}
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/site/{site-id}/small-cell
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/small-cell/{small-cell-id})
func (w *ServerImpl) PostEnterprisesEnterpriseSiteSmallCell(ctx echo.Context, target Target, enterpriseId string, siteId string, smallCellId string) error {
	var flag bool
	var entIndex, siteIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, sc := range *st.SmallCell {
						if sc.SmallCellId == smallCellId {
							return ctx.String(http.StatusConflict, "Small Cell with Id: "+smallCellId+" already exists")
						}
					}
					flag = true
					entIndex = eIndex
					siteIndex = sIndex
					break
				}
			}
		}
	}

	var sc EnterprisesEnterpriseSiteSmallCell
	err = json.Unmarshal(body, &sc)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site

		*st1[siteIndex].SmallCell = append(*st1[siteIndex].SmallCell, sc)
		return ctx.String(http.StatusOK, "Small Cell with Id: "+smallCellId+" added successfully")
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/site/{site-id}/upf
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/upf/{upf-id})
func (w *ServerImpl) DeleteEnterprisesEnterpriseSiteUpf(ctx echo.Context, target Target, enterpriseId string, siteId string, upfId string) error {
	var flag bool
	var entIndex, siteIndex, upfIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for uIndex, upf := range *st.Upf {
						if upf.UpfId == upfId {
							flag = true
							entIndex = eIndex
							siteIndex = sIndex
							upfIndex = uIndex
							break
						}
					}
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site
		upf1 := *st1[siteIndex].Upf

		upf1 = append(upf1[:upfIndex], upf1[upfIndex+1:]...)
		st1[siteIndex].Upf = &upf1
		return ctx.String(http.StatusOK, "Upf with Id: "+upfId+" deleted successfully")
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/site/{site-id}/upf
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/upf/{upf-id})
func (w *ServerImpl) GetEnterprisesEnterpriseSiteUpf(ctx echo.Context, target Target, enterpriseId string, siteId string, upfId string) error {
	for _, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, upf := range *st.Upf {
						if upf.UpfId == upfId {
							return ctx.JSON(http.StatusOK, upf)
						}
					}
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/site/{site-id}/upf
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}/upf/{upf-id})
func (w *ServerImpl) PostEnterprisesEnterpriseSiteUpf(ctx echo.Context, target Target, enterpriseId string, siteId string, upfId string) error {
	var flag bool
	var entIndex, siteIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for sIndex, st := range *ent.Site {
				if st.SiteId == siteId {
					for _, upf := range *st.Upf {
						if upf.UpfId == upfId {
							return ctx.String(http.StatusConflict, "Upf with Id: "+upfId+" already exists")
						}
					}
					flag = true
					entIndex = eIndex
					siteIndex = sIndex
					break
				}
			}
		}
	}

	var upf EnterprisesEnterpriseSiteUpf
	err = json.Unmarshal(body, &upf)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise
		st1 := *e1[entIndex].Site

		*st1[siteIndex].Upf = append(*st1[siteIndex].Upf, upf)
		return ctx.String(http.StatusOK, "Upf with Id: "+upfId+" added successfully")
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/template
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/template/{template-id})
func (w *ServerImpl) DeleteEnterprisesEnterpriseTemplate(ctx echo.Context, target Target, enterpriseId string, templateId string) error {
	var flag bool
	var entIndex, temIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for tIndex, tem := range *ent.Template {
				if tem.TemplateId == templateId {
					flag = true
					entIndex = eIndex
					temIndex = tIndex
					break
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		tem1 := *e1[entIndex].Template

		tem1 = append(tem1[:temIndex], tem1[temIndex+1:]...)
		e1[entIndex].Template = &tem1
		return ctx.String(http.StatusOK, "Template with Id:"+templateId+" deleted from Enteprise Id: "+enterpriseId)
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/template
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/template/{template-id})
func (w *ServerImpl) GetEnterprisesEnterpriseTemplate(ctx echo.Context, target Target, enterpriseId string, templateId string) error {
	for _, e := range *w.config.Enterprise {
		if e.EnterpriseId == enterpriseId {
			for _, tem := range *e.Template {
				if tem.TemplateId == templateId {
					return ctx.JSON(http.StatusOK, tem)
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/template
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/template/{template-id})
func (w *ServerImpl) PostEnterprisesEnterpriseTemplate(ctx echo.Context, target Target, enterpriseId string, templateId string) error {
	var flag bool
	var entIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, tem := range *ent.Template {
				if tem.TemplateId == templateId {
					return ctx.String(http.StatusConflict, "Template with Id: "+templateId+"already exists with enterprise id "+
						enterpriseId)
				}
			}
			flag = true
			entIndex = eIndex
			break
		}
	}

	var tem EnterprisesEnterpriseTemplate
	err = json.Unmarshal(body, &tem)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise

		*e1[entIndex].Template = append(*e1[entIndex].Template, tem)
		return ctx.String(http.StatusOK, "Template with Id:"+templateId+" added successfully to Enterprise Id: "+
			enterpriseId)
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/template/{template-id}/mbr
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/template/{template-id}/mbr)
func (w *ServerImpl) DeleteEnterprisesEnterpriseTemplateMbr(ctx echo.Context, target Target, enterpriseId string, templateId string) error {
	var flag bool
	var entIndex, temIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for tIndex, tem := range *ent.Template {
				if tem.TemplateId == templateId {
					flag = true
					entIndex = eIndex
					temIndex = tIndex
					break
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		tem1 := *e1[entIndex].Template

		if tem1[temIndex].Mbr == nil {
			return ctx.String(http.StatusBadRequest, "Mbr is not present")
		}
		tem1[temIndex].Mbr = nil
		return ctx.String(http.StatusOK, "Template with Id:"+templateId+" deleted from Enteprise Id: "+enterpriseId)
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/template/{template-id}/mbr
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/template/{template-id}/mbr)
func (w *ServerImpl) GetEnterprisesEnterpriseTemplateMbr(ctx echo.Context, target Target, enterpriseId string, templateId string) error {
	for _, e := range *w.config.Enterprise {
		if e.EnterpriseId == enterpriseId {
			for _, tem := range *e.Template {
				if tem.TemplateId == templateId {
					if tem.Mbr != nil {
						return ctx.JSON(http.StatusOK, tem.Mbr)
					}
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/template/{template-id}/mbr
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/template/{template-id}/mbr)
func (w *ServerImpl) PostEnterprisesEnterpriseTemplateMbr(ctx echo.Context, target Target, enterpriseId string, templateId string) error {
	var flag bool
	var entIndex, temIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for tIndex, tem := range *ent.Template {
				if tem.TemplateId == templateId {
					flag = true
					entIndex = eIndex
					temIndex = tIndex
					break
				}
			}

		}
	}

	var temMbr EnterprisesEnterpriseTemplateMbr
	err = json.Unmarshal(body, &temMbr)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise
		tem1 := *e1[entIndex].Template

		tem1[temIndex].Mbr = &temMbr
		return ctx.String(http.StatusOK, "Template with Id:"+templateId+" added successfully to Enterprise Id: "+
			enterpriseId)
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /enterprises/enterprise/{enterprise-id}/traffic-class
// (DELETE /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/traffic-class/{traffic-class-id})
func (w *ServerImpl) DeleteEnterprisesEnterpriseTrafficClass(ctx echo.Context, target Target, enterpriseId string, trafficClassId string) error {
	var flag bool
	var entIndex, tcIndex int

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for tIndex, tc := range *ent.TrafficClass {
				if tc.TrafficClassId == trafficClassId {
					flag = true
					entIndex = eIndex
					tcIndex = tIndex
					break
				}
			}
		}
	}

	if flag {
		e1 := *w.config.Enterprise
		tc1 := *e1[entIndex].TrafficClass

		tc1 = append(tc1[:tcIndex], tc1[tcIndex+1:]...)
		e1[entIndex].TrafficClass = &tc1
		return ctx.String(http.StatusOK, "TrafficClass with Id:"+trafficClassId+" deleted from Enteprise Id: "+enterpriseId)
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GET /enterprises/enterprise/{enterprise-id}/traffic-class
// (GET /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/traffic-class/{traffic-class-id})
func (w *ServerImpl) GetEnterprisesEnterpriseTrafficClass(ctx echo.Context, target Target, enterpriseId string, trafficClassId string) error {
	for _, e := range *w.config.Enterprise {
		if e.EnterpriseId == enterpriseId {
			for _, tc := range *e.TrafficClass {
				if tc.TrafficClassId == trafficClassId {
					return ctx.JSON(http.StatusOK, tc)
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// POST /enterprises/enterprise/{enterprise-id}/traffic-class
// (POST /aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/traffic-class/{traffic-class-id})
func (w *ServerImpl) PostEnterprisesEnterpriseTrafficClass(ctx echo.Context, target Target, enterpriseId string, trafficClassId string) error {
	var flag bool
	var entIndex int
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for eIndex, ent := range *w.config.Enterprise {
		if ent.EnterpriseId == enterpriseId {
			for _, tc := range *ent.TrafficClass {
				if tc.TrafficClassId == trafficClassId {
					return ctx.String(http.StatusConflict, "Traffic Class with Id: "+trafficClassId+"already exists with enterprise id "+
						enterpriseId)
				}
			}
			flag = true
			entIndex = eIndex
			break
		}
	}

	var tc EnterprisesEnterpriseTrafficClass
	err = json.Unmarshal(body, &tc)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if flag {
		e1 := *w.config.Enterprise

		*e1[entIndex].TrafficClass = append(*e1[entIndex].TrafficClass, tc)
		return ctx.String(http.StatusOK, "Traffic Class with Id:"+trafficClassId+" added successfully to Enterprise Id: "+
			enterpriseId)
	}

	return echo.NewHTTPError(http.StatusNotFound)
}
