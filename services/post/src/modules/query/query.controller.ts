import { Controller, Get, Res, Req, Param, Logger } from '@nestjs/common';
import { Request, Response } from 'express';
import { PostQueryService } from './post.query.service';
import { CommentQueryService } from './comment.query.service';
import { Post } from '@prisma/client';

@Controller('/post/query')
export class QueryController {
  constructor(
    private readonly postQueryService: PostQueryService,
    private readonly commentQueryService: CommentQueryService,
  ) {}

  @Get()
  getHello(): string {
    return this.postQueryService.getHello();
  }

  /**
   * gets user ID and skip number from url params and returns 10 posts from the skip number.
   * @param req
   * @param res
   * @param uid
   * @param skip
   * @returns {Post[]}
   */
  @Get('/post/:uid/:skip')
  getPosts(@Param('uid') uid: any, @Param('skip') skip: any): Promise<Post[]> {
    skip = skip || 0;
    return this.postQueryService.getUserPosts(Number(uid), Number(skip));
  }
}
