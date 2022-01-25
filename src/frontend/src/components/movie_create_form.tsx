import * as React from "react"; 
import {Movie} from "../models/movie"; 
import {FormRow} from "./form"; 
 
export interface MovieCreatingFormProps { 
  movie: Movie; 
  onSubmit: (seats: number) => any 
} 
 
export interface MovieCreatingFormState { 
  seats: number; 
} 
 
export class MovieCreatingForm
  extends React.Component<MovieCreatingFormProps, MovieCreatingFormState> { 
  constructor(p:  MovieCreatingFormProps) { 
    super(p); 
 
    this.state = {seats: 1}; 
  } 
} 