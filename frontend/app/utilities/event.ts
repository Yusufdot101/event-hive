import type { Location } from "./api";
import { fetchWithRefreshTokenRetry } from "./api";

export type EventItem = {
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

export type UserItem = {
    ID: string;
    CreatedAt: string;
    LastUpdatedAt: string;
    Name: string;
    Email: string;
};

export const createEvent = async (
    title: string,
    description: string,
    startDate: string,
    startTime: string,
    endDate: string,
    endTime: string,
    address: string,
    lnglat: Location,
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

export const getEvents = async (): Promise<EventItem[] | undefined> => {
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

export const getAttendingStatus = async (eventID: string): Promise<boolean> => {
    try {
        const res = await fetchWithRefreshTokenRetry(
            `events/${eventID}/attend`,
            {
                method: "GET",
            },
        );
        if (!res) return false;
        const data = await res.json();
        const attendingStatus = data.userIsAttending;
        return attendingStatus;
    } catch (error) {
        console.log(error);
        return false;
    }
};

export const changeAttendingStatus = async (
    eventID: string,
    action: "attend" | "unattend",
): Promise<boolean> => {
    try {
        const res = await fetchWithRefreshTokenRetry(
            `events/${eventID}/attend`,
            {
                method: action === "attend" ? "POST" : "DELETE",
            },
        );
        if (!res || !res.ok) return false;
        const data = await res.json();
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

export const getEventAttendees = async (
    eventID: string,
): Promise<UserItem[] | undefined> => {
    try {
        const res = await fetchWithRefreshTokenRetry(
            `events/${eventID}/attendees`,
            { method: "GET" },
        );
        if (!res) return;

        const data = await res?.json();
        const error = data.error;
        if (error) {
            alert(error);
            return;
        }

        const users = data.users;
        return users;
    } catch (error) {
        console.log(error);
        return;
    }
};

export const getEventCreator = async (
    eventID: string,
): Promise<UserItem | undefined> => {
    try {
        const res = await fetchWithRefreshTokenRetry(
            `events/${eventID}/creator`,
            { method: "GET" },
        );
        if (!res) return;

        const data = await res?.json();
        const error = data.error;
        if (error) {
            alert(error);
            return;
        }

        const user = data.user;
        return user;
    } catch (error) {
        console.log(error);
        return;
    }
};

export const deleteEvent = async (eventID: string): Promise<boolean> => {
    try {
        const res = await fetchWithRefreshTokenRetry(`events/${eventID}`, {
            method: "DELETE",
        });
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
