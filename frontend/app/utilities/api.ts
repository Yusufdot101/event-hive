import { useAuthStore } from "../store/useAuthStore";
import { refreshToken } from "./refreshToken";

export const BASE_API_URL = process.env.NEXT_PUBLIC_BASE_API_URL;
export const GEOAPIFY_API_KEY = process.env.NEXT_PUBLIC_GEOAPIFY_API_KEY;

export interface LocationResponse {
    status: string;
    country: string;
    countryCode: string;
    region: string;
    regionName: string;
    city: string;
    zip: string;
    lat: number;
    lon: number;
    timezone: string;
    isp: string;
    org: string;
    as: string;
    query: string;
}

export type Location = {
    lng: number;
    lat: number;
};
export const defaultLocation: Location = { lng: 0, lat: 0 };

export const getUserLocation = async (): Promise<Location> => {
    try {
        const res = await fetch("http://ip-api.com/json/");
        const json = (await res.json()) as LocationResponse;
        if (typeof json.lat === "number" && typeof json.lon === "number") {
            return { lng: json.lon, lat: json.lat };
        }
    } catch (error) {
        console.log(error);
    }
    return defaultLocation;
};

type locationInfo = {
    country: string;
    city: string;
    street: string;
};

export const getAddressFromLngLat = async (
    location: Location,
): Promise<locationInfo | null> => {
    try {
        const res = await fetch(
            `https://api.geoapify.com/v1/geocode/reverse?lat=${location.lat}&lon=${location.lng}&apiKey=${GEOAPIFY_API_KEY}`,
        );
        const data = await res.json();

        const { country, city, street } = data.features[0].properties;
        return { country, city, street };
    } catch (error) {
        console.log(error);
        return null;
    }
};

export const fetchWithRefreshTokenRetry = async (
    path: string,
    options: RequestInit,
): Promise<Response | undefined> => {
    const accessToken = useAuthStore.getState().accessToken;
    const url = path.startsWith("http") ? path : BASE_API_URL + "/" + path;
    try {
        let res = await fetch(url, {
            ...options,
            credentials: "include",
            headers: {
                ...(options.headers || {}),
                Authorization: accessToken ? `Bearer ${accessToken}` : "",
            },
        });

        if (res.status === 401) {
            await refreshToken();
            const newAccessToken = useAuthStore().accessToken;
            let res = await fetch(url, {
                ...options,
                credentials: "include",
                headers: {
                    ...(options.headers || {}),
                    Authorization: newAccessToken
                        ? `Bearer ${newAccessToken}`
                        : "",
                },
            });

            if (res.status === 401) {
                useAuthStore.getState().clearAccessToken(); // because the refresh token didn't refresh access token successfully
                // alert("please login to use this feature.");
            }
        }
        return res;
    } catch (error) {
        console.log(error);
        return undefined;
    }
};
