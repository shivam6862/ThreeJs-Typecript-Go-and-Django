type ChildrenProps = {
  children: string;
};

export const Children = (props: ChildrenProps) => {
  return (
    <div>
      <h1>Children</h1>
      <div>{props.children}</div>
    </div>
  );
};
