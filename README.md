![logo](logo.png)

# Invoice CLI

CLI tool to generate invoices from a hledger
[timedot file](https://hledger.org/dev/hledger.html#timedot).

## Requirements

- parse billing table for combination of user/account
  - match the most specific entry
  - rates can have part of an account. For example, a rate for `cust:a` would
    match time entries for `cust:a:proj1`
- parse all timedot files in dir and build collection of entries and add amounts
- generate an invoice for a particular month
- generate ledger entries
- weekly status reports
