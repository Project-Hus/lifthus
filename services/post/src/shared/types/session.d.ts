export interface LifthusSessionJWTPayload {
  pps: 'lifthus_session';
  iss: 'https://auth.lifthus.com';
  sid: string;
  tid: string; // maybe added later
  uid: string;
  exp: number;
}
