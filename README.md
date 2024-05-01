![logo](logo.png)

Kairos is a tool that generates beautiful invoices from time log files stored in
Git. This makes it easy for your team to log their time using their standard
tools (Git/Text Editor) and then generate invoices and other reports using
[hledger](https://hledger.org/).

## Features

- time is logged in [timedot](https://hledger.org/dev/hledger.html#timedot)
  format (easy and convenient for developers)
- supports multiple projects/activities (you can track at any granularity your
  activities)
- multi-user support with different rates for different users
- full power of hledger for generating reports and integration with your
  accounting system.

## Example

See the [example](example/) for information on how to use this tool.

Example invoice:

<img src="example/kairos-example-invoice.png" width="400">

## Requirements

- parse billing table for combination of user/account
  - match the most specific entry
  - rates can have part of an account. For example, a rate for `cust:a` would
    match time entries for `cust:a:proj1`
- parse all timedot files in dir and build collection of entries and add amounts
- generate an invoice for a particular month
- generate ledger entries
- weekly status reports
