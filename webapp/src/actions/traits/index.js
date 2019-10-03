export const ON_OFF = 'ON_OFF';
export const TEMPERATURE_SETTING = 'TEMPERATURE_SETTING'

export const temperatureSetting = params => ({
    type: `ADD_${params.type}_TRAIT`,
    data: {
        id: params.id,
        type: TEMPERATURE_SETTING,
        temperature: params.temperature,
        humidity: params.humidity
    }
})