import React, {Component} from 'react';

import {Area, AreaChart, ResponsiveContainer, Tooltip} from "recharts";
import Card from "../Card/Card";

import './GasGraphs.scss';
import * as _ from "lodash";
import BigNumber from "bignumber.js";

class GasGraphs extends Component {
    render() {
        const {batches} = this.props;

        const last50 = _.takeRight(batches, 50);

        const graphData = last50.map(d => ({
            txHash: d.id,
            gasCost: d.totalGas,
            gasCostPerTx: new BigNumber(d.totalGas / d.transactions).toNumber(),
            transactions: d.transactions,
        }));

        return (
            <div className="GasGraphs">
                <Card>
                    <div className="CardHeading">
                        <h3>Gas Cost History</h3>
                    </div>
                    <div className="CardBody"><p>Average batched transaction cost and how much on average a transaction inside spent gas.</p></div>
                    <div className="ChartWrapper">
                        <ResponsiveContainer>
                            <AreaChart margin={{top: 0, bottom:0,left: 0, right: 0}} data={graphData}>
                                <defs>
                                    <linearGradient id="yellowGradient" x1="0" y1="0" x2="0" y2="1">
                                        <stop offset="15%" stopColor="#F4B83B" stopOpacity={0.75}/>
                                        <stop offset="100%" stopColor="#F4B83B" stopOpacity={0.1}/>
                                    </linearGradient>
                                    <linearGradient id="blueGradient" x1="0" y1="0" x2="0" y2="1">
                                        <stop offset="15%" stopColor="#3B77F4" stopOpacity={0.75}/>
                                        <stop offset="100%" stopColor="#3B77F4" stopOpacity={0.1}/>
                                    </linearGradient>
                                </defs>
                                <Tooltip/>
                                <Area type="monotone" xAxisId="name" dataKey="gasCost" stroke="#F4B83B" fill="url(#yellowGradient)" />
                                <Area type="monotone" xAxisId="name" dataKey="gasCostPerTx" stroke="#3B77F4" fill="url(#blueGradient)" />
                            </AreaChart>
                        </ResponsiveContainer>
                    </div>
                </Card>
            </div>
        );
    }
}

export default GasGraphs;
