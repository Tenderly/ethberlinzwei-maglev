import React from 'react';
import {applyMiddleware, compose, createStore} from "redux";
import thunk from "redux-thunk";
import {Provider} from "react-redux";

import './App.css';

import reducers from './core/reducer';

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
        <div className="App">
            test app
        </div>
    </Provider>
  );
}

export default App;
