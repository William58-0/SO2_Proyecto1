import { useState } from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import { Table, Card, Button } from 'react-bootstrap';
import { Greet, WailsInit } from "../wailsjs/go/main/App";
import { CircularProgressbar, CircularProgressbarWithChildren, buildStyles } from "react-circular-progressbar";

/*
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
                    {}
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
*/
function App() {
    const [percentageCPU, setPercentageCPU] = useState(0);
    const [percentageDisk, setPercentageDisk] = useState(0);

    function greet() {
        // Greet(name).then(updateResultText);
        setPercentageCPU(50)
        setPercentageDisk(10)
    }

    function colorPorcentaje(valor) {
        if (valor <= 33) {
            return '#6495ED'
        } else if (valor >= 34 && valor <= 66) {
            return '#FF8C00'
        } else {
            return '#FF0000'
        }
    }

    return (
        <>
            <h1 style={{ fontSize: 40 }}>Administrador de tareas</h1>
            <div class='row' style={{ padding: '2%' }}>
                <div class='col' >
                    <Card>
                        <Card.Header as="h5" style={{ fontSize: 30 }}>
                            Uso de CPU
                        </Card.Header>
                        <Card.Body style={{ overflowY: 'auto' }}>
                            <Card.Text>
                                <div style={{ width: 250, height: 250, margin: 'auto' }}>
                                    <CircularProgressbarWithChildren value={percentageCPU}
                                        styles={buildStyles({
                                            pathColor: colorPorcentaje(percentageCPU),
                                        })}>
                                        { }
                                        <img style={{ width: 110 }} src="https://i.imgur.com/b9NyUGm.png" alt="doge" />
                                        <div style={{ fontSize: 30 }}>
                                            <strong>{percentageCPU}%</strong>
                                        </div>
                                    </CircularProgressbarWithChildren>
                                </div><br />
                                <Table striped bordered hover >
                                    <thead>
                                        <tr>
                                            <th>Núcleos CPU</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <tr>
                                            <td >
                                                2
                                            </td>
                                        </tr>
                                    </tbody>
                                </Table>
                            </Card.Text>
                        </Card.Body>
                    </Card>
                </div>
                <div class='col' >
                    <Card>
                        <Card.Header as="h5" style={{ fontSize: 30 }}>
                            Uso de disco duro
                        </Card.Header>
                        <Card.Body style={{ overflowY: 'auto' }}>
                            <Card.Text>
                                <div style={{ width: 250, height: 250, margin: 'auto' }}>
                                    <CircularProgressbarWithChildren value={percentageDisk}
                                        styles={buildStyles({
                                            pathColor: colorPorcentaje(percentageDisk),
                                        })}>
                                        { }
                                        <img style={{ width: 110 }} src="https://i.imgur.com/b9NyUGm.png" alt="doge" />
                                        <div style={{ fontSize: 30 }}>
                                            <strong>{percentageDisk}%</strong>
                                        </div>
                                    </CircularProgressbarWithChildren>
                                </div>
                                <br />
                                <Table striped bordered hover >
                                    <thead>
                                        <tr>
                                            <th>Espacio utilizado</th>
                                            <th>Espacio libre</th>
                                            <th>Espacio total</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <tr>
                                            <td >
                                                2
                                            </td>
                                            <td >
                                                2
                                            </td>
                                            <td >
                                                2
                                            </td>
                                        </tr>
                                    </tbody>
                                </Table>
                            </Card.Text>
                        </Card.Body>
                    </Card>
                </div>
            </div>
            <button className="btn" onClick={greet}>dfafafda</button>
        </>
    )
}

export default App
