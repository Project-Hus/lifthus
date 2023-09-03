import {
  Controller,
  Get,
  Query,
  BadRequestException,
  Inject,
} from '@nestjs/common';
import { CommentQueryService } from 'src/modules/query/comment/comment.query.service';
import { Comment } from 'src/modules/domain/aggregates/comment/comment.model';

@Controller('/post/query/comment')
export class CommentQueryController {
  constructor(
    @Inject(CommentQueryService)
    private readonly commentQueryService: CommentQueryService,
  ) {}

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
