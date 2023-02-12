package main

import (
	"reflect"
	"strings"
	"testing"

	"github.com/kr/pretty"
)

var testRateData = `
account,user,rate
time:cust:a,,80
time:cust:a,fred,90
time:cust:a:onsite,,150
time:cust:b,,70
time:cust:c,,65
`

func TestRateParser(t *testing.T) {
	exp := []rate{
		{"time:cust:a", "", 80},
		{"time:cust:a", "fred", 90},
		{"time:cust:a:onsite", "", 150},
		{"time:cust:b", "", 70},
		{"time:cust:c", "", 65},
	}

	rates, err := parseRates(strings.NewReader(testRateData))
	if err != nil {
		t.Fatal("Error parsing rates: ", err)
	}
	if !reflect.DeepEqual(rates, exp) {
		pretty.Println("exp: ", exp)
		pretty.Println("rates: ", rates)
		t.Fatal("Did not get expected rates: ", err)
	}
}
