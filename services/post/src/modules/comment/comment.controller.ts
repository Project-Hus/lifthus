import { Body, Controller, Post, Req, UseGuards } from '@nestjs/common';
import { CommentService } from './comment.service';
import { UserGuard } from 'src/common/guards/post.guard';
import { CommentDto } from './comment.dto';
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
  wirtePost(
    @Req() req: Request,
    @Body() comment: CommentDto,
  ): Promise<Comment> {
    const uid: number = req.uid; // embedded user id
    return this.commentService.wirteComment({
      author: uid, // whatever the author is signed user.
      content: comment.content,
    });
  }

  /**
   * updates the post by the form data if the user is signed.
   * @param req
   * @param post
   * @returns
   */
  @UseGuards(UserGuard)
  @Put()
  updatePost(
    @Req() req: Request,
    @Body() post: PostDto,
  ):
    | Prisma.PrismaPromise<Prisma.BatchPayload>
    | { code: number; message: string } {
    const uid: number = req.uid;
    const aid: number = Number(post.author);
    if (uid !== aid) return { code: 403, message: 'Forbidden' };
    return this.postService.updatePost({
      author: aid,
      slug: '',
      content: post.content,
    });
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
    @Param('pid') pid: string,
  ): Prisma.PrismaPromise<Prisma.BatchPayload> {
    const uid: number = req.uid;
    return this.postService.deletePost(uid, Number(pid));
  }

  /**
   * likes the post by the pid in the body if the user is signed.
   * @param req
   * @param post
   * @returns
   */
  @UseGuards(UserGuard)
  @Post('/like')
  likePost(@Req() req: Request, @Body() post: any): any {
    const uid: number = req.uid;
    return this.postService.likePost(uid, post);
  }

  /**
   * unlikes the post by the pid in the body if the user is signed.
   * @param req
   * @param post
   * @returns
   */
  @UseGuards(UserGuard)
  @Post('/unlike')
  unlikePost(@Req() req: Request, @Body() post: any): any {
    const uid: number = req.uid;
    return this.postService.unlikePost(uid, post);
  }
}
