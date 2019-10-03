import React from 'react'
import PropTypes from 'prop-types'
import { Toggle } from 'office-ui-fabric-react/lib/Toggle';

const Switch = ({ label, on = false, onChange }) => (
    <Toggle
        defaultChecked={on}
        label={label}
        onText="On"
        offText="Off"
        onFocus={() => console.log('onFocus called')}
        onBlur={() => console.log('onBlur called')}
        onChange={onChange}
    />
)

Switch.propTypes = {
    on: PropTypes.bool,
    onChange: PropTypes.func
}

export default Switch