// task.service.ts
import { Injectable } from '@nestjs/common';
import { User } from '../user/user.model';
import { UpdateCommentDto } from '../../dto(later put out)/comment.dto';

export type InsertCommentInput = {};
interface IImage {
  // id BigInt @id @default(autoincrement()) @db.UnsignedBigInt
  // author    BigInt   @db.UnsignedBigInt
  // createdAt DateTime @default(now())
  // updatedAt DateTime @updatedAt
  // post   Post?   @relation(fields: [postId], references: [id], onDelete: Cascade)
  // postId BigInt? @db.UnsignedBigInt
  // parentId BigInt?   @db.UnsignedBigInt
  // parent   Comment?  @relation("replies", fields: [parentId], references: [id], onDelete: Cascade)
  // replies  Comment[] @relation("replies")
  // content String @db.VarChar(531)
  // likenum  Int           @default(0)
  // mentions Mention[]     @relation("mentions")
  // likes    CommentLike[]
  // @@unique([postId, order])
}

export type CreateCommentInput = {};

interface IComment {
  getID(): bigint;
  getAuthor(): User;
  update(updateData: UpdateCommentDto): Comment;

  like(user: User): void;
  unlike(user: User): void;
  isLikedBy(user: User): boolean;
}

@Injectable()
export class Comment {
  private id: bigint;
  private author: User;

  private content: string;

  private likenum: number;
  private likers: bigint[];
  constructor() {}

  getID() {
    return this.id;
  }

  getAuthor() {
    return this.author;
  }

  update(updateData: UpdateCommentDto): Comment {
    this.content = updateData.content;
    return this;
  }

  like(user: User): void {
    if (!this.isLikedBy(user)) {
      this.likers.push(user.getID());
    }
  }

  unlike(user: User): void {
    if (this.isLikedBy(user)) {
      this.likers = this.likers.filter((id) => id !== user.getID());
    }
  }

  isLikedBy(user: User): boolean {
    return this.likers.includes(user.getID());
  }
}
