import React from 'react';
import humanizeDuration from 'humanize-duration';
import Card from "../Card/Card";

import './BatchesList.scss';
import moment from "moment";
import RerenderableTime from "../RerenderableTime/RerenderableTime";
import BatchStatus from "../Batch/BatchStatus";

const BatchesList = ({batches, label, synced}) => {
    return (
        <Card className="BatchesList">
            <div className="CardHeading">
                <h3>{label}</h3>
                <div className="MarginLeftAuto">
                    <RerenderableTime date={synced} format={date => date.fromNow()} className="SmallText"/>
                </div>
            </div>
            <div className="CardBody">
                {batches.map(batch => {
                    return (
                        <div key={batch.id} className="BatchesList__Item">
                            <div>
                                <div>
                                    <strong className="MonoText">{batch.id}</strong>
                                </div>
                                <div>
                                    <RerenderableTime date={batch.startedAt} format={(date, now) => humanizeDuration(moment.duration(now.diff(date, 'seconds'), 'seconds').asMilliseconds(), { largest: 2 })}/>
                                </div>
                            </div>
                            <div>
                                <div>
                                    {batch.transactions} {batch.transactions === 1 ? 'Transaction' : 'Transactions'}
                                </div>
                                <div>
                                    <BatchStatus batch={batch}/>
                                </div>
                            </div>
                        </div>
                    )
                })}
            </div>
        </Card>
    );
};

export default BatchesList;
