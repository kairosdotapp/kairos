![logo](logo.png)

Kairos is a tool that generates beautiful invoices from time log files stored in
Git. This makes it easy for your team to log their time using their standard
tools (Git/Text Editor) and then generate invoices and other reports using
[hledger](https://hledger.org/).

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

## Example

See the [example](example/) for information on how to use this tool.

Example invoice:

![example invoice](example/kairos-example-invoice.png =400x300)
