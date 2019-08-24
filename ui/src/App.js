import React, {Component} from 'react';
import {applyMiddleware, compose, createStore} from "redux";
import thunk from "redux-thunk";
import {Provider} from "react-redux";
import {Route, BrowserRouter as Router} from "react-router-dom";
import axios from 'axios';

import reducers from './core/reducer';
import OverviewPage from "./pages/OverviewPage";
import Header from "./components/layout/Header";
import BatchPage from "./pages/BatchPage";
import * as _ from "lodash";
import * as actions from './core/actions';

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

const store = createStore(
    reducers,
    composeEnhancers(
        applyMiddleware(thunk),
    ),
);

class App extends Component {
    componentDidMount() {
        this.startPolling();
    }

    startPolling = _.throttle(async () => {
        try {
            const {data} = await axios.get('http://165.22.18.163:80/find');

            await store.dispatch(actions.setWorld(data));
        } catch (e) {
            console.error(e);
        }

        this.startPolling();
    }, 2000);

    render() {
        return (
            <Provider store={store}>
                <Router>
                    <div className="App">
                        <Header/>
                        <Route path="/" exact component={OverviewPage}/>
                        <Route path="/batch/:id" exact component={BatchPage}/>
                    </div>
                </Router>
            </Provider>
        );
    }
}

export default App;
