import { Injectable } from '@nestjs/common';
import { Post, Prisma } from '@prisma/client';
import { PrismaService } from 'src/prisma/prisma.service';

@Injectable()
export class PostService {
  constructor(private readonly prisma: PrismaService) {}

  wirtePost(data: Prisma.PostCreateInput): Promise<Post> {
    return this.prisma.post.create({
      data,
    });
  }

  /**
   * @param aid
   * @param pid
   * @returns {count:number} only expected 0 or 1
   */
  updatePost(
    data: Prisma.PostUpdateInput,
  ): Prisma.PrismaPromise<Prisma.BatchPayload> {
    const pid = Number(data.id);
    return this.prisma.post.updateMany({
      data,
      where: { id: pid },
    });
  }

  /**
   *
   * @param aid
   * @param pid
   * @returns {count:number} only expected 0 or 1
   */
  deletePost(
    aid: number,
    pid: number,
  ): Prisma.PrismaPromise<Prisma.BatchPayload> {
    return this.prisma.post.deleteMany({
      where: { id: pid, author: aid },
    });
  }

  likePost(uid: number, where: Prisma.PostWhereUniqueInput): Promise<Post> {
    return this.prisma.postLike
      .create({
        data: { user: uid, post: { connect: where } },
      })
      .then((res) => {
        return this.prisma.post.update({
          data: { likenum: { increment: 1 } },
          where,
        });
      });
  }

  unlikePost(uid: number, where: Prisma.PostWhereUniqueInput): Promise<Post> {
    return this.prisma.postLike
      .delete({
        where: { postId_user: { user: uid, postId: where.id } },
      })
      .then((res) => {
        return this.prisma.post.update({
          data: { likenum: { decrement: 1 } },
          where,
        });
      });
  }
}
