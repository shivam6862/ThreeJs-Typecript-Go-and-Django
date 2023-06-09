type ChangeInputProps = {
  value: string;
  handleChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
};

export const ChangeInput = ({ value, handleChange }: ChangeInputProps) => {
  return (
    <div>
      <h1>Change Input</h1>
      <div>
        <input type="text" value={value} onChange={handleChange} />
      </div>
    </div>
  );
};

// type ChangeInputProps = {
//   value: string;
//   handleChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
// };

// export const ChangeInput = (props: ChangeInputProps) => {
//   return (
//     <div>
//       <h1>Change Input</h1>
//       <div>
//         <input type="text" value={props.value} onChange={props.handleChange} />
//       </div>
//     </div>
//   );
// };
