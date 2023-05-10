import { JwtService } from '@nestjs/jwt';
import { Request, Response, NextFunction } from 'express';
import { LifthusSessionJWTPayload } from '../types/session';
import { Logger } from '@nestjs/common';

/**
 * a middleware to set uid to req if user is signed.
 * @param req
 * @param res
 * @param next
 */
export async function uidMiddleware(
  req: Request,
  res: Response,
  next: NextFunction,
) {
  const jwtService = new JwtService();
  Logger.log(req.cookies);
  if (req.cookies['lifthus_st']) {
    // set uid to req with cookie lifthus_st
    const lstSigned = req.cookies['lifthus_st'];
    Logger.log(`lstSigned: ${lstSigned}`);
    try {
      const lst = await jwtService.verifyAsync<LifthusSessionJWTPayload>(
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
