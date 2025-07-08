import './style.css';
import './app.css';

import {GetJSONFromFile} from '../wailsjs/go/main/App';
import { mergeAndplotData, clearData } from './JSONGrapher/src/index.js';

document.querySelector('#app').innerHTML = ``;

//////Start of JSONGrapher graphDivGraph2 Injection
//First read in a file... we are intentionaly using Go do this since the real application
//would be sending a json from go.
let providedJson1 = await GetJSONFromFile("./json_records/SrTiO3_rainbow.json")
let providedJson2 = await GetJSONFromFile("./json_records/SrTiO3_rainbow.json")
//console.log("The providedJson received by main.js", providedJson1)

// Create container elements
const graphContainer = document.createElement('div');
graphContainer.innerHTML = `
  <center>
    <h2>This is a simple JSONGrapher Wails App to Demonstrate Showing Graphs from JSON</h2>
    <div id="graphDivGraph2"></div>
    <div id="errorMessagesDivGraph2"></div>
    Click on 'Change Graph' to switch between graphs. <button class="btn" onclick="changeGraph()">Change Graph</button> 
  </center>
`;
//I did consider using this as well:   <button class="btn" onclick="clearData(graphDivGraph2)">Clear Graph</button>

// Append to desired location in the DOM
document.querySelector('#app').appendChild(graphContainer);

//call mergeAndplotData. requires getting some arguments ready, first, as these are the args:
//    mergeAndplotData(existingFigDict, newFigDict, newFigDictFileName, graphDivName, messagesToUserDiv, errorDiv)
const errorMessagesDivGraph2 = document.getElementById("errorMessagesDivGraph2");
const graphDivGraph2 = document.getElementById("graphDivGraph2");
const messagesToUserDiv = "This is a message for the user upon loading.";

//let's plot a graph. The path used will be relative to main.go. The string below is the graph filename and does not load it again.
let mergedFigDict = await mergeAndplotData(null, providedJson1, "SrTiO3_rainbow.json", graphDivGraph2, messagesToUserDiv, errorMessagesDivGraph2);
let currentGraph = "SrTiO3_rainbow.json";

//We'll make a little function to switch between graphs, then populate the window with it.
//This function will utilize this module's 'currentGraph' variable to know which graph is currently up.
async function changeGraph(){
  if (currentGraph === "SrTiO3_rainbow.json"){
    clearData(graphDivGraph2, document)
    //change to the other graph. Though we don't need to read file in again... we will, since we are 
    // trying to create behavior of when a new json will come from Go, and this does that.
    providedJson2 = await GetJSONFromFile("./json_records/DRIFTS_CO_Adsorption_onAu22_offset2d.json")
    mergedFigDict = await mergeAndplotData(null, providedJson2, "DRIFTS_CO_Adsorption_onAu22_offset2d.json", graphDivGraph2, messagesToUserDiv, errorMessagesDivGraph2);
    currentGraph = "DRIFTS_CO_Adsorption_onAu22_offset2d.json"     //update the currentGraph variable, which is just a name.
  } else if (currentGraph === "DRIFTS_CO_Adsorption_onAu22_offset2d.json"){
    //change to the other graph. Though we don't need to read file in again... we will, since we are 
    // trying to create behavior of when a new json will come from Go, and this does that.
    providedJson1 = await GetJSONFromFile("./json_records/SrTiO3_rainbow.json")
    mergedFigDict = await mergeAndplotData(null, providedJson1, "SrTiO3_rainbow.json", graphDivGraph2, messagesToUserDiv, errorMessagesDivGraph2);
    currentGraph = "SrTiO3_rainbow.json"     //update the currentGraph variable
  }

}

//Add the new function to the window so that the changeGraph button can access it.
window.changeGraph = changeGraph;



//////End of JSONGrapher graphDivGraph2 Injection

