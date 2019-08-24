import React, {Component} from 'react';
import {connect} from 'react-redux';

import BatchesList from "../components/BatchesList/BatchesList";
import Page from "../components/Page/Page";
import {getFinishedBatches, getPendingBatches} from "../core/selectors";

import './OverviewPage.scss';
import StationAnimation from "../StationAnimation/StationAnimation";
import GasGraphs from "../components/GasGraphs/GasGraphs";
import GasMetrics from "../components/GasGraphs/GasMetrics";

function mapStateToProps({app}) {
    return {
        currentBatches: getPendingBatches(app),
        finishedBatches: getFinishedBatches(app),
        lastSync: app.lastSync,
    };
}

function mapDispatchToProps(dispatch) {
    return {};
}



class OverviewPage extends Component {
    render() {
        const {currentBatches, finishedBatches, lastSync} = this.props;

        return (
            <Page>
                <div className="OverviewPage">
                    <div className="BatchesWrapper">
                        <BatchesList synced={lastSync} label="Current Batches" batches={currentBatches}/>
                        <BatchesList synced={lastSync} label="Latest Batches" batches={finishedBatches}/>
                        <GasMetrics/>
                    </div>
                    <div className="AnimationsWrapper">
                        <StationAnimation/>
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
