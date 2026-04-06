import { createContext, useMemo, useState, useEffect, useContext, type ReactNode } from "react";
import type { User } from "../types/chat";

export interface UserContextValue {
    user: User | null;
    setUser: React.Dispatch<React.SetStateAction<User | null>>;
    isAuthReady: boolean;
}

const baseApi = import.meta.env.VITE_API_BASE;

export const UserContext = createContext<UserContextValue | null>(null);

export function UserProvider({children}: {children: ReactNode}){
    const [user, setUser] = useState<User | null>(null);
    const [isAuthReady, setIsAuthReady] = useState(false);

    useEffect(() => {
        const token = localStorage.getItem("token");
        if (!token) {
            setIsAuthReady(true);
            return;
        }

        console.log("start get user /v1/me");
        console.log(`${baseApi}/v1/me`)
        fetch(`${baseApi}/v1/me`, {
            headers: { Authorization: `Bearer ${token}` },
        })
        .then((res) => {
            if (!res.ok) throw new Error("unauthorized");
            return res.json();
        })
        .then((data) => {
            setUser(data);
        })
        .catch(()=>{
            //localStorage.removeItem("token");
            setUser(null);
        })
        .finally(()=>{
            setIsAuthReady(true);
        })
    }, [])

    const value = useMemo<UserContextValue>(
        () => ({
            user,
            setUser,
            isAuthReady,
        }),
        [user, isAuthReady]
    );

    return (
        <UserContext.Provider value={value}>{children}</UserContext.Provider>
    );
}

export function useUser(): UserContextValue{
    const ctx = useContext(UserContext);
    if (ctx == null){
        throw new Error("useUser must be used within UserProvider");
    }
    return ctx
}
