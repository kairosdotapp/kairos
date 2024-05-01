# Kairos Example

To run this example (in a Linux/MacOS shell):

[`./run-example.sh`](run-example.sh)

This will generate two invoices:

- [AB Corp](https://kairosdotapp.github.io/kairos/example/2024-04_time-cust-abcorp.html)
- [XYZ Inc.](https://kairosdotapp.github.io/kairos/example/2024-04_time-cust-xyzinc.html)

These can be printed or copy/pasted into an email.

The following files are required:

- [`rates.csv`](rates.csv): defines rates for customers/projects/tasks/users.
  The shortest match with the timelog entires is used.
- [`customers.csv`](customers.csv): customer information
- An invoice template ([example](../invoice.tpl)): Template that defines how the
  invoice looks. This can be modifiy to include/remove columns, change logo,
  etc.
- timelog files ([fred](fred.timedot), [joe](joe.timedot)): time log files in
  [timedot](https://hledger.org/dev/hledger.html#timedot) format.

example invoice:

<img src="kairos-example-invoice.png" width="400">
