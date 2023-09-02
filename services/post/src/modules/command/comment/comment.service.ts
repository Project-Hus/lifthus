import { Injectable, Logger } from '@nestjs/common';
import { Comment, CommentLike, Prisma } from '@prisma/client';
import { PrismaService } from 'src/prisma/prisma.service';
import { CreateCommentDto, UpdateCommentDto } from './comment.dto';

@Injectable()
export class CommentService {
  constructor(private readonly prisma: PrismaService) {}

  createComment(comment: CreateCommentDto): Promise<Comment> {
    const newComment: Prisma.CommentCreateInput = {
      author: comment.author, // whatever the author is signed user.
      content: comment.content,
    };
    if (comment.postId) {
      newComment.post = { connect: { id: comment.postId } };
    } else if (comment.parentId) {
      newComment.parent = { connect: { id: comment.parentId } };
    }

    return this.prisma.comment.create({
      data: newComment,
    });
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
