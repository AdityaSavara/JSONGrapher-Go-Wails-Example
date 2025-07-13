import './style.css';
import './app.css';

import {GetJSONFromFile, CreateExampleJSON} from '../wailsjs/go/main/App';
import { mergeAndplotData, clearData } from './JSONGrapher/src/index.js';
import { createCSV } from './JSONGrapher/src/fileUtils.js';
import {addOpeningURLonButtonClick} from './JSONGrapher/src/linkUtils.js';
document.querySelector('#app').innerHTML = ``;


//////Start of block of code for upper JSONGrapher example which changes graphs with button click. 

//This graph will be named 'graph2' and will be in DIV graphDivGraph2

//First read in a file, using Go to demonstrate receiving the json from GO.
let providedJson1 = await GetJSONFromFile("./json_records/SrTiO3_rainbow.json")
let providedJson2 = await GetJSONFromFile("./json_records/SrTiO3_rainbow.json")

// Create a div element for the graph.
// An separate outer div is required for each graph.
const upperGraphOuterDiv = document.createElement('div');
upperGraphOuterDiv.innerHTML = `
  <center>
    <h2><font color='black'>This is a simple JSONGrapher Wails App to Demonstrate Showing Graphs from JSON</font></h2>
    <div id="graphDivGraph2"></div>
    <div id="errorMessagesDivGraph2"></div>
  Click on 'Change Graph' to switch between graphs.<br>
	<button class="btn" onclick="changeGraph()">Change Graph</button> 
	<button class="btn" onclick="clearUpperGraph()">Clear Upper Graph</button> <br>
  <button class="btn" id="downloadUpperGraphJSON-btn" onclick="downloadUpperGraphJSON()">Download JSON</button> 
  <button class="btn" id="downloadUpperGraphCSV-btn" onclick="downloadUpperGraphCSV()">Download CSV</button>   
  </center>
`;
//I did consider using this as well:   <button class="btn" onclick="clearData(graphDivGraph2)">Clear Graph</button>

// Append to desired location in the DOM
document.querySelector('#app').appendChild(upperGraphOuterDiv);

//call mergeAndplotData. requires getting some arguments ready, first, as these are the args:
//    mergeAndplotData(existingFigDict, newFigDict, newFigDictFileName, graphDivName, messagesToUserDiv, errorDiv)
const errorMessagesDivGraph2 = document.getElementById("errorMessagesDivGraph2");
const graphDivGraph2 = document.getElementById("graphDivGraph2");
const messagesToUserDiv = "This is a message for the user upon loading.";

//let's plot a graph. The path used will be relative to main.go. The string below is the graph filename and does not load it again.
let mergedFigDict = await mergeAndplotData(null, providedJson1, "SrTiO3_rainbow.json", graphDivGraph2, messagesToUserDiv, errorMessagesDivGraph2);
let currentUpperGraphDataSource = "SrTiO3_rainbow.json";
let currentUpperGraphJson = providedJson1 //This is the initial json.

//We'll make a little function to switch between graphs, then populate the window with it.
//This function will utilize this module's 'currentGraphDataSource' variable to know which graph is currently up.
//NOTE: for the download links, we are using variables that are global across this file.
async function changeGraph(){
	//This function checks which of two prechosen graphs is loaded, then clears and loads the other graph.
  if (currentUpperGraphDataSource === "SrTiO3_rainbow.json"){
    //change to the other graph, start with clearing.
    //clearing a graph just requires passing in its Div and optionally the document it is in.
    await clearData(graphDivGraph2, document);
    //now load the other graph. Though we don't need to read file in again... we will, since we are 
    // trying to create behavior of when a new json will come from Go, and this does that.
    providedJson2 = await GetJSONFromFile("./json_records/DRIFTS_CO_Adsorption_onAu22_offset2d.json");
    currentUpperGraphJson = providedJson2
    mergedFigDict = await mergeAndplotData(null, currentUpperGraphJson, "DRIFTS_CO_Adsorption_onAu22_offset2d.json", graphDivGraph2, messagesToUserDiv, errorMessagesDivGraph2);
    currentUpperGraphDataSource = "DRIFTS_CO_Adsorption_onAu22_offset2d.json";     //update the currentGraphDataSource variable, which is just a name.
  } else if (currentUpperGraphDataSource === "DRIFTS_CO_Adsorption_onAu22_offset2d.json"){
    //change to the other graph, start with clearing.
    //clearing a graph just requires passing in its Div and optionally the document it is in.
	  await clearData(graphDivGraph2, document);
	  //now load the other graph. Though we don't need to read file in again... we will, since we are 
    // trying to create behavior of when a new json will come from Go, and this does that.
    providedJson1 = await GetJSONFromFile("./json_records/SrTiO3_rainbow.json");
    currentUpperGraphJson = providedJson1
    mergedFigDict = await mergeAndplotData(null, currentUpperGraphJson, "SrTiO3_rainbow.json", graphDivGraph2, messagesToUserDiv, errorMessagesDivGraph2);
    currentUpperGraphDataSource = "SrTiO3_rainbow.json";     //update the currentGraphDataSource variable
  }

}

