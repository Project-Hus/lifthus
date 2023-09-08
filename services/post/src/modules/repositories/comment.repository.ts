import {
  HttpException,
  Inject,
  Injectable,
  InternalServerErrorException,
} from '@nestjs/common';
import {
  Comment,
  CreateCommentInput,
  CreateReplyInput,
  UpdateCommentInput,
} from '../../domain/aggregates/comment/comment.model';

import { Post } from '../../domain/aggregates/post/post.model';
import { PrismaService } from 'src/modules/repositories/prisma/prisma.service';
import { CommentRepository } from './abstract/comment.repository';
import { Comment as PComment } from '@prisma/client';

@Injectable()
export class PrismaCommentRepository extends CommentRepository {
  constructor(
    @Inject(PrismaService) private readonly prismaService: PrismaService,
  ) {
    super();
  }

  async _getComment(cid: bigint): Promise<Comment> {
    try {
      const comm = await this.prismaService.comment.findUnique({
        where: {
          id: cid,
        },
      });
      return this.createModel(comm);
    } catch (e) {
      return Promise.reject(e);
    }
  }

  async _getComments(pid: bigint): Promise<Comment[]> {
    try {
      const comms = await this.prismaService.comment.findMany({
        where: {
          postId: pid,
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

  async _getCommentsNum(pid: bigint): Promise<number> {
    try {
      return this.prismaService.comment.count({
        where: {
          postId: pid,
        },
      });
    } catch (err) {
      throw new InternalServerErrorException();
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
      return this.createModel(deletedComment);
    } catch (e) {
      return Promise.reject(e);
    }
  }

  async _save(cid: bigint, updates: UpdateCommentInput): Promise<Comment> {
    try {
      const target = await this.prismaService.comment.update({
        where: {
          id: cid,
        },
        data: {
          content: updates.content,
        },
      });
      return this.createModel(target);
    } catch (e) {
      return Promise.reject(e);
    }
  }

  private createModel(comm: PComment): Comment {
    if (!comm) return null;
    else if (!comm.parentId) {
      return Comment.queryComment({
        id: comm.id,
        author: comm.author,
        content: comm.content,
        postId: comm.postId,
        replies: undefined,
        createdAt: comm.createdAt,
        updatedAt: comm.updatedAt,
      });
    } else {
      return Comment.queryReply({
        id: comm.id,
        author: comm.author,
        postId: comm.postId,
        parentId: comm.id,
        content: comm.content,
        createdAt: comm.createdAt,
        updatedAt: comm.updatedAt,
      });
    }
  }
}
