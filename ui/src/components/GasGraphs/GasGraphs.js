import React, {Component} from 'react';

import {Area, AreaChart, ResponsiveContainer, Tooltip} from "recharts";
import Card from "../Card/Card";

import './GasGraphs.scss';

const dummyData = [
    {
        batchId: 'asdqwe1231',
        gasCost: 210000,
        gasCostPerTx: 54,
        transactions: 234
    },
    {
        batchId: 'asdqwe1232',
        gasCost: 210000,
        gasCostPerTx: 54,
        transactions: 234
    },
    {
        batchId: 'asdqwe1233',
        gasCost: 210000,
        gasCostPerTx: 54,
        transactions: 234
    },
    {
        batchId: 'asdqwe1234',
        gasCost: 210000,
        gasCostPerTx: 54,
        transactions: 234
    },
    {
        batchId: 'asdqwe1235',
        gasCost: 210000,
        gasCostPerTx: 54,
        transactions: 234
    },
    {
        batchId: 'asdqwe1236',
        gasCost: 210000,
        gasCostPerTx: 54,
        transactions: 234
    },
    {
        batchId: 'asdqwe1237',
        gasCost: 210000,
        gasCostPerTx: 54,
        transactions: 234
    },
    {
        batchId: 'asdqwe1238',
        gasCost: 210000,
        gasCostPerTx: 54,
        transactions: 234
    },
    {
        batchId: 'asdqwe1239',
        gasCost: 210000,
        gasCostPerTx: 100000,
        transactions: 234
    },
];

class GasGraphs extends Component {
    render() {
        return (
            <div className="GasGraphs">
                <Card>
                    <div className="CardHeading">
                        <h3>Gas Cost History</h3>
                    </div>
                    <div className="CardBody"><p>Average batched transaction cost and how much on average a transaction inside spent gas.</p></div>
                    <div className="ChartWrapper">
                        <ResponsiveContainer>
                            <AreaChart margin={{top: 0, bottom:0,left: 0, right: 0}} data={dummyData}>
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
