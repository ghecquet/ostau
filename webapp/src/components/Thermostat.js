import React from 'react'
import Nest from 'react-nest-thermostat'

const Thermostat = ({ }) => (
    <Nest
        minValue={16}
        maxValue={30}
        ambientTemperature={20}
        targetTemperature={21}
    />
)

export default Thermostat