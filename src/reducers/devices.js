export default (state = [], action) => {
    switch (action.type) {
        case 'ADD_DEVICE':
            return [...state, action.payload]
        default:
            return state
    }
}