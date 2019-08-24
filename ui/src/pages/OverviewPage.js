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
