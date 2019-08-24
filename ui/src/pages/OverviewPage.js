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
                        <div>
                            <button onClick={() => this.handleAddBatch('batch1')}>Add 1</button>
                            <button onClick={() => this.handleFinishBatch('batch1')}>Finish 1</button>
                            <button onClick={() => this.handleAddBatch('batch2')}>Add 2</button>
                            <button onClick={() => this.handleFinishBatch('batch2')}>Finish 2</button>
                            <button onClick={() => this.handleAddBatch('batch3')}>Add 3</button>
                            <button onClick={() => this.handleAddBatch('batch4')}>Add 4</button>
                            <button onClick={() => this.handleAddBatch('batch5')}>Add 5</button>
                        </div>
                        <StationAnimation batches={currentBatches}/>
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
