import React from "react";

type ButtonProps = {
  variant: "primary" | "secondary";
  children: string;
} & Omit<React.ComponentProps<"button">, "children">;

export const CustomButton = ({ variant, children, ...rest }: ButtonProps) => {
  return (
    <div>
      <h1>Wrapping Html element</h1>
      <button className={`class-with-${variant}`} {...rest}>
        {children}
      </button>
    </div>
  );
};

// type ButtonProps = {
//   variant: "primary" | "secondary";
// } & React.ComponentProps<"button">;

// export const CustomButton = ({ variant, children, ...rest }: ButtonProps) => {
//   return (
//     <div>
//       <h1>Wrapping Html element</h1>
//       <button className={`class-with-${variant}`} {...rest}>
//         {children}
//       </button>
//     </div>
//   );
// };
