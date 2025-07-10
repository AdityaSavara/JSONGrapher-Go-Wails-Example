package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	jgr "github.com/AdityaSavara/jsongrapher-go/jsonrecord"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
// This does *not* print, it just creates a string suitable for printing.
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}


//This function gets a JSON record from a file, for pre-existing records.
func (a *App) GetJSONFromFile(filename string) (interface{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file) // replaces ioutil.ReadAll
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var data interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return data, nil
}



//This function shows an example of how to create a JSONGrapherRecord.
//In practice, one would typically get xValues and yValues from a program
// rather than hardcoding them or loading them from file. 
func (a *App) CreateExampleJSON() (interface{}, error) {

	//Typically, user would have code here to generate the xValues and yValues.
	xValues := []interface{}{
		nil, 0.167, 0.33496, 0.50598, 0.67361, 0.84123, 1.00652, 1.17088, 1.33617, 1.49799,
		1.65619, 1.81733, 1.97414, 2.13134, 2.2892, 2.44419, 2.5978, 2.75083, 2.90336,
		3.05666, 3.20982, 3.35977, 3.51159, 3.66332, 3.81277, 3.96417, 4.11437, 4.26234,
		4.41177, 4.55836, 4.69914, 4.83091, 4.95142, 5.06196, 5.16176, 5.25171, 5.33179,
		5.40588, 5.47535, 5.60704, 5.7311, 5.83834, 5.92758, 5.99835, 6.05958, 6.12115,
		6.17853, 6.23654, 6.2809, 6.3319, 6.3937, 6.45549, 6.51578, 6.57668, 6.62612,
		6.67915,
	}
	yValues := []interface{}{
		nil, 136.55906, 131.45221, 120.59915, 118.16312, 127.37727, 129.73358, 133.77259,
		137.97628, 122.40547, 122.89972, 131.9398, 133.27169, 134.09176, 129.50373,
		126.03748, 124.21117, 124.68017, 126.28798, 125.05354, 129.32094, 126.03506,
		123.28353, 120.96402, 125.23938, 115.82581, 118.56249, 121.57265, 108.84001,
		86.15691, 83.9003, 78.60271, 65.59399, 36.99054, 36.41493, 26.2636, 14.7489,
		14.71689, 10.64583, 15.86853, 16.84498, 9.3199, 15.2712, 10.77233, 14.83896,
		12.38018, 10.94238, 15.22293, 4.3628, 6.23571, 6.96898, 12.52826, 18.08212,
		22.37963, 25.72857, 18.84665,
	}



	record := jgr.CreateNewJSONGrapherRecord()
	record.SetComments("CH4 Activation over Perovskite Catalysts")
	record.SetDatatype("CO2_Adsorption___Differential_Enthalpy")

	record.Layout.Title.Text = "CO2 Differential Enthalpy of Adsorption"
	record.SetXAxisLabelIncludingUnits("Quantity Adsorbed (µmol*m^(-2))")
	record.SetYAxisLabelIncludingUnits("Adsorption_Enthalpy (kJ*mol**(-1))")

	dataSeries := jgr.JSONGrapherDataSeries{
		Name:       "LaFeO3",
		TraceStyle: "scatter_spline",
		X:          xValues,
		Y:          yValues,
	}

	record.AddDataSeries(dataSeries)
	record.SetLayoutStyle("Nature")
	record.SetTraceStylesCollection("default")
	return record, nil
}
