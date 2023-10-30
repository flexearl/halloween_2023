import { ReactElement, useContext, useState } from 'react';
import React from 'react';
import './App.css';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import NavBar from './components/NavBar/NavBar';
import Home from './components/Home/Home';
import Login from './components/Login/Login';
import Register from './components/Register/Register';
import { TypeWriter } from './components/Typewriter';
import ReactTyped from "react-typed";
import DayPage from './components/DayPage/DayPage';
import { createContext } from 'vm';
import { UserContext, UserContextProps, User } from './components/UserContext';

function App():ReactElement {
  const defaultUser : User = {emailaddress:"", username:"", pumpkinsAchieved:[]}
  const [user, setUser] = useState<User >(defaultUser)
  const value = {
    user,
    setUser
  }
  
  return (
    <div className="App">
      <UserContext.Provider value={value}>
        <BrowserRouter>
        <NavBar />
          <Routes>
            
            <Route index element={<Home/>}/>
            <Route path={'/day/:number'} element={<DayPage/>}/>
            <Route path="/login" element={<Login/>}/>
            
          </Routes>

        </BrowserRouter>
      
      </UserContext.Provider>
    </div>
  );
}

export default App;
