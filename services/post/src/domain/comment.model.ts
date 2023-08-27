// task.service.ts
import { Injectable } from '@nestjs/common';

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

@Injectable()
export class Comment {
  private images: IImage[] = [];

  private userGroup: bigint;
  private author: bigint;
  private createdAt: Date;
  private updatedAt: Date;
  private slug: string;

  private content: string;
  private mentions: bigint[];
  private likenum: number;
  private likes: bigint[];

  private comments: Comment[] = [];

  constructor() {
    this.images = [];
    this.comments = [];
  }
}
