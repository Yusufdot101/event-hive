import { promises } from "dns";
import { useAuthStore } from "../store/useAuthStore";
import { BASE_API_URL, fetchWithRefreshTokenRetry, location } from "./api";

export const createEvent = async (
    title: string,
    description: string,
    startDate: string,
    startTime: string,
    endDate: string,
    endTime: string,
    address: string,
    lnglat: location,
): Promise<boolean> => {
    const [startHour, startMinute] = startTime.split(":").map(Number);
    const [endHour, endMinute] = endTime.split(":").map(Number);
    let date = new Date(startDate);
    date.setHours(startHour, startMinute);
    startDate = date.toISOString();

    date = new Date(endDate);
    date.setHours(endHour, endMinute);
    endDate = date.toISOString();

    try {
        const res = await fetchWithRefreshTokenRetry(`${BASE_API_URL}/events`, {
            method: "POST",
            credentials: "include",
            body: JSON.stringify({
                title,
                description,
                address,
                longitude: lnglat.lng,
                latitude: lnglat.lat,
                startsAt: startDate,
                endsAt: endDate,
            }),
            headers: {
                "Content-Type": "application/json",
            },
        });
        //
        if (!res) return false;
        const data = await res?.json();
        const error = data.error;
        if (error) {
            alert(error);
            return false;
        }
        return true;
    } catch (error) {
        console.log(error);
        return false;
    }
};
