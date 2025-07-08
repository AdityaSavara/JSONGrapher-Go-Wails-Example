# README

## About

This is a simple JSONGrapher and Wails example made by modifying the official Wails Vanilla template.
This is not intended to be a template for building JSONGrapher with Wails projects,
it is intended to be an example for people to see how to plot graphs and clear them, with JSONGrapher.

To see how the example works as a demonstration, download, go to the main directory (with app.go)
and then type in 'wails dev' and press enter. Then click on "change graph".
This demonstrates showing a graph from JSONGrapher using a Wails GUI.

The basic idea is as follows:
(a) Wails can provide a front-end GUI for Golang ("Go") projects, and is made with javascript.
(b) JSONGrapher has a javascript version, so has a function that can be imported to plot graphs in Wails.
(c) JSON objects can be passed from Go to Wails that are JSONGrapher compatible, and then plotted.
(d) Typically, the GUI will call Go functions (by a user clicking a button etc.) to request data to display.

Accordingly, this small demonstration example has the following features.
 - The Go code passes json objects to Wails, by the function GetJSONFromFile in ./app.go
 - The Wails code calls this Go function from ./frontend/src/main.js
 - The json files are located in ./json_records/ (note that location is relative to app.go even in the javascript code)
 - The main.js file imports functions from JSONGrapher, and has a button to change which graph is shown.
 - The graph is located in a div, and clicking the button uses the clearData function then plots the other file.

Additional details to know for development:
- It was needed to copy the JSONGrapher source code into ./frontend/JSONGrapher 
- Only two JSONGrapher directores, 'src' and 'styles' needed to be pasted in, from https://github.com/AdityaSavara/JSONGrapher
- It is not necessary, or even recommended, for you Go code to require a .json, normally. 
- The .json files were for demonstration purposes, here. Normally, your code would generate and send JSON
- directly to Wails upon request, in a JSONGrapher compatible format (rather than the read and send, here).
- In creating this example from the Wails Vanilla template, changes were made to the css and other files,
- in part because JSONGrapher is easiest to use on a page (or frame) with a white background.
- The JSONGrapher directories include third party open source software, as noted in LICENSE.txt and NOTICE.txt
- Use of JSONGrapher in your Wails app, like this, will require including these licenses with your source code.- 

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
