type ListProps<T> = {
  items: T[];
  onClick: (value: T) => void;
};

export const List = <T extends { id: number; first: string; last: string }>({
  items,
  onClick,
}: ListProps<T>) => {
  // export const List = <T extends {}>({ items, onClick }: ListProps<T>) => {

  // export const List = <T extends React.ReactNode>({
  //   items,
  //   onClick,
  // }: ListProps<T>) => {

  // export const List = <T extends string | number>({
  //   items,
  //   onClick,
  // }: ListProps<T>) => {

  return (
    <div>
      <h1>Generics type</h1>
      <div>
        {items.map((item, index) => {
          return (
            <div key={item.id} onClick={() => onClick(item)}>
              {item.first} {item.last}
            </div>
          );
        })}
      </div>
    </div>
  );
};
