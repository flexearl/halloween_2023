import React, {ReactElement, useEffect, useState} from "react";
import { DayPagesContent } from "./DayPage-Content";
import DayPageContent from "./DayPage-Content";
import './DayPage.css'
import { useParams } from "react-router-dom";

interface DayPageProps{
    DayNumb: number;
}





export default function DayPage():ReactElement{

    
    const {number} = useParams()
    let dayNumb = Number(number)
    
    
    return(
        <div className="DayPage-Container">
            <h1>{"Day "+number}</h1>
            <p>{dayNumb&&  DayPagesContent[dayNumb-1].Description}</p>
            <input className="Answer-Input"/>
        </div>
    )
} 