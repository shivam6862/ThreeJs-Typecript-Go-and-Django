import "./App.css";
import { Basic } from "./components/Basic";
import { Object } from "./components/Object";
import { ArrayOfObject } from "./components/ArrayOfObjec";
import { Status } from "./components/Status";
import { Children } from "./components/Children";
import { OscarChildren } from "./components/OscarChildren";
import { Click } from "./components/Click";
import { ChangeInput } from "./components/ChangeInput";

import { nameList } from "./Data/nameList";
import { personName } from "./Data/personName";

function App() {
  return (
    <div className="App">
      <div>
        <h1>Typecript</h1>
        <Basic name="Shivam" messageCount={20} isLoggedIn={true} />
        <Object name={personName} />
        <ArrayOfObject nameList={nameList} />
        <Status status={"success"} />
        <Children>Children</Children>
        <OscarChildren>
          <Children>OscarChildren</Children>
        </OscarChildren>
        <Basic name="Shivam" isLoggedIn={true} />
        <Click
          handleClick={(event, id) => {
            console.log("Button Clicked!", event, id);
          }}
        />
        <ChangeInput value="" handleChange={(event) => console.log(event.target.value)} />
      </div>
    </div>
  );
}

export default App;
