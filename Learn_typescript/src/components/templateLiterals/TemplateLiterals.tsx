type HorizontalPosition = "left" | "center" | "right";
type VerticalPosition = "top" | "center" | "bottom";

type TemplatePosition = {
  position:
    | Exclude<`${HorizontalPosition}-${VerticalPosition}`, "center-center">
    | "center";
};

export const TemplateLiterals = ({ position }: TemplatePosition) => {
  return (
    <div>
      <h1>Template Literals</h1>
      <div>Notification Position : {position}</div>
    </div>
  );
};
