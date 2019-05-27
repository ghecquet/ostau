import { combineReducers } from 'redux';
import roomsReducer from './rooms';
import devicesReducer from './devices';

export default combineReducers({
    roomsReducer,
    devicesReducer
});