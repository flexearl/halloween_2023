import React, {ReactElement, useEffect} from 'react'
import DayPreview from '../DayPreview/DayPreview'
import { Link } from 'react-router-dom'
import ReactTyped from 'react-typed'
import './Home.css'
import axios from 'axios'
import { UserContext } from '../UserContext'


export default function Home(): ReactElement{
    const {user, setUser} = React.useContext(UserContext)
    let dayNumbs = [1, 2, 3, 4]
    
    useEffect(() => {
        //Axios request
        if(user?.pumpkinsAchieved.length != 0){
        axios.post("http://localhost:8000//check_user_pumpkins", {
            username: user?.username,
        }).then(response => {
            console.log(response)
        }
            
        )    
        }

    }, [])
    
    return(
        <div className='Home-Container'>
            <ReactTyped strings={["Welcome to CodeSpook"]} typeSpeed={80} />
            {dayNumbs.map(numb => {
                return <DayPreview DayNumb = {numb}/>
            })}  
        </div>

    )
}