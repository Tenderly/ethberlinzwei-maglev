import React from 'react';
import {applyMiddleware, compose, createStore} from "redux";
import thunk from "redux-thunk";
import {Provider} from "react-redux";
import {Route, BrowserRouter as Router} from "react-router-dom";

import reducers from './core/reducer';
import OverviewPage from "./pages/OverviewPage";
import Header from "./components/layout/Header";

import './App.css';

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
                <div className="PageContent">
                    <Route path="/" exact component={OverviewPage}/>
                </div>
            </div>
        </Router>
    </Provider>
  );
}

export default App;
