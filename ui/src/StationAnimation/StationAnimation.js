import React, {Component} from 'react';

import StationBack from './StationBack.svg';
import Rails from './Rails.svg';
import TrainFront from './TrainFront.svg';
import TrainWagon from './TrainWagon.svg';
import StationFront from './StationFront.svg';

import './StationAnimation.scss'

class StationAnimation extends Component {
    constructor(props) {
        super(props);

        this.state = {
            trains: [],
        };
    }

    componentDidMount() {
        const {batches} = this.props;

        console.log(batches);

        this.setState({
            trains: batches.reduce((data, batch) => {
                return data;
            }, {}),
        })
    }

    render() {
        const {batches} = this.props;

        return (
            <div className="StationAnimation">
                <div className="Fade"/>
                <div className="FrontLayer Layer">
                    <img src={StationFront} alt="" className="StationFront" data-batches={batches.length}/>
                </div>
                <div className="MiddleLayer Layer">
                    {batches.length === 0 && <div className="TrainTrack">
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
                    {batches.map(batch => <div className="TrainTrack" key={batch.id}>
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
                    </div>)}
                </div>
                <div className="BackLayer Layer">
                    <img src={StationBack} alt="" className="StationBack"/>
                </div>
            </div>
        );
    }
}

export default StationAnimation;
