package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

var devicesTemplate = `
<h1>Devices</h1>
<table>
<tr style='text-align: left'>
	<th>Id</th>
	<th>Name</th>
	<th>Price</th>
</tr>
{{range .}}
<tr>
	<td>{{.Id}}</td>
	<td>{{.Name}}</td>
	<td>{{.Price}}</td>
</tr>
{{end}}
</table>
`

type Devices struct {
	Id        int
	Name      string
	Price     int
	CreatedAt time.Time
}

func main() {
	var devices = template.Must(template.New("devices").
		Parse(devicesTemplate))

	// Create a data to pass to the template
	data := []Devices{
		{Id: 1, Name: "Apple Mac M1", Price: 3000, CreatedAt: time.Date(2021, time.Month(2), 21, 10, 00, 0, 0, time.UTC)},
		{Id: 2, Name: "Apple Ipad M1", Price: 2000, CreatedAt: time.Date(2022, time.Month(5), 21, 10, 00, 0, 0, time.UTC)},
	}
	// Execute the template using the device list as data and
	// os.Stdout as the destination
	if err := devices.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
