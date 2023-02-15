package main

import (
	"reflect"
	"strings"
	"testing"

	"github.com/kr/pretty"
)

var customerTestData = `
account,name,address,city,state,zip
time:cust:a,Customer A,12343 S. Main St,Columbus,OH,12345
time:cust:b,Customer B,120 Central Ave,Dayton,OH,12345
time:cust:c,Customer C,12343 S. Main St,Wheeling,WV,12345
`

func testCustomers(t *testing.T) customers {
	custs, err := parseCustomers(strings.NewReader(customerTestData))
	if err != nil {
		t.Fatal("Error parsing customer: ", err)
	}

	return custs
}

func TestCustomerParser(t *testing.T) {
	exp := customers{
		{"time:cust:a", "Customer A", "12343 S. Main St", "Columbus", "OH", "12345"},
		{"time:cust:b", "Customer B", "120 Central Ave", "Dayton", "OH", "12345"},
		{"time:cust:c", "Customer C", "12343 S. Main St", "Wheeling", "WV", "12345"},
	}

	custs := testCustomers(t)

	if !reflect.DeepEqual(custs, exp) {
		pretty.Println("exp: ", exp)
		pretty.Println("custs: ", custs)
		t.Fatal("Did not get expected customers")
	}
}

func TestCustomerFind(t *testing.T) {
	custs := testCustomers(t)

	c, ok := custs.find("time:cust:a:proj1")
	if !ok {
		t.Fatal("Did not find customer")
	}

	exp := customer{"time:cust:a", "Customer A", "12343 S. Main St", "Columbus", "OH", "12345"}
	if c != exp {
		pretty.Println("Did not get correct customer exp: %v, got %v", exp, c)
		t.Fatal()
	}
}
