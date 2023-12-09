package vc

import (
	"encoding/json"
	"io"
)

const (
	DEVICEINFO = "DeviceInfo"
)

func getDeviceInfo(vc *vc) (DeviceInfo, error) {

	resp, err := vc.client.Get(vc.url + DEVICEINFO)
	if err != nil {
		return emptyDeviceInfo(), err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return emptyDeviceInfo(), err
	}

	if resp.StatusCode != 200 {
		return emptyDeviceInfo(), newResponseError(resp.StatusCode)
	}

	var deviceData Device
	err = json.Unmarshal(body, &deviceData)
	if err != nil {
		return emptyDeviceInfo(), err
	}

	return deviceData.DeviceInfo, nil
}

func emptyDeviceInfo() DeviceInfo {
	return DeviceInfo{
		DeviceKey:  "",
		MacAddress: "00:00:00:00:00:00",
	}
}
