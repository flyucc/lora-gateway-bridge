package structs

import (
	"github.com/brocaar/loraserver/api/gw"
	"github.com/brocaar/lorawan"
)

// DownlinkTransmitted implements the downlink transmitted message.
type DownlinkTransmitted struct {
	MessageType MessageType `json:"msgtype"`

	DIID uint32 `json:"diid"`
}

// DownlinkTransmittedToProto converts the DownlinkTransmitted to the protobuf struct.
func DownlinkTransmittedToProto(gatewayID lorawan.EUI64, dt DownlinkTransmitted) (gw.DownlinkTXAck, error) {
	return gw.DownlinkTXAck{
		GatewayId: gatewayID[:],
		Token:     dt.DIID,
	}, nil
}
