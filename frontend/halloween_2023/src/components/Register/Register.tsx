import React, {ReactElement, useState} from "react";

const [emailAddress, setEmailAddress] = useState("")
const[username, setUsername] = useState("")


function handleEmailAdressChange(e : React.FormEvent<HTMLInputElement>){
    let newEmailAddress = e.currentTarget.value
    setEmailAddress(newEmailAddress)
}

function

export default function Register(): ReactElement{
    return(
        <div className="Login-Container">
            <input placeholder="Email Address" onChange={handleEmailAdressChange} value={emailAddress}/>
            <input placeholder="Username" onChange={}
        </div>
    )
}