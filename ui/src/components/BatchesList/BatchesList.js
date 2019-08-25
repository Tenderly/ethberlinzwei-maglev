import React from 'react';
import humanizeDuration from 'humanize-duration';
import Card from "../Card/Card";

import './BatchesList.scss';
import moment from "moment";
import RerenderableTime from "../RerenderableTime/RerenderableTime";
import BatchStatus from "../Batch/BatchStatus";
import {Link} from "react-router-dom";
import {generateShortAddress} from "../../core/Ethereum";

const BatchesList = ({batches, label, synced, inProgress}) => {
    return (
        <Card className="BatchesList">
            <div className="CardHeading">
                <h3>{label}</h3>
                <div className="MarginLeftAuto">
                    <RerenderableTime date={synced} format={date => date.fromNow()} className="SmallText"/>
                </div>
            </div>
            <div className="CardBody BatchesListScroll">
                {inProgress && !!inProgress.length && <div className="BatchesList__Item">
                    <div>
                        <div className="mb5">
                            <strong className="MonoText">Batching Transactions</strong>
                        </div>
                        <div className="SmallText">
                            In batch pool: <strong>{inProgress.length}</strong>
                        </div>
                    </div>
                    <div className="OtherData">
                        <div className="mb10">
                            <strong>{inProgress.length} tx</strong> in batch
                        </div>
                        <div>
                            <BatchStatus batch={{
                                status: 'batching',
                            }}/>
                        </div>
                    </div>
                </div>}
                {batches.map(batch => {
                    return (
                        <Link to={`/batch/${batch.id}`} key={batch.id} className="BatchesList__Item">
                            <div>
                                <div className="mb5">
                                    <strong className="MonoText">{generateShortAddress(batch.id, 16, 6)}</strong>
                                </div>
                                {batch.status === 'pending' && <div className="SmallText">
                                    Mining duration: <strong>
                                    <RerenderableTime date={batch.sentAt} format={(date, now) => humanizeDuration(moment.duration(now.diff(date, 'seconds'), 'seconds').asMilliseconds(), { largest: 2 })}/>
                                    </strong>
                                </div>}
                                {batch.status === 'success' && <div className="SmallText">
                                    Mined in: <strong>{humanizeDuration(moment.duration(batch.miningTime, 'seconds').asMilliseconds(), {
                                    largest: 2,
                                })}</strong>
                                </div>}
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
