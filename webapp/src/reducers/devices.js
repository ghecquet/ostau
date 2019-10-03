import addTrait from './traits'

export default (state = [], action) => {
    switch (action.type) {
        case 'ADD_DEVICE':
            return {
                ...state, 
                [action.id]: action.payload
            }
        case 'ADD_DEVICE_TRAIT':
            return {
                ...state,
                [action.data.id]: addTrait(state[action.data.id], action.data)
            }
        default:
            return state
    }
}