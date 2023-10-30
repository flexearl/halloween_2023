import React, {ReactElement, useEffect} from 'react'
import './DayPreview.css'
import { Link, Route, Routes } from 'react-router-dom';
import DayPage from '../DayPage/DayPage';
import { ReactComponent as Pumpkin} from "../pumpkin-pixel.svg"
import pumpkin from "../pumpkin-pixel.png"

interface DayPreviewProps {
    DayNumb: number;
    Pumpkin:boolean
}

export default function DayPreview(props: DayPreviewProps): ReactElement {

    return(
        <div className="DayPreview-Container">
            <Link to ={"/day/"+props.DayNumb} className='DayPreview-Link'>{"Day "+props.DayNumb}
            
            </Link>
            <img src ={pumpkin}className='Pumpkin-Picture'/>
        </div>

       
    )
}