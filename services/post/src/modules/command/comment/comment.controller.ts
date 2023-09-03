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
import { CreateCommentDto, UpdateCommentDto } from './comment.dto';
import { Comment, CommentLike, Prisma } from '@prisma/client';

import { Request } from 'express';

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
    @Req() req: Request,
    @Body() comment: CreateCommentDto,
  ): Promise<Comment> {
    const uid: number = req.uid; // embedded user id
    if (comment.author !== uid) throw new ForbiddenException();
    return this.commentService.createComment(comment);
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
    @Req() req: Request,
    @Body() comment: UpdateCommentDto,
  ): Prisma.PrismaPromise<Prisma.BatchPayload> {
    if (req.uid !== comment.author) throw new ForbiddenException();
    return this.commentService.updateComment(comment);
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
      aid: req.uid,
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
    return this.commentService.likeComment({ uid: req.uid, cid: Number(cid) });
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
