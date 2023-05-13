import { Injectable } from '@nestjs/common';
import { Post, Prisma } from '@prisma/client';
import { PrismaService } from 'src/prisma/prisma.service';

@Injectable()
export class PostService {
  constructor(private readonly prisma: PrismaService) {}

  getHello(): string {
    return 'Hello World!';
  }

  wirtePost(data: Prisma.PostCreateInput): Promise<Post> {
    return this.prisma.post.create({
      data,
    });
  }

  updatePost(data: Prisma.PostUpdateInput): Promise<Post> {
    const pid = Number(data.id);
    return this.prisma.post.update({
      data,
      where: { id: pid },
    });
  }

  deletePost(where: Prisma.PostWhereUniqueInput): Promise<Post> {
    return this.prisma.post.delete({
      where,
    });
  }

  likePost(uid: number, where: Prisma.PostWhereUniqueInput): Promise<Post> {
    this.prisma.postLike.create({
      data: { user: uid, post: { connect: where } },
    });
    return this.prisma.post.update({
      data: { likenum: { increment: 1 } },
      where,
    });
  }

  unlikePost(uid: number, where: Prisma.PostWhereUniqueInput): Promise<Post> {
    this.prisma.postLike.delete({
      where: { postId_user: { user: uid, postId: where.id } },
    });
    return this.prisma.post.update({
      data: { likenum: { decrement: 1 } },
      where,
    });
  }
}
