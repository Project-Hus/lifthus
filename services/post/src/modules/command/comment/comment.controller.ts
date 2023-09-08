import {
  Body,
  Controller,
  Delete,
  ForbiddenException,
  Inject,
  Logger,
  Param,
  Post,
  Put,
  Req,
  UseGuards,
} from '@nestjs/common';
import { CommentService } from './comment.service';
import { UserGuard } from 'src/common/guards/post.guard';
import { UpdateCommentDto } from './comment.dto';
import { Prisma } from '@prisma/client';

import { Request } from 'express';
import {
  CreateCommentRequestDto,
  CreateCommentServiceDto,
  UpdateCommentRequestDto,
  UpdateCommentServiceDto,
} from 'src/dto/inbound/comment.dto';
import { Uid } from 'src/common/decorators/authParam.decorator';
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
    @Req() req: Request,
    @Param('cid') cid: any,
  ): Prisma.PrismaPromise<Prisma.BatchPayload> {
    return this.commentService.deleteComment({
      cid: Number(cid),
      aid: Number(req.uid),
    });
  }

  /**
   * likes the post by the pid in the body if the user is signed.
   * @param req
   * @param post
   * @returns
   */
  @UseGuards(UserGuard)
  @Post('/like/:cid')
  likeComment(@Req() req: Request, @Param('cid') cid: any): Promise<number> {
    return this.commentService.likeComment({
      uid: Number(req.uid),
      cid: Number(cid),
    });
  }

  // /**
  //  * unlikes the post by the pid in the body if the user is signed.
  //  * @param req
  //  * @param post
  //  * @returns
  //  */
  // @UseGuards(UserGuard)
  // @Post('/unlike/:cid')
  // unlikeComment(
  //   @Req() req: Request,
  //   @Param('cid') cid: number,
  // ): Promise<[CommentLike, Comment]> {
  //   return this.commentService.unlikeComment(req.uid, { id: cid });
  // }
}
