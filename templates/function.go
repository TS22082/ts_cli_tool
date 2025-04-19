package templates

const Function = `/**
 * @function %s
 * @param {string} arg
 * @returns {string}
 */

const %s = (arg) => {
  return arg;
};

export default %s;
`
