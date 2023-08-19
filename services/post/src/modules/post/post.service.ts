import { Injectable } from '@nestjs/common';
import { Post, PostLike, Prisma } from '@prisma/client';
import { PrismaService } from 'src/prisma/prisma.service';
import { CreatePostDto, UpdatePostDto } from './post.dto';
import crypto from 'crypto';
import { slugify } from 'src/common/utils/utils';

@Injectable()
export class PostService {
  constructor(private readonly prisma: PrismaService) {}

  createPost(post: CreatePostDto): Promise<Post> {
    // first, set the range of slug and get it.
    let slug: string;
    const slugEnd: number = post.content.indexOf('\n');
    if (slugEnd == -1 || slugEnd > 30) {
      slug = post.content.slice(0, 30);
    } else {
      slug = post.content.slice(0, slugEnd);
    }
    // get slug
    slug = encodeURIComponent(slug + crypto.randomBytes(8).toString('hex'));

    // Post create form
    let data: Prisma.PostCreateInput = {
      author: post.author,
      slug,
      content: post.content,
    };

    return this.prisma.post.create({
      data,
    });
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
