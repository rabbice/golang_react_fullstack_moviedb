import * as React from "react"; 
import * as ReactDOM from "react-dom"; 
import {MoviesContainer} from "./components/movie_container";

class App extends React.Component { 
  render() { 
    return <div className="container"> 
      <h1>The MovieDB</h1> 
      <MoviesContainer movieListURL="http://localhost:8000/movies"/> 
    </div> 
  } 
} 
 
ReactDOM.render( 
  <App/>, document.getElementById('root')); 