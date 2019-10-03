import React from 'react'
import PropTypes from 'prop-types'
import {Grid} from 'grommet'
import Device from './Device'

const Devices = ({ devices }) => (
    <Grid>
        {devices.map(device => (
          <Device key={device.id} {...device} />
        ))}
    </Grid>
)

Devices.propTypes = {
    devices: PropTypes.arrayOf(
        PropTypes.shape({
            id: PropTypes.number.isRequired
        }).isRequired
    ).isRequired,
}

export default Devices