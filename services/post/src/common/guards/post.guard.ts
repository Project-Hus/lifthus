import { Injectable, CanActivate, ExecutionContext } from '@nestjs/common';
import { Observable } from 'rxjs';
import { getCookie } from '../util/util';

@Injectable()
export class UserGuard implements CanActivate {
  canActivate(
    context: ExecutionContext,
  ): boolean | Promise<boolean> | Observable<boolean> {
    const lst = getCookie('lifthus_st', context);
    if (!lst) {
      return false;
    }

    return true;
  }
}
