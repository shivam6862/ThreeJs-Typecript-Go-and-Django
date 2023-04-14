import { useState } from "react";

type Authuser = {
  name: string;
  email: string;
};

export const UseStateHook = () => {
  const [user, setUser] = useState<Authuser | null>(null);
  const handleLogin = () => {
    setUser({
      name: "shivam",
      email: "Shivam@gmail.com",
    });
  };
  const handleLogOut = () => {
    setUser(null);
  };
  return (
    <div>
      <h1>UseState</h1>
      <div>
        <button onClick={handleLogin}>Login</button>
        <button onClick={handleLogOut}>Logout</button>
        <div>User name is {user?.name}</div>
        <div>User email is {user?.email}</div>
      </div>
    </div>
  );
};

// import { useState } from "react";

// type Authuser = {
//   name: string;
//   email: string;
// };

// export const UseStateHook = () => {
//   const [user, setUser] = useState<Authuser>({}as Authuser);
//   const handleLogin = () => {
//     setUser({
//       name: "shivam",
//       email: "Shivam@gmail.com",
//     });
//   };
//   return (
//     <div>
//       <h1>UseState</h1>
//       <div>
//         <button onClick={handleLogin}>Login</button>
//         <div>User name is {user.name}</div>
//         <div>User email is {user.email}</div>
//       </div>
//     </div>
//   );
// };
