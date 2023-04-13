type ClickProps = {
  //   handleClick: () => void;
  handleClick: (
    event: React.MouseEvent<HTMLButtonElement>,
    id: number
  ) => void;
};
export const Click = (props: ClickProps) => {
  return (
    <div>
      <h1>Click Event</h1>
      <div>
        <button onClick={(event)=>props.handleClick(event,1)}>Click me</button>
      </div>
    </div>
  );
};
