import { Inject, Injectable } from '@nestjs/common';
import { Prisma } from '@prisma/client';
import { PrismaService } from 'src/modules/repositories/prisma/prisma.service';
import {
  CreateCommentServiceDto,
  UpdateCommentServiceDto,
} from 'src/dto/inbound/comment.dto';
import { CommentDto } from 'src/dto/outbound/comment.dto';
import { UserRepository } from 'src/modules/repositories/abstract/user.repository';
import { CommentRepository } from 'src/modules/repositories/abstract/comment.repository';
import {
  CommentParents,
  CommentUpdates,
} from 'src/domain/aggregates/comment/comment.vo';

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
    const parents = new CommentParents(comment.postId, comment.parentId);
    const newComment = author.createComment(
      comment.author,
      parents,
      comment.content,
    );
    const createdComment = await this.commentRepo.createComment(newComment);
    return new CommentDto(createdComment);
  }

  async updateComment({
    clientId,
    updates,
  }: {
    clientId: bigint;
    updates: UpdateCommentServiceDto;
  }): Promise<CommentDto> {
    const author = this.userRepo.getUser(clientId);
    const target = await this.commentRepo.getComment(updates.id);
    const commentUpdates = new CommentUpdates(updates.content);
    const updatedTarget = author.updateComment(target, commentUpdates);
    const savedTarget = await this.commentRepo.save(updatedTarget);
    return new CommentDto(savedTarget);
  }

  async deleteComment({
    clientId,
    cid,
  }: {
    clientId: bigint;
    cid: bigint;
  }): Promise<CommentDto> {
    const author = this.userRepo.getUser(clientId);
    const target = await this.commentRepo.getComment(cid);
    const deletionVerifiedTarget = author.deleteComment(target);
    const deletedTarget = await this.commentRepo.deleteComment(
      deletionVerifiedTarget,
    );
    return new CommentDto(deletedTarget);
  }
}
