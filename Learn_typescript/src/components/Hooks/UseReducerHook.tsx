import { useReducer } from "react";

type CounterSate = {
  count: number;
};

type UdateAction = {
  type: "increment" | "decrement";
  payload: number;
};

type ResetAction = {
  type: "reset";
};

type CouunterAction = UdateAction | ResetAction;

const initialState = { count: 0 };

function reducer(state: CounterSate, action: CouunterAction) {
  switch (action.type) {
    case "increment":
      return { count: state.count + action.payload };
    case "decrement":
      return { count: state.count - action.payload };
    case "reset":
      return initialState;
    default:
      return state;
  }
}

export const UseReducerHook = () => {
  const [state, dispatch] = useReducer(reducer, initialState);
  return (
    <div>
      <h1>UseReducer</h1>
      <div>
        <div> Count :{state.count}</div>
        <button onClick={() => dispatch({ type: "increment", payload: 10 })}>
          Increment
        </button>
        <button onClick={() => dispatch({ type: "decrement", payload: 10 })}>
          Decrement
        </button>
        <button onClick={() => dispatch({ type: "reset" })}>Reset</button>
      </div>
    </div>
  );
};

// import { useReducer } from "react";

// type CounterSate = {
//   count: number;
// };

// type CouunterAction = {
//   type: string;
//   payload: number;
// };

// const initialState = { count: 0 };

// function reducer(state: CounterSate, action: CouunterAction) {
//   switch (action.type) {
//     case "increment":
//       return { count: state.count + action.payload };
//     case "decrement":
//       return { count: state.count - action.payload };
//     default:
//       return state;
//   }
// }

// export const UseReducerHook = () => {
//   const [state, dispatch] = useReducer(reducer, initialState);
//   return (
//     <div>
//       <h1>UseReducer</h1>
//       <div>
//         <div> Count :{state.count}</div>
//         <button onClick={() => dispatch({ type: "increment", payload: 10 })}>
//           Increment
//         </button>
//         <button onClick={() => dispatch({ type: "decrement", payload: 10 })}>
//           Decrement
//         </button>
//       </div>
//     </div>
//   );
// };
