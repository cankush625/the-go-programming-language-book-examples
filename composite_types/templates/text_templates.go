package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

// Create a template to parse the data into format
const deviceTemplate = `Device Information:
{{range .}}--------------------------------
Device Id: 	  {{.Id}}
Device Name:  {{.Name}}
Price: 		  {{.Price}}
Age:		  {{.CreatedAt | daysAgo}} days
Is Available: {{.Id | isAvailable}}
{{end}}
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func isAvailable(deviceId int) bool {
	deviceCount := map[int]int{1: 10, 2: 0, 3: 45}
	return deviceCount[deviceId] > 0
}

type Device struct {
	Id        int
	Name      string
	Price     int
	CreatedAt time.Time
}

func main() {
	// template.Must helper function makes error handling more convenient: it accepts
	// the template and an error, checks that the error is nil (and panics otherwise)
	// and then returns the template
	var devices = template.Must(template.New("devices").
		Funcs(template.FuncMap{"isAvailable": isAvailable, "daysAgo": daysAgo}).
		Parse(deviceTemplate))

	// Create a data to pass to the template
	data := []Device{
		{Id: 1, Name: "Apple Mac M1", Price: 3000, CreatedAt: time.Date(2021, time.Month(2), 21, 10, 00, 0, 0, time.UTC)},
		{Id: 2, Name: "Apple Ipad M1", Price: 2000, CreatedAt: time.Date(2022, time.Month(5), 21, 10, 00, 0, 0, time.UTC)},
	}
	// Execute the template using the device list as data and
	// os.Stdout as the destination
	if err := devices.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}

// Output:
// Device Information:
// --------------------------------
// Device Id:        1
// Device Name:  Apple Mac M1
// Price:            3000
// Age:              1025 days
// Is Available: true
// --------------------------------
// Device Id:        2
// Device Name:  Apple Ipad M1
// Price:            2000
// Age:              571 days
// Is Available: false
