import * as React from "react"; 
import * as ReactDOM from "react-dom"; 
import {HashRouter as Router, Route} from "react-router-dom";
import {MoviesContainer} from "./components/movie_container";
import { Navigation } from "./components/navigation";
import 'bootstrap/dist/css/bootstrap.min.css';

class App extends React.Component<{}, {}> { 
  render() {
    const movieList = () => <MoviesContainer movieListURL="http://localhost:8000/v1/movies"/>
    return <Router>
      <div className="container">
        <Navigation brandName="Movies"/>
        <h1>The MovieDB</h1> 
        <Route exact path="/"component={movieList}/>
      </div> 
    </Router>
  } 
} 
 
ReactDOM.render( 
  <App/>, document.getElementById('root')); 