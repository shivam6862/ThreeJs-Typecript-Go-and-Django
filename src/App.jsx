import React from "react";
import classes from "./App.module.css";

import Ball from "./component/threeJs/Ball";
import Header from "./component/Header";

function App() {
  return (
    <div className={classes.App}>
      <div className={classes.Header}>
        <Header />
      </div>
      <Ball />
    </div>
  );
}

export default App;
