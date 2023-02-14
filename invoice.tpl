<table>
  <tr>
    <th>Date</th>
    <th>Who</th>
    <th>Log</th>
    <th>Hours</th>
    <th>Amount</th>
  </tr>
  {{- range $entry := .Entries }}
    <tr>
      <th>{{ $entry.Date }}</th>
    </tr>
  {{- end }}
</table>
