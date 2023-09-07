import { Inject, Injectable, Logger } from '@nestjs/common';
import { Comment, CommentLike, Prisma } from '@prisma/client';
import { PrismaService } from 'src/prisma/prisma.service';
import { CreateCommentDto, UpdateCommentDto } from './comment.dto';
import { CreateCommentServiceDto } from 'src/dto/inbound/comment.dto';
import { CommentDto } from 'src/dto/outbound/comment.dto';
import { UserRepository } from 'src/modules/repositories/abstract/user.repository';
import { CommentRepository } from 'src/modules/repositories/abstract/comment.repository';

@Injectable()
export class CommentService {
  constructor(
    @Inject(PrismaService) private readonly prisma: PrismaService,
    @Inject(UserRepository) private readonly userRepo: UserRepository,
    @Inject(CommentRepository) private readonly commentRepo: CommentRepository,
  ) {}

  async createComment({
    clientId,
    comment,
  }: {
    clientId: bigint;
    comment: CreateCommentServiceDto;
  }): Promise<CommentDto> {
    const author = this.userRepo.getUser(clientId);
    const newComment = author.createComment(comment);
    const createdComment = await this.commentRepo.createComment(newComment);
    return new CommentDto(createdComment);
  }
  adfsafds;
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
    return this.prisma.comment.deleteMany({
      where: { id: cid, author: aid },
    });
  }

  likeComment({ uid, cid }: { uid: number; cid: number }): Promise<number> {
    const likeComment = this.prisma.commentLike.create({
      data: { user: uid, comment: { connect: { id: cid } } },
    });
    const increaseCommentLikeNum = this.prisma.comment.update({
      data: { likenum: { increment: 1 } },
      where: { id: cid },
    });
    return this.prisma
      .$transaction([likeComment, increaseCommentLikeNum])
      .then((res) => {
        // if liking is successful, return the new likenum
        return res[1].likenum;
      })
      .catch((err) => {
        // if failed try unliking
        const unlikeComment = this.prisma.commentLike.delete({
          where: { commentId_user: { user: uid, commentId: cid } },
        });
        const decreaseCommnetLikeNum = this.prisma.comment.update({
          data: { likenum: { decrement: 1 } },
          where: { id: cid },
        });
        return this.prisma
          .$transaction([unlikeComment, decreaseCommnetLikeNum])
          .then((res) => {
            // if unliking is successful, return the new likenum
            return res[1].likenum;
          })
          .catch((err) => {
            return err;
          });
      });
  }
}
