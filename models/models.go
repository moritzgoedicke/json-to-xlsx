package models

type DeviceData struct {
	Identifiers []struct {
		DeviceIds struct {
			DeviceID  string `json:"device_id"`
			DeviceEUI string `json:"dev_eui"`
		} `json:"device_ids"`
	} `json:"identifiers"`
	Data struct {
		ReceivedAt    string `json:"received_at"`
		UplinkMessage struct {
			DecodedPayload DecodedPayload `json:"decoded_payload"`
		} `json:"uplink_message"`
	} `json:"data"`
}

type DecodedPayload struct {
	BAT float64 `json:"BAT"`
	H1  float64 `json:"H1"`
	H2  float64 `json:"H2"`
	T1  float64 `json:"T1"`
}
