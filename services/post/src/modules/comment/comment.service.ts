import { Injectable } from '@nestjs/common';
import { Comment, Prisma } from '@prisma/client';
import { PrismaService } from 'src/prisma/prisma.service';

@Injectable()
export class CommentService {
  constructor(private readonly prisma: PrismaService) {}

  wirteComment(data: Prisma.CommentCreateInput): Promise<Comment> {
    return this.prisma.comment.create({
      data,
    });
  }

  /**
   * @param aid
   * @param cid
   * @returns {count:number} only expected 0 or 1
   */
  updateComment(
    data: Prisma.CommentUpdateInput,
  ): Prisma.PrismaPromise<Prisma.BatchPayload> {
    const cid = Number(data.id);
    const aid = Number(data.author);
    return this.prisma.comment.updateMany({
      data,
      where: { id: cid, author: aid },
    });
  }

  /**
   *
   * @param aid
   * @param cid
   * @returns {count:number} only expected 0 or 1
   */
  deleteComment(
    aid: number,
    cid: number,
  ): Prisma.PrismaPromise<Prisma.BatchPayload> {
    return this.prisma.post.deleteMany({
      where: { id: cid, author: aid },
    });
  }

  likeComment(
    uid: number,
    where: Prisma.CommentWhereUniqueInput,
  ): Promise<Comment> {
    return this.prisma.commentLike
      .create({
        data: { user: uid, comment: { connect: where } },
      })
      .then((res) => {
        return this.prisma.comment.update({
          data: { likenum: { increment: 1 } },
          where,
        });
      });
  }

  unlikeComment(
    uid: number,
    where: Prisma.CommentWhereUniqueInput,
  ): Promise<Comment> {
    return this.prisma.commentLike
      .delete({
        where: { commentId_user: { user: uid, commentId: where.id } },
      })
      .then((res) => {
        return this.prisma.comment.update({
          data: { likenum: { decrement: 1 } },
          where,
        });
      });
  }
}
