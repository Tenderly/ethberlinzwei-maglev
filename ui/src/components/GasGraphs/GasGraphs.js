import React, {Component} from 'react';

import {Area, AreaChart, ResponsiveContainer, Tooltip} from "recharts";
import Card from "../Card/Card";

import './GasGraphs.scss';
import * as _ from "lodash";
import BigNumber from "bignumber.js";

const dummyData = [
    {
        batchId: 'asdqwe1231',
        gasCost: 210000,
        gasCostPerTx: 17433,
        transactions: 234
    },
    {
        batchId: 'asdqwe1232',
        gasCost: 311230,
        gasCostPerTx: 15302,
        transactions: 234
    },
    {
        batchId: 'asdqwe1233',
        gasCost: 243000,
        gasCostPerTx: 12439,
        transactions: 234
    },
    {
        batchId: 'asdqwe1234',
        gasCost: 180000,
        gasCostPerTx: 16000,
        transactions: 234
    },
    {
        batchId: 'asdqwe1235',
        gasCost: 220000,
        gasCostPerTx: 56000,
        transactions: 234
    },
    {
        batchId: 'asdqwe1236',
        gasCost: 213000,
        gasCostPerTx: 17000,
        transactions: 234
    },
    {
        batchId: 'asdqwe1237',
        gasCost: 207000,
        gasCostPerTx: 19000,
        transactions: 234
    },
    {
        batchId: 'asdqwe1238',
        gasCost: 210000,
        gasCostPerTx: 28000,
        transactions: 234
    },
    {
        batchId: 'asdqwe1239',
        gasCost: 210000,
        gasCostPerTx: 14000,
        transactions: 234
    },
];

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
