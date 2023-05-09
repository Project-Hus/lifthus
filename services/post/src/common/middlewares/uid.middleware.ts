import { Injectable, NestMiddleware } from '@nestjs/common';
import { Request, Response, NextFunction } from 'express';

@Injectable()
export class UidMiddleware implements NestMiddleware {
  use(req: Request, res: Response, next: NextFunction) {
    if (req.cookies['lifthus_st']) {
      // set uid to req with cookie lifthus_st
      req.uid = req.cookies['lifthus_st'];
    }
    next();
  }
}
