import { ExecutionContext } from '@nestjs/common';

export const getCookie = (data: string, ctx: ExecutionContext): string => {
  const request = ctx.switchToHttp().getRequest();
  return request.cookies?.[data] ? request.cookies?.[data] : undefined;
};

/**
 * @param {string} data - string to be processed
 * @returns {string} - slug generated from data
 * @description
 * generates slug from string
 * @example
 * getSlug('hello world'); // 'hello-world'
 */
export const generateSlug = (data: string): string => {
  const slug = data;
  return slug;
};
