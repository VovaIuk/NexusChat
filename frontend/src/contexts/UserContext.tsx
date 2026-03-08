import { createContext, useMemo, useState, useEffect, useContext, type ReactNode } from "react";
import type { User } from "../types/chat";

export interface UserContextValue {
    user: User | null;
    setUser: React.Dispatch<React.SetStateAction<User | null>>;
}

export const UserContext = createContext<UserContextValue | null>(null);

export function UserProvider({children}: {children: ReactNode}){
    const [user, setUser] = useState<User | null>(null);

    useEffect(() => {
        setUser({id: 1, tag: "alice_dev", name: "alice"});
    }, [])

    const value = useMemo<UserContextValue>(
        () => ({
            user,
            setUser,
        }),
        [user]
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
