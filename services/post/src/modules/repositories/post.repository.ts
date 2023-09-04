import { Inject, Injectable } from '@nestjs/common';

import { PostRepository } from 'src/domain/repositories/post.repository';

import { PrismaService } from 'src/prisma/prisma.service';

import { User } from 'src/domain/aggregates/user/user.model';
import { Post, UpdatePostInput } from 'src/domain/aggregates/post/post.model';
import { PostSummary } from 'src/domain/aggregates/post/postSummary.model';

import { Prisma } from '@prisma/client';

type PostSummWithImage = {
  id: bigint;
  userGroup: bigint | null;
  author: bigint;
  createdAt: Date;
  updatedAt: Date;
  slug: string;
  images: PrismaPostImage[];
};

type PostWithImages = {
  id: bigint;
  userGroup: bigint | null;
  author: bigint;
  createdAt: Date;
  updatedAt: Date;
  slug: string;
  content: string;
  images: PrismaPostImage[];
};

type PrismaPostImage = {
  id: bigint;
  src: string;
  order: number;
};

@Injectable()
export class PrismaPostRepository extends PostRepository {
  constructor(
    @Inject(PrismaService) private readonly prismaService: PrismaService,
  ) {
    super();
  }

  async _getAllPostSumms(skip: number): Promise<PostSummary[]> {
    try {
      const allPosts: PostSummWithImage[] =
        await this.prismaService.post.findMany({
          include: {
            images: {
              select: {
                order: true,
                id: true,
                src: true,
              },
              orderBy: {
                order: 'asc',
              },
              take: 1,
            },
          },
          orderBy: {
            createdAt: 'desc',
          },
          take: 10,
          skip: skip,
        });

      return allPosts.map((post: PostSummWithImage) => {
        return PostSummary.create({
          id: post.id,
          author: post.author,
          createdAt: post.createdAt,
          images: post.images.map((image) => image.src),
          slug: post.slug,
        });
      });
    } catch (e) {
      return Promise.reject(e);
    }
  }

  async _getUsersPostSumms(
    users: User[],
    skip: number,
  ): Promise<PostSummary[]> {
    try {
      const usersPosts = await this.prismaService.post.findMany({
        where: {
          author: {
            in: users.map((user) => user.getID()),
          },
        },
        include: {
          images: {
            select: {
              order: true,
              id: true,
              src: true,
            },
            orderBy: {
              order: 'asc',
            },
            take: 1,
          },
        },
        orderBy: {
          createdAt: 'desc',
        },
        take: 10,
        skip: skip,
      });
      return usersPosts.map((post: PostSummWithImage) => {
        return PostSummary.create({
          id: post.id,
          author: post.author,
          createdAt: post.createdAt,
          images: post.images.map((image) => image.src),
          slug: post.slug,
        });
      });
    } catch (e) {
      return Promise.reject(e);
    }
  }

  async _getPostByID(pid: bigint): Promise<Post | null> {
    try {
      const post: PostWithImages = await this.prismaService.post.findUnique({
        where: {
          id: pid,
        },
        include: {
          images: {
            select: {
              order: true,
              id: true,
              src: true,
            },
            orderBy: {
              order: 'asc',
            },
          },
        },
      });
      if (!post) return null;
      return Post.query({
        slug: post.slug,
        author: post.author,
        images: post.images.map((image) => image.src),
        content: post.content,
        id: post.id,
        createdAt: post.createdAt,
        updatedAt: post.updatedAt,
      });
    } catch (e) {
      return Promise.reject(e);
    }
  }

  async _getPostBySlug(slug: string): Promise<Post | null> {
    try {
      const post: PostWithImages = await this.prismaService.post.findUnique({
        where: {
          slug: encodeURIComponent(slug),
        },
        include: {
          images: {
            select: {
              order: true,
              id: true,
              src: true,
            },
            orderBy: {
              order: 'asc',
            },
          },
        },
      });
      if (!post) return null;
      return Post.query({
        slug: post.slug,
        author: post.author,
        images: post.images.map((image) => image.src),
        content: post.content,
        id: post.id,
        createdAt: post.createdAt,
        updatedAt: post.updatedAt,
      });
    } catch (e) {
      return Promise.reject(e);
    }
  }

  async _createPost(post: Post): Promise<Post> {
    if (!post.isPre()) return undefined;
    try {
      let createInput: Prisma.PostCreateInput = {
        author: post.getAuthor(),
        slug: post.getSlug(),
        content: post.getContent(),
        images: {
          create: post.getImageSrcs().map((src, idx) => {
            return {
              src,
              order: idx,
            };
          }),
        },
      };

      const newPost = await this.prismaService.post.create({
        data: createInput,
        include: {
          images: {
            select: {
              order: true,
              id: true,
              src: true,
            },
            orderBy: {
              order: 'asc',
            },
          },
        },
      });

      return Post.query({
        slug: newPost.slug,
        author: newPost.author,
        images: newPost.images.map((image) => image.src),
        content: newPost.content,
        id: newPost.id,
        createdAt: newPost.createdAt,
        updatedAt: newPost.updatedAt,
      });
    } catch (e) {
      return Promise.reject(e);
    }
  }

  async _deletePost(target: Post): Promise<Post> {
    try {
      const deletedPost = await this.prismaService.post.delete({
        where: {
          id: target.getID(),
        },
        include: {},
      });
      return target;
    } catch (e) {
      return Promise.reject(e);
    }
  }

  async _save(pid: bigint, updates: UpdatePostInput): Promise<Post> {
    try {
      const updatedPost = await this.prismaService.post.update({
        where: {
          id: pid,
        },
        include: {
          images: {
            select: {
              order: true,
              id: true,
              src: true,
            },
            orderBy: {
              order: 'asc',
            },
          },
        },
        data: updates,
      });

      return Post.query({
        slug: updatedPost.slug,
        author: updatedPost.author,
        images: updatedPost.images.map((image) => image.src),
        content: updatedPost.content,
        id: updatedPost.id,
        createdAt: updatedPost.createdAt,
        updatedAt: updatedPost.updatedAt,
      });
    } catch (e) {
      return Promise.reject(e);
    }
  }
}
