import { useAuthStore } from "../store/useAuthStore";
import { BASE_API_URL } from "./api";

export const refreshToken = async () => {
    try {
        const res = await fetch(BASE_API_URL + "/auth/refreshtoken", {
            method: "PUT",
            credentials: "include",
        });

        const data = await res.json();
        if (data.error && res.status !== 400) {
            alert("An alert occured. Please try again later");
            return;
        }

        const accessToken = data.accessToken;
        if (!accessToken) {
            alert("An alert occured. Please try again later");
            return;
        }

        useAuthStore.getState().setAcessToken(accessToken);
        useAuthStore.getState().setIsLoggedIn(true);
    } catch (error) {
        console.log(error);
        alert("An alert occured. Please try again later");
    }
};
