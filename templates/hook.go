package templates

const Hook = `import { useCallback, useState } from "react";

const use%sData = () => {
  const [count, setCount] = useState(0);

  const incrementCount = useCallback(() => {
    setCount(count + 1);
  }, [count, setCount]);

  return { incrementCount, count };
};

export default use%sData;
`
