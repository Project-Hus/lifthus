import { Injectable } from '@nestjs/common';
import { User } from '../user/user.model';

@Injectable()
export class Like<T> {
  private liker: User;
  private target: T;
  private liked: boolean;

  static create<T>(liker: User, target: T, liked: boolean): Like<T> {
    return new Like<T>().setLiker(liker).setrTarget(target);
  }

  private setLiker(liker: User): Like<T> {
    this.liker = liker;
    return this;
  }

  private setrTarget(target: T): Like<T> {
    this.target = target;
    return this;
  }

  getLiker(): User {
    return this.liker;
  }

  getTarget(): T {
    return this.target;
  }

  isLiked(): boolean {
    return this.liked;
  }

  like(u: User): Like<T> | undefined {
    if (u.getID() !== this.liker.getID() || this.liked) return undefined;
    this.liked = true;
    return this;
  }

  unlike(u: User): Like<T> | undefined {
    if (u.getID() !== this.liker.getID() || !this.liked) return undefined;
    this.liked = false;
    return this;
  }
}
