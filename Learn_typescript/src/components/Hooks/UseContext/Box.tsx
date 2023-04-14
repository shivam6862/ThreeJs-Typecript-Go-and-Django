import { useContext } from "react";
import { ThemeContext } from "./UseContextHook";

export const Box = () => {
  const theme = useContext(ThemeContext);
  return (
    <div>
      <h1>UseContextHook</h1>
      <div
        style={{
          backgroundColor: theme.primary.main,
          color: theme.primary.text,
        }}
      >
        UseContextHook example
      </div>
    </div>
  );
};
