import React from 'react';
import Card from "../Card/Card";
import classNames from 'classnames';

import './GasMetrics.scss';
import * as _ from "lodash";
import BigNumber from "bignumber.js";

const GasMetrics = ({transactions, batches}) => {
    const txAverage = _.meanBy(transactions, tx => tx.gas);

    const batchAverage = _.meanBy(batches, b => b.totalGas);

    const txPerBatchAverage = _.meanBy(batches, b => b.transactions);

    const gasSavedPerc = new BigNumber(txAverage / 21000 * 100).toFixed(2);

    return (
        <div className="GasMetrics">
            <div className="MetricsRow">
                <Card>
                    <div className="CardHeading">
                        <h3>Avg. Batch Resolution</h3>
                    </div>
                    <div className="CardBody">
                        <div className="MetricNumber">
                            2h 23min
                        </div>
                        <div className="MetricDescription">
                            Based on the execution time of the previous {batches.length} batches.
                        </div>
                    </div>
                </Card>
                <Card>
                    <div className="CardHeading">
                        <h3>Avg. Batch Gas Cost</h3>
                    </div>
                    <div className="CardBody">
                        <div className="MetricNumber">
                            {new BigNumber(batchAverage).toFormat()}
                        </div>
                        <div className="MetricDescription">
                            Based on the average gas cost of the previous {batches.length} batches.
                        </div>
                    </div>
                </Card>

                <Card>
                    <div className="CardHeading">
                        <h3>Avg. Tx Gas Cost</h3>
                    </div>
                    <div className="CardBody">
                        <div className={classNames(
                            "MetricNumber"
                        )}>
                            {new BigNumber(txAverage).toFormat()}
                        </div>
                        <div className="MetricDescription">
                            Based on the previous {transactions.length} transactions from {batches.length} batches.
                        </div>
                    </div>
                </Card>
            </div>
            <div className="MetricsRow">
                <Card>
                    <div className="CardHeading">
                        <h3>Tx Gas Saved</h3>
                    </div>
                    <div className="CardBody">
                        <div className={classNames(
                            "MetricNumber",
                            {
                                'SuccessText': gasSavedPerc < 100,
                                'DangerText': gasSavedPerc > 100,
                            }
                        )}>
                            {gasSavedPerc}%
                        </div>
                        <div className="MetricDescription">
                            How much a batched tx uses gas compared to the average tx gas used.
                        </div>
                    </div>
                </Card>
                <Card>
                    <div className="CardHeading">
                        <h3>Avg. Tx per Batch</h3>
                    </div>
                    <div className="CardBody">
                        <div className="MetricNumber">
                            {txPerBatchAverage}
                        </div>
                        <div className="MetricDescription">
                            Average transactions that are batched together based on the previous {batches.length} batches.
                        </div>
                    </div>
                </Card>

                <Card>
                    <div className="CardHeading">
                        <h3>Current Balance</h3>
                    </div>
                    <div className="CardBody">
                        <div className="MetricNumber">
                            0.371294 ETH
                        </div>
                        <div className="MetricDescription">
                            Base on average batch cost this will allow you to run <strong>5</strong> more batches.
                        </div>
                    </div>
                </Card>
            </div>
        </div>
    );
};

export default GasMetrics;
