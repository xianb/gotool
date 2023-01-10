package utils

import "testing"

func TestIsRealNumberString(t *testing.T) {

	tests := []struct {
		text string
		tag  string
	}{
		{
			text: "-90000000",
			tag:  "zheng shu",
		},
		{
			text: "-1212",
			tag:  "zheng shu",
		},
		{
			text: "1.232",
			tag:  "xiao shu",
		},
		{
			text: "-1.232",
			tag:  "xiao shu",
		},
		{
			text: "dfsfsdfsdf",
			tag:  "charactor",
		},
	}

	for i, test := range tests {

		t.Logf("TestIsRealNumber[%d] = %v", i, IsRealNumberString(test.text))

	}

}
