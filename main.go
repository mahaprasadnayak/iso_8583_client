package main

import (
	"fmt"
	"iso/service"
	"net"
	"time"

	"github.com/moov-io/iso8583"
	"github.com/moov-io/iso8583/encoding"
	"github.com/moov-io/iso8583/field"
	"github.com/moov-io/iso8583/network"
	"github.com/moov-io/iso8583/prefix"
)

func main() {
	fmt.Println("In Main")
	client()
}
func client() {
	tcpAddr, err := net.ResolveTCPAddr("tcp","192.168.1.13:9091")
    if err != nil {
        println("ResolveTCPAddr failed:", err.Error())
            
    }
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
        println("Dial failed:", err.Error())
           
    }
	spec:=&iso8583.MessageSpec{
		Name: "ISO 8583 v1987 ASCII",
		Fields: map[int]field.Field{
			0: field.NewString(&field.Spec{
				Length:      4,
				Description: "Message Type Indicator",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
			}),
			1: field.NewBitmap(&field.Spec{
				Length:      16,
				Description: "Bitmap",
				Enc:         encoding.BytesToASCIIHex,
				Pref:        prefix.Hex.Fixed,
			}),
			7: field.NewString(&field.Spec{
				Length:      10,
				Description: "Transmission Date & Time",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
			}),
			11: field.NewString(&field.Spec{
				Length:      6,
				Description: "Systems Trace Audit Number (STAN)",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
			}),
			39: field.NewString(&field.Spec{
				Length:      2,
				Description: "Response Code",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
			}),
			70: field.NewString(&field.Spec{
				Length:      3,
				Description: "Network management information code",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
			}),
		},
	}
	message := iso8583.NewMessage(spec)
	err1 := message.Marshal(&service.NetworkManagementRequest{
		MTI:                  field.NewStringValue("0800"),
		TransmissionDateTime: field.NewStringValue(time.Now().UTC().Format("18022022")),
		STAN:                 field.NewStringValue("000001"),
		InformationCode:      field.NewStringValue("001"),
	})
	if err1!=nil{
		fmt.Println("Error in marshalling",err1)
	}
	
	header := network.NewBCD2BytesHeader()
	rawMessage, err := message.Pack()
	if err != nil {
		fmt.Println("error is ", err)
	}
	fmt.Println("Raw message is ", string(rawMessage))
	header.SetLength(len(rawMessage))
	_, err = header.WriteTo(conn)
	if err != nil {
		fmt.Println("error in write message: ", err)
	}
	n, err := conn.Write(rawMessage)
	if err != nil {
		fmt.Println("error in write message: ", err)
	}
	fmt.Println("rev to server",n)

// 	//reading
	header1 := network.NewBCD2BytesHeader()
	it, rederr := header1.ReadFrom(conn)
	if rederr != nil {
	fmt.Println("error",err)
	}
	fmt.Println("it",it)

// // Make a buffer to hold message
// 	//buf := make([]byte, header1.Length())
// // Read the incoming message into the buffer.
// 	read, err := io.ReadFull(conn, rawMessage)
// 	if err != nil {
// 		fmt.Println("error",err)
// 	}
// 	fmt.Println("read",read)
// 	message1 := iso8583.NewMessage(specs.Spec87ASCII)
// 	message1.Unpack(rawMessage)
	

	
}


