import {Movie} from "../models/movie"; 
import * as React from "react"; 
 
export interface MovieItemProps { 
  movie: Movie; 
} 
 
export class MovieItem extends React.Component<MovieItemProps, {}> { 
  render() { 
 
  return (
    <tr> 
      <td>{this.props.movie.Title}</td>  
      <td>{this.props.movie.Year}</td>
      <td>{this.props.movie.Overview}</td>
      <td>{this.props.movie.Directors}</td>
      <td>{this.props.movie.Budget}</td>
      <td>{this.props.movie.Gross}</td>
    </tr> 
  )
  } 
}