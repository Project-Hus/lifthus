import { Inject, Injectable } from '@nestjs/common';
import {
  Comment,
  CreateReplyInput,
} from '../domain/aggregates/comment/comment.model';

import { Post } from '../domain/aggregates/post/post.model';
import { PrismaService } from 'src/prisma/prisma.service';
import { CommentRepository } from '../domain/repositories/comment.repository';
import { Prisma } from '@prisma/client';

@Injectable()
export class PrismaCommentRepository extends CommentRepository {
  constructor(
    @Inject(PrismaService) private readonly prismaService: PrismaService,
  ) {
    super();
  }

  async _getCommentByID(cid: bigint): Promise<Comment | null> {
    try {
      const comm = await this.prismaService.comment.findUnique({
        where: {
          id: cid,
        },
        include: {
          replies: {
            select: {
              id: true,
              author: true,
              createdAt: true,
              updatedAt: true,
              content: true,
            },
          },
        },
      });
      if (!comm) return null;
      return Comment.queryComment({
        id: comm.id,
        author: comm.author,
        content: comm.content,
        postId: comm.postId,
        createdAt: comm.createdAt,
        updatedAt: comm.updatedAt,
        replies: comm.replies.map((r) => {
          return Comment.queryReply({
            id: r.id,
            postId: comm.postId,
            parentId: comm.id,
            author: r.author,
            content: r.content,
            createdAt: r.createdAt,
            updatedAt: r.updatedAt,
          });
        }),
      });
    } catch (e) {
      return Promise.reject(e);
    }
  }

  async _getComments(post: Post): Promise<Comment[]> {
    try {
      const comms = await this.prismaService.comment.findMany({
        where: {
          postId: post.getID(),
          parentId: null,
        },
        include: {
          replies: {
            select: {
              id: true,
              author: true,
              createdAt: true,
              updatedAt: true,
              content: true,
            },
          },
        },
      });
      const comments: Comment[] = comms.map((comm) => {
        const replies: Comment[] = comm.replies.map((r) => {
          return Comment.queryReply({
            id: r.id,
            author: r.author,
            postId: comm.postId,
            parentId: comm.id,
            content: r.content,
            createdAt: r.createdAt,
            updatedAt: r.updatedAt,
          });
        });
        return Comment.queryComment({
          id: comm.id,
          author: comm.author,
          postId: comm.postId,
          content: comm.content,
          createdAt: comm.createdAt,
          updatedAt: comm.updatedAt,
          replies: replies,
        });
      });
      return comments;
    } catch (e) {
      return Promise.reject(e);
    }
  }

  async _createComment(comment: Comment): Promise<Comment> {
    if (!comment.isPre()) return Promise.reject('Comment already created');
    try {
      const newComm = await this.prismaService.comment.create({
        data: {
          postId: comment.getPostID(),
          parentId: comment.getParentID(),
          author: comment.getAuthor(),
          content: comment.getContent(),
        },
      });
      if (!newComm.parentId) {
        return Comment.queryComment({
          id: newComm.id,
          author: newComm.author,
          postId: newComm.postId,
          content: newComm.content,
          createdAt: newComm.createdAt,
          updatedAt: newComm.updatedAt,
          replies: [],
        });
      }
      return Comment.queryReply({
        id: newComm.id,
        author: newComm.author,
        postId: newComm.postId,
        parentId: newComm.parentId,
        content: newComm.content,
        createdAt: newComm.createdAt,
        updatedAt: newComm.updatedAt,
      });
    } catch (e) {
      return Promise.reject(e);
    }
  }
  async _deleteComment(target: Comment): Promise<Comment> {
    try {
      const deletedComment = await this.prismaService.comment.delete({
        where: {
          id: target.getID(),
        },
      });
      return target;
    } catch (e) {
      return Promise.reject(e);
    }
  }

  async _save(changes: Set<Comment>): Promise<void> {
    try {
      const changeList = Array.from(changes);
      // start prisma transaction
      await this.prismaService.$transaction([
        ...changeList.map((change) => {
          return this.prismaService.comment.update({
            where: {
              id: change.getID(),
            },
            data: {
              content: change.getContent(),
            },
          });
        }),
      ]);
      return;
    } catch (e) {
      return Promise.reject(e);
    }
  }
}
