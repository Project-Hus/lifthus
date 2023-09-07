import { Inject, Injectable } from '@nestjs/common';
import { Prisma } from '@prisma/client';
import { PrismaService } from 'src/prisma/prisma.service';
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

@Injectable()
export class PostService {
  constructor(
    @Inject(PrismaService) private readonly prisma: PrismaService,
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
      const userPost = author.createPost(post);
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
    const updatedPost = author.updatePost(originalPost, postUpdates);
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

  likePost({ uid, pid }: { uid: number; pid: number }): Promise<number> {
    const likePost = this.prisma.postLike.create({
      data: { user: uid, post: { connect: { id: pid } } },
    });
    const increasePostLikeNum = this.prisma.post.update({
      data: { likenum: { increment: 1 } },
      where: { id: pid },
    });
    return this.prisma
      .$transaction([likePost, increasePostLikeNum])
      .then((res) => {
        // if liking is successful, return the new likenum
        return res[1].likenum;
      })
      .catch(async (err) => {
        // if failed, try unliking
        const unlikePost = this.prisma.postLike.delete({
          where: { postId_user: { user: uid, postId: pid } },
        });
        const decreasePostLikeNum = this.prisma.post.update({
          data: { likenum: { decrement: 1 } },
          where: { id: pid },
        });
        return this.prisma
          .$transaction([unlikePost, decreasePostLikeNum])
          .then((res) => {
            // if unliking is successful, return the new likenum
            return res[1].likenum;
          })
          .catch((err) => {
            return Promise.reject(err);
          });
      });
  }
}
