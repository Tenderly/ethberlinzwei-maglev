export const ADD_BATCH = 'ADD_BATCH';
export const FINISH_BATCH = 'FINISH_BATCH';
export const UPDATE_WORLD = 'UPDATE_WORLD';

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

export const setWorld = (data) => {
    return dispatch => {
        const world = {
            batches: {},
            virtualTransactions: {},
            batchTransactions: {},
        };

        const batches = data.filter(d => !!d.tx_hash);

        batches.forEach(b => {
            world.batches[b.tx_hash] = {
                id: b.tx_hash,
                status: !!b.block_number ? 'success' : 'pending',
                transactions: !world.batches[b.tx_hash] ? 1 : world.batches[b.tx_hash].transactions + 1
            };

            world.virtualTransactions[b.v_hash] = {
                id: b.v_hash,
                batchId: b.tx_hash,
                from: b.identity.toLowerCase(),
                to: b.to.toLowerCase(),
            };

            if (world.batchTransactions[b.tx_hash]) {
                world.batchTransactions[b.tx_hash] = [
                    ...world.batchTransactions[b.tx_hash],
                    b.v_hash
                ];
            } else {
                world.batchTransactions[b.tx_hash] = [b.v_hash];
            }
        });

        dispatch({
            type: UPDATE_WORLD,
            batches: world.batches,
            virtualTransactions: world.virtualTransactions,
            batchTransactions: world.batchTransactions,
        });
    };
};
