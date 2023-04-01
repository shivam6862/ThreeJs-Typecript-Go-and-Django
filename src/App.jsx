import React from "react";
import classes from "./App.module.css";

import Ball from "./component/threeJs/Ball";
import Header from "./component/Header";
import Line from "./component/threeJs/Line";
import Text from "./component/threeJs/Text";

function App() {
  return (
    <div className={classes.App}>
      <div className={classes.Header}>
        <Header />
      </div>
      <Ball />
      <Line />
      <Text />
    </div>
  );
}

export default App;
