package jsonrecord

import (
	"encoding/json"
	"fmt"

	"github.com/AdityaSavara/jsongrapher-go/styles"
)

// Structs related to layout and trace styling definitions
type Layout struct {
	Title        Title                  `json:"title,omitempty"`
	XAxis        Axis                   `json:"xaxis,omitempty"`
	YAxis        Axis                   `json:"yaxis,omitempty"`
	Font         map[string]interface{} `json:"font,omitempty"`
	PaperBgColor string                 `json:"paper_bgcolor,omitempty"`
	PlotBgColor  string                 `json:"plot_bgcolor,omitempty"`
}

type Title struct {
	Text string                 `json:"text,omitempty"`
	Font map[string]interface{} `json:"font,omitempty"`
	X    float64                `json:"x,omitempty"`
}

type Axis struct {
	Title     Title                  `json:"title,omitempty"`
	TickFont  map[string]interface{} `json:"tickfont,omitempty"`
	ShowGrid  bool                   `json:"showgrid,omitempty"`
	GridColor string                 `json:"gridcolor,omitempty"`
	GridWidth float64                `json:"gridwidth,omitempty"`
	LineColor string                 `json:"linecolor,omitempty"`
	LineWidth float64                `json:"linewidth,omitempty"`
	Ticks     string                 `json:"ticks,omitempty"`
	TickWidth float64                `json:"tickwidth,omitempty"`
	TickColor string                 `json:"tickcolor,omitempty"`
}

type Line struct {
	Shape string  `json:"shape,omitempty"`
	Width float64 `json:"width,omitempty"`
}

type Marker struct {
	Size float64 `json:"size,omitempty"`
}

// ApplyLayoutStyle applies the layout style previously set in PlotStyle.
func (r *JSONGrapherRecord) ApplyLayoutStyle() error {
	styleName := r.PlotStyle.LayoutStyle
	if styleName == "" {
		return fmt.Errorf("no layout style name has been set in PlotStyle")
	}

	// LayoutStylesLibrary is defined in layout_styles_library.go
	styleData, ok := styles.LayoutStylesLibrary[styleName]
	if !ok {
		return fmt.Errorf("layout style '%s' not found", styleName)
	}

	layoutMap, ok := styleData.(map[string]interface{})["layout"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid structure for layout style '%s'", styleName)
	}

	// This is a deep copy operation for the layout map to avoid modifying the global library directly
	jsonBytes, err := json.Marshal(layoutMap)
	if err != nil {
		return fmt.Errorf("error marshaling layout style: %w", err)
	}

	var newLayout Layout
	if err := json.Unmarshal(jsonBytes, &newLayout); err != nil {
		return fmt.Errorf("error unmarshaling layout style to struct: %w", err)
	}

	r.Layout = newLayout
	return nil
}

// applyTraceStyle is a helper function to apply a specific trace style from the
func (r *JSONGrapherRecord) ApplyTraceStyle(dataSeries *JSONGrapherDataSeries) error {
	collectionName := r.PlotStyle.TraceStylesCollection
	if collectionName == "" {
		return fmt.Errorf("no trace style collection name has been set in PlotStyle")
	}

	// TraceStylesCollectionLibrary is defined in trace_styles_collection_library.go
	collectionData, ok := styles.TraceStylesCollectionLibrary[collectionName]
	if !ok {
		return fmt.Errorf("trace style collection '%s' not found", collectionName)
	}

	collectionMap, ok := collectionData.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid structure for trace style collection '%s'", collectionName)
	}

	traceStyleName := dataSeries.TraceStyle
	if traceStyleName == "" {
		// If no trace_style is explicitly set on the data series, use a default from the collection if available
		if _, exists := collectionMap["scatter_spline"]; exists {
			traceStyleName = "scatter_spline"
		} else if _, exists := collectionMap["scatter"]; exists {
			traceStyleName = "scatter"
		} else if _, exists := collectionMap["line"]; exists {
			traceStyleName = "line"
		} else {
			return fmt.Errorf("no default trace style found in collection '%s' for data series without explicit style", collectionName)
		}
	}

	styleAttrs, ok := collectionMap[traceStyleName].(map[string]interface{})
	if !ok {
		return fmt.Errorf("trace style '%s' not found in collection '%s'", traceStyleName, collectionName)
	}

	// Marshal and unmarshal to deep copy and apply to JSONGrapherDataSeries fields
	jsonBytes, err := json.Marshal(styleAttrs)
	if err != nil {
		return fmt.Errorf("error marshaling trace style for '%s': %w", traceStyleName, err)
	}

	var tempSeries JSONGrapherDataSeries
	if err := json.Unmarshal(jsonBytes, &tempSeries); err != nil {
		return fmt.Errorf("error unmarshaling trace style to struct for '%s': %w", traceStyleName, err)
	}

	// Apply the unmarshaled style attributes to the current data series
	// Only update if the temporary series has a non-zero value for the field
	if tempSeries.Type != "" {
		dataSeries.Type = tempSeries.Type
	}
	if tempSeries.Mode != "" {
		dataSeries.Mode = tempSeries.Mode
	}
	if tempSeries.Line != nil {
		dataSeries.Line = tempSeries.Line
	}
	if tempSeries.Marker != nil {
		dataSeries.Marker = tempSeries.Marker
	}
	return nil
}

// ApplyTraceStylesToAllDataSeries applies the currently set trace style collection to all data series.
func (r *JSONGrapherRecord) ApplyTraceStylesToAllDataSeries() error {
	for i := range r.Data {
		err := r.ApplyTraceStyle(&r.Data[i])
		if err != nil {
			return fmt.Errorf("failed to apply style to data series at index %d: %w", i, err)
		}
	}
	return nil
}

// ApplyTraceStyleToDataSeriesAtIndex applies the currently set trace style collection to a single data series by index.
func (r *JSONGrapherRecord) ApplyTraceStyleToDataSeriesAtIndex(index int) error {
	if index < 0 || index >= len(r.Data) {
		return fmt.Errorf("index %d out of bounds for data series (0 to %d)", index, len(r.Data)-1)
	}
	return r.ApplyTraceStyle(&r.Data[index])
}
