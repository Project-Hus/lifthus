import { Injectable } from '@nestjs/common';
import { User } from '../user/user.model';

@Injectable()
export class Like<T> {
  private liker: bigint;
  private target: bigint;
  private liked: boolean;

  static create<T>(liker: bigint, target: bigint, liked: boolean): Like<T> {
    return new Like<T>().setLiker(liker).setrTarget(target);
  }

  private setLiker(liker: bigint): Like<T> {
    this.liker = liker;
    return this;
  }

  private setrTarget(target: bigint): Like<T> {
    this.target = target;
    return this;
  }

  getLiker(): bigint {
    return this.liker;
  }

  getTarget(): bigint {
    return this.target;
  }

  isLiked(): boolean {
    return this.liked;
  }

  like(u: User): Like<T> | undefined {
    if (u.getID() !== this.liker || this.liked) return undefined;
    this.liked = true;
    return this;
  }

  unlike(u: User): Like<T> | undefined {
    if (u.getID() !== this.liker || !this.liked) return undefined;
    this.liked = false;
    return this;
  }
}
