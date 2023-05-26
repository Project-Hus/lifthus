import {
  Body,
  Controller,
  Delete,
  Param,
  Post,
  Put,
  Req,
  UseGuards,
} from '@nestjs/common';
import { UserGuard } from 'src/common/guards/post.guard';
import { Request } from 'express';
import { PostService } from './post.service';
import { Post as PPost, PostLike, Prisma } from '@prisma/client';
import { CreatePostDto, UpdatePostDto } from './post.dto';

/**
 * Mutation Controller
 * @description this controller is responsible for handling all mutation requests for posts.
 * @class MutationController
 */
@Controller('/post/post')
export class PostController {
  constructor(private readonly postService: PostService) {}

  /**
   * generates new post by the form data if the user is signed.
   * @param req
   * @returns
   */
  @UseGuards(UserGuard)
  @Post()
  wirtePost(@Req() req: Request, @Body() post: CreatePostDto): Promise<PPost> {
    const uid: number = req.uid; // embedded user id
    // whatever, this endpoint is for currently signed user.
    // it would be better to check if the author is signed user.
    // but for now, there is no logic that deals with the uid in frontend.
    // so just embedding the uid to the author field.
    post.author = uid;
    return this.postService.createPost(post);
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
    @Body() post: UpdatePostDto,
  ):
    | Prisma.PrismaPromise<Prisma.BatchPayload>
    | { code: number; message: string } {
    const uid: number = req.uid;
    const aid: number = post.author;
    // if the author is not signed user, return 403 Forbidden.
    if (uid !== aid) return { code: 403, message: 'Forbidden' };
    return this.postService.updatePost(post);
  }

  /**
   * deletes the post by the pid in the body if the user is signed.
   * @param req
   * @param pid
   * @returns
   */
  @UseGuards(UserGuard)
  @Delete('/:pid')
  deletePost(
    @Req() req: Request,
    @Param('pid') pid: any,
  ): Prisma.PrismaPromise<Prisma.BatchPayload> {
    return this.postService.deletePost({ aid: req.uid, pid: Number(pid) });
  }

  /**
   * likes the post by the pid in the body if the user is signed.
   * @param req
   * @param post
   * @returns
   */
  @UseGuards(UserGuard)
  @Post('/like/:pid')
  likePost(
    @Req() req: Request,
    @Param('pid') pid: number,
  ): Promise<[PostLike, PPost]> {
    return this.postService.likePost(req.uid, pid);
  }

  /**
   * unlikes the post by the pid in the body if the user is signed.
   * @param req
   * @param post
   * @returns
   */
  @UseGuards(UserGuard)
  @Post('/unlik/:pid')
  unlikePost(
    @Req() req: Request,
    @Param('pid') pid: number,
  ): Promise<[PostLike, PPost]> {
    return this.postService.unlikePost(req.uid, pid);
  }
}
