import React from 'react';
import {CheckCircle} from 'react-feather';

import Logo from './MaglevLogo.svg';

import './Header.scss';
import {Link} from "react-router-dom";

const Header = () => {
    return (
        <header className="Header">
            <Link to="/" className="LogoWrapper">
                <img src={Logo} className="LogoImage" alt="Maglev Logo"/>
                <div className="AppName">
                    Maglev
                </div>
                <div className="AppDesc">An Ethereum Tx Station</div>
            </Link>
            <div className="ConnectionBar">
                <CheckCircle size={16}/>
                <span>Connected</span>
            </div>
        </header>
    );
};

export default Header;
