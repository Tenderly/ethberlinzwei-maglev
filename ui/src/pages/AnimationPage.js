import React, {Component} from 'react';
import {connect} from 'react-redux';
import {getAllBatches} from "../core/selectors";
import * as actions from "../core/actions";
import StationAnimation from "../StationAnimation/StationAnimation";
import Page from "../components/Page/Page";
import {bindActionCreators} from "redux";

function mapStateToProps({app}) {
    return {
        batches: getAllBatches(app),

    };
}

function mapDispatchToProps(dispatch) {
    return {
        actions: bindActionCreators(actions, dispatch),
    };
}

class AnimationPage extends Component {
    handleAddBatch = (batchId) => {
        const {actions} = this.props;

        actions.addBatch(batchId);
    };

    handleFinishBatch = (batchId) => {
        const {actions} = this.props;

        actions.finishBatch(batchId);
    };

    render() {
        const {batches} = this.props;
        return (
            <Page>
                <div>
                    <button onClick={() => this.handleAddBatch('batchOne')}>Add One</button>
                    <button onClick={() => this.handleFinishBatch('batchOne')}>Finish One</button>
                    <button onClick={() => this.handleAddBatch('batchTwo')}>Add Two</button>
                    <button onClick={() => this.handleFinishBatch('batchTwo')}>Finish Two</button>
                </div>
                <StationAnimation batches={batches}/>
            </Page>
        );
    }
}

export default connect(
    mapStateToProps,
    mapDispatchToProps,
)(AnimationPage);
