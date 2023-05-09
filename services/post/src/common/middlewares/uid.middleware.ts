import { Injectable, NestMiddleware, Logger } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { Request, Response, NextFunction } from 'express';
import { LifthusSessionJWTPayload } from '../types/session';

@Injectable()
export class UidMiddleware implements NestMiddleware {
  constructor(private readonly jwtService: JwtService) {}
  async use(req: Request, res: Response, next: NextFunction) {
    if (req.cookies['lifthus_st']) {
      // set uid to req with cookie lifthus_st
      const lstSigned = req.cookies['lifthus_st'];
      try {
        const lst = await this.jwtService.verifyAsync<LifthusSessionJWTPayload>(
          lstSigned,
        );
        // if lifthus session token is valid, set uid to req
        // lst.uid to number
        req.uid = parseInt(lst.uid);
      } catch (e) {
        // else just nothing
      }
    }
    next();
  }
}
