export interface AuthResponse{

}

export interface LoginRequest{

}

export interface RegistrationRequest{

}

export async function loginUser(req: LoginRequest): Promise<AuthResponse>{
    await new Promise(resolve => setTimeout(resolve, 1000));

    return {

    }
}

export async function registrationUser(req: RegistrationRequest): Promise<AuthResponse>{
    await new Promise(resolve => setTimeout(resolve, 1000));

    return{

    }
}


