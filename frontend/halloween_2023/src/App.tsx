import { ReactElement } from 'react';
import React from 'react';
import './App.css';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import NavBar from './components/NavBar/NavBar';

function App():ReactElement {
  return (
    <div className="App">
        <BrowserRouter >
          <Routes>
            <Route path="/" element={<NavBar/>}>

            </Route>
          </Routes>
        </BrowserRouter>
    </div>
  );
}

export default App;
