import { useState } from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import { Greet, WailsInit } from "../wailsjs/go/main/App";

function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [name, setName] = useState('');
    const updateName = (e) => setName(e.target.value);
    const updateResultText = (result) => setResultText(result);

    runtime.EventsOn("cpu_usage", (msg) => updateResultText(msg))

    function greet() {
        // Greet(name).then(updateResultText);
        runtime.EventsOn("cpu_usage", (msg) => updateResultText(msg))
    }

    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo" />
            <div id="result" className="result">{resultText}</div>
            <div id="input" className="input-box">
                <input id="name" className="input" onChange={updateName} autoComplete="off" name="input" type="text" />
                <button className="btn" onClick={greet}>Greet</button>
            </div>
        </div>
    )
}

export default App
