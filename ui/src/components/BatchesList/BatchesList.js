import React from 'react';
import humanizeDuration from 'humanize-duration';
import Card from "../Card/Card";

import './BatchesList.scss';
import moment from "moment";
import RerenderableTime from "../RerenderableTime/RerenderableTime";
import BatchStatus from "../Batch/BatchStatus";
import {Link} from "react-router-dom";
import {generateShortAddress} from "../../core/Ethereum";

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
                        <Link to={`/batch/${batch.id}`} key={batch.id} className="BatchesList__Item">
                            <div>
                                <div className="mb5">
                                    <strong className="MonoText">{generateShortAddress(batch.id, 16, 6)}</strong>
                                </div>
                                <div className="SmallText">
                                    <RerenderableTime date={batch.startedAt} format={(date, now) => humanizeDuration(moment.duration(now.diff(date, 'seconds'), 'seconds').asMilliseconds(), { largest: 2 })}/>
                                </div>
                            </div>
                            <div className="OtherData">
                                <div className="mb10">
                                    <strong>{batch.transactions} tx</strong> in batch
                                </div>
                                <div>
                                    <BatchStatus batch={batch}/>
                                </div>
                            </div>
                        </Link>
                    )
                })}
            </div>
        </Card>
    );
};

export default BatchesList;
