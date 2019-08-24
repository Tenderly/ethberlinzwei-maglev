import {combineReducers} from "redux";

const initialState = {};

const reducer = (state = initialState, action) => {
    switch (action.type) {
        default:
            return state;
    }
};

const reducers = combineReducers({
    app: reducer,
});

export default reducers;
