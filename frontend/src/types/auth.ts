import type { User } from "./chat";

export interface LoginToken{
    refresh: string;
    expiresIn: number;
    type: string;
}

export interface LoginResponse{
    user: User;
    token: LoginToken;
}