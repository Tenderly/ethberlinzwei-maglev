import React, {Component} from 'react';
import {connect} from 'react-redux';

import BatchesList from "../components/BatchesList/BatchesList";
import Page from "../components/Page/Page";
import {getFinishedBatches, getPendingBatches} from "../core/selectors";

import './OverviewPage.scss';
import StationAnimation from "../StationAnimation/StationAnimation";
import GasGraphs from "../components/GasGraphs/GasGraphs";
import GasMetrics from "../components/GasGraphs/GasMetrics";
import {bindActionCreators} from "redux";

import * as actions from '../core/actions';
import Loader from "../components/Loader/Loader";
import moment from "moment";
import RerenderableTime from "../components/RerenderableTime/RerenderableTime";

function mapStateToProps({app}) {
    return {
        currentBatches: getPendingBatches(app),
        finishedBatches: getFinishedBatches(app),
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
    handleAddBatch = (batchId) => {
        const {actions} = this.props;

        actions.addBatch(batchId);
    };

    handleFinishBatch = (batchId) => {
        const {actions} = this.props;

        actions.finishBatch(batchId);
    };

    render() {
        const {currentBatches, finishedBatches, lastSync, appLoaded} = this.props;

        if (!appLoaded) {
            return <Page>
                <div className="OverviewLoader">
                    <Loader/>
                    <RerenderableTime className="LoaderText" date={lastSync} format={(date, now) => {
                        const diff = now.diff(date, 'seconds');

                        if (diff > 120) return 'One block at a time :)';

                        if (diff > 90) return 'It\'s just a ether of time...';

                        if (diff > 60) return 'We\'ve got our best engineer, Nebojsa, working on it!';

                        if (diff > 30) return 'No worries, just a bit longer..';

                        return '';
                    }}/>
                </div>
            </Page>
        }

        return (
            <Page>
                <div className="OverviewPage">
                    <div className="BatchesWrapper">
                        <BatchesList synced={lastSync} label="Current Batches" batches={currentBatches}/>
                        <BatchesList synced={lastSync} label="Latest Batches" batches={finishedBatches}/>
                        <StationAnimation batches={currentBatches}/>
                    </div>
                    <div className="AnimationsWrapper">
                        <GasMetrics/>
                        <GasGraphs/>
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
