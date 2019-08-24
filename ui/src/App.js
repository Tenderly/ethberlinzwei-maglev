import React from 'react';
import {applyMiddleware, compose, createStore} from "redux";
import thunk from "redux-thunk";
import {Provider} from "react-redux";
import {Route, BrowserRouter as Router} from "react-router-dom";
import * as moment from "moment";

import reducers from './core/reducer';
import OverviewPage from "./pages/OverviewPage";
import Header from "./components/layout/Header";
import BatchPage from "./pages/BatchPage";

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

const store = createStore(
    reducers,
    composeEnhancers(
        applyMiddleware(thunk),
    ),
);

function App() {
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

export default App;
