import { useRef, useEffect } from "react";

export const UseRefHook = () => {
  const InputRef = useRef<HTMLInputElement>(null!);
  const interValueRef = useRef<number | null>(null);

  useEffect(() => {
    InputRef.current.focus();
  }, []);
  
  return (
    <div>
      <h1>UseRef</h1>
      <div>
        <input type="text" ref={InputRef} />
      </div>
    </div>
  );
};
