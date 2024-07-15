package adresses

import "testing"

func TestTypeOfAddress( t *testing.T) {
	addressForTest := "street 123"
	expectedAddressType := "Street"
	typeAddressReceived := TypeAdresses(addressForTest)

	if typeAddressReceived != expectedAddressType {
		t.Errorf("Expected %s but received %s", expectedAddressType, typeAddressReceived)
	}
	
}
