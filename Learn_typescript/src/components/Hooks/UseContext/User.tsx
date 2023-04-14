import { useContext } from "react";
import { UserContext } from "./UserContext";

export const User = () => {
  const userContext = useContext(UserContext);

  const handleLogin = () => {
    userContext.setUser({
      name: "shivam",
      email: "Shivam@gmail.com",
    });
  };
  const handleLogOut = () => {
    userContext.setUser(null);
  };
  return (
    <div>
      <h1>User</h1>
      <div>
        <button onClick={handleLogin}>Login</button>
        <button onClick={handleLogOut}>Logout</button>
        <div>User name is {userContext.user?.name}</div>
        <div>User email is {userContext.user?.email}</div>
      </div>
    </div>
  );
};
