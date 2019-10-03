export default (state = [], action) => {
    switch (action.type) {
        case 'ADD_ROOM':
            return [...state, action.payload]
        default:
            return state
    }
}