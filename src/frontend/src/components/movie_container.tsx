import * as React from "react"; 
import {MovieList} from "./movie_list";
import {Loader} from "./loader";
import {Movie} from "../models/movie"; 
 
export interface MoviesContainerProps { 
  movieListURL: string; 
} 
 
export interface MoviesContainerState { 
  loading: boolean; 
  movies: Movie[] 
}

export class MoviesContainer extends React.Component <MoviesContainerProps, MoviesContainerState> { 
  constructor(p: MoviesContainerProps) { 
    super(p); 
 
    this.state = { 
      loading: true, 
      movies: [] 
    }; 
 
    fetch(p.movieListURL) 
      .then<Movie[]>((response) => response.json()) 
      .then(movies => { 
        this.setState({ 
          loading: false, 
          movies: movies 
        })
      }) 
  }
  render() {
    return <Loader loading={this.state.loading} message="Loading movies...">
        <MovieList movies={this.state.movies} />
    </Loader>
    }
} 