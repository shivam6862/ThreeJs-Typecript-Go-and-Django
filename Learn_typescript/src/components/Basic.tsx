type BasicProps = {
  name: string;
  // messageCount: number;
  messageCount?: number;
  isLoggedIn: boolean;
};

export const Basic = (props: BasicProps) => {
  const { messageCount = 0 } = props;
  return (
    <div>
      <h1>Basic Problem</h1>
      <div>
        {props.isLoggedIn ? <>1</> : <>2</>}
        <h2>
          Welcome to Learn TypeScript {props.name} {props.messageCount}
        </h2>
        <h3>
          Welcome to Learn TypeScript {props.name} {messageCount}
        </h3>
      </div>
    </div>
  );
};
