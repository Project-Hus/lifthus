export {};

declare global {
  interface BigInt {
    toJSON(): number;
  }
}
