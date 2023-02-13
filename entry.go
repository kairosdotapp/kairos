package main

type entry struct {
	date    string
	account string
	logs    []string
	hours   float32
	user    string
	cost    float32
}

func (e *entry) setDate(d string) {
	e.date = d
	e.account = ""
	e.logs = []string{}
	e.hours = 0
}

func (e *entry) setAccount(a string) {
	e.account = a
	e.logs = []string{}
	e.hours = 0
}

func (e *entry) hasDate() bool {
	if e.date != "" {
		return true
	}

	return false
}

func (e *entry) hasAccount() bool {
	if e.account != "" {
		return true
	}

	return false
}

func (e *entry) clearDate() {
	e.date = ""
	e.account = ""
	e.logs = []string{}
	e.hours = 0
}

func (e *entry) clearAccount() {
	e.account = ""
	e.logs = []string{}
	e.hours = 0
}

type entries []entry

func (es *entries) populateCost(r rates, user string) error {
	for i, e := range *es {
		(*es)[i].cost = r.find(e.account, user) * e.hours
		(*es)[i].user = user
	}

	return nil
}
