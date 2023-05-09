import { Injectable, CanActivate, ExecutionContext } from '@nestjs/common';
import { Observable } from 'rxjs';

// only signed user can access with req.uid
@Injectable()
export class UserGuard implements CanActivate {
  canActivate(
    context: ExecutionContext,
  ): boolean | Promise<boolean> | Observable<boolean> {
    const uid = context.switchToHttp().getRequest().uid;
    if (!uid && uid != 0) {
      return false;
    }
    return true;
  }
}
