import { useState } from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import { Greet, WailsInit } from "../wailsjs/go/main/App";
import { CircularProgressbar, CircularProgressbarWithChildren } from "react-circular-progressbar";

function App() {
    const [percentage, setPercentage] = useState(0);

    function greet() {
        // Greet(name).then(updateResultText);
        setPercentage(50)
    }

    return (
        <div id="App">
            <div style={{ width: 200, height: 200 }}>
                <CircularProgressbar
                    value={percentage}
                    text={`${percentage}%`}
                />

            </div>
            <div style={{ width: 200, height: 200 }}>
                <CircularProgressbarWithChildren value={percentage}>
                    {/* Put any JSX content in here that you'd like. It'll be vertically and horizonally centered. */}
                    <img style={{ width: 60 }} src="https://i.imgur.com/b9NyUGm.png" alt="doge" />
                    <div style={{fontSize: 50}}>
                        <strong>{percentage}%</strong>
                    </div>
                </CircularProgressbarWithChildren>;

            </div>
            <button className="btn" onClick={greet}>dfafafda</button>
        </div>
    )
}

export default App
