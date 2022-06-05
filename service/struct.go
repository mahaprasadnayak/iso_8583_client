package service

import (
	"net"

	"github.com/moov-io/iso8583/field"
)

type Data struct {
	//Mode string `json:"mode"` //client mode
	// PAN  string `json:"pan"`  //PAN number
	// PC   int64  `json:"pc"`   //processing code number
	// TA   string `json:"ta"`   //txn amount
	// SA   string `json:"sa"`   //settlement amount
	// BA   string `json:"ba"`   //billing amount
}

type Client struct {
	socket net.Conn
	data   chan []byte
}

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}
type NetworkManagementRequest struct {
	MTI                  *field.String `index:"0"`
	TransmissionDateTime *field.String `index:"7"`
	STAN                 *field.String `index:"11"`
	InformationCode      *field.String `index:"70"`
	
}
type NetworkManagementResponse struct {
	MTI                  *field.String `index:"0"`
	TransmissionDateTime *field.String `index:"7"`
	STAN                 *field.String `index:"11"`
	InformationCode      *field.String `index:"70"`
	ResponseCode 		 *field.String `index:"39"`
}