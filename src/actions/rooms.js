let id = 0

export const addRoom = (label) => dispatch => {
    dispatch({
        type: 'ADD_ROOM',
        payload: {
            id: id++,
            label: label
        }
    })
}