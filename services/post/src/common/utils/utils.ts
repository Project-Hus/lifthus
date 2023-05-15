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
export const slugify = (data: string): string => {
  const specialChars = '@#$%&*?,'; //일부 특수문자 utf-8 인코딩
  let slug = data.replace(/[\s]+/g, '-'); // 공백을 '-'로 치환

  for (let i = 0; i < specialChars.length; i++) {
    const char = specialChars[i];
    const encodedChar = encodeURIComponent(char);
    slug = slug.split(char).join(encodedChar);
  }

  return slug;
};
