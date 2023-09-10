export class CommentParents {
  constructor(
    public readonly postId: bigint,
    public readonly parentId?: bigint,
  ) {}
}

export class CommentUpdates {
  constructor(public readonly content: string) {}
}
