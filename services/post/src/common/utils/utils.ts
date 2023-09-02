import { ExecutionContext } from '@nestjs/common';

import crypto from 'crypto';

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
export const slugify = (data: string): string => {
  return encodeURIComponent(data + crypto.randomBytes(8).toString('hex'));
};

export const stringifyAny = (data: any): string => {
  return JSON.stringify(data);
};
