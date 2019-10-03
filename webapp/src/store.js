import { createStore, applyMiddleware } from 'redux'
import { composeWithDevTools } from 'redux-devtools-extension';
import thunk from 'redux-thunk'
import reducer from './reducers/root'

import tasmota from './tasmota';

const mqttConfig = {
  host: "192.168.1.92",
  port: 9001
}

export default function configureStore(initialState={}) {
    return createStore(
        reducer,
        initialState,
        composeWithDevTools(applyMiddleware(thunk, tasmota(mqttConfig))),
    );
}