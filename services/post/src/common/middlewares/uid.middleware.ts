import { JwtService } from '@nestjs/jwt';
import { LifthusSessionJWTPayload } from '../types/session';
import { Injectable, Logger, NestMiddleware } from '@nestjs/common';
import { Request, Response, NextFunction } from 'express';

// class jsdoc template
/**
 * @class UidMiddleware
 * @implements {NestMiddleware}
 * @description
 * checks lifthus_st in cookie and sets uid to req if it is.
 */
@Injectable()
export class UidMiddleware implements NestMiddleware {
  constructor(private readonly jwtService: JwtService) {}
  async use(req: Request, res: Response, next: NextFunction) {
    Logger.log(req.cookies);
    if (req.cookies['lifthus_st']) {
      // set uid to req with cookie lifthus_st
      const lstSigned = req.cookies['lifthus_st'];
      Logger.log(`lstSigned: ${lstSigned}`);
      try {
        const lst = await this.jwtService.verifyAsync<LifthusSessionJWTPayload>(
          lstSigned,
        );
        // if lifthus session token is valid, set uid to req
        // lst.uid to number
        if (lst.uid) {
          Logger.log(`uid: ${lst.uid}`);
          req.uid = parseInt(lst.uid);
        }
      } catch (e) {
        Logger.log('JWT VAL FAILS:', e);
      }
    }
    next();
  }
}
