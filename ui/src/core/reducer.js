import {combineReducers} from "redux";
import moment from "moment";
import {ADD_BATCH, FINISH_BATCH} from "./actions";

const initialState = {
    batches: {
        'batch1': {
            id: 'batch1',
            status: 'pending',
            startedAt: moment('2019-08-24 12:31'),
            transactions: 6,
        },
    },
    virtualTransactions: {
        '0xvirt1': {
            vHash: '0xvirt1',
            from: 'qwe',
            to: '0xGoToAddress',
        },
        '0xvirt2': {
            vHash: '0xvirt2',
            from: 'qwe',
            to: '0xGoToAddress',
        },
        '0xvirt3': {
            vHash: '0xvirt3',
            from: 'qwe',
            to: '0xGoToAddress',
        },
        '0xvirt4': {
            vHash: '0xvirt4',
            from: 'qwe',
            to: '0xGoToAddress',
        },
        '0xvirt5': {
            vHash: '0xvirt5',
            from: 'qwe',
            to: '0xGoToAddress',
        },
        '0xvirt6': {
            vHash: '0xvirt6',
            from: 'qwe',
            to: '0xGoToAddress',
        },
        '0xvirt7': {
            vHash: '0xvirt7',
            from: 'qwe',
            to: '0xGoToAddress',
        },
    },
    batchTransactions: {
        '0xae7b283ac112e89379f': [
            '0xvirt1',
            '0xvirt2',
            '0xvirt3',
            '0xvirt4',
            '0xvirt5',
            '0xvirt6',
            '0xvirt7',
        ],
    },
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
        default:
            return state;
    }
};

const reducers = combineReducers({
    app: reducer,
});

export default reducers;
