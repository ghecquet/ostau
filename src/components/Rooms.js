import React from 'react'
import PropTypes from 'prop-types'
import {Grid} from 'grommet'
import Room from './Room'

const Rooms = ({ rooms }) => (
    <Grid>
        {rooms.map(room => (
          <Room key={room.id} {...room} />
        ))}
    </Grid>
)

Rooms.propTypes = {
    rooms: PropTypes.arrayOf(
        PropTypes.shape({
            id: PropTypes.number.isRequired,
            label: PropTypes.string.isRequired
        }).isRequired
    ).isRequired,
}

export default Rooms