import { ExecutionContext } from '@nestjs/common';

export const getCookie = (data: string, ctx: ExecutionContext): string => {
  const request = ctx.switchToHttp().getRequest();
  return request.cookies?.[data] ? request.cookies?.[data] : undefined;
};

