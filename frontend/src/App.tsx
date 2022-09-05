import React from 'react'
import './App.css'
import {BrowserRouter, Route, Routes} from 'react-router-dom'
import {HelloWorld} from "./components/HelloWorld";

export const App: React.FC = () => <BrowserRouter>
    <Routes>
        <Route path="/" element={<HelloWorld/>}/>
    </Routes>
</BrowserRouter>

export default App