// to make the file a module and avoid the TypeScript error
export {};

declare global {
  namespace Express {
    type Uid = string;
    interface Request {
      uid?: Uid;
    }
  }
}
