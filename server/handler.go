package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

func handleSmartHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := getRequest(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}

		resp, err := getResponse(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}

		resp.RequestId = req.RequestID
		resp.Payload.AgentUserId = "123"

		for _, input := range req.Inputs {
			// payload := bytes.NewBuffer(input.Payload)

			switch input.Intent {
			case ActionDevicesSYNC:
				resp.Payload.Devices = handleSmartHomeSync()
			case ActionDevicesQUERY:
				resp.Payload.Devices = handleSmartHomeQuery(input.Payload)
			case ActionDeviceEXECUTE:
				resp.Payload.Commands = handleSmartHomeExecute(input.Payload)
			default:
				http.Error(w, fmt.Errorf("not implemented").Error(), http.StatusNotImplemented)
				return
			}
		}
	}
}

func handleSmartHomeSync() interface{} {

	// TODO - get devices from database

	return []interface{}{
		SmartHomeV1SyncDevices{
			ID:         "Washer",
			DeviceType: "action.devices.types.WASHER",
			Traits: []string{
				"action.devices.traits.OnOff",
				"action.devices.traits.StartStop",
				"action.devices.traits.RunCycle",
			},
			Name: SmartHomeV1SyncName{
				DefaultNames: []string{"My Washer"},
				Name:         "Washer",
				Nicknames:    []string{"Washer"},
			},
			DeviceInfo: SmartHomeV1SyncDeviceInfo{
				Manufacturer: "Test",
				Model:        "Test",
				HwVersion:    "1.0",
				SwVersion:    "1.0.1",
			},
			WillReportState: true,
			Attributes: Map{
				"pausable": true,
			},
			OtherDeviceIds: []SmartHomeV1SyncOtherDeviceIds{
				SmartHomeV1SyncOtherDeviceIds{
					DeviceId: "deviceid123",
				},
			},
		},
		SmartHomeV1SyncDevices{
			ID:         "Light",
			DeviceType: "action.devices.types.LIGHT",
			Traits: []string{
				"action.devices.traits.OnOff",
			},
			Name: SmartHomeV1SyncName{
				DefaultNames: []string{"My Light"},
				Name:         "My Light",
				Nicknames:    []string{"My Light"},
			},
			DeviceInfo: SmartHomeV1SyncDeviceInfo{
				Manufacturer: "Test",
				Model:        "Test",
				HwVersion:    "1.0",
				SwVersion:    "1.0.1",
			},
			WillReportState: true,
			Attributes: Map{
				"pausable": true,
			},
			OtherDeviceIds: []SmartHomeV1SyncOtherDeviceIds{
				SmartHomeV1SyncOtherDeviceIds{
					DeviceId: "my-light",
				},
			},
		},
	}
}

func handleSmartHomeQuery(input interface{}) interface{} {
	payload := new(SmartHomeV1QueryRequestPayload)

	err := mapstructure.Decode(input, payload)
	if err != nil {
		return nil
	}

	// TODO - get status from database

	return Map{
		"Washer": Map{
			"on":     false,
			"online": true,
		},
		"Light": Map{
			"on":     false,
			"online": true,
		},
	}
}

func handleSmartHomeExecute(input interface{}) interface{} {
	payload := new(SmartHomeV1ExecutePayload)

	err := mapstructure.Decode(input, payload)
	if err != nil {
		fmt.Println("Error is ", err)
		return nil
	}

	return []interface{}{
		Map{
			"ids":    []string{"Washer"},
			"status": "SUCCESS",
			"states": Map{
				"on":     true,
				"online": true,
			},
		},
	}
}

type ContextKey string

const ContextRequestKey ContextKey = "request"
const ContextResponseKey ContextKey = "response"

func wrapHandler(h http.HandlerFunc) http.HandlerFunc {
	return wrapRequest(wrapResponse(h))
}

func wrapRequest(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := new(FullfillmentRequest)

		data, _ := ioutil.ReadAll(r.Body)
		fmt.Println("Request ", string(data))

		decoder := json.NewDecoder(bytes.NewBuffer(data))

		if err := decoder.Decode(&request); err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}

		ctx := context.WithValue(r.Context(), ContextRequestKey, request)

		h(w, r.WithContext(ctx))
	}
}

func getRequest(ctx context.Context) (*FullfillmentRequest, error) {
	value := ctx.Value(ContextRequestKey)

	if response, ok := value.(*FullfillmentRequest); ok {
		return response, nil
	}

	return nil, fmt.Errorf("not found")
}

func wrapResponse(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := new(FullfillmentResponse)

		ctx := context.WithValue(r.Context(), ContextResponseKey, response)

		h(w, r.WithContext(ctx))

		js, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("Response ", string(js))

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func getResponse(ctx context.Context) (*FullfillmentResponse, error) {
	value := ctx.Value(ContextResponseKey)

	if response, ok := value.(*FullfillmentResponse); ok {
		return response, nil
	}

	return nil, fmt.Errorf("not found")
}
