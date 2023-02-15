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
            <h1>Invoice</h1>
            Invoice {{ .Number }}
        </header>
        <main>
            <section>
                <h2>To:</h2>
                Customer: CustA
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
                Total cost: <b>${{ .Cost }}</b>
                <br/>
            </section>
        </main>
        <footer>
            <b>Thank you for your valued business!</b>
        </footer>
    </body>
</html>
