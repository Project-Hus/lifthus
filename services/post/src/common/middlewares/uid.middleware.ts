import { JwtService } from '@nestjs/jwt';
import { LifthusSessionJWTPayload } from 'src/common/types/session';
import { Inject, Injectable, Logger, NestMiddleware } from '@nestjs/common';
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
  constructor(@Inject(JwtService) private readonly jwtService: JwtService) {}
  async use(req: Request, res: Response, next: NextFunction) {
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
