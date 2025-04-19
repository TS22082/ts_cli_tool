package templates

const Component = `import React from "react";
import PropTypes from "prop-types";
import use%sData from "./use%sData";

/**
 *
 * @param props
 * @returns {Element}
 * @constructor
 */
const %s = ({ text }) => {
  const { incrementCount, count } = use%sData();

  return (
    <div>
      <p>{text}</p>
      <p>{count}</p>
      <button type="button" onClick={incrementCount}>
        Increment
      </button>
    </div>
  );
};

%s.propTypes = {
  text: PropTypes.string.isRequired,
};

export default %s;
`
