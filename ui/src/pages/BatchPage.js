import React from 'react';
import Page from "../components/Page/Page";

import './BatchPage.scss';
import {ArrowLeft} from "react-feather";
import {Link} from "react-router-dom";

const BatchPage = () => {
    return (
        <Page>
            <div className="PageHeading">
                <Link to="/">
                    <ArrowLeft size={32}/>
                </Link>
                <h1>Batch <span className="BlueText">0x123123qwe123</span></h1>
            </div>
            <div className="BatchDetailsWrapper">

            </div>
            <div className="BatchTxWrapper">

            </div>
        </Page>
    );
};

export default BatchPage;
