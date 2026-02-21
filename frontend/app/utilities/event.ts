import { fetchWithRefreshTokenRetry, location } from "./api";

export type event = {
    ID: string;
    CreatedAt: string;
    StartsAt: string;
    EndsAt: string;
    LastUpdatedAt: string;
    CreatorID: string;
    Title: string;
    Description: string;
    Latitude: number;
    Longitude: number;
    Address: string;
};

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
        const res = await fetchWithRefreshTokenRetry("events", {
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
            if (error === "invalid or expired access token") {
                alert("please login to access this feature");
                return false;
            }
            alert(error);
            return false;
        }
        return true;
    } catch (error) {
        console.log(error);
        return false;
    }
};

export const getEvents = async (): Promise<event[] | undefined> => {
    try {
        const res = await fetchWithRefreshTokenRetry("events", {
            method: "GET",
        });
        if (!res) return undefined;
        const data = await res.json();
        const events = data.events;
        return events;
    } catch (error) {
        console.log(error);
        return undefined;
    }
};
