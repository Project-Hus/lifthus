import {
  Body,
  Controller,
  Delete,
  Inject,
  Param,
  Post,
  Put,
  UseGuards,
} from '@nestjs/common';
import { CommentService } from './comment.service';
import { UserGuard } from 'src/shared/guards/post.guard';
import {
  CreateCommentRequestDto,
  CreateCommentServiceDto,
  UpdateCommentRequestDto,
  UpdateCommentServiceDto,
} from 'src/dto/inbound/comment.dto';
import { Uid } from 'src/shared/decorators/authParam.decorator';
import { CommentDto } from 'src/dto/outbound/comment.dto';

@Controller('/post/comment')
export class CommentController {
  constructor(
    @Inject(CommentService) private readonly commentService: CommentService,
  ) {}

  /**
   * generates new post by the form data if the user is signed.
   * @param req
   * @returns
   */
  @UseGuards(UserGuard)
  @Post()
  createComment(
    @Uid() clientId: bigint,
    @Body() commentForm: CreateCommentRequestDto,
  ): Promise<CommentDto> {
    const comment = new CreateCommentServiceDto(commentForm);
    return this.commentService.createComment({ clientId, comment });
  }

  /**
   * updates the post by the form data if the user is signed.
   * @param req
   * @param post
   * @returns
   */
  @UseGuards(UserGuard)
  @Put()
  updateComment(
    @Uid() clientId,
    @Body() updatesForm: UpdateCommentRequestDto,
  ): Promise<CommentDto> {
    const updates = new UpdateCommentServiceDto(updatesForm);
    return this.commentService.updateComment({ clientId, updates });
  }

  /**
   * deletes the post by the pid in the body if the user is signed.
   * @param req
   * @param pid
   * @returns
   */
  @UseGuards(UserGuard)
  @Delete('/:cid')
  deleteComment(
    @Uid() clientId: bigint,
    @Param('cid') cid: string,
  ): Promise<CommentDto> {
    return this.commentService.deleteComment({ clientId, cid: BigInt(cid) });
  }
}
