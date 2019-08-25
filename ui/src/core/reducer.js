import {combineReducers} from "redux";
import moment from "moment";
import {ADD_BATCH, FINISH_BATCH, UPDATE_WORLD} from "./actions";
import * as _ from "lodash";
import humanizeDuration from "humanize-duration";

const initialState = {
    batches: {},
    virtualTransactions: {},
    batchTransactions: {},
    lastSync: moment(),
    appLoaded: false,
};

const reducer = (state = initialState, action) => {
    switch (action.type) {
        case ADD_BATCH:
            return {
                ...state,
                batches: {
                    ...state.batches,
                    [action.batchId]: {
                        id: action.batchId,
                        status: 'pending',
                        startedAt: moment(),
                        transactions: 120,
                    },
                }
            };
        case FINISH_BATCH:
            return {
                ...state,
                batches: {
                    ...state.batches,
                    [action.batchId]: {
                        ...state.batches[action.batchId],
                        status: 'success'
                    },
                }
            };
        case UPDATE_WORLD:
            return {
                ...state,
                lastSync: moment(),
                appLoaded: true,
                batches: action.batches,
                virtualTransactions: action.virtualTransactions,
                batchTransactions: action.batchTransactions,
                batchingTransactions: action.batchingTransactions,
            };
        default:
            return state;
    }
};

const reducers = combineReducers({
    app: reducer,
});

export default reducers;
