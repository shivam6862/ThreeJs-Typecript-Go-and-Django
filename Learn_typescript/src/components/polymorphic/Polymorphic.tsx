import React, { Children } from "react";

type polymorphicOwnProps<E extends React.ElementType> = {
  size?: "sm" | "md" | "lg";
  color?: "primary" | "secondary";
  children: React.ReactNode;
  as?: E;
};

type polymorphicProps<E extends React.ElementType> = polymorphicOwnProps<E> &
  Omit<React.ComponentProps<E>, keyof polymorphicOwnProps<E>>;

export const Polymorphic = <E extends React.ElementType = "div">({
  size,
  color,
  children,
  as,
}: polymorphicProps<E>) => {
  const Component = as || "div";
  return (
    <div>
      <h1>polymorphic</h1>
      <Component className={`class-with-${color}`}>{children}</Component>
    </div>
  );
};

// import React, { Children } from "react";

// type polymorphicProps = {
//   size?: "sm" | "md" | "lg";
//   color?: "primary" | "secondary";
//   children: React.ReactNode;
//   as?: React.ElementType;
// };

// export const Polymorphic = ({
//   size,
//   color,
//   children,
//   as,
// }: polymorphicProps) => {
//   const Component = as || "div";
//   return (
//     <div>
//       <h1>polymorphic</h1>
//       <Component className={`class-with-${color}`}>{children}</Component>
//     </div>
//   );
// };
