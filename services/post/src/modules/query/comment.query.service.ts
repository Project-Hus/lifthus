import { Injectable } from '@nestjs/common';
import { Comment } from '@prisma/client';

@Injectable()
export class CommentQueryService {
  /**
   * Get comments of specified post.
   * @param param0
   * @returns
   */
  getComments({
    pid,
    skip,
  }: {
    pid: number;
    skip: number;
  }): Promise<Comment[]> {
    return Promise.reject('Not implemented');
  }
}
