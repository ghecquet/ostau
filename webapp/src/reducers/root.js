import { combineReducers } from 'redux';
import devicesReducer from './devices';
import virtualReducer from './virtual';

export default combineReducers({
    virtual: virtualReducer,
    devices: devicesReducer,
});