type LearnStylingProps = {
  styles: React.CSSProperties;
};

export const LearnStyling = (props: LearnStylingProps) => {
  return (
    <div>
      <h1>LearnStyling</h1>
      <div style={props.styles}>LearnStyling</div>
    </div>
  );
};
