package templates

const FunctionTest = `import %s from ".";

describe("%s", () => {
  test("tests function", () => {
    expect(%s("hello")).toBe("hello");
  });
});
`
