<!DOCTYPE html>
<html lang="en">
    <head>
        <title>Invoice</title>
        <meta name="keywords" content="invoice" />
        <meta name="description" content="foo"/>
        <meta charset="UTF-8"/>
        <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
        <link rel="stylesheet" href="https://cdn.simplecss.org/simple.min.css"/>
    </head>
    <body>
        <header>
            <br/>
            <img src="https://raw.githubusercontent.com/kairosdotapp/kairos/main/logo.png"/>
        </header>
        <main>
            <section>
                <div style="display: flex;">
                    <div>
                        K Consulting
                        <br/>
                        2342 West Ave.
                        <br/>
                        City, ST  12345
                        <br/>
                    </div>
                </div>
                <br/>
                <div style="display: flex; justify-content: space-between">
                    <b>Invoice #{{ .Number }}</b>
                    {{ .Date }}
                </div>
                <br/>
                To:
                <br/>
                {{ .Customer.Name }}
                <br/>
                {{ .Customer.Address }}
                <br/>
                {{ .Customer.City }}, {{ .Customer.State }} {{ .Customer.Zip }}
            </section>
            <section>
                <table class="table-auto">
                    <thead>
                        <tr>
                            <th>Project</th>
                            <th>Hours</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{- range $entry := .ProjectTotals }}
                        <tr>
                            <td>{{ $entry.Name }}</td>
                            <td>{{ $entry.Hours }}</td>
                        </tr>
                        {{- end }}
                    </tbody>
                </table>
                Total hours: <b>{{ .Hours }}</b>
            </section>
            <section>
                <table>
                    <col style="width: 17%;">
                    <thead>
                        <tr>
                            <th>Date</th>
                            <th>Who</th>
                            <th>Project</th>
                            <th>Description</th>
                            <th>Hours</th>
                            <th>Cost</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{- range $entry := .Entries }}
                        <tr>
                            <td>{{ $entry.Date }}</td>
                            <td>{{ $entry.User }}</td>
                            <td>{{ $entry.Project }}</td>
                            <td>
                                <ul>
                                    {{- range $log := $entry.Logs }}
                                    <li>{{ $log }}</li>
                                    {{- end }}
                                </ul>
                            </td>
                            <td>{{ $entry.Hours }}</td>
                            <td>${{ $entry.Cost }}</td>
                        </tr>
                        {{- end }}
                    </tbody>
                </table>
                <br/>
                <br/>
                {{- if .Expenses }}
                <b>Time subtotal: ${{ .Cost }}</b>
                <br/>
                {{- end }}
            </section>
            {{- if .Expenses }}
            <section>
                <h3>Supplies</h3>
                <table>
                    <col style="width: 17%;">
                    <thead>
                        <tr>
                            <th>Date</th>
                            <th>Vendor</th>
                            <th>Description</th>
                            <th>Amount</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{- range $expense := .Expenses }}
                        <tr>
                            <td>{{ $expense.Date.Format "2006-01-02" }}</td>
                            <td>{{ $expense.Vendor }}</td>
                            <td>{{ $expense.Description }}</td>
                            <td>${{ printf "%.2f" $expense.Amount }}</td>
                        </tr>
                        {{- end }}
                    </tbody>
                </table>
                <br/>
                <b>Supplies subtotal: ${{ printf "%.2f" .ExpenseTotal }}</b>
                <br/>
                <br/>
                <hr/>
                <b>Total amount due: ${{ printf "%.2f" .TotalWithExpenses }}</b>
                <br/>
            </section>
            {{- else }}
            <section>
                <b>Total amount due: ${{ .Cost }}</b>
                <br/>
            </section>
            {{- end }}
        </main>
        <footer>
            <b>Thank you for your valued business!</b>
            <br/>
            <br/>
        </footer>
    </body>
</html>
