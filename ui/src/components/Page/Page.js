import React from 'react';

import './Page.scss';

const Page = ({children}) => {
    return (
        <div className="Page">
            {children}
        </div>
    );
};

export default Page;
