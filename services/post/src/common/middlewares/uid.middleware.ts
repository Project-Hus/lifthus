import { Injectable, NestMiddleware } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { Request, Response, NextFunction } from 'express';
import { LifthusSessionJWTPayload } from '../types/session';

@Injectable()
export class UidMiddleware implements NestMiddleware {
  constructor(private jwtService: JwtService) {}

  use(req: Request, res: Response, next: NextFunction) {
    if (req.cookies['lifthus_st']) {
      // set uid to req with cookie lifthus_st
      const lstSigned = req.cookies['lifthus_st'];
      const lst = this.jwtService.verify<LifthusSessionJWTPayload>(lstSigned);

      req.uid = req.cookies['lifthus_st'];
    }
    next();
  }
}
