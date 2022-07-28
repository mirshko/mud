/**
 * Pads start of a hex string with 0 to create a 160 bit hex string
 * which can be used as an Ethereum address
 * @param input Hex string
 * @returns 160 bit hex string
 */
export function toEthAddress(input: string) {
  // Cut off 0x prefix
  if (input.substring(0, 2) == "0x") input = input.substring(2);
  // Pad start with 0 to get length of 160 bit
  input = input.padStart(40, "0");
  // Prefix with 0x
  return `0x${input}`;
}

export function extractEncodedArguments(input: string) {
  // Cutting off the first 4 bytes, which represent the function selector
  if (input[0] !== "0" && input[1] !== "x") throw new Error("Invalid hex string");
  return "0x" + input.substring(10);
}
