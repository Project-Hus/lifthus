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
import { CommentParents } from 'src/domain/aggregates/comment/comment.vo';
import { Timestamps } from 'src/domain/vo';

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
          const parents = new CommentParents(pid, comm.id);
          const timestamps = new Timestamps(r.createdAt, r.updatedAt);
          return Comment.from(r.author, r.id, parents, r.content, timestamps);
        });
        const parents = new CommentParents(pid, null);
        const timestamps = new Timestamps(comm.createdAt, comm.updatedAt);
        return Comment.from(
          comm.author,
          comm.id,
          parents,
          comm.content,
          timestamps,
          replies,
        );
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
    if (comment.isPersisted()) return Promise.reject('Comment already created');
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
        const parents = new CommentParents(newComm.postId, null);
        const timestamps = new Timestamps(newComm.createdAt, newComm.updatedAt);
        return Comment.from(
          newComm.author,
          newComm.id,
          parents,
          newComm.content,
          timestamps,
          [],
        );
      }
      const parents = new CommentParents(newComm.postId, newComm.parentId);
      const timestamps = new Timestamps(newComm.createdAt, newComm.updatedAt);
      return Comment.from(
        newComm.author,
        newComm.id,
        parents,
        newComm.content,
        timestamps,
      );
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
    const parents = new CommentParents(comm.postId, comm.parentId);
    const timestamps = new Timestamps(comm.createdAt, comm.updatedAt);
    return Comment.from(
      comm.author,
      comm.id,
      parents,
      comm.content,
      timestamps,
    );
  }
}
