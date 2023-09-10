import { Inject, Injectable } from '@nestjs/common';
import { Prisma } from '@prisma/client';
import { PrismaService } from 'src/modules/repositories/prisma/prisma.service';
import { UserRepository } from 'src/modules/repositories/abstract/user.repository';
import { PostRepository } from 'src/modules/repositories/abstract/post.repository';
import { Post } from 'src/domain/aggregates/post/post.model';
import { PostDto } from 'src/dto/outbound/post.dto';
import {
  CreatePostServiceDto,
  UpdatePostServiceDto,
} from 'src/dto/inbound/post.dto';
import { PostLikeRepository } from 'src/modules/repositories/abstract/like.repository';
import { CommentRepository } from 'src/modules/repositories/abstract/comment.repository';
import { PostContents, PostUpdates } from 'src/domain/aggregates/post/post.vo';

@Injectable()
export class PostService {
  constructor(
    @Inject(UserRepository) private readonly userRepo: UserRepository,
    @Inject(PostRepository) private readonly postRepo: PostRepository,
    @Inject(PostLikeRepository)
    private readonly postLikeRepo: PostLikeRepository,
    @Inject(CommentRepository) private readonly commentRepo: CommentRepository,
  ) {}

  async createPost({
    clientId,
    post,
  }: {
    clientId: bigint;
    post: CreatePostServiceDto;
  }): Promise<PostDto> {
    try {
      const author = this.userRepo.getUser(clientId);
      const contents = new PostContents(post.imageSrcs, post.content);
      const userPost = author.createPost(post.author, contents);
      const newPost: Post = await this.postRepo.createPost(userPost);
      return new PostDto(newPost, 0, false, 0);
    } catch (err) {
      throw err;
    }
  }

  async updatePost({
    clientId,
    postUpdates,
  }: {
    clientId: bigint;
    postUpdates: UpdatePostServiceDto;
  }): Promise<PostDto> {
    const author = this.userRepo.getUser(clientId);
    const originalPost = await this.postRepo.getPostByID(postUpdates.id);
    const updates = new PostUpdates(postUpdates.content);
    const updatedPost = author.updatePost(originalPost, updates);
    const savedPost = await this.postRepo.save(updatedPost);
    return new PostDto(savedPost);
  }

  async deletePost({
    clientId,
    pid,
  }: {
    clientId: bigint;
    pid: bigint;
  }): Promise<PostDto> {
    const author = this.userRepo.getUser(clientId);
    const targetPost = await this.postRepo.getPostByID(pid);
    const deletionVerifiedPost = author.deletePost(targetPost);
    const deletedPost = await this.postRepo.deletePost(deletionVerifiedPost);
    return new PostDto(deletedPost);
  }
}
