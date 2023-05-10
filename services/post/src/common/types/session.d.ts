export interface LifthusSessionJWTPayload {
  purpose: 'lifthus_session';
  sid: string;
  tid?: string;
  uid?: string;
  exp: number;
}
