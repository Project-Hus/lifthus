import { JwtService } from '@nestjs/jwt';
import { LifthusSessionJWTPayload } from '../types/session';
import { Injectable, NestMiddleware } from '@nestjs/common';
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
    // get the origin of request
    const origin = req.headers['origin'];
    if (origin === 'http://localhost:3000') {
      const authHeader = req.headers['authorization'];
      if (
        authHeader &&
        authHeader.startsWith('Bearer ') &&
        authHeader.length > 7
      ) {
        const lstSigned = authHeader.slice(7);
        try {
          const lst =
            await this.jwtService.verifyAsync<LifthusSessionJWTPayload>(
              lstSigned,
            );
          if (lst.uid) {
            req.uid = parseInt(lst.uid);
          }
        } catch (e) {}
      }
    } else {
      if (req.cookies['lifthus_st']) {
        // set uid to req with cookie lifthus_st
        const lstSigned = req.cookies['lifthus_st'];
        try {
          const lst =
            await this.jwtService.verifyAsync<LifthusSessionJWTPayload>(
              lstSigned,
            );
          req.exp = false;
          // if lifthus session token is valid, set uid to req
          // lst.uid to number
          if (lst.uid) {
            req.uid = parseInt(lst.uid);
          }
        } catch (e: any) {
          if (e.name === 'TokenExpiredError') {
            req.exp = true;
          }
        }
      }
    }
    next();
  }
  async useV2(req: Request, res: Response, next: NextFunction) {
    // get Authorization header
    const authHeader = req.headers['authorization'];
    if (authHeader && authHeader.startsWith('Bearer ')) {
      const lstSigned = authHeader.slice(7);
      try {
        const lst = await this.jwtService.verifyAsync<LifthusSessionJWTPayload>(
          lstSigned,
        );
        if (lst.uid) {
          req.uid = parseInt(lst.uid);
        }
      } catch (e: any) {
        if (e.name === 'TokenExpiredError') {
          req.exp = true;
        }
      }
    }
    next();
  }
}
