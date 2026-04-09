import type {LoginResponse} from "../types/auth";
import type {User} from "../types/chat";

const apiBase = import.meta.env.VITE_API_BASE;


export async function login(tag: string, password: string): Promise<LoginResponse>{
    const res = await fetch(`${apiBase}/v1/login`, {
        method: "POST",
        headers: {"Content-type": "application/json"},
        body: JSON.stringify({tag, password}),
    });

    if (!res.ok) {
        const err = await res.json().catch(()=>{});
        console.log(err);
        throw new Error(err);
    }

    return res.json();
}

export async function register(tag: string, name: string, password: string): Promise<User> {
    const res = await fetch(`${apiBase}/v1/registration`, {
        method: "POST",
        headers: {"Content-type": "application/json"},
        body: JSON.stringify({tag, name, password}),
    });

    if (!res.ok){
        const err = await res.json().catch(()=>{});
        console.log(err);
        throw new Error(err);
    }

    return res.json();
}