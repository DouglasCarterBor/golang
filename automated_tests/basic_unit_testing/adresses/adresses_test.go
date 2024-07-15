package adresses

import "testing"

type scenario struct { 
	address string
	expectedType string
}

func TestTypeOfAddress( t *testing.T) {
	
	// addressForTest := "street 123"
	// expectedAddressType := "Street"
	// typeAddressReceived := TypeAdresses(addressForTest)

	// if typeAddressReceived != expectedAddressType {
	// 	t.Errorf("Expected %s but received %s", expectedAddressType, typeAddressReceived)
	// }

	scenarios := []scenario{
		{ "street 123", "Street" },
		{ "avenue 123", "Avenue" },
		{ "boulevard 123", "Boulevard" },
		{ "road 123", "Road" },
		{ "place 123", "Place" },
		{ "court 123", "Court" },
	}

	for _, scenario := range scenarios {
		typeAddressReceived := TypeAdresses(scenario.address)
		if typeAddressReceived != scenario.expectedType {
			t.Errorf("Expected %s but received %s", scenario.expectedType, typeAddressReceived)
		}
	}
}
