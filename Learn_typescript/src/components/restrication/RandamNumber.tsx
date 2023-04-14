type RandamNumberType = {
  value: number;
};

type PositiveNumber = RandamNumberType & {
  isPositive: boolean;
  isNegative?: never;
  isZero?: never;
};

type NegativeNumber = RandamNumberType & {
  isPositive?: never;
  isNegative: boolean;
  isZero?: never;
};

type ZeroNumber = RandamNumberType & {
  isPositive?: never;
  isNegative?: never;
  isZero: boolean;
};

type RandamNumberProps = PositiveNumber | NegativeNumber | ZeroNumber;

export const RandamNumber = ({
  value,
  isPositive,
  isNegative,
  isZero,
}: RandamNumberProps) => {
  return (
    <div>
      <h1>Randam Number</h1>
      <div>
        {value} {isPositive && "positive"} {isNegative && "negative"}
        {isZero && "zero"}
      </div>
    </div>
  );
};
