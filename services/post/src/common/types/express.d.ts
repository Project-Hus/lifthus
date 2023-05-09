// to make the file a module and avoid the TypeScript error
export {};

declare global {
  namespace Express {
    type Uid = number;
    interface Request {
      uid?: Uid;
    }
  }
}
