import React, { Fragment } from "react";
import { Provider } from "react-redux";
import { BrowserRouter as Router } from "react-router-dom";
import Store from "../store";
import HomePage from "./home";

function App() {
  return (
    <Fragment>
      <Router>
        <Provider store={Store}>
          <HomePage />
        </Provider>
      </Router>
    </Fragment>
  );
}

export default App;
