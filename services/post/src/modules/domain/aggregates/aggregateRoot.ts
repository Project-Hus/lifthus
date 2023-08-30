import { Injectable } from '@nestjs/common';

/**
 * @description
 * AggregateRoot is the base class for all aggregate roots, this class is used to track the status of each aggregate root.
 */
@Injectable()
export class AggregateRoot {
  private created: boolean = false;
  private updated: boolean = false;
  private deleted: boolean = false;
  private setCreated() {
    this.created = true;
  }
  private setUpdated() {
    this.updated = true;
  }
  private setDeleted() {
    this.deleted = true;
  }
  isCreated() {
    return this.created;
  }
  isUpdated() {
    return this.updated;
  }
  isDeleted() {
    return this.deleted;
  }
}
