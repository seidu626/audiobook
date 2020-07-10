export interface Authenticate {
  name?: string;
  username: string;
  password: string;
}

export class User {
  id: number;
  email: string;
  username: string;
  provider: string;
  uid: string;
  name: string;
  nickname: string | null;
  image: string | null;
  role: string;
  jwt: string;
  isAdmin: boolean;
}