async function clearUpperGraph(){
	//clearing a graph just requires passing in its Div and optionally the document it is in.
	await clearData(graphDivGraph2, document);
};


async function downloadUpperGraphJSON(){
  let downloadButton = document.getElementById("downloadUpperGraphJSON-btn");    
  const filename ="upperGraph.json"
  const downloadLink = createDownloadJSONLink(currentUpperGraphJson, filename);
  window.open(downloadLink, "_blank");
  //downloadButton = addOpeningURLonButtonClick (downloadButton, downloadLink)
}

async function downloadUpperGraphCSV(){
  let downloadButton = document.getElementById("downloadUpperGraphCSV-btn");
  //The csv String is in csvContent.csv
  const csvContent = createCSV(currentUpperGraphJson);
  const filename = "upperGraph.csv"
  const downloadLink = createDownloadCSVLink(csvContent.csv, filename);  
  window.open(downloadLink, "_blank");
  //downloadButton = addOpeningURLonButtonClick (downloadButton, downloadLink)
}


//Add the new functions to the window so that the buttons can access them.
window.changeGraph = changeGraph;
window.clearUpperGraph = clearUpperGraph;
window.downloadUpperGraphCSV = downloadUpperGraphCSV;
window.downloadUpperGraphJSON = downloadUpperGraphJSON;


//////Start of block of code for upper JSONGrapher example which changes graphs with button click. 


//////Start of block of code for lower JSONGrapher example 
// which involves populating the individual JSONGrapher fields

// Create a div element for the graph.
// An separate outer div is required for each graph.
const lowerGraphOuterDiv = document.createElement('div');
lowerGraphOuterDiv.innerHTML = `
  <center>
    <h2><font color='black'>Click on the various buttons to proceed with the demonstration.<br> You can also zoom by mouse and hover over individual points.</font></h2>
    <div id="graphDivGraph3"></div>
    <div id="errorMessagesDivGraph3"></div>
  Click on 'Create Lower Graph' to create an additional graph, and 'Clear Lower Graph' to remove it. <br>
	<button class="btn" onclick="createLowerGraph()">Create Lower Graph</button> 
	<button class="btn" onclick="clearLowerGraph()">Clear Lower Graph</button> <br>
  <button class="btn" id="downloadLowerGraphJSON-btn" onclick="downloadLowerGraphJSON()">Download JSON</button> 
  <button class="btn" id="downloadLowerGraphCSV-btn" onclick="downloadLowerGraphCSV()">Download CSV</button>   
  </center>
`;


// Append to desired location in the DOM
document.querySelector('#app').appendChild(lowerGraphOuterDiv);
let currentLowerGraphJson

//This time, instead of making variables directly in main.js, we'll make them inside the createLowerGraph function.
async function createLowerGraph(){
	//We will retrieve the JSONGrapher record for plotting.
	const graph3JSON = await CreateExampleJSON()
	//to plot, we'll call mergeAndplotData. requires getting some arguments ready, first, as these are the args:
	//    mergeAndplotData(existingFigDict, newFigDict, newFigDictFileName, graphDivName, messagesToUserDiv, errorDiv)
	const errorMessagesDivGraph3 = document.getElementById("errorMessagesDivGraph3");
	const graphDivGraph3 = document.getElementById("graphDivGraph3");
	const messagesToUserDiv = "Click to create the lower graph or to clear it. Clicking to create multiple times will load the data multiple times.";
  currentLowerGraphJson = graph3JSON
	//let's plot a graph. We'll provide null for the filename.
	let mergedFigDict = await mergeAndplotData(null, currentLowerGraphJson, null, graphDivGraph3, messagesToUserDiv, errorMessagesDivGraph3);	
};

async function clearLowerGraph(){
	//clearing a graph just requires passing in its Div and optionally the document it is in.
	clearData(graphDivGraph3, document);
};


//Add the new functions to the window so that the changeGraph button can access it.
window.createLowerGraph = createLowerGraph;
window.clearLowerGraph = clearLowerGraph;

//////End of block of code for lower JSONGrapher example 