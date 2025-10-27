import React, {ReactElement, useState} from "react";
import axios from "axios";



interface NewUser {
    email: string,
    username: string
    password: string,

}

export default function Register(): ReactElement{
    const [emailAddress, setEmailAddress] = useState("")
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")

    function handleEmailAdressChange(e : React.FormEvent<HTMLInputElement>){
        let newEmailAddress = e.currentTarget.value
        setEmailAddress(newEmailAddress)
    }

    function handleUsernameChange(e: React.FormEvent<HTMLInputElement>){
        let newUsername = e.currentTarget.value
        setUsername(newUsername)
    }

    function handlePasswordChange(e: React.FormEvent<HTMLInputElement>){
        let newPassword = e.currentTarget.value 
        setPassword(newPassword)
    }

    function handleRegisterClick(e: React.MouseEvent<HTMLElement>){

        const newUser: NewUser = {
            email: emailAddress,
            username: username,
            password: password
        }

        const response = axios.post("http://localhost:8000/register_user", newUser)
        console.log(response)
    }

    return(
        <div className="Login-Container">
            <input placeholder="Email Address" onChange={handleEmailAdressChange} value={emailAddress}/>
            <input placeholder="Username" onChange={handleUsernameChange} value = {username}/>
            <input placeholder="Password" onChange={handlePasswordChange} value ={password}/>
            <button onClick={handleRegisterClick} >Register</button>
        </div>
    )
}