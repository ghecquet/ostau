import { createStore, applyMiddleware } from 'redux'
import { composeWithDevTools } from 'redux-devtools-extension';
import thunk from 'redux-thunk'
import mqtt from 'redux-mqtt'
import reducer from './reducers/root'
import PahoMQTT from 'paho-mqtt'

const client = new PahoMQTT.Client("localhost", 9001, "");

client.onMessageArrived = onMessageArrived

client.connect({onSuccess:onConnect});

// called when a message arrives
function onMessageArrived(message) {
    console.log("onMessageArrived:"+message.payloadString);
  }

// called when the client connects
function onConnect() {
    // Once a connection has been made, make a subscription and send a message.
    console.log("onConnect");
    client.subscribe("#");
    const message = new PahoMQTT.Message("Hello");
    message.destinationName = "World";
    client.send(message);
  }

export default function configureStore(initialState={}) {
    return createStore(
        reducer,
        initialState,
        composeWithDevTools(applyMiddleware(thunk, mqtt(client))),
    );
}