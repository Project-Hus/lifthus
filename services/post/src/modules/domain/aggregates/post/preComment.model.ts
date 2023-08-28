import { Injectable } from '@nestjs/common';

export type CreatePreCommentInput = {};
@Injectable()
export class PreComment {
  private author: bigint;

  private comments: Comment[] = [];

  constructor() {}
}
