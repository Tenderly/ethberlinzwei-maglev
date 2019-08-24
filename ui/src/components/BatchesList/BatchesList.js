import React from 'react';
import Card from "../Card/Card";

import './BatchesList.scss';

const BatchesList = ({batches, label}) => {
    return (
        <Card>
            <div className="CardHeading">
                <h3>{label}</h3>
            </div>
            <div className="CardBody">
                <div>

                </div>
            </div>
        </Card>
    );
};

export default BatchesList;
