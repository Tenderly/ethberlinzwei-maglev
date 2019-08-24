import React from 'react';

import './Card.scss';

const Card = ({children, className}) => {
    return (
        <div className="Card">
            {children}
        </div>
    );
};

export default Card;
