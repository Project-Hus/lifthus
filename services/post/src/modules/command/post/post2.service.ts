import { Injectable } from '@nestjs/common';
import { Prisma } from '@prisma/client';
import { PrismaService } from 'src/prisma/prisma.service';
import { CreatePostDto, UpdatePostDto } from './post.dto';
import { slugify } from 'src/common/utils/utils';
import { UserRepository } from 'src/modules/domain/repositories/user.repository';
import { PostRepository } from 'src/modules/domain/repositories/post.repository';
import { Post } from 'src/modules/domain/aggregates/post/post.model';

@Injectable()
export class PostService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly userRepo: UserRepository,
    private readonly postRepo: PostRepository,
  ) {}

  async createPost({
    post,
    imageSrcs,
  }: {
    post: CreatePostDto;
    imageSrcs: string[];
  }): Promise<Post> {
    try {
      const author = this.userRepo.getUser(BigInt(post.author));
      const images: string[] = post.images.map((file) => file.location);
      const userPost = author.createPost({ images, content: post.content });
      return await this.postRepo.createPost(userPost);
    } catch (err) {
      throw err;
    }
  }

  /**
   * @param aid
   * @param pid
   * @returns {count:number} only expected 0 or 1
   */
  updatePost(data: UpdatePostDto): Prisma.PrismaPromise<Prisma.BatchPayload> {
    return this.prisma.post.updateMany({
      data,
      where: { id: data.id },
    });
  }

  /**
   *
   * @param aid
   * @param pid
   * @returns {count:number} only expected 0 or 1
   */
  deletePost({
    aid,
    pid,
  }: {
    aid: number;
    pid: number;
  }): Prisma.PrismaPromise<Prisma.BatchPayload> {
    return this.prisma.post.deleteMany({
      where: { id: pid, author: aid },
    });
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
