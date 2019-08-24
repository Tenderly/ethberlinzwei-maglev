import * as _ from "lodash";

export function getAllBatches(state, last) {
    const batches = Object.values(state.batches);

    if (last) {
        return _.takeRight(batches, last);
    }

    return batches;
}

export function getPendingBatches(state) {
    const batches =  Object.values(state.batches).filter(b => b.status === 'pending');

    if (!batches || !batches.length) {
        return [];
    }

    return batches;
}

export function getFinishedBatches(state) {
    const batches =  Object.values(state.batches).filter(b => b.status !== 'pending');

    if (!batches || !batches.length) {
        return [];
    }

    return batches;
}

export function getFinishedTx(state) {
    const txs =  Object.values(state.virtualTransactions).filter(b => b.finished);

    if (!txs || !txs.length) {
        return [];
    }

    return txs;
}

export function getBatchById(state, id) {
    const batch = state.batches[id];

    if (!batch) {
        return null;
    }

    return batch;
}

export function getTransactionsForBatch(state, id) {
    const txIds = state.batchTransactions[id];

    if (!txIds) return [];

    return txIds.map(tx => state.virtualTransactions[tx]);
}
