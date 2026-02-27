import { useAuthStore } from "../store/useAuthStore";
import { BASE_API_URL, decodeJWT } from "./api";

export const refreshToken = async () => {
    try {
        const res = await fetch(BASE_API_URL + "/auth/refreshtoken", {
            method: "PUT",
            credentials: "include",
        });

        if (res.status === 400) return;

        const data = await res.json();
        if (data.error) {
            alert("An alert occured. Please try again later");
            return;
        }

        const accessToken = data.accessToken;
        if (!accessToken) {
            alert("An alert occured. Please try again later");
            return;
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
    } catch (error) {
        console.log(error);
        alert("An alert occured. Please try again later");
    }
};
