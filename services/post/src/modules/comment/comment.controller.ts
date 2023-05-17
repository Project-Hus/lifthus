import {
  Body,
  Controller,
  Delete,
  Post,
  Put,
  Req,
  UseGuards,
} from '@nestjs/common';
import { CommentService } from './comment.service';
import { UserGuard } from 'src/common/guards/post.guard';
import { CreateCommentDto, UpdateCommentDto } from './comment.dto';
import { Comment, Prisma } from '@prisma/client';
import { Request } from 'express';

@Controller('/post/comment')
export class CommentController {
  constructor(private readonly commentService: CommentService) {}

  /**
   * generates new post by the form data if the user is signed.
   * @param req
   * @returns
   */
  @UseGuards(UserGuard)
  @Post()
  wirteComment(
    @Req() req: Request,
    @Body() comment: CreateCommentDto,
  ): Promise<Comment> {
    const uid: number = req.uid; // embedded user id
    // whatever, this endpoint is for currently signed user.
    // it would be better to check if the author is signed user.
    // but for now, there is no logic that deals with the uid in frontend.
    // so just embedding the uid to the author field.
    comment.author = uid;
    const commentInput: Prisma.CommentCreateInput = {
      author: uid, // whatever the author is signed user.
      content: comment.content,
      post: { connect: { id: comment.postId } },
    };
    if (comment.parentId) {
      commentInput.parent = { connect: { id: comment.parentId } };
    }
    return this.commentService.wirteComment(commentInput);
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
  ):
    | Prisma.PrismaPromise<Prisma.BatchPayload>
    | { code: number; message: string } {
    if (req.uid !== comment.author) return { code: 403, message: 'Forbidden' };
    return this.commentService.updateComment(comment);
  }

  /**
   * deletes the post by the pid in the body if the user is signed.
   * @param req
   * @param pid
   * @returns
   */
  @UseGuards(UserGuard)
  @Delete()
  deletePost(
    @Req() req: Request,
    @Body('pid') pid: number,
  ): Prisma.PrismaPromise<Prisma.BatchPayload> {
    const uid: number = req.uid;
    return this.commentService.deleteComment(uid, pid);
  }

  /**
   * likes the post by the pid in the body if the user is signed.
   * @param req
   * @param post
   * @returns
   */
  @UseGuards(UserGuard)
  @Post('/like')
  likePost(@Req() req: Request, @Body('pid') pid: number): Promise<Comment> {
    const uid: number = req.uid;
    return this.commentService.likeComment(uid, { id: pid });
  }

  /**
   * unlikes the post by the pid in the body if the user is signed.
   * @param req
   * @param post
   * @returns
   */
  @UseGuards(UserGuard)
  @Post('/unlike')
  unlikePost(@Req() req: Request, @Body('pid') pid: number): Promise<Comment> {
    const uid: number = req.uid;
    return this.commentService.unlikeComment(uid, { id: pid });
  }
}
