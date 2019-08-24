import React from 'react';
import {CheckCircle} from 'react-feather';

import Logo from './MaglevLogo.svg';

import './Header.scss';

const Header = () => {
    return (
        <header className="Header">
            <div className="LogoWrapper">
                <img src={Logo} className="LogoImage" alt="Maglev Logo"/>
                <div className="AppName">
                    ETH Maglev
                </div>
            </div>
            <div className="ConnectionBar">
                <CheckCircle size={16}/>
                <span>Connected</span>
            </div>
        </header>
    );
};

export default Header;
