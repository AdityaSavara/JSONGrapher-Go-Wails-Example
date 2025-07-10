# README

## About

This is a simple JSONGrapher and Wails example made by modifying the official Wails Vanilla template.
This is not intended to be a template for building JSONGrapher with Wails projects,
it is intended to be an example for people to see how to plot graphs and clear them, with JSONGrapher.

To see how the example works as a demonstration, download, go to the main directory (with app.go)
and then type in 'wails dev' and press enter. Then click on "change graph".
This demonstrates showing a graph from JSONGrapher using a Wails GUI.
You may also click to create the lower graph.

The basic algorithmic concepts are as follows:
(a) Wails can provide a front-end GUI for Golang ("Go") projects, and is made with javascript.
(b) JSONGrapher has a javascript version, so has a function that can be imported to plot graphs in Wails.
(c) JSON objects can be passed from Go to Wails that are JSONGrapher compatible, and then plotted.
(d) Typically, the GUI will call Go functions (by a user clicking a button etc.) to request data to display.

Accordingly, this small demonstration example has the following features.
 - The Go code passes json objects to Wails by functions in ./app.go (GetJSONFromFile or CreateExampleJSON)
 - The Wails code has javascript calls to the Go functions from ./frontend/src/main.js
 - The json files are located in ./json_records/ (note that location is relative to app.go even in the javascript code)
 - The main.js file imports functions from JSONGrapher, and has a button to change which graph is shown.
 - The graph is located in a div, and clicking the appropriate button uses the clearData function.

A developer would implement all of the above things in their app, mimicking what is in this code.

Additional details to know for development:
- It was needed to copy the JSONGrapher source code into ./frontend/JSONGrapher 
- Only two JSONGrapher javascript directores, 'src' and 'styles' are needed
- the javascript directores are originally from https://github.com/AdityaSavara/JSONGrapher
- The jsongrapher-go directory enables the Go based JSONGrapher record creation,
- and the go.mod file has been filled in to use the local version of jsongrapher-go
-  (however, a person can retrieve the most up to date version of jsongrapher-go if desired).
- It is not necessary, or even recommended, to use GetJSONFromFile, under normal circumstances.
- The .json files were for demonstration purposes, here. Normally, your code would generate and send JSON
- directly to Wails upon request, in a JSONGrapher compatible format (more similar to the CreateExampleJSON function).
- In creating this example from the Wails Vanilla template, changes were made to the css and other files,
- in part because JSONGrapher is easiest to use on a page (or frame) with a white background.
- The JSONGrapher directories include third party open source software, as noted in LICENSE.txt and NOTICE.txt
- Use of JSONGrapher in your Wails app, like this, will require including the licenses with your source code.


## Project Config

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: https://wails.io/docs/reference/project-config

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
