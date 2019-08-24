import React from 'react';
import Card from "../Card/Card";

import './GasMetrics.scss';

const GasMetrics = () => {
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
                            Based on previous 14 batches in the last 24h
                        </div>
                    </div>
                </Card>
                <Card>
                    <div className="CardHeading">
                        <h3>Avg. Batch Gas Cost</h3>
                    </div>
                    <div className="CardBody">
                        <div className="MetricNumber">
                            210,000
                        </div>
                        <div className="MetricDescription">
                            Based on previous 14 batches in the last 24h
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
                            Based on previous 14 batches in the last 24h
                        </div>
                    </div>
                </Card>
                <Card>
                    <div className="CardHeading">
                        <h3>Avg. Batch Gas Cost</h3>
                    </div>
                    <div className="CardBody">
                        <div className="MetricNumber">
                            210,000
                        </div>
                        <div className="MetricDescription">
                            Based on previous 14 batches in the last 24h
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
