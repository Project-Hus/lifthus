import { Injectable } from '@nestjs/common';
import { CreatePostInput, Post } from '../domain/aggregates/post/post.model';
import { PostSummary } from '../domain/aggregates/post/postSummary.model';
import { User } from '../domain/aggregates/user/user.model';
import { PostRepository } from '../domain/repositories/post.repository';
import { PrismaService } from 'src/prisma/prisma.service';

import { Prisma, Post as PrismaPost } from '@prisma/client';

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
  constructor(private readonly prismaService: PrismaService) {
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
        return new PostSummary(
          post.id,
          post.author,
          post.createdAt,
          post.images.map((image) => image.src),
          post.slug,
        );
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
        return new PostSummary(
          post.id,
          post.author,
          post.createdAt,
          post.images.map((image) => image.src),
          post.slug,
        );
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
      return Post.reconstitue({
        slug: post.slug,
        author: new User(post.author),
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
      return Post.reconstitue({
        slug: post.slug,
        author: new User(post.author),
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
        author: post.getAuthor().getID(),
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
      });
      const newPostID = newPost.id;
      return await this._getPostByID(newPostID);
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
      });
      return target;
    } catch (e) {
      return Promise.reject(e);
    }
  }

  async _save(changes: Set<Post>): Promise<void> {
    return;
  }
}
