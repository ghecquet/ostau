import { TEMPERATURE_SETTING } from '../actions/traits';

export default (state = {}, action) => {

    console.log(action)
    switch (action.type) {
        case TEMPERATURE_SETTING: {

            return {
                ...state,
                temperatureSetting: {
                    thermostatTemperatureAmbient: action.temperature,
                    thermostatHumidityAmbient: action.humidity
                }
            }
        }
        default:
            return state;
    }
}