import { createParamDecorator, ExecutionContext } from '@nestjs/common';

export const Uid = createParamDecorator((data: any, ctx: ExecutionContext) => {
  const request = ctx.switchToHttp().getRequest();
  const uid: BigInt = request.uid;
  return uid || undefined;
});

export const Exp = createParamDecorator((data: any, ctx: ExecutionContext) => {
  const request = ctx.switchToHttp().getRequest();
  const exp: boolean = request.exp;
  return exp;
});
