package templates

const ComponentTest = `/* eslint-disable react/jsx-filename-extension */

import React from "react";
import { render } from "%s";

import %s from ".";

describe("<%s />", () => {
  it("snapshot of %s", () => {
    const tree = render(<%s text="test text" />);
    expect(tree).toMatchSnapshot();
  });
});
`
