import React, {Component} from 'react';

import StationBack from './StationBack.svg';
import Rails from './Rails.svg';
import TrainFront from './TrainFront.svg';
import TrainWagon from './TrainWagon.svg';
import StationFront from './StationFront.svg';

import './StationAnimation.scss'

class StationAnimation extends Component {
    render() {
        return (
            <div className="StationAnimation">
                <div className="Fade"/>
                <div className="FrontLayer Layer">
                    <img src={StationFront} alt="" className="StationFront"/>
                </div>
                <div className="MiddleLayer Layer">
                    <div className="TrainTrack">
                        <div className="Rails">
                            <img src={Rails} alt="" className="Rail"/>
                            <img src={Rails} alt="" className="Rail"/>
                            <img src={Rails} alt="" className="Rail"/>
                            <img src={Rails} alt="" className="Rail"/>
                            <img src={Rails} alt="" className="Rail"/>
                            <img src={Rails} alt="" className="Rail"/>
                            <img src={Rails} alt="" className="Rail"/>
                            <img src={Rails} alt="" className="Rail"/>
                            <img src={Rails} alt="" className="Rail"/>
                        </div>
                        <div className="Train">
                            <img src={TrainFront} alt="" className="TrainFront"/>
                            <img src={TrainWagon} alt="" className="TrainWagon"/>
                            <img src={TrainWagon} alt="" className="TrainWagon"/>
                            <img src={TrainWagon} alt="" className="TrainWagon"/>
                        </div>
                    </div>
                </div>
                <div className="BackLayer Layer">
                    <img src={StationBack} alt="" className="StationFront"/>
                </div>
            </div>
        );
    }
}

export default StationAnimation;
