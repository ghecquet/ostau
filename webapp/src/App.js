import React from 'react'
import {Provider} from 'react-redux'
import {Grommet} from 'grommet'
import {grommet} from 'grommet/themes'
import logo from './logo.svg'
import configureStore from './store'
import './App.css';
import Devices from './components/Devices'
import Switch from './components/Switch'
import Thermostat from './components/Thermostat'
import AddTemperatureSetting from './containers/AddTemperatureSetting'

function App() {
  return (
    <Provider store={configureStore()}>
      <Grommet theme={grommet} full={true}>
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
            Edit <code>src/App.js</code> and save to reload.
          </p>
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>
        </header>

        <Devices devices={[{id: 1, label: "Test Device"}]} />
        <Switch label="Device 1" on="false" />
        <div className="Thermostat">
          <Thermostat />
        </div>

        <AddTemperatureSetting id="whatever" />
      </div>
      </Grommet>
    </Provider>
  );
}

export default App;
