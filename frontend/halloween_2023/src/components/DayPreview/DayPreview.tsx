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
        <>
            <Link to ={"/day/"+props.DayNumb} className='DayPreview-Link'>
               
                {"Day "+props.DayNumb}
                <img src ={pumpkin} className='Pumpkin-Picture'/>
                
            </Link>
        </>
       
    )
}