import React, {Component} from 'react';

import StationBack from './StationBack.svg';
import Rails from './Rails.svg';
import TrainFront from './TrainFront.svg';
import TrainWagon from './TrainWagon.svg';
import StationFront from './StationFront.svg';

import './StationAnimation.scss'
import {Transition} from "react-transition-group";

class StationAnimation extends Component {
    constructor(props) {
        super(props);

        this.state = {
            activeBatches: props.batches ? props.batches.filter(b => b.status === 'pending').length : 0,
        };
    }

    incrementBatches = () => {
        this.setState({
            activeBatches: this.state.activeBatches + 1,
        });
    };

    decrementBatches = () => {
        this.setState({
            activeBatches: this.state.activeBatches - 1,
        });
    };

    render() {
        const {batches} = this.props;
        const {activeBatches} = this.state;

        return (
            <div className="StationAnimation">
                <div className="Fade"/>
                <div className="FrontLayer Layer">
                    <img src={StationFront} alt="" className="StationFront" data-batches={activeBatches}/>
                </div>
                <div className="MiddleLayer Layer">
                    {activeBatches === 0 && <div className="TrainTrack">
                        <div className="Rails NoAnim">
                            <div className="RailWrapper"><img src={Rails} alt="" className="Rail"/></div>
                            <div className="RailWrapper"><img src={Rails} alt="" className="Rail"/></div>
                            <div className="RailWrapper"><img src={Rails} alt="" className="Rail"/></div>
                            <div className="RailWrapper"><img src={Rails} alt="" className="Rail"/></div>
                            <div className="RailWrapper"><img src={Rails} alt="" className="Rail"/></div>
                            <div className="RailWrapper"><img src={Rails} alt="" className="Rail"/></div>
                            <div className="RailWrapper"><img src={Rails} alt="" className="Rail"/></div>
                            <div className="RailWrapper"><img src={Rails} alt="" className="Rail"/></div>
                            <div className="RailWrapper"><img src={Rails} alt="" className="Rail"/></div>
                        </div>
                    </div>}
                    {batches.map(batch => <Transition appear className key={batch.id} in={batch.status === 'pending'} mountOnEnter unmountOnExit timeout={1000} onEntering={this.incrementBatches} onExited={this.decrementBatches}>
                        {state => <div className={`TrainTrack TrainTrack--${state}`}>
                            <div className="Rails">
                                <div className="RailWrapper"><img src={Rails} alt="" className="Rail"/></div>
                                <div className="RailWrapper"><img src={Rails} alt="" className="Rail"/></div>
                                <div className="RailWrapper"><img src={Rails} alt="" className="Rail"/></div>
                                <div className="RailWrapper"><img src={Rails} alt="" className="Rail"/></div>
                                <div className="RailWrapper"><img src={Rails} alt="" className="Rail"/></div>
                                <div className="RailWrapper"><img src={Rails} alt="" className="Rail"/></div>
                                <div className="RailWrapper"><img src={Rails} alt="" className="Rail"/></div>
                                <div className="RailWrapper"><img src={Rails} alt="" className="Rail"/></div>
                                <div className="RailWrapper"><img src={Rails} alt="" className="Rail"/></div>
                            </div>
                            <div className="Train">
                                <img src={TrainFront} alt="" className="TrainFront"/>
                                <div className="TrainWagonWrapper"><img src={TrainWagon} alt="" className="TrainWagon"/></div>
                                <div className="TrainWagonWrapper"><img src={TrainWagon} alt="" className="TrainWagon"/></div>
                                <div className="TrainWagonWrapper"><img src={TrainWagon} alt="" className="TrainWagon"/></div>
                            </div>
                        </div>}
                    </Transition>)}
                </div>
                <div className="BackLayer Layer">
                    <img src={StationBack} alt="" className="StationBack"/>
                </div>
            </div>
        );
    }
}

export default StationAnimation;
