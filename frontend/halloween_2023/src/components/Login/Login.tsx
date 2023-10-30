import React, {ReactElement, useState, useContext} from 'react'
import axios from 'axios'
import { Link } from 'react-router-dom'
import './Login.css'
import { UserContext, User } from '../UserContext'



export default function Login(): ReactElement{
    const {user, setUser} = React.useContext(UserContext)
    
    const [emailAddress, setEmailAddress] = useState("")
    const [password, setPassword] = useState("")

    function handlePasswordChange(e: React.FormEvent<HTMLInputElement>){
        let newPassword = e.currentTarget.value
        setPassword(newPassword)
    }

    function handleEmailAdressChange(e: React.FormEvent<HTMLInputElement>){
        let newEmailAddress = e.currentTarget.value
        setEmailAddress(newEmailAddress)
    }

    function handleLoginClick(e: React.MouseEvent<HTMLElement>){
        e.preventDefault()
        setUser((prevState:User) => {
            prevState.username = "valid"
            return {
                ...prevState,
            }
            
        })
        /*
        axios.post("http://localhost:8000/login", {
            emailaddress: emailAddress,
            password: password
        }).then(response => {
            console.log(response)
        }
            
        )
        */
    }

    return(
        <div className='flex-container'>
            <div className='content-container'>
                <div className="Login-Container">
                    <input placeholder="Email Address" value={emailAddress} onChange={handleEmailAdressChange}/>
                    <input placeholder="Password" value={password} onChange={handlePasswordChange}/>
                    <button onClick={handleLoginClick} className='submit-btn'>Login</button>
                    <Link to ="/register"/>
                </div>
            </div>
        </div>
        
    )
}