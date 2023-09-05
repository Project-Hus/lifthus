import {
  Controller,
  Get,
  Query,
  BadRequestException,
  Inject,
} from '@nestjs/common';
import { CommentQueryService } from 'src/modules/query/comment/comment.query.service';
import { CommentDto } from 'src/dto/outbound/comment.dto';
import { Uid } from 'src/common/decorators/authParam.decorator';

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
    @Uid() client: BigInt | undefined,
  ): Promise<CommentDto[]> {
    const skip = Number(skipStr) || 0;
    return this.commentQueryService.getComments({ pid: pidStr, skip, client });
  }
}
