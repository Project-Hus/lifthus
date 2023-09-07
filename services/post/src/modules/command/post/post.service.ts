import { Inject, Injectable } from '@nestjs/common';
import { Post, PostLike, Prisma } from '@prisma/client';
import { PrismaService } from 'src/prisma/prisma.service';

@Injectable()
export class PostService {
  constructor(@Inject(PrismaService) private readonly prisma: PrismaService) {}

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
