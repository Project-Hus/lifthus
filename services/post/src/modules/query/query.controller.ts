import { Controller, Get, Res, Req, UseGuards } from '@nestjs/common';
import { Request, Response } from 'express';
import { UserGuard } from 'src/common/guards/post.guard';
import { PostQueryService } from './post.query.service';
import { CommentQueryService } from './comment.query.service';
import { Post, Prisma } from '@prisma/client';

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

  @Get('/post/:uid')
  async getPosts(@Req() req: Request, @Res() res: Response): Promise<Post[]> {
    return Promise.reject('Not Implemented');
  }
}
