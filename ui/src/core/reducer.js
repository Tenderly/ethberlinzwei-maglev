import {combineReducers} from "redux";
import moment from "moment";

const initialState = {
    batches: {
        '0xae7b283ac112e89379f': {
            id: '0xae7b283ac112e89379f',
            status: 'pending',
            startedAt: moment('2019-08-24 12:31'),
            transactions: 6,
        },
    },
    virtualTransactions: {},
    batchTransactions: {},
    lastSync: moment(),
    appLoaded: false,
};

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
