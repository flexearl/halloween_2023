import {ReactElement, useContext} from "react";
import React from "react";
import { Link } from "react-router-dom";
import './NavBar.css'
import { UserContext } from "../UserContext";

export default function NavBar():ReactElement{
    const {user, setUser} = useContext(UserContext)
    return(
        <div className="NavBar-Container">
            <Link to="/">Home</Link>
            <Link to="/login">{user?.username? user.username: "Account"}</Link>
        </div>
    )
}