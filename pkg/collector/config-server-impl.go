// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package collector

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/onosproject/chronos-exporter/pkg/utils"
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
	for index, app := range s.config.Applications {
		if app.ApplicationId == applicationId {
			s.config.Applications = append(s.config.Applications[:index], s.config.Applications[index+1:]...)
			return ctx.String(http.StatusOK, "Application "+app.DisplayName+" deleted successfully")
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GetApplicationsApplication (GET /config/applications/{application-id})
func (s *ServerImpl) GetApplicationsApplication(ctx echo.Context, applicationId string) error {
	for _, app := range s.config.Applications {
		if app.ApplicationId == applicationId {
			return ctx.JSON(http.StatusOK, app)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// PostApplicationsApplication (POST /config/applications/{application-id})
func (s *ServerImpl) PostApplicationsApplication(ctx echo.Context, applicationId string) error {
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	//TO DO: check if the applicationId in param and application id in payload are same or not
	for _, app := range s.config.Applications {
		if app.ApplicationId == applicationId {
			return ctx.String(http.StatusConflict, "Application "+app.DisplayName+"already exists with application id "+applicationId)
		}
	}

	var app Application
	err = json.Unmarshal(body, &app)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	s.config.Applications = append(s.config.Applications, app)
	return ctx.String(http.StatusOK, "Application "+app.DisplayName+"added successfully")
}

// DeleteSitesSite (DELETE /config/sites/{site-id})
func (s *ServerImpl) DeleteSitesSite(ctx echo.Context, siteId string) error {
	for index, site := range s.config.Sites {
		if site.SiteId == siteId {
			s.config.Sites = append(s.config.Sites[:index], s.config.Sites[index+1:]...)
			return ctx.String(http.StatusOK, "Site "+*site.DisplayName+" deleted successfully")
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GetSitesSite (GET /config/sites/{site-id})
func (s *ServerImpl) GetSitesSite(ctx echo.Context, siteId string) error {
	for _, site := range s.config.Sites {
		if site.SiteId == siteId {
			return ctx.JSON(http.StatusOK, site)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// PostSitesSite (POST /config/sites/{site-id})
func (s *ServerImpl) PostSitesSite(ctx echo.Context, siteId string) error {
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	for _, st := range s.config.Sites {
		if st.SiteId == siteId {
			return ctx.String(http.StatusConflict, "Site "+*st.DisplayName+"already exists with site id "+siteId)
		}
	}

	var site Site
	err = json.Unmarshal(body, &site)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	s.config.Sites = append(s.config.Sites, site)
	return ctx.String(http.StatusOK, "Site "+*site.DisplayName+"added successfully")
}

// DeleteSitesSiteDeviceGroupsDeviceGroup (DELETE /config/sites/{site-id}/device-groups/{device-group-id})
func (s *ServerImpl) DeleteSitesSiteDeviceGroupsDeviceGroup(ctx echo.Context, siteId string, deviceGroupId string) error {
	s1 := s.config.Sites
	var flag bool
	var siteIndex, dgIndex int
	for sIndex, site := range s.config.Sites {
		if site.SiteId == siteId {
			for dIndex, dg := range site.DeviceGroups {
				if dg.DeviceGroupId == deviceGroupId {
					flag = true
					siteIndex = sIndex
					dgIndex = dIndex
					break
					//site.DeviceGroups = append(site.DeviceGroups[:index], site.DeviceGroups[index+1:]...)
					//return ctx.String(http.StatusOK, "Device Group  "+dg.DisplayName+" deleted successfully from Site"+
					//	*site.DisplayName)
				}
			}
		}
		if flag {
			s1[siteIndex].DeviceGroups = append(s1[siteIndex].DeviceGroups[:dgIndex], s1[siteIndex].DeviceGroups[dgIndex+1:]...)
			return ctx.String(http.StatusOK, "Device Group  "+deviceGroupId+" deleted successfully from Site"+
				*s1[siteIndex].DisplayName)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GetSitesSiteDeviceGroupsDeviceGroup (GET /config/sites/{site-id}/device-groups/{device-group-id})
func (s *ServerImpl) GetSitesSiteDeviceGroupsDeviceGroup(ctx echo.Context, siteId string, deviceGroupId string) error {
	for _, site := range s.config.Sites {
		if site.SiteId == siteId {
			for _, dg := range site.DeviceGroups {
				if dg.DeviceGroupId == deviceGroupId {
					return ctx.JSON(http.StatusOK, dg)
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// PostSitesSiteDeviceGroupsDeviceGroup (POST /config/sites/{site-id}/device-groups/{device-group-id})
func (s *ServerImpl) PostSitesSiteDeviceGroupsDeviceGroup(ctx echo.Context, siteId string, deviceGroupId string) error {
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	var flag bool
	var dg DeviceGroup
	err = json.Unmarshal(body, &dg)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	s1 := s.config.Sites
	for sIndex, site := range s.config.Sites {
		if site.SiteId == siteId {
			for _, devG := range site.DeviceGroups {
				if devG.DeviceGroupId == deviceGroupId {
					return ctx.String(http.StatusConflict, "Device Group "+devG.DisplayName+
						" already exists with device group id  "+deviceGroupId+" in site "+*site.DisplayName)
				}
			}
			flag = true
		}
		if flag {
			s1[sIndex].DeviceGroups = append(s1[sIndex].DeviceGroups, dg)
			return ctx.String(http.StatusOK, "Device Group  "+dg.DisplayName+" added successfully to Site"+
				*site.DisplayName)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// DeleteSitesSiteDeviceGroupsDeviceGroupDevicesDevice (DELETE /config/sites/{site-id}/device-groups/{device-group-id}/devices/{dev-id})
func (s *ServerImpl) DeleteSitesSiteDeviceGroupsDeviceGroupDevicesDevice(ctx echo.Context, siteId string, deviceGroupId string, devId string) error {
	s1 := s.config.Sites
	var flag bool
	var siteIndex, dgIndex, deviceIndex int
	for sIndex, site := range s.config.Sites {
		if site.SiteId == siteId {
			for dIndex, dg := range site.DeviceGroups {
				if dg.DeviceGroupId == deviceGroupId {
					for devIndex, dev := range *dg.Devices {
						if dev == devId {
							flag = true
							siteIndex = sIndex
							dgIndex = dIndex
							deviceIndex = devIndex
							break
							//dg1 := *dg.Devices
							//*dg.Devices = append(dg1[:index], dg1[index+1:]...)
							//return ctx.JSON(http.StatusOK, "Deleted device "+dev+ " from Device Group "+ dg.DeviceGroupId)
						}
					}
				}
			}
		}
		if flag {
			dg1 := *s1[siteIndex].DeviceGroups[dgIndex].Devices
			*s1[siteIndex].DeviceGroups[dgIndex].Devices = append(dg1[:deviceIndex], dg1[deviceIndex+1:]...)
			return ctx.JSON(http.StatusOK, "Deleted device "+devId+" from Device Group "+s1[siteIndex].DeviceGroups[dgIndex].DisplayName)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GetSitesSiteDeviceGroupsDeviceGroupDevicesDevice (GET /config/sites/{site-id}/device-groups/{device-group-id}/devices/{dev-id})
func (s *ServerImpl) GetSitesSiteDeviceGroupsDeviceGroupDevicesDevice(ctx echo.Context, siteId string, deviceGroupId string, devId string) error {
	for _, site := range s.config.Sites {
		if site.SiteId == siteId {
			for _, dg := range site.DeviceGroups {
				if dg.DeviceGroupId == deviceGroupId {
					for _, dev := range *dg.Devices {
						if dev == devId {
							return ctx.String(http.StatusOK, "Found device "+dev)
						}
					}
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// PostSitesSiteDeviceGroupsDeviceGroupDevicesDevice (POST /config/sites/{site-id}/device-groups/{device-group-id}/devices/{dev-id})
func (s *ServerImpl) PostSitesSiteDeviceGroupsDeviceGroupDevicesDevice(ctx echo.Context, siteId string, deviceGroupId string, devId string) error {
	//body, err := utils.ReadRequestBody(ctx.Request().Body)
	//if err != nil {
	//	return ctx.String(http.StatusInternalServerError, err.Error())
	//}

	//Check device should be taken from json or from param
	/*
		err = json.Unmarshal(body, &device)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}*/

	s1 := s.config.Sites
	var flag bool
	var siteIndex, dgIndex int
	for sIndex, site := range s.config.Sites {
		if site.SiteId == siteId {
			for dIndex, dg := range site.DeviceGroups {
				if dg.DeviceGroupId == deviceGroupId {
					for _, dev := range *dg.Devices {
						if dev == devId {
							return ctx.JSON(http.StatusConflict, "Device "+dev+" already exist in Device Group "+dg.DeviceGroupId)
						}
						//dg1 := *dg.Devices
						//*dg.Devices = append(dg1, devId)
						//return ctx.JSON(http.StatusOK, "Added device "+dev+ " in Device Group "+ dg.DeviceGroupId)
					}
					flag = true
					siteIndex = sIndex
					dgIndex = dIndex
					break
				}
			}
		}
		if flag {
			dg1 := *s1[siteIndex].DeviceGroups[dgIndex].Devices
			*s1[siteIndex].DeviceGroups[dgIndex].Devices = append(dg1, devId)
			return ctx.JSON(http.StatusOK, "Added device "+devId+" in Device Group "+s1[siteIndex].DeviceGroups[dgIndex].DisplayName)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// DeleteSitesSiteDevicesDevice (DELETE /config/sites/{site-id}/devices/{serial-number})
func (s *ServerImpl) DeleteSitesSiteDevicesDevice(ctx echo.Context, siteId string, serialNumber string) error {
	s1 := s.config.Sites
	var flag bool
	var siteIndex, deviceIndex int
	for sIndex, site := range s.config.Sites {
		if site.SiteId == siteId {
			for dIndex, dev := range site.Devices {
				if dev.SerialNumber == serialNumber {

					flag = true
					siteIndex = sIndex
					deviceIndex = dIndex
					break
					//site.Devices = append(site.Devices[:index], site.Devices[index+1:]...)
					//return ctx.String(http.StatusOK, "Device "+dev.DisplayName+ " deleted successfully to Site "+ *site.DisplayName)
				}
			}
		}
		if flag {
			s1[siteIndex].Devices = append(s1[siteIndex].Devices[:deviceIndex], s1[siteIndex].Devices[deviceIndex+1:]...)
			return ctx.String(http.StatusOK, "Device "+serialNumber+" deleted successfully to Site "+*s1[siteIndex].DisplayName)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GetSitesSiteDevicesDevice (POST /config/sites/{site-id}/devices/{serial-number})
func (s *ServerImpl) GetSitesSiteDevicesDevice(ctx echo.Context, siteId string, serialNumber string) error {
	for _, site := range s.config.Sites {
		if site.SiteId == siteId {
			for _, dev := range site.Devices {
				if dev.SerialNumber == serialNumber {
					return ctx.JSON(http.StatusOK, dev)
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// PostSitesSiteDevicesDevice (POST /config/sites/{site-id}/devices/{serial-number})
func (s *ServerImpl) PostSitesSiteDevicesDevice(ctx echo.Context, siteId string, serialNumber string) error {
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	var device Device
	err = json.Unmarshal(body, &device)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	var flag bool
	var siteIndex int
	s1 := s.config.Sites
	for sIndex, site := range s.config.Sites {
		if site.SiteId == siteId {
			for _, dev := range site.Devices {
				if dev.SerialNumber == serialNumber {
					return ctx.String(http.StatusConflict, "Device "+dev.DisplayName+" already present at Site "+*site.DisplayName)
				}
				/*de := site.Devices
				site.Devices = append(de, device)
				return ctx.String(http.StatusOK, "Device "+device.DisplayName+ " added successfully at Site "+ *site.DisplayName)*/
			}
			flag = true
			siteIndex = sIndex
		}
		if flag {
			s1[siteIndex].Devices = append(s1[siteIndex].Devices, device)
			return ctx.String(http.StatusOK, "Device "+device.DisplayName+" added successfully at Site "+*s1[siteIndex].DisplayName)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// DeleteSitesSiteSimsSim (DELETE /config/sites/{site-id}/sims/{iccid})
func (s *ServerImpl) DeleteSitesSiteSimsSim(ctx echo.Context, siteId string, iccid string) error {
	var flag bool
	var siteIndex, simIndex int
	s1 := s.config.Sites
	for sIndex, site := range s.config.Sites {
		if site.SiteId == siteId {
			for siIndex, sim := range site.Sims {
				if sim.Iccid == iccid {
					flag = true
					siteIndex = sIndex
					simIndex = siIndex
					break
					/*site.Sims = append(site.Sims[:index], site.Sims[index+1:]...)
					return ctx.String(http.StatusOK, "Sim "+ *sim.DisplayName+" deleted successfully from Site"+*site.DisplayName)*/
				}
			}
		}
		if flag {
			s1[siteIndex].Sims = append(s1[siteIndex].Sims[:simIndex], s1[siteIndex].Sims[simIndex+1:]...)
			return ctx.String(http.StatusOK, "Sim "+iccid+" deleted successfully from Site"+*s1[siteIndex].DisplayName)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GetSitesSiteSimsSim (GET /config/sites/{site-id}/sims/{iccid})
func (s *ServerImpl) GetSitesSiteSimsSim(ctx echo.Context, siteId string, iccid string) error {
	for _, site := range s.config.Sites {
		if site.SiteId == siteId {
			for _, sim := range site.Sims {
				if sim.Iccid == iccid {
					return ctx.JSON(http.StatusOK, sim)
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// PostSitesSiteSimsSim (POST /config/sites/{site-id}/sims/{iccid})
func (s *ServerImpl) PostSitesSiteSimsSim(ctx echo.Context, siteId string, iccid string) error {
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	var si Sim
	err = json.Unmarshal(body, &si)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	var flag bool
	var siteIndex int
	s1 := s.config.Sites
	for sIndex, site := range s.config.Sites {
		if site.SiteId == siteId {
			for _, sim := range site.Sims {
				if sim.Iccid == iccid {
					return ctx.String(http.StatusConflict, "Sim "+sim.Iccid+" already exist at site "+site.SiteId)
				}
				//site.Sims = append(site.Sims, si)
				//return ctx.String(http.StatusOK, "Sim "+ *si.DisplayName+ " added to site "+ *site.DisplayName)
			}
			flag = true
			siteIndex = sIndex
		}
		if flag {
			s1[siteIndex].Sims = append(s1[siteIndex].Sims, si)
			return ctx.String(http.StatusOK, "Sim "+*si.DisplayName+" added to site "+*s1[siteIndex].DisplayName)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// DeleteSitesSiteSlicesSlice (DELETE /config/sites/{site-id}/slices/{slice-id})
func (s *ServerImpl) DeleteSitesSiteSlicesSlice(ctx echo.Context, siteId string, sliceId string) error {
	var flag bool
	var siteIndex, sliceIndex int
	s1 := s.config.Sites
	for sIndex, site := range s.config.Sites {
		if site.SiteId == siteId {
			for slIndex, slice := range site.Slices {
				if slice.SliceId == sliceId {
					flag = true
					siteIndex = sIndex
					sliceIndex = slIndex
					break
				}
			}
		}
		if flag {
			s1[siteIndex].Slices = append(s1[siteIndex].Slices[:sliceIndex], s1[siteIndex].Slices[sliceIndex+1:]...)
			return ctx.String(http.StatusOK, "Slice"+sliceId+" deleted from site "+*s1[siteIndex].DisplayName)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GetSitesSiteSlicesSlice (GET /config/sites/{site-id}/slices/{slice-id})
func (s *ServerImpl) GetSitesSiteSlicesSlice(ctx echo.Context, siteId string, sliceId string) error {
	for _, site := range s.config.Sites {
		if site.SiteId == siteId {
			for _, slice := range site.Slices {
				if slice.SliceId == sliceId {
					return ctx.JSON(http.StatusOK, slice)
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// PostSitesSiteSlicesSlice (POST /config/sites/{site-id}/slices/{slice-id})
func (s *ServerImpl) PostSitesSiteSlicesSlice(ctx echo.Context, siteId string, sliceId string) error {
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	var sl Slice
	err = json.Unmarshal(body, &sl)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	var flag bool
	var siteIndex int
	s1 := s.config.Sites
	for sIndex, site := range s.config.Sites {
		if site.SiteId == siteId {
			for _, slice := range site.Slices {
				if slice.SliceId == sliceId {
					return ctx.String(http.StatusConflict, "Slice "+slice.DisplayName+" already exists at site "+*site.DisplayName)
				}
			}
			flag = true
			siteIndex = sIndex
		}
		if flag {
			s1[siteIndex].Slices = append(s1[siteIndex].Slices, sl)
			return ctx.String(http.StatusOK, "Slice"+sliceId+" added at site "+*s1[siteIndex].DisplayName)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// DeleteSitesSiteSlicesSliceApplicationsApplication (DELETE /config/sites/{site-id}/slices/{slice-id}/applications/{app-id})
func (s *ServerImpl) DeleteSitesSiteSlicesSliceApplicationsApplication(ctx echo.Context, siteId string, sliceId string, appId string) error {
	s1 := s.config.Sites
	var flag bool
	var siteIndex, sliceIndex, appIndex int
	for sIndex, site := range s.config.Sites {
		if site.SiteId == siteId {
			for slIndex, sl := range site.Slices {
				if sl.SliceId == sliceId {
					for apIndex, app := range *sl.Applications {
						if app == appId {
							flag = true
							siteIndex = sIndex
							sliceIndex = slIndex
							appIndex = apIndex
							break
						}
					}
				}
			}
		}
		if flag {
			app1 := *s1[siteIndex].Slices[sliceIndex].Applications
			*s1[siteIndex].Slices[sliceIndex].Applications = append(app1[:appIndex], app1[appIndex+1:]...)
			return ctx.JSON(http.StatusOK, "Deleted application "+appId+" from Slice "+s1[siteIndex].Slices[sliceIndex].DisplayName)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GetSitesSiteSlicesSliceApplicationsApplication (GET /config/sites/{site-id}/slices/{slice-id}/applications/{app-id})
func (s *ServerImpl) GetSitesSiteSlicesSliceApplicationsApplication(ctx echo.Context, siteId string, sliceId string, appId string) error {
	for _, site := range s.config.Sites {
		if site.SiteId == siteId {
			for _, sl := range site.Slices {
				if sl.SliceId == sliceId {
					for _, app := range *sl.Applications {
						if app == appId {
							return ctx.String(http.StatusOK, "Found application "+appId)
						}
					}
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// PostSitesSiteSlicesSliceApplicationsApplication (POST /config/sites/{site-id}/slices/{slice-id}/applications/{app-id})
func (s *ServerImpl) PostSitesSiteSlicesSliceApplicationsApplication(ctx echo.Context, siteId string, sliceId string, appId string) error {
	//body, err := utils.ReadRequestBody(ctx.Request().Body)
	//if err != nil {
	//	return ctx.String(http.StatusInternalServerError, err.Error())
	//}

	//Check application should be taken from json or from param
	/*
		err = json.Unmarshal(body, &device)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}*/

	s1 := s.config.Sites
	var flag bool
	var siteIndex, sliceIndex int
	for sIndex, site := range s.config.Sites {
		if site.SiteId == siteId {
			for slIndex, sl := range site.Slices {
				if sl.SliceId == sliceId {
					for _, app := range *sl.Applications {
						if app == appId {
							return ctx.JSON(http.StatusConflict, "Application "+app+" already exist in Slice "+sl.DisplayName)
						}
					}
					flag = true
					siteIndex = sIndex
					sliceIndex = slIndex
					break
				}
			}
		}
		if flag {
			app1 := *s1[siteIndex].Slices[sliceIndex].Applications
			*s1[siteIndex].Slices[sliceIndex].Applications = append(app1, appId)
			return ctx.JSON(http.StatusOK, "Added Application "+appId+" in Slice "+s1[siteIndex].Slices[sliceIndex].DisplayName)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// DeleteSitesSiteSlicesSliceDeviceGroupsDeviceGroup (DELETE /config/sites/{site-id}/slices/{slice-id}/device-groups/{dg-id})
func (s *ServerImpl) DeleteSitesSiteSlicesSliceDeviceGroupsDeviceGroup(ctx echo.Context, siteId string, sliceId string, dgId string) error {
	s1 := s.config.Sites
	var flag bool
	var siteIndex, sliceIndex, dgIndex int
	for sIndex, site := range s.config.Sites {
		if site.SiteId == siteId {
			for slIndex, sl := range site.Slices {
				if sl.SliceId == sliceId {
					for devIndex, dg := range *sl.DeviceGroups {
						if dg == dgId {
							flag = true
							siteIndex = sIndex
							sliceIndex = slIndex
							dgIndex = devIndex
							break
						}
					}
				}
			}
		}
		if flag {
			dg1 := *s1[siteIndex].Slices[sliceIndex].DeviceGroups
			*s1[siteIndex].Slices[sliceIndex].DeviceGroups = append(dg1[:dgIndex], dg1[dgIndex+1:]...)
			return ctx.JSON(http.StatusOK, "Deleted Device Group "+dgId+" from Slice "+s1[siteIndex].Slices[sliceIndex].DisplayName)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GetSitesSiteSlicesSliceDeviceGroupsDeviceGroup (GET /config/sites/{site-id}/slices/{slice-id}/device-groups/{dg-id})
func (s *ServerImpl) GetSitesSiteSlicesSliceDeviceGroupsDeviceGroup(ctx echo.Context, siteId string, sliceId string, dgId string) error {
	for _, site := range s.config.Sites {
		if site.SiteId == siteId {
			for _, sl := range site.Slices {
				if sl.SliceId == sliceId {
					for _, dg := range *sl.DeviceGroups {
						if dg == dgId {
							return ctx.String(http.StatusOK, "Device Group exists "+dgId)
						}
					}
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// PostSitesSiteSlicesSliceDeviceGroupsDeviceGroup (POST /config/sites/{site-id}/slices/{slice-id}/device-groups/{dg-id})
func (s *ServerImpl) PostSitesSiteSlicesSliceDeviceGroupsDeviceGroup(ctx echo.Context, siteId string, sliceId string, dgId string) error {
	//body, err := utils.ReadRequestBody(ctx.Request().Body)
	//if err != nil {
	//	return ctx.String(http.StatusInternalServerError, err.Error())
	//}

	//Check Device Group should be taken from json or from param
	/*
		err = json.Unmarshal(body, &device)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}*/

	s1 := s.config.Sites
	var flag bool
	var siteIndex, sliceIndex int
	for sIndex, site := range s.config.Sites {
		if site.SiteId == siteId {
			for slIndex, sl := range site.Slices {
				if sl.SliceId == sliceId {
					for _, dg := range *sl.DeviceGroups {
						if dg == dgId {
							return ctx.JSON(http.StatusConflict, "Device Group "+dg+" already exist in Slice "+sl.DisplayName)
						}
					}
					flag = true
					siteIndex = sIndex
					sliceIndex = slIndex
					break
				}
			}
		}
		if flag {
			app1 := *s1[siteIndex].Slices[sliceIndex].DeviceGroups
			*s1[siteIndex].Slices[sliceIndex].DeviceGroups = append(app1, dgId)
			return ctx.JSON(http.StatusOK, "Added Device Group "+dgId+" in Slice "+s1[siteIndex].Slices[sliceIndex].DisplayName)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// DeleteSitesSiteSmallCellsSmallCell (DELETE /config/sites/{site-id}/small-cells/{small-cell-id})
func (s *ServerImpl) DeleteSitesSiteSmallCellsSmallCell(ctx echo.Context, siteId string, smallCellId string) error {
	var flag bool
	var siteIndex, scIndex int
	s1 := s.config.Sites
	for sIndex, site := range s.config.Sites {
		if site.SiteId == siteId {
			for smIndex, sc := range site.SmallCells {
				if sc.SmallCellId == smallCellId {
					flag = true
					siteIndex = sIndex
					scIndex = smIndex
					break
				}
			}
		}
		if flag {
			s1[siteIndex].SmallCells = append(s1[siteIndex].SmallCells[:scIndex], s1[scIndex].SmallCells[scIndex+1:]...)
			return ctx.String(http.StatusOK, "SmallCell "+smallCellId+" deleted from site "+*s1[siteIndex].DisplayName)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// GetSitesSiteSmallCellsSmallCell (GET /config/sites/{site-id}/small-cells/{small-cell-id})
func (s *ServerImpl) GetSitesSiteSmallCellsSmallCell(ctx echo.Context, siteId string, smallCellId string) error {
	for _, site := range s.config.Sites {
		if site.SiteId == siteId {
			for _, sc := range site.SmallCells {
				if sc.SmallCellId == smallCellId {
					return ctx.JSON(http.StatusOK, sc)
				}
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// PostSitesSiteSmallCellsSmallCell (POST /config/sites/{site-id}/small-cells/{small-cell-id})
func (s *ServerImpl) PostSitesSiteSmallCellsSmallCell(ctx echo.Context, siteId string, smallCellId string) error {
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	var smc SmallCell
	err = json.Unmarshal(body, &smc)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	var flag bool
	var siteIndex int
	s1 := s.config.Sites
	for sIndex, site := range s.config.Sites {
		if site.SiteId == siteId {
			for _, sm := range site.SmallCells {
				if sm.SmallCellId == smallCellId {
					return ctx.String(http.StatusConflict, "SmallCell "+sm.DisplayName+" already exists at site "+*site.DisplayName)
				}
			}
			flag = true
			siteIndex = sIndex
		}
		if flag {
			s1[siteIndex].SmallCells = append(s1[siteIndex].SmallCells, smc)
			return ctx.String(http.StatusOK, "Slice"+smallCellId+" added at site "+*s1[siteIndex].DisplayName)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}
