export const convertToInt = (val: any) => {
  const parsed = parseInt(typeof val === "string" ? val : "", 10);

  if (Number.isInteger(parsed)) return parsed;

  return undefined;
};

/** allowed hosts */
export const hosts = ["githubusercontent", "github"].map(val => `gist.${val}.com`);
