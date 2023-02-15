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
            <h2>Invoice</h2>
        </header>
        <main>
            <section>
                <div style="display: flex;">
                    <div>
                        BEC Systems
                        <br/>
                        15000 Warwick Rd.
                        <br/>
                        Marshallville, OH  44645
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
                            <th>Date</th>
                            <th>Who</th>
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
                Total hours: <b>{{ .Hours }}</b>
                <br/>
                <br/>
                <b>Amount due: ${{ .Cost }}</b>
                <br/>
            </section>
        </main>
        <footer>
            <b>Thank you for your valued business!</b>
            <br/>
            <br/>
            <img src="http://bec-systems.com/site/wp-content/uploads/2020/03/bec_logo_blue_146x48-1.png"/>
        </footer>
    </body>
</html>
