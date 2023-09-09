export class PostIds {
  constructor(public readonly id: bigint, public readonly slug: string) {}
}

export class PostContents {
  constructor(
    public readonly imageSrcs: string[],
    public readonly content: string,
  ) {}
}

export class PostUpdates {
  constructor(public readonly content: string) {}
}
