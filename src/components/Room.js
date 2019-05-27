import React from 'react'
import PropTypes from 'prop-types'
import {Box} from 'grommet'

const Room = ({ label }) => (
    <Box border={{ color: 'brand', size: 'large' }} pad='xlarge'>
        {label}
    </Box>
)

Room.propTypes = {
    label: PropTypes.string.isRequired
}

export default Room