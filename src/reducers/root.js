import { combineReducers } from 'redux';
import roomsReducer from './rooms';
import devicesReducer from './devices';
import { reduxMqttReducer } from 'redux-mqtt'

export default combineReducers({
    roomsReducer,
    devicesReducer,
    reduxMqttReducer,
});