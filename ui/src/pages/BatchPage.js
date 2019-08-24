import React from 'react';
import Page from "../components/Page/Page";

import './BatchPage.scss';
import {ArrowLeft} from "react-feather";
import {Link} from "react-router-dom";
import {connect} from "react-redux";
import {getBatchById, getTransactionsForBatch} from "../core/selectors";
import Card from "../components/Card/Card";
import BatchStatus from "../components/Batch/BatchStatus";
import {generateShortAddress} from "../core/Ethereum";
import Loader from "../components/Loader/Loader";
import RerenderableTime from "../components/RerenderableTime/RerenderableTime";
import BigNumber from "bignumber.js";

const BatchPage = ({batch, transactions, appLoaded, lastSync}) => {
    if (!batch || !appLoaded) {
        return <Page>
            <div className="PageLoader">
                <Loader/>
                <RerenderableTime className="LoaderText" date={lastSync} format={(date, now) => {
                    const diff = now.diff(date, 'seconds');

                    if (diff > 90) return 'One block at a time :)';

                    if (diff > 60) return 'It\'s just a ether of time...';

                    if (diff > 30) return 'We\'ve got our best engineer, Nebojsa, working on it!';

                    if (diff > 10) return 'No worries, just a bit longer..';

                    return '';
                }}/>
            </div>
        </Page>
    }

    return (
        <Page>
            <div className="PageHeading">
                <Link to="/" className="BackButton">
                    <ArrowLeft size={32}/>
                </Link>
                <h1>Batch <span className="BlueText">{generateShortAddress(batch.id, 12, 6)}</span></h1>
            </div>
            <div className="BatchPageContent">
                <div className="BatchDetailsWrapper">
                    <Card>
                        <div className="CardHeading">
                            <h3>Batch Information</h3>
                        </div>
                        <div className="CardBody">
                            <div>
                                <div>Tx Hash</div>
                                <div>
                                    <span className="MonoText">{batch.id}</span>
                                </div>
                            </div>
                            <div>
                                <div>Status</div>
                                <div>
                                    <BatchStatus batch={batch}/>
                                </div>
                            </div>
                            <div>
                                <div>Gas</div>
                                <div>
                                    {new BigNumber(batch.totalGas).toFormat(0)}
                                </div>
                            </div>
                            <div>
                                <div>Avg. gas per tx</div>
                                <div>
                                    {new BigNumber(batch.totalGas / batch.transactions).toFormat(0)}
                                </div>
                            </div>
                        </div>
                    </Card>
                </div>
                <div className="BatchTxWrapper">
                    <Card>
                        <div className="CardHeading">
                            <h3>Transactions</h3>
                        </div>
                        <div className="CardBody BatchTransactionsList">
                            {transactions.map(tx => <div key={tx.id} className="BatchTransactionsList">
                                <div><strong>Virtual Hash:</strong> <span className="MonoText">{tx.id}</span></div>
                                <div>{tx.from} to {tx.to}</div>
                                <div><strong>Gas:</strong> {tx.gas} ({new BigNumber(tx.gas / batch.totalGas * 100).toFixed(2)}% of batch)</div>
                            </div>)}
                        </div>
                    </Card>
                </div>
            </div>
        </Page>
    );
};

const mapStateToProps = (state, ownProps) => {
    const {match: {params: {id}}} = ownProps;



    return {
        batch: getBatchById(state.app, id),
        transactions: getTransactionsForBatch(state.app, id),
        appLoaded: state.app.appLoaded,
        lastSync: state.app.lastSync,
    };
};

export default connect(
    mapStateToProps,
)(BatchPage);
