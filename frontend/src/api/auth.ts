import type {LoginResponse} from "../types/auth";

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

    return res.json()
}