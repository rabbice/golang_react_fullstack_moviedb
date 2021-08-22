import {Movie} from "../models/movie"; 
import {MovieItem} from "./movie_item"; 
import * as React from "react"; 
 
export interface MovieListProps { 
  movies: Movie[]; 
} 
 
export class MovieList extends React.Component<MovieListProps, {}> {
  render() { 
    const items = this.props.movies.map(m => <MovieItem movie={m} /> ); 
 
    return <table className="table"> 
        <thead> 
            <tr> 
                <th>Name</th>
                <th>Year</th>
                <th>Overview</th> 
                <th>Directors</th> 
                <th>Budget</th> 
                <th>Gross</th> 
            </tr> 
        </thead> 
        <tbody> 
            {items} 
        </tbody> 
    </table> 
  }   
} 