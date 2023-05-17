import { Injectable } from '@nestjs/common';
import { Comment, CommentLike, Prisma } from '@prisma/client';
import { PrismaService } from 'src/prisma/prisma.service';
import { UpdateCommentDto } from './comment.dto';

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
    data: UpdateCommentDto,
  ): Prisma.PrismaPromise<Prisma.BatchPayload> {
    return this.prisma.comment.updateMany({
      data,
      where: { id: data.id, author: data.author },
    });
  }

  /**
   *
   * @param aid
   * @param cid
   * @returns {count:number} only expected 0 or 1
   */
  deleteComment({
    cid,
    aid,
  }: {
    cid: number;
    aid: number;
  }): Prisma.PrismaPromise<Prisma.BatchPayload> {
    return this.prisma.post.deleteMany({
      where: { id: cid, author: aid },
    });
  }

  likeComment(
    uid: number,
    where: Prisma.CommentWhereUniqueInput,
  ): Promise<[CommentLike, Comment]> {
    const likeComment = this.prisma.commentLike.create({
      data: { user: uid, comment: { connect: where } },
    });
    const increaseCommentLikeNum = this.prisma.comment.update({
      data: { likenum: { increment: 1 } },
      where,
    });
    return this.prisma.$transaction([likeComment, increaseCommentLikeNum]);
  }

  unlikeComment(
    uid: number,
    where: Prisma.CommentWhereUniqueInput,
  ): Promise<[CommentLike, Comment]> {
    const unlikeComment = this.prisma.commentLike.delete({
      where: { commentId_user: { user: uid, commentId: where.id } },
    });
    const decreaseCommnetLikeNum = this.prisma.comment.update({
      data: { likenum: { decrement: 1 } },
      where,
    });
    return this.prisma.$transaction([unlikeComment, decreaseCommnetLikeNum]);
  }
}
