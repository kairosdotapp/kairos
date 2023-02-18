package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fmt.Println("invoice-cli")
	flags.Usage = func() {
		fmt.Println("usage: siot [OPTION]... COMMAND [OPTION]...")
		fmt.Println("Global options:")
		flags.PrintDefaults()
		fmt.Println()
		fmt.Println("Available commands:")
		fmt.Println("  - invoice (create an invoice)")
	}

	flags.Parse(os.Args[1:])

	// extract sub command and its arguments
	args := flags.Args()

	if len(args) < 1 {
		// run serve command by default
		args = []string{"invoice"}
	}

	switch args[0] {
	case "invoice":
		if err := runInvoice(args[1:]); err != nil {
			log.Println("invoice error: ", err)
		}
	default:
		log.Fatal("Unknown command; options: invoice")
	}
}

func runInvoice(args []string) error {
	flags := flag.NewFlagSet("invoice", flag.ExitOnError)
	flagAccount := flags.String("account", "", "Account to process")
	flagYearMonth := flags.String("month", "", "specify year and month: YYYY-MM")
	flagDate := flags.String("data", "", "invoice date (YYYY-MM-DD). If blank, current date will be used.")
	flagUser := flags.String("user", "", "user time entries to use. If blank, all users will be processed.")
	flagNumber := flags.Int("num", 0, "invoice #")

	flags.Parse(args)

	if *flagYearMonth == "" {
		flags.Usage()
		return fmt.Errorf("Error, must specify invoice month")
	}

	fCustomers, err := os.Open("customers.csv")
	if err != nil {
		return fmt.Errorf("Error opening customers.csv: %v", err)
	}

	if *flagNumber <= 0 {
		flags.Usage()
		return fmt.Errorf("Error, Invoice number must be set")
	}

	defer fCustomers.Close()

	customers, err := parseCustomers(fCustomers)

	if err != nil {
		return fmt.Errorf("Error parsing customers file: %v", err)
	}

	fRates, err := os.Open("rates.csv")
	if err != nil {
		return fmt.Errorf("Error opening rates.csv: %v", err)
	}

	rates, err := parseRates(fRates)
	if err != nil {
		return fmt.Errorf("Error parsing rates: %v", err)
	}

	var userLogFiles []string

	if *flagUser != "" {
		userLogFiles = []string{*flagUser + ".timedot"}
	} else {
		userLogFiles, err = filepath.Glob("*.timedot")
		if err != nil {
			return fmt.Errorf("Error searching for timedot files: %v", err)
		}
	}

	var es entries

	for _, u := range userLogFiles {
		f, err := os.Open(u)
		if err != nil {
			return fmt.Errorf("Error opening %v: %v", u, err)
		}

		e, err := parseTimedot(f)
		f.Close()
		if err != nil {
			return fmt.Errorf("Error parsing %v: %v", u, err)
		}

		user := strings.TrimSuffix(u, ".timedot")

		err = e.populateCost(rates, user)
		if err != nil {
			return fmt.Errorf("Error populating cost for user %v: %v", user, err)
		}

		es = append(es, e...)
	}

	sort.Sort(es)

	if len(es) <= 0 {
		return fmt.Errorf("No log entries found")
	}

	inv, err := invoice(*flagNumber, es, customers, *flagAccount, *flagYearMonth, *flagDate)

	if err != nil {
		return fmt.Errorf("Error creating invoice: %v", err)
	}

	accountDashes := strings.ReplaceAll(*flagAccount, ":", "-")

	invoiceName := fmt.Sprintf("%v_%v.html", *flagYearMonth, accountDashes)

	err = os.WriteFile(invoiceName, []byte(inv), 0644)

	if err != nil {
		return fmt.Errorf("Error writing invoice file: %v", err)
	}

	log.Println("Invoice generated: ", invoiceName)

	return nil
}
