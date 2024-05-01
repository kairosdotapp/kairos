#!/bin/sh

go run ../ invoice -account time:cust:abcorp -num 1002 -month 2024-04 -template ../invoice.tpl
go run ../ invoice -account time:cust:xyzinc -num 1003 -month 2024-04 -template ../invoice.tpl
