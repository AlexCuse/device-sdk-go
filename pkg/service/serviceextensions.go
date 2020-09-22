package service

import (
	"github.com/edgexfoundry/device-sdk-go/internal/common"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/metadata"
	dsModels "github.com/edgexfoundry/go-mod-core-contracts/models"
)

func (s *DeviceService) GetDeviceClient() metadata.DeviceClient {
	return common.DeviceClient
}

func (s *DeviceService) GetCurrentConfig() []DeviceConfig {
	res := make([]DeviceConfig, 0)

	for _, d := range common.CurrentConfig.DeviceList {
		res = append(res, DeviceConfig{
			Name:        d.Name,
			Profile:     d.Profile,
			Description: d.Description,
			Labels:      d.Labels,
			Protocols:   d.Protocols,
			AutoEvents:  d.AutoEvents,
		})
	}

	return res
}

type DeviceConfig struct {
	// Name is the Device name
	Name string
	// Profile is the profile name of the Device
	Profile string
	// Description describes the device
	Description string
	// Other labels applied to the device to help with searching
	Labels []string
	// Protocols for the device - stores protocol properties
	Protocols map[string]dsModels.ProtocolProperties
	// AutoEvent supports auto-generated events sourced from a device service
	AutoEvents []dsModels.AutoEvent
}
