import "./App.css";
import { Basic } from "./components/Basic";
import { Object } from "./components/Object";
import { ArrayOfObject } from "./components/ArrayOfObjec";
import { Status } from "./components/Status";
import { Children } from "./components/Children";
import { OscarChildren } from "./components/OscarChildren";
import { Click } from "./components/Click";
import { ChangeInput } from "./components/ChangeInput";
import { LearnStyling } from "./components/LearnStyling";

import { UseStateHook } from "./components/Hooks/UseStateHook";
import { UseReducerHook } from "./components/Hooks/UseReducerHook";
import { ThemeContextProvider } from "./components/Hooks/UseContext/UseContextHook";
import { Box } from "./components/Hooks/UseContext/Box";
import { UserContextProvider } from "./components/Hooks/UseContext/UserContext";
import { User } from "./components/Hooks/UseContext/User";
import { UseRefHook } from "./components/Hooks/UseRefHook";
import { Couter } from "./components/Hooks/ClassWay";

import { Private } from "./components/ComponentAsProps/Private";
import { Profile } from "./components/ComponentAsProps/Profile";

import { List } from "./components/generics/List";
import { RandamNumber } from "./components/restrication/RandamNumber";
import { TemplateLiterals } from "./components/templateLiterals/TemplateLiterals";

import { CustomButton } from "./components/wrappingHtmlElement/Button";
import { CustomCompoentCall } from "./components/wrappingHtmlElement/CustomCompoent";
import { Polymorphic } from "./components/polymorphic/Polymorphic";

import { nameList } from "./Data/nameList";
import { personName } from "./Data/personName";

function App() {
  return (
    <div className="App">
      <h1 className="heading">Typecript</h1>
      <div>
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
        <ChangeInput
          value=""
          handleChange={(event) => console.log(event.target.value)}
        />
        <LearnStyling styles={{ border: "2px solid green", padding: "1rem" }} />
      </div>
      <div>
        <UseStateHook />
        <UseReducerHook />
        <ThemeContextProvider>
          <Box />
        </ThemeContextProvider>
        <UserContextProvider>
          <User />
        </UserContextProvider>
        <UseRefHook />
        <Couter message="The count value is : " />
      </div>
      <div>
        <Private isLoggedIn={true} component={Profile} />
        {/* <List items={["Shyam" , "pushPendra" ,"Shashank" , "Sarvagye"]} onClick={(item)=>console.log(item)}/> */}
        <List items={nameList} onClick={(item) => console.log(item)} />
        <RandamNumber value={5} isPositive />
        <TemplateLiterals position="center-top" />
        <CustomButton variant="primary" onClick={() => console.log("Clicked!")}>
          Primary Button
        </CustomButton>
        <CustomCompoentCall name="Shivam" messageCount={20} isLoggedIn={true} />
        <Polymorphic as="label" size="lg" htmlFor="someId">
          Label
        </Polymorphic>
        <Polymorphic as="h1" size="lg">
          Hello World
        </Polymorphic>
      </div>
    </div>
  );
}

export default App;
