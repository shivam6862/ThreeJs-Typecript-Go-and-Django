type OscarChildrenProps = {
  children: React.ReactNode;
};

export const OscarChildren = (props: OscarChildrenProps) => {
  return (
    <div>
      <h1>OscarChildren</h1>
      <div>{props.children}</div>
    </div>
  );
};
