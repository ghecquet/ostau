import React from 'react'
import {Provider} from 'react-redux'
import {Grommet} from 'grommet'
import {grommet} from 'grommet/themes'
import logo from './logo.svg'
import configureStore from './store'
import './App.css';
import Rooms from './components/Rooms'

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

        <Rooms rooms={[{id: 1, label: "Test Room"}]} />
      </div>
      </Grommet>
    </Provider>
  );
}

export default App;
