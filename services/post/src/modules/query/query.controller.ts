import {
  Controller,
  Get,
  Res,
  Req,
  Param,
  Logger,
  Body,
  Query,
  BadRequestException,
} from '@nestjs/common';
import { Request, Response } from 'express';
import { PostQueryService } from './post.query.service';
import { CommentQueryService } from './comment.query.service';
import { Post } from '@prisma/client';

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
