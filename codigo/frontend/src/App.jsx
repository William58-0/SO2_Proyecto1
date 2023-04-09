import { useState } from 'react';
import './App.css';
import { Table, Card, Button } from 'react-bootstrap';
import { CircularProgressbarWithChildren, buildStyles } from "react-circular-progressbar";


function App() {
    const [percentageCPU, setPercentageCPU] = useState(0);
    const [nucleosCPU, setNucleosCPU] = useState(0);

    const [percentageDisk, setPercentageDisk] = useState(0);
    const [usadoDisk, setUsadoDisk] = useState(0);
    const [libreDisk, setLibreDisk] = useState(0);
    const [totalDisk, setTotalDisk] = useState(0);

    const [fase2, setFase2] = useState(true);
    const [usadoMemoria, setUsadoMemoria] = useState(0);
    const [totalMemoria, setTotalMemoria] = useState(0);
    const [percentageMemoria, setPercentageMemoria] = useState(0);

    runtime.EventsOn("recursos", (content) => {
        // console.log(content)
        var json = JSON.parse(content)
        console.log(json)

        setPercentageCPU(json['CPU']['Porcentaje'])
        setNucleosCPU(json['CPU']['Nucleos'])

        setPercentageDisk(json['Disco']['Porcentaje'])
        setUsadoDisk(json['Disco']['Usado'])
        setLibreDisk(json['Disco']['Disponible'])
        setTotalDisk(json['Disco']['Total'])

        setUsadoMemoria(json['Memoria']['Usado'])
        setTotalMemoria(json['Memoria']['Total'])
        setPercentageMemoria(json['Memoria']['Porcentaje'])
    })

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
            <h1 style={{ fontSize: 40, marginTop: '2%' }}>Administrador de tareas</h1>
            <div className="d-flex justify-content-around align-items-around">
                <Button onClick={() => setFase2(false)} variant="success">CPU y Disco</Button>
                <Button onClick={() => setFase2(true)}>Memoria</Button>
            </div>
            <div class='row' style={{ padding: '2%' }}>
                {fase2 ? <>
                    <Card>
                        <Card.Header as="h5" style={{ fontSize: 30 }}>
                            Uso de Memoria
                        </Card.Header>
                        <Card.Body style={{ overflowY: 'auto' }}>
                            <Card.Text>
                                <div style={{ width: 250, height: 250, margin: 'auto' }}>
                                    <CircularProgressbarWithChildren value={percentageMemoria}
                                        styles={buildStyles({
                                            pathColor: colorPorcentaje(percentageMemoria),
                                        })}>
                                        { }
                                        <img style={{ width: 110 }} src="https://img.freepik.com/iconos-gratis/ram_318-446826.jpg" alt="doge" />
                                        <div style={{ fontSize: 30 }}>
                                            <strong>{percentageMemoria}%</strong>
                                        </div>
                                    </CircularProgressbarWithChildren>
                                </div><br />
                                <Table striped bordered hover >
                                    <thead>
                                        <tr>
                                            <th>En uso</th>
                                            <th>Total</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <tr>
                                            <td >
                                                {usadoMemoria} GB
                                            </td>
                                            <td >
                                                {totalMemoria} GB
                                            </td>
                                        </tr>
                                    </tbody>
                                </Table>
                            </Card.Text>
                        </Card.Body>
                    </Card>
                </>
                    :
                    <>
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
                                                <img style={{ width: 110 }} src="https://cdn-icons-png.flaticon.com/512/4617/4617522.png" alt="doge" />
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
                                                        {nucleosCPU}
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
                                                <img style={{ width: 110 }} src="https://cdn-icons-png.flaticon.com/512/287/287441.png" alt="doge" />
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
                                                        {usadoDisk} GB
                                                    </td>
                                                    <td >
                                                        {libreDisk} GB
                                                    </td>
                                                    <td >
                                                        {totalDisk} GB
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </Table>
                                    </Card.Text>
                                </Card.Body>
                            </Card>
                        </div>
                    </>
                }
            </div>
            <div style={{ textAlign: 'right', padding: '2%' }}>
                William Alejandro Borrayo Alarcón - 201909103
            </div>
        </>
    )
}

export default App
