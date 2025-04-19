package templates

const Story = `import React from "react";
import %s from ".";

export default {
  title: "Tests/%s",
  component: %s,
  args: {
    text: "Test Component",
  },
};

const Template = (args) => <%s {...args} />;

export const Default = Template.bind({});
`
