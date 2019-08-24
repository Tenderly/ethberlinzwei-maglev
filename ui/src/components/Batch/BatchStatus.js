import React, {Fragment} from 'react';

import {CheckCircle, PauseCircle} from 'react-feather';

import './BatchStatus.scss';

const BatchStatus = ({batch}) => {
    return (
        <div className="BatchStatus">
            {batch.status === 'pending' && <Fragment>
                <PauseCircle size={18} color="#3B77F4"/>
                <span className="BlueText">Pending</span>
            </Fragment>}
            {batch.status === 'success' && <Fragment>
                <CheckCircle color="#16ba43" size={18}/>
                <span className="SuccessText">Success</span>
            </Fragment>}
        </div>
    );
};

export default BatchStatus;
