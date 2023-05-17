import { Controller, Get, Res, Req, Param } from '@nestjs/common';
import { Request, Response } from 'express';
import { UserGuard } from 'src/common/guards/post.guard';
import { PostQueryService } from './post.query.service';
import { CommentQueryService } from './comment.query.service';
import { Post, Prisma } from '@prisma/client';
import { PostQueryDto } from './post.query.dto';

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
  @Get('/post/:uid')
  @Get('/post/:uid/:skip')
  async getPosts(
    @Req() req: Request,
    @Res() res: Response,
    @Param('uid') uid: number,
    @Param('skip') skip: number,
  ): Promise<Post[]> {
    if (skip === undefined) {
      skip = 0;
    }
    return this.postQueryService.getUserPosts(uid, skip);
  }
}
