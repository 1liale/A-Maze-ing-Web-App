export type User = {
  given_name?: string;
  family_name?: string;
  nickname?: string;
  name?: string;
  picture?: string;
  locale?: string;
  updated_at?: string;
  email: string;
  email_verified: boolean;
  sub: string;
};

export type UserCardProfile = {
  picture?: string;
  name?: string;
  email?: string;
  bio?: string;
};
