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
