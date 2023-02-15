<table>
  <tr>
    <th>Date</th>
    <th>Who</th>
    <th>Log</th>
    <th>Hours</th>
    <th>Cost</th>
  </tr>
  {{- range $entry := .Entries }}
    <tr>
      <th>{{ $entry.Date }}</th>
      <th>{{ $entry.User }}</th>
      <th>
        <ul>
        {{- range $log := $entry.Logs }}
          <li>{{ $log }}</li>
        {{- end }}
        </ul>
      </th>
      <th>{{ $entry.Hours }}</th>
      <th>{{ $entry.Cost }}</th>
    </tr>
  {{- end }}
</table>
