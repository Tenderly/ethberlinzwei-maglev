import React, {Component} from 'react';
import moment from "moment";

class RerenderableTime extends Component {
    constructor(props) {
        super(props);

        this.state = {
            now: moment(),
        }
    }

    componentDidMount() {
        const interval = setInterval(() => {
            this.setState({
                now: moment(),
            });
        }, 3000);

        this.setState({
            interval,
        });
    }

    componentWillUnmount() {
        clearInterval(this.state.interval);
    }

    render() {
        const {now} = this.state;
        const {date, format, className} = this.props;

        return (
            <span className={className}>{format(date, now)}</span>
        );
    }
}

export default RerenderableTime;
