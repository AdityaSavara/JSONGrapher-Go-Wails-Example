package styles

// TraceStylesCollectionLibrary holds various predefined collections of trace styles for graphs.
// Mirrors the styles_library from trace_styles_collection_library.py.
var TraceStylesCollectionLibrary = map[string]interface{}{ //
	"default": map[string]interface{}{ //
		"scatter_spline": map[string]interface{}{ //
			"type":   "scatter",                                               //
			"mode":   "lines+markers",                                         //
			"line":   map[string]interface{}{"shape": "spline", "width": 2.0}, //
			"marker": map[string]interface{}{"size": 10.0},                    //
		},
		"scatter_line": map[string]interface{}{ //
			"type":   "scatter",                                               //
			"mode":   "lines+markers",                                         //
			"line":   map[string]interface{}{"shape": "linear", "width": 2.0}, //
			"marker": map[string]interface{}{"size": 10.0},                    //
		},
		"line": map[string]interface{}{ //
			"type":   "scatter",                                               //
			"mode":   "lines",                                                 //
			"line":   map[string]interface{}{"shape": "linear", "width": 2.0}, //
			"marker": map[string]interface{}{"size": 10.0},                    //
		},
		"lines": map[string]interface{}{ //
			"type":   "scatter",                                               //
			"mode":   "lines",                                                 //
			"line":   map[string]interface{}{"shape": "linear", "width": 2.0}, //
			"marker": map[string]interface{}{"size": 10.0},                    //
		},
		"scatter": map[string]interface{}{ //
			"type":   "scatter",                            //
			"mode":   "markers",                            //
			"marker": map[string]interface{}{"size": 10.0}, //
		},
		"spline": map[string]interface{}{ //
			"type":   "scatter",                            //
			"mode":   "markers",                            //
			"marker": map[string]interface{}{"size": 10.0}, //
		},
		"bar": map[string]interface{}{ //
			"type":   "scatter",                            //
			"mode":   "markers",                            //
			"marker": map[string]interface{}{"size": 10.0}, //
		},
		"heatmap": map[string]interface{}{ //
			"type":       "heatmap", //
			"colorscale": "Viridis", //
		},
	},
	"scatter_spline": map[string]interface{}{ //
		"scatter_spline": map[string]interface{}{ //
			"type":   "scatter",                                               //
			"mode":   "lines+markers",                                         //
			"line":   map[string]interface{}{"shape": "spline", "width": 2.0}, //
			"marker": map[string]interface{}{"size": 0.0},                     //
		},
		"scatter": map[string]interface{}{ //
			"type":   "scatter",                                               //
			"mode":   "lines+markers",                                         //
			"line":   map[string]interface{}{"shape": "spline", "width": 2.0}, //
			"marker": map[string]interface{}{"size": 0.0},                     //
		},
		"spline": map[string]interface{}{ //
			"type":   "scatter",                                               //
			"mode":   "lines+markers",                                         //
			"line":   map[string]interface{}{"shape": "spline", "width": 2.0}, //
			"marker": map[string]interface{}{"size": 0.0},                     //
		},
	},
}
