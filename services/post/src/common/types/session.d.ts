export interface LifthusSessionJWTPayload {
  purpose: 'lifthus_session';
  sid: string;
  uid: string;
  exp: number;
}
