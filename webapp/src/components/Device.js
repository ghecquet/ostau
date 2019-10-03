import React from 'react'
import PropTypes from 'prop-types'
import {Box} from 'grommet'

const Device = ({ label }) => (
    <Box border={{ color: 'brand', size: 'large' }} pad='xlarge'>
        {label}
    </Box>
)

Device.propTypes = {
    label: PropTypes.string.isRequired
}

export default Device