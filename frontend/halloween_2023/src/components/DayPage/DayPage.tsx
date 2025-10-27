import React, {ReactElement, useEffect, useState} from "react";
import { DayPagesContent } from "./DayPage-Content";
import DayPageContent from "./DayPage-Content";
import './DayPage.css'
import { useParams } from "react-router-dom";
import axios from "axios";

interface DayPageProps{
    DayNumb: number;
}





export default function DayPage():ReactElement{
    
    const {number} = useParams()

    const [puzzleContent, setPuzzleContent] = useState<any>()
    const [puzzleInput, setPuzzleInput] = useState("")

    let dayNumb = Number(number)
    useEffect(() => {
        console.log("Sending request")
        axios.get(`http://localhost:8000/get_puzzle_content/${dayNumb}`).then((res) => setPuzzleContent(res.data)).catch((err) => console.error(err))
    }, [number])

    useEffect(() => {
    axios
        .get(`http://localhost:8000/get_user_puzzle_input?daynumber=${dayNumb}&userid=${9}`)
        .then((res) => {
        setPuzzleInput(res.data.puzzle_input)
        })
        .catch((err) => {
        console.error("Error fetching puzzle input:", err)
        })
    }, [dayNumb])


    
    
    return(
        <div className="DayPage-Container">
            <h1>{"Day "+number}</h1>
            <p>{dayNumb &&  DayPagesContent[dayNumb-1].Description}</p>
            <p>{puzzleContent}</p>
            <p>{puzzleInput}</p>
            <input className="Answer-Input"/>
        </div>
    )
} 