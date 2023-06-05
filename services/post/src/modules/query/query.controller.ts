import {
  Controller,
  Get,
  Param,
  Query,
  BadRequestException,
} from '@nestjs/common';
import { PostQueryService } from './post.query.service';
import { CommentQueryService } from './comment.query.service';
import { Post, Comment } from '@prisma/client';

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

  @Get('/comment')
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

  @Get('/post/slug/:slug')
  getPostBySlug(@Param('slug') slug: string): Promise<Post> {
    return this.postQueryService.getPostBySlug(slug);
  }

  @Get('/post/id/:id')
  getPostById(@Param('id') idStr: string): Promise<Post> {
    const id = Number(idStr);
    if (isNaN(id)) {
      throw new BadRequestException();
    }
    return this.postQueryService.getPostById(id);
  }

  /**
   * gets all posts from the database and returns them.
   * @param skip
   * @returns Post[]
   * @example /post/query/post/all/0
   */
  @Get('/post/all/:skip')
  getAllPosts(@Param('skip') skipStr: string): Promise<Post[]> {
    const skip = Number(skipStr) || 0;
    return this.postQueryService.getAllPosts(Number(skip));
  }

  /**
   * gets users query param and returns posts from the users.
   * @param users
   * @returns Post[]
   * @example /post/query/post?users=1,2,3
   */
  @Get('/post')
  getUsersPosts(
    @Query('users') usersStr: string,
    @Query('skip') skipStr: string,
  ): Promise<Post[]> {
    const users: number[] = usersStr
      .split(',')
      .map((userStr) => Number(userStr));
    if (users.some((user) => isNaN(user))) {
      throw new BadRequestException();
    }
    const skip = Number(skipStr) || 0;
    return this.postQueryService.getUsersPosts({ users, skip });
  }

  @Get('/postnc')
  getUsersPostsNoComments(
    @Query('users') usersStr: string,
    @Query('skip') skipStr: string,
  ): Promise<Post[]> {
    const users: number[] = usersStr
      .split(',')
      .map((userStr) => Number(userStr));
    if (users.some((user) => isNaN(user))) {
      throw new BadRequestException();
    }
    const skip = Number(skipStr) || 0;
    return this.postQueryService.getUsersPostsNoComments({ users, skip });
  }

  /**
   * gets user ID and skip number from url params and returns 10 posts from the skip number.
   * @param req
   * @param res
   * @param uid
   * @param skip
   * @returns {Post[]}
   */
  @Get('/post/user/:uid/:skip')
  getUserPosts(
    @Param('uid') uid: any,
    @Param('skip') skip: any,
  ): Promise<Post[]> {
    skip = skip || 0;
    return this.postQueryService.getUserPosts(Number(uid), Number(skip));
  }
}
