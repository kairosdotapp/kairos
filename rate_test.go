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
time:cust:a:onsite,fred,200
time:cust:b,,70
time:cust:c,,65
`

func TestRateParser(t *testing.T) {
	exp := rates{
		{"time:cust:a", "", 80},
		{"time:cust:a", "fred", 90},
		{"time:cust:a:onsite", "", 150},
		{"time:cust:a:onsite", "fred", 200},
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
		t.Fatal("Did not get expected rates")
	}
}

func TestRateFind(t *testing.T) {
	tests := []struct {
		account string
		user    string
		exp     float32
	}{
		{"time:cust:a", "", 80},
		{"time:cust:a", "fred", 90},
		{"time:cust:a:onsite", "", 150},
		{"time:cust:b", "", 70},
		{"time:cust:c", "cliff", 65},
		{"time:cust:a:onsite", "fred", 200},
	}

	rates, err := parseRates(strings.NewReader(testRateData))
	if err != nil {
		t.Fatal("Error parsing rates: ", err)
	}

	for _, test := range tests {
		r, ok := rates.find(test.account, test.user)
		if !ok {
			pretty.Printf("did not find expected rate for %v\n", test)
			t.Fatal()
		}

		if r != test.exp {
			pretty.Printf("rate for %v did not match, got: %v\n",
				test, r)
			t.Fatal()
		}
	}
}
