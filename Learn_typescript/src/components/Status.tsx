type StatusProps = {
  //   status: string;
  status: "Loading" | "success" | "error";
};
export const Status = (props: StatusProps) => {
  let message = "";
  if (props.status == "Loading") message = "Loading...";
  else if (props.status == "success") message = "Data fetched Sucessfully!";
  else if (props.status == "error") message = "Error Fetching data";
  return (
    <div>
      <h1>Status</h1>
      <div>
        <h2>{message}</h2>
      </div>
    </div>
  );
};
