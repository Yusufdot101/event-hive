import { useAuthStore } from "../store/useAuthStore";
import { BASE_API_URL } from "./api";

export const logout = async () => {
    try {
        const res = await fetch(BASE_API_URL + "/auth/logout", {
            method: "PUT",
            credentials: "include",
        });
        if (!res.ok) {
            alert("An error occurred. Please try again later");
        }

        useAuthStore.getState().clearAccessToken();
    } catch (error) {
        alert("An error occurred. Please try again later");
        console.log(error);
    }
};
