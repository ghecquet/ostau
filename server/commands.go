package server

// make that a proto

// TODO - probably move this to its own dir
type Map map[string]interface{}

type SmartHomeV1Intents string

const (
	ActionDevicesSYNC      SmartHomeV1Intents = "action.devices.SYNC"
	ActionDevicesQUERY     SmartHomeV1Intents = "action.devices.QUERY"
	ActionDeviceEXECUTE    SmartHomeV1Intents = "action.devices.EXECUTE"
	ActionDeviceDISCONNECT SmartHomeV1Intents = "action.devices.DISCONNECT"
)

type SmartHomeV1ExecuteStatus string

const (
	SmartHomeV1ExecuteStatusSUCCESS SmartHomeV1ExecuteStatus = "SUCCESS"
	SmartHomeV1ExecuteStatusPENDING SmartHomeV1ExecuteStatus = "PENDING"
	SmartHomeV1ExecuteStatusOFFLINE SmartHomeV1ExecuteStatus = "OFFLINE"
	SmartHomeV1ExecuteStatusERROR   SmartHomeV1ExecuteStatus = "ERROR"
)

type SmartHomeV1ExecuteErrors string

type SmartHomeV1SyncRequestInputs struct {
	intent SmartHomeV1Intents
}

type SmartHomeV1SyncRequest struct {
	requestId string
	inputs    []SmartHomeV1SyncRequestInputs
}

type SmartHomeV1QueryRequestDevices struct {
	id         string
	customData Map
}

type SmartHomeV1QueryRequestPayload struct {
	devices []SmartHomeV1QueryRequestDevices
}

type SmartHomeV1QueryRequestInputs struct {
	intent  SmartHomeV1Intents
	payload SmartHomeV1QueryRequestPayload
}

type SmartHomeV1QueryRequest struct {
	requestId string
	inputs    []SmartHomeV1QueryRequestInputs
}

type SmartHomeV1ExecuteRequestExecution struct {
	command   string
	params    Map
	challenge struct {
		pin string
		ack bool
	}
}

type SmartHomeV1ExecuteRequestCommands struct {
	devices   []SmartHomeV1QueryRequestDevices
	execution []SmartHomeV1ExecuteRequestExecution
}

type SmartHomeV1ExecuteRequestPayload struct {
	commands []SmartHomeV1ExecuteRequestCommands
}

type SmartHomeV1ExecuteRequestInputs struct {
	intent  SmartHomeV1Intents
	payload SmartHomeV1ExecuteRequestPayload
}

type SmartHomeV1DisconnectRequestInputs struct {
	intent  SmartHomeV1Intents
	payload SmartHomeV1ExecuteRequestPayload
}

type SmartHomeV1ExecuteRequest struct {
	requestId string
	inputs    []SmartHomeV1ExecuteRequestInputs
}

type SmartHomeV1DisconnectRequest struct {
	requestId string
	inputs    []SmartHomeV1DisconnectRequestInputs
}

// export type SmartHomeV1Request = SmartHomeV1SyncRequest | SmartHomeV1QueryRequest |
//   SmartHomeV1ExecuteRequest | SmartHomeV1DisconnectRequest

type SmartHomeV1SyncName struct {
	defaultNames []string
	name         string
	nicknames    []string
}

type SmartHomeV1SyncDeviceInfo struct {
	manufacturer string
	model        string
	hwVersion    string
	swVersion    string
}

type SmartHomeV1SyncOtherDeviceIds struct {
	agentId  string
	deviceId string
}

type SmartHomeV1SyncDevices struct {
	id              string
	deviceType      string
	traits          []string
	name            SmartHomeV1SyncName
	willReportState bool
	deviceInfo      SmartHomeV1SyncDeviceInfo
	attributes      Map
	customData      Map
	roomHint        string
	otherDeviceIds  []SmartHomeV1SyncOtherDeviceIds
}

type SmartHomeV1SyncPayload struct {
	agentUserId string
	errorCode   string
	debugString string
	devices     []SmartHomeV1SyncDevices
}

type SmartHomeV1SyncResponse struct {
	requestId string
	payload   SmartHomeV1SyncPayload
}

type SmartHomeV1QueryPayload struct {
	devices Map
}

type SmartHomeV1QueryResponse struct {
	requestId string
	payload   SmartHomeV1QueryPayload
}

type SmartHomeV1ExecuteResponseCommands struct {
	ids             []string
	status          SmartHomeV1ExecuteStatus
	errorCode       SmartHomeV1ExecuteErrors
	debugString     string
	states          Map
	challengeNeeded struct {
		challengeType ChallengeType
	}
}

type ChallengeType string

const (
	ChallengeTypeACKNEEDED       = "ackNeeded"
	ChallengeTypePINNEEDED       = "pinNeeded"
	ChallengeTypeFailedPINNEEDED = "challengeFailedPinNeeded"
)

type SmartHomeV1ExecutePayload struct {
	commands    []SmartHomeV1ExecuteResponseCommands
	errorCode   SmartHomeV1ExecuteErrors
	debugString string
}

type SmartHomeV1ExecuteResponse struct {
	requestId string
	payload   SmartHomeV1ExecutePayload
}

type SmartHomeV1DisconnectResponse struct{}

// export type SmartHomeV1Response = SmartHomeV1SyncResponse | SmartHomeV1QueryResponse |
//   SmartHomeV1ExecuteResponse | SmartHomeV1DisconnectResponse

type SmartHomeV1ReportStateRequest struct {
	requestId   string
	agentUserId string
	payload     SmartHomeV1ReportStatePayload
}

type SmartHomeV1ReportStatePayload struct {
	Devices SmartHomeV1ReportStateDevices
}

type SmartHomeV1ReportStateDevices struct {
	States Map
}
