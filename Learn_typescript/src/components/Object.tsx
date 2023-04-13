// type ObjectProps = {
//   name: {
//     first: string;
//     last: string;
//   };
// };

import { ObjectProps } from "../type/Object";

export const Object = (props: ObjectProps) => {
  return (
    <div>
      <h1>Object</h1>
      <div>
        {props.name.first} {props.name.last}
      </div>
    </div>
  );
};
