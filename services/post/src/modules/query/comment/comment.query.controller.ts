import { Controller, Get, Query, BadRequestException } from '@nestjs/common';
import { CommentQueryService } from './comment.query.service';
import { Comment } from '@prisma/client';

@Controller('/post/query/comment')
export class CommentQueryController {
  constructor(private readonly commentQueryService: CommentQueryService) {}

  @Get()
  getComments(
    @Query('pid') pidStr: string,
    @Query('skip') skipStr: string,
  ): Promise<Comment[]> {
    const pid = Number(pidStr);
    if (isNaN(pid)) {
      throw new BadRequestException();
    }
    const skip = Number(skipStr) || 0;
    return this.commentQueryService.getComments({ pid, skip });
  }
}
