/**
 * @param {string} address
 * @param {number} [leftOffset]
 * @param {number} [rightOffset]
 * @returns {string}
 */
export function generateShortAddress(address, leftOffset = 6, rightOffset = 4) {
    return `${address.slice(0, leftOffset)}...${address.slice(address.length - rightOffset)}`;
}
