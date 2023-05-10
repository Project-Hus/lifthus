export interface LifthusSessionJWTPayload {
  purpose: 'lifthus_session';
  sid: string;
  tid?: string; // maybe added later
  uid: string;
  exp: number;
}
