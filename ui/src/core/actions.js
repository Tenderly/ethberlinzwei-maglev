export const ADD_BATCH = 'ADD_BATCH';
export const FINISH_BATCH = 'FINISH_BATCH';

export const addBatch = (batchId) => {
    return dispatch => {
        dispatch({
            type: ADD_BATCH,
            batchId,
        });
    }
};

export const finishBatch = (batchId) => {
    return dispatch => {
        dispatch({
            type: FINISH_BATCH,
            batchId,
        });
    }
};
