import * as React from "react"; 
import { NavLink } from "react-router-dom"; 
 
export interface NavigationProps { 
  brandName: string; 
} 
 
export class Navigation extends React.Component<NavigationProps, {}> {
    render() {
        return <nav className="navbar navbar-default">
            <div className="container">
                <div className="navbar-header">
                    <NavLink to="/" className="navbar-brand">
                        {this.props.brandName}
                    </NavLink>
                </div>
                <ul className="nav navbar-nav">
                    <li><NavLink to="/">Movies</NavLink></li>
                </ul>
            </div>
        </nav>
    }   
}

