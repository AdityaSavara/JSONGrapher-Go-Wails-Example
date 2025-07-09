package jsonrecord

import (
	"encoding/json"
	"os"
)

// JSONGrapherRecord represents the main structure of a JSONGrapher record.
type JSONGrapherRecord struct {
	Comments    string                  `json:"comments,omitempty"`
	JsonGrapher string                  `json:"jsongrapher,omitempty"`
	Datatype    string                  `json:"datatype,omitempty"`
	Data        []JSONGrapherDataSeries `json:"data,omitempty"`
	Layout      Layout                  `json:"layout,omitempty"` // Layout struct is defined in style.go but used here
	PlotStyle   PlotStyle               `json:"plot_style,omitempty"`
}

// PlotStyle represents the structure for plot-wide styling, including layout and trace style collections.
type PlotStyle struct {
	LayoutStyle           string `json:"layout_style,omitempty"`
	TraceStylesCollection string `json:"trace_styles_collection,omitempty"`
}

type JSONGrapherDataSeries struct {
	Name       string        `json:"name,omitempty"`
	TraceStyle string        `json:"trace_style,omitempty"`
	X          []interface{} `json:"x,omitempty"`
	Y          []interface{} `json:"y,omitempty"`
	Line       *Line         `json:"line,omitempty"`   // Line struct defined in style.go
	Marker     *Marker       `json:"marker,omitempty"` // Marker struct defined in style.go
	Mode       string        `json:"mode,omitempty"`
	Type       string        `json:"type,omitempty"`
}

// createNewJSONGrapherRecord creates and returns a new empty JSONGrapherRecord
// with default values, including the JsonGrapher string.
func CreateNewJSONGrapherRecord() *JSONGrapherRecord {
	return &JSONGrapherRecord{
		JsonGrapher: "To plot this file, go to www.jsongrapher.com and drag this file into your browser, or use the python version of JSONGrapher. File created with python Version 3.0",
		Data:        make([]JSONGrapherDataSeries, 0),
		Layout: Layout{ // Layout struct used here, defined in style.go
			Title: Title{Text: "Graph Title"},               // Title struct used here, defined in style.go
			XAxis: Axis{Title: Title{Text: "X-axis Label"}}, // Axis struct used here, defined in style.go
			YAxis: Axis{Title: Title{Text: "Y-axis Label"}}, // Axis struct used here, defined in style.go
		},
		PlotStyle: PlotStyle{}, // Initialize PlotStyle
	}
}

func (r *JSONGrapherRecord) AddDataSeries(series JSONGrapherDataSeries) {
	r.Data = append(r.Data, series)
}

// SetComments sets the comments field of the JSONGrapherRecord.
func (r *JSONGrapherRecord) SetComments(comments string) {
	r.Comments = comments
}

// SetDatatype sets the datatype field of the JSONGrapherRecord.
func (r *JSONGrapherRecord) SetDatatype(datatype string) {
	r.Datatype = datatype
}

// SetXAxisLabelIncludingUnits sets the title text for the X-axis.
func (r *JSONGrapherRecord) SetXAxisLabelIncludingUnits(labelWithUnits string) {
	r.Layout.XAxis.Title.Text = labelWithUnits
}

// SetYAxisLabelIncludingUnits sets the title text for the Y-axis.
func (r *JSONGrapherRecord) SetYAxisLabelIncludingUnits(labelWithUnits string) {
	r.Layout.YAxis.Title.Text = labelWithUnits
}

// SetLayoutStyle sets the 'layout_style' field within PlotStyle.
// This function only sets the string name, it does not apply the style immediately.
func (r *JSONGrapherRecord) SetLayoutStyle(styleName string) {
	r.PlotStyle.LayoutStyle = styleName
}

// SetTraceStylesCollection sets the 'trace_styles_collection' field within PlotStyle.
// This function only sets the string name, it does not apply the styles immediately.
func (r *JSONGrapherRecord) SetTraceStylesCollection(collectionName string) {
	r.PlotStyle.TraceStylesCollection = collectionName
}

// This function writes the JSONGrapherRecord to file.
// Typically, file name string should include ".json" at the end.
func (r *JSONGrapherRecord) WriteToFile(filename string) error {
	jsonOutput, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, jsonOutput, 0644)
}

// This function pretty prints the JSONGrapherRecord
func (r *JSONGrapherRecord) Print() error {
	jsonOutput, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		return err
	}
	os.Stdout.Write(jsonOutput)
	os.Stdout.Write([]byte("\n"))
	return nil
}
