import { useAuthStore } from "../store/useAuthStore";
import { BASE_API_URL, decodeJWT } from "./api";

export const signin = async (
    handleError: (error: string) => void,
    email: string,
    password: string,
): Promise<boolean> => {
    try {
        const res = await fetch(BASE_API_URL + "/auth/signin", {
            method: "PUT",
            credentials: "include",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ email, password }),
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

        const { payload } = decodeJWT(accessToken);
        if (!payload || !payload.sub) {
            console.error("invalid JWT payload");
            useAuthStore.getState().clearAccessToken();
            return false;
        }

        const userID = payload.sub;
        if (userID === "") {
            console.error("invalid user ID in JWT");
            useAuthStore.getState().clearAccessToken();
            return false;
        }

        useAuthStore.getState().setUserID(userID);
        useAuthStore.getState().setAcessToken(accessToken);
        useAuthStore.getState().setIsLoggedIn(true);
        return true;
    } catch (error) {
        alert("An error occurred. Please try again later");
        console.log(error);
        return false;
    }
};
