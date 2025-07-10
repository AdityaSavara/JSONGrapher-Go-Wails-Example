package styles

// LayoutStylesLibrary holds various predefined layout styles for graphs.
// Mirrors the styles_library from layout_styles_library.py.
var LayoutStylesLibrary = map[string]interface{}{ //
	"default": map[string]interface{}{ //
		"layout": map[string]interface{}{ //
			"title": map[string]interface{}{"font": map[string]interface{}{"size": 32.0}, "x": 0.5},                                                                          //
			"xaxis": map[string]interface{}{"title": map[string]interface{}{"font": map[string]interface{}{"size": 27.0}}, "tickfont": map[string]interface{}{"size": 23.0}}, //
			"yaxis": map[string]interface{}{"title": map[string]interface{}{"font": map[string]interface{}{"size": 27.0}}, "tickfont": map[string]interface{}{"size": 23.0}}, //
			"legend": map[string]interface{}{ //
				"title": map[string]interface{}{"font": map[string]interface{}{"size": 22.0}}, //
				"font":  map[string]interface{}{"size": 22.0},                                 //
			},
		},
	},
	"offset2d": map[string]interface{}{ //
		"layout": map[string]interface{}{ //
			"title": map[string]interface{}{"font": map[string]interface{}{"size": 32.0}, "x": 0.5},                                                                          //
			"xaxis": map[string]interface{}{"title": map[string]interface{}{"font": map[string]interface{}{"size": 27.0}}, "tickfont": map[string]interface{}{"size": 23.0}}, //
			"yaxis": map[string]interface{}{"title": map[string]interface{}{"font": map[string]interface{}{"size": 27.0}}, "tickfont": map[string]interface{}{"size": 23.0}}, //
		},
	},
	"Nature": map[string]interface{}{ //
		"layout": map[string]interface{}{ //
			"title":         map[string]interface{}{"font": map[string]interface{}{"size": 32.0, "family": "Times New Roman", "color": "black"}}, //
			"font":          map[string]interface{}{"size": 25.0, "family": "Times New Roman"},                                                   //
			"paper_bgcolor": "white",                                                                                                             //
			"plot_bgcolor":  "white",                                                                                                             //
			"xaxis": map[string]interface{}{ //
				"showgrid":  true,      //
				"gridcolor": "#ddd",    //
				"gridwidth": 1.0,       //
				"linecolor": "black",   //
				"linewidth": 2.0,       //
				"ticks":     "outside", //
				"tickwidth": 2.0,       //
				"tickcolor": "black",   //
			},
			"yaxis": map[string]interface{}{ //
				"showgrid":  true,      //
				"gridcolor": "#ddd",    //
				"gridwidth": 1.0,       //
				"linecolor": "black",   //
				"linewidth": 2.0,       //
				"ticks":     "outside", //
				"tickwidth": 2.0,       //
				"tickcolor": "black",   //
			},
		},
	},
	"Science": map[string]interface{}{ //
		"layout": map[string]interface{}{ //
			"title":         map[string]interface{}{"font": map[string]interface{}{"size": 32.0, "family": "Arial", "color": "black"}}, //
			"font":          map[string]interface{}{"size": 25.0, "family": "Arial"},                                                   //
			"paper_bgcolor": "white",                                                                                                   //
			"plot_bgcolor":  "white",                                                                                                   //
			"xaxis": map[string]interface{}{ //
				"showgrid":  true,      //
				"gridcolor": "#ccc",    //
				"gridwidth": 1.0,       //
				"linecolor": "black",   //
				"linewidth": 2.0,       //
				"ticks":     "outside", //
				"tickwidth": 2.0,       //
				"tickcolor": "black",   //
			},
			"yaxis": map[string]interface{}{ //
				"showgrid":  true,      //
				"gridcolor": "#ccc",    //
				"gridwidth": 1.0,       //
				"linecolor": "black",   //
				"linewidth": 2.0,       //
				"ticks":     "outside", //
				"tickwidth": 2.0,       //
				"tickcolor": "black",   //
			},
		},
	},
}
