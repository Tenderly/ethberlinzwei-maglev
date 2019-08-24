import React from 'react';

import BatchesList from "../components/BatchesList/BatchesList";
import Page from "../components/Page/Page";

const OverviewPage = () => {
    return (
        <Page>
            <div>
                <BatchesList label="Current Batches"/>
                <BatchesList label="Latest Batches" className=""/>
            </div>
        </Page>
    );
};

export default OverviewPage;
