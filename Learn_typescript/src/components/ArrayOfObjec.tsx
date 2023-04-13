// type ArrayOfObjectProps = {
//   nameList: {
//     first: string;
//     last: string;
//   }[];
// };

import { Name } from "../type/Object";

type ArrayOfObjectProps = {
  nameList: Name[];
};

export const ArrayOfObject = (props: ArrayOfObjectProps) => {
  return (
    <div>
      <h1>Array of Object</h1>
      <div>
        {props.nameList.map((name, index) => {
          return (
            <div key={index}>
              {name.first}
              {name.last}
            </div>
          );
        })}
      </div>
    </div>
  );
};
