import React from 'react'
import { connect } from 'react-redux'
import { temperatureSetting } from '../actions/traits'

import {Stack, TextField, PrimaryButton, Label} from 'office-ui-fabric-react'

let AddTemperatureSetting = ({ id, type = "VIRTUAL", dispatch }) => {
  let temperature

  return (
      <Stack>
        <Stack.Item align="start">
          <TextField label="Temperature" onChange={(e, value) => temperature = value } />
        </Stack.Item>
        
        <Stack.Item align="start">
          <Label>Primary</Label>
          <PrimaryButton
            data-automation-id="test"
            text="Button"
            onClick={() => dispatch(temperatureSetting({id, type, temperature}))}
          />
        </Stack.Item>
      </Stack>
  )
}
AddTemperatureSetting = connect()(AddTemperatureSetting)

export default AddTemperatureSetting