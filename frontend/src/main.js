import './style.css';
import './app.css';

import {GetJSONFromFile, CreateExampleJSON} from '../wailsjs/go/main/App';
import { mergeAndplotData, clearData } from './JSONGrapher/src/index.js';

document.querySelector('#app').innerHTML = ``;

//////Start of block of code for upper JSONGrapher example which changes graphs with button click. 

//This graph will be named 'graph2' and will be in DIV graphDivGraph2

//First read in a file, using Go to demonstrate receiving the json from GO.
let providedJson1 = await GetJSONFromFile("./json_records/SrTiO3_rainbow.json")
let providedJson2 = await GetJSONFromFile("./json_records/SrTiO3_rainbow.json")
//console.log("The providedJson received by main.js", providedJson1)

// Create div element for the graphs. This isn't necessary but will be convenient.
const allGraphsOuterDiv = document.createElement('div');
allGraphsOuterDiv.innerHTML = `
  <center>
    <h2>This is a simple JSONGrapher Wails App to Demonstrate Showing Graphs from JSON</h2>
    <div id="graphDivGraph2"></div>
    <div id="errorMessagesDivGraph2"></div>
    Click on 'Change Graph' to switch between graphs. <button class="btn" onclick="changeGraph()">Change Graph</button> 
  </center>
`;
//I did consider using this as well:   <button class="btn" onclick="clearData(graphDivGraph2)">Clear Graph</button>

// Append to desired location in the DOM
document.querySelector('#app').appendChild(allGraphsOuterDiv);

//call mergeAndplotData. requires getting some arguments ready, first, as these are the args:
//    mergeAndplotData(existingFigDict, newFigDict, newFigDictFileName, graphDivName, messagesToUserDiv, errorDiv)
const errorMessagesDivGraph2 = document.getElementById("errorMessagesDivGraph2");
const graphDivGraph2 = document.getElementById("graphDivGraph2");
const messagesToUserDiv = "This is a message for the user upon loading.";

//let's plot a graph. The path used will be relative to main.go. The string below is the graph filename and does not load it again.
let mergedFigDict = await mergeAndplotData(null, providedJson1, "SrTiO3_rainbow.json", graphDivGraph2, messagesToUserDiv, errorMessagesDivGraph2);
let currentGraphDataSource = "SrTiO3_rainbow.json";

//We'll make a little function to switch between graphs, then populate the window with it.
//This function will utilize this module's 'currentGraphDataSource' variable to know which graph is currently up.
async function changeGraph(){
	//This function checks which of two prechosen graphs is loaded, then clears and loads the other graph.
  if (currentGraphDataSource === "SrTiO3_rainbow.json"){
    //change to the other graph, start with clearing.
	//clearing a graph just requires passing in its Div and optionally the document it is in.
	clearData(graphDivGraph2, document)
    //now load the other graph. Though we don't need to read file in again... we will, since we are 
    // trying to create behavior of when a new json will come from Go, and this does that.
    providedJson2 = await GetJSONFromFile("./json_records/DRIFTS_CO_Adsorption_onAu22_offset2d.json")
    mergedFigDict = await mergeAndplotData(null, providedJson2, "DRIFTS_CO_Adsorption_onAu22_offset2d.json", graphDivGraph2, messagesToUserDiv, errorMessagesDivGraph2);
    currentGraphDataSource = "DRIFTS_CO_Adsorption_onAu22_offset2d.json"     //update the currentGraphDataSource variable, which is just a name.
  } else if (currentGraphDataSource === "DRIFTS_CO_Adsorption_onAu22_offset2d.json"){
    //change to the other graph, start with clearing.
    //clearing a graph just requires passing in its Div and optionally the document it is in.
	clearData(graphDivGraph2, document)
	//now load the other graph. Though we don't need to read file in again... we will, since we are 
    // trying to create behavior of when a new json will come from Go, and this does that.
    providedJson1 = await GetJSONFromFile("./json_records/SrTiO3_rainbow.json")
    mergedFigDict = await mergeAndplotData(null, providedJson1, "SrTiO3_rainbow.json", graphDivGraph2, messagesToUserDiv, errorMessagesDivGraph2);
    currentGraphDataSource = "SrTiO3_rainbow.json"     //update the currentGraphDataSource variable
  }

}

//Add the new function to the window so that the changeGraph button can access it.
window.changeGraph = changeGraph;


//////Start of block of code for upper JSONGrapher example which changes graphs with button click. 


//////Start of block of code for lower JSONGrapher example 
// which involves populating the individual JSONGrapher fields

//We have an existing outerDiv which we will add to.
allGraphsOuterDiv.innerHTML += `
  <center>
    <h2>This is a simple JSONGrapher Wails App to Demonstrate Showing Graphs from JSON</h2>
    <div id="graphDivGraph3"></div>
    <div id="errorMessagesDivGraph3"></div>
    Click on 'Create Lower Graph' to create an additional graph, and 'Clear Lower Graph' to remove it. 
	<button class="btn" onclick="createLowerGraph()">Create Lower Graph</button> 
	<button class="btn" onclick="clearLowerGraph()">Clear Lower Graph</button> 
  </center>
`;


//This time, instead of making variables directly in main.js, we'll make them inside the createLowerGraph function.
async function createLowerGraph(){
	//We will retrieve the JSONGrapher record for plotting.
	const graph3JSON = await CreateExampleJSON()
	//to plot, we'll call mergeAndplotData. requires getting some arguments ready, first, as these are the args:
	//    mergeAndplotData(existingFigDict, newFigDict, newFigDictFileName, graphDivName, messagesToUserDiv, errorDiv)
	const errorMessagesDivGraph3 = document.getElementById("errorMessagesDivGraph3");
	const graphDivGraph3 = document.getElementById("graphDivGraph3");
	const messagesToUserDiv = "Click to create the lower graph or to clear it. Clicking to create multiple times will load the data multiple times.";

	//let's plot a graph. We'll provide null for the filename.
	let mergedFigDict = await mergeAndplotData(null, graph3JSON, null, graphDivGraph3, messagesToUserDiv, errorMessagesDivGraph3);	
}

async function clearLowerGraph(){
	//clearing a graph just requires passing in its Div and optionally the document it is in.
	clearData(graphDivGraph3, document);
}


//Add the new functions to the window so that the changeGraph button can access it.
window.createLowerGraph = createLowerGraph;
window.clearLowerGraph = clearLowerGraph;

