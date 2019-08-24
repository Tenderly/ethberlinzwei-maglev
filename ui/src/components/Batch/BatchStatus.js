import React, {Fragment} from 'react';

import {CheckCircle} from 'react-feather';

import './BatchStatus.scss';

const BatchStatus = ({batch}) => {
    return (
        <div className="BatchStatus">
            {batch.status === 'pending' && <Fragment>
                <CheckCircle/>
                <span className="BlueText">Pending</span>
            </Fragment>}
            {batch.status === 'success' && <Fragment>
                <CheckCircle/>
                <span className="SuccessText">Success</span>
            </Fragment>}
        </div>
    );
};

export default BatchStatus;
