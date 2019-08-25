import React, {Component} from 'react';
import {connect} from 'react-redux';

import BatchesList from "../components/BatchesList/BatchesList";
import Page from "../components/Page/Page";
import {getAllBatches, getFinishedBatches, getFinishedTx, getPendingBatches} from "../core/selectors";

import './OverviewPage.scss';
import StationAnimation from "../StationAnimation/StationAnimation";
import GasGraphs from "../components/GasGraphs/GasGraphs";
import GasMetrics from "../components/GasGraphs/GasMetrics";
import {bindActionCreators} from "redux";

import * as actions from '../core/actions';
import Loader from "../components/Loader/Loader";
import RerenderableTime from "../components/RerenderableTime/RerenderableTime";

function mapStateToProps({app}) {
    return {
        currentBatches: getPendingBatches(app),
        finishedBatches: getFinishedBatches(app),
        allBatches: getAllBatches(app, 5),
        finishedTx: getFinishedTx(app),
        batchingTransactions: app.batchingTransactions,
        lastSync: app.lastSync,
        appLoaded: app.appLoaded,
    };
}

function mapDispatchToProps(dispatch) {
    return {
        actions: bindActionCreators(actions, dispatch)
    };
}



class OverviewPage extends Component {
    render() {
        const {currentBatches, batchingTransactions, allBatches, finishedBatches, finishedTx, lastSync, appLoaded} = this.props;

        if (!appLoaded) {
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
                <div className="OverviewPage">
                    <div className="BatchesWrapper">
                        <BatchesList synced={lastSync} label="Current Batches" batches={currentBatches} inProgress={batchingTransactions}/>
                        <BatchesList synced={lastSync} label="Latest Batches" batches={finishedBatches}/>
                        <StationAnimation batches={allBatches}/>
                    </div>
                    <div className="AnimationsWrapper">
                        <GasMetrics batches={finishedBatches} transactions={finishedTx}/>
                        <GasGraphs batches={finishedBatches} transactions={finishedTx}/>
                    </div>
                </div>
            </Page>
        );
    }
}

export default connect(
    mapStateToProps,
    mapDispatchToProps,
)(OverviewPage);
