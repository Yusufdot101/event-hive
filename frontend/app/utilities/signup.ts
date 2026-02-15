import { useAuthStore } from "../store/useAuthStore";
import { BASE_API_URL } from "./api";

export const signup = async (
    handleError: (error: string) => void,
    name: string,
    email: string,
    password: string,
): Promise<boolean> => {
    log;
    try {
        const res = await fetch(BASE_API_URL + "/auth/signup", {
            method: "POST",
            credentials: "include",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ name, email, password }),
        });

        const data = await res.json();
        if (data.error) {
            handleError(data.error);
            return false;
        }

        const accessToken = data.accessToken;
        if (!accessToken) {
            alert("An error occurred. Please try again later");
            return false;
        }

        useAuthStore.getState().setAcessToken(accessToken);
        useAuthStore.getState().setIsLoggedIn(true);
        return true;
    } catch (error) {
        alert("An error occurred. Please try again later");
        console.log(error);
        return false;
    }
};
