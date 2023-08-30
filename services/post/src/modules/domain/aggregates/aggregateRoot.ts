// import { Injectable } from '@nestjs/common';

// /**
//  * @description
//  * AggregateRoot is the base class for all aggregate roots, this class is used to track the status of each aggregate root.
//  */
// @Injectable()
// export class AggregateRoot {
//   private _created: boolean = false;
//   private _updated: boolean = false;
//   private _deleted: boolean = false;
//   constructor({ created }: { created: boolean }) {
//     this._created = created;
//   }
//   protected setCreated() {
//     this._created = true;
//   }
//   protected setUpdated() {
//     this._updated = true;
//   }
//   protected setDeleted() {
//     this._deleted = true;
//   }
//   protected cancel() {
//     this._created = false;
//     this._updated = false;
//     this._deleted = false;
//   }
//   isCreated() {
//     return this._created;
//   }
//   isUpdated() {
//     return this._updated;
//   }
//   isDeleted() {
//     return this._deleted;
//   }
// }
