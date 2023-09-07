export class CreateCommentRequestDto {
  constructor(
    public postId: string,
    public author: string,
    public content: string,
    public parentId?: string,
  ) {}
}

export class CreateCommentServiceDto {
  postId: bigint;
  author: bigint;
  content: string;
  parentId?: bigint;
  constructor(c: CreateCommentRequestDto) {
    this.postId = BigInt(c.postId);
    this.author = BigInt(c.author);
    this.content = c.content;
    this.parentId = !!c.parentId ? BigInt(c.parentId) : undefined;
  }
}

export class UpdateCommentRequestDto {
  constructor(
    public id: string,
    public author: string,
    public content: string,
  ) {}
}

export class UpdateCommentServiceDto {
  id: bigint;
  author: bigint;
  content: string;
  constructor(up: UpdateCommentRequestDto) {
    this.id = BigInt(up.id);
    this.author = BigInt(up.author);
    this.content = up.content;
  }
}
