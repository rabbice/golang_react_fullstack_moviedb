import * as React from "react"; 
import * as ReactDOM from "react-dom"; 
import {Hello} from "./components/hello"; 
 
ReactDOM.render( 
  <div className="container"> 
    <h1>MovieDB</h1> 
    <Hello name="World"/> 
  </div>, 
  document.getElementById("mymovie-app") 
);