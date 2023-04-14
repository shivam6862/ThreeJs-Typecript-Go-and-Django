import React from "react";
import { Basic } from "../Basic";

export const CustomCompoentCall = (props: React.ComponentProps<typeof Basic>) => {
  return (
    <div>
      <h1>Calling other Component Data</h1>
      <div>{props.name}</div>
      <div>{props.messageCount}</div>
      <div>{props.isLoggedIn}</div>
    </div>
  );
};
