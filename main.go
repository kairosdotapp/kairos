package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func main() {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fmt.Println("kairos")
	flags.Usage = func() {
		fmt.Println("usage: siot [OPTION]... COMMAND [OPTION]...")
		fmt.Println("Global options:")
		flags.PrintDefaults()
		fmt.Println()
		fmt.Println("Available commands:")
		fmt.Println("  - invoice (create an invoice)")
	}

	_ = flags.Parse(os.Args[1:])

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
	flagDate := flags.String("date", "", "invoice date (YYYY-MM-DD). If blank, current date will be used.")
	flagUser := flags.String("user", "", "user time entries to use. If blank, all users will be processed.")
	flagNumber := flags.Int("num", 0, "invoice #")
	flagBegin := flags.String("begin", "", "beginning date")
	flagEnd := flags.String("end", "", "ending date")
	flagTemplate := flags.String("template", "invoice.tpl", "invoice template to use")

	_ = flags.Parse(args)

	if *flagYearMonth == "" && *flagBegin == "" {
		flags.Usage()
		return fmt.Errorf("Error, must specify invoice month or beginning date")
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

	accountDashes := strings.ReplaceAll(*flagAccount, ":", "-")

	var start time.Time
	var end time.Time
	var invoiceName string

	if *flagYearMonth != "" {
		start, err = time.Parse(time.DateOnly, *flagYearMonth+"-01")
		if err != nil {
			return fmt.Errorf("Error parsing invoice month: %v", err)
		}
		end = start.AddDate(0, 1, 0).Add(-time.Second)
		invoiceName = fmt.Sprintf("%v_%v.html", *flagYearMonth, accountDashes)
	} else if *flagBegin != "" {
		start, err = time.Parse(time.DateOnly, *flagBegin)
		if err != nil {
			return fmt.Errorf("Error parsing begin: %v", err)
		}
		if *flagEnd != "" {
			end, err = time.Parse(time.DateOnly, *flagEnd)
			if err != nil {
				return fmt.Errorf("Error parsing end: %v", err)
			}
		} else {
			end = time.Now()
		}

		startT := start.Format("2006-01-02")
		endT := end.Format("2006-01-02")

		invoiceName = fmt.Sprintf("%v_%v_%v.html", startT, endT, accountDashes)
	}

	inv, err := invoice(*flagNumber, es, customers, *flagAccount, start, end, *flagDate, *flagTemplate)

	if err != nil {
		return fmt.Errorf("Error creating invoice: %v", err)
	}

	err = os.WriteFile(invoiceName, []byte(inv), 0644)

	if err != nil {
		return fmt.Errorf("Error writing invoice file: %v", err)
	}

	log.Println("Invoice generated: ", invoiceName)

	return nil
}
