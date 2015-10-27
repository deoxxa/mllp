package mllp

import (
	"bytes"
	"fmt"
	"testing"
)

const sampleMessage = `MSH|^~\&|ZIS|1^AHospital|||199605141144||ADT^A01|20031104082400|P|2.3|||
AL|NE|||8859/15|<CR>EVN|A01|20031104082400.0000+0100|20031104082400
PID||""|10||Vries^Danny^D.^^de||19951202|M|||Rembrandlaan^7^Leiden^^7301TH^""
^^P||""|""||""|||||||""|""<CR>PV1||I|3w^301^""^01|S|||100^van den Berg^^A.S.
^^""^dr|""||9||||H||||20031104082400.0000+0100`

func wrapWithMarkers(b []byte) []byte {
	return append(append([]byte{0x0b}, b...), 0x1c, 0x0d)
}

func TestReadMessage(t *testing.T) {
	r := NewReader(bytes.NewReader(wrapWithMarkers([]byte(sampleMessage))))

	m, err := r.ReadMessage()
	if err != nil {
		panic(err)
	}

	if string(m) != sampleMessage {
		panic(fmt.Errorf("data was corrupted"))
	}
}
