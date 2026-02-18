import { LngLatLike } from "maplibre-gl";

export const BASE_API_URL = process.env.NEXT_PUBLIC_BASE_API_URL;

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

export const defaultLocation = { lng: 0, lat: 0 };

export const getLocation = async (): Promise<LngLatLike> => {
    try {
        const res = await fetch("http://ip-api.com/json/");
        const json = (await res.json()) as LocationResponse;
        if (typeof json.lat === "number" && typeof json.lon === "number") {
            return [json.lon, json.lat];
        }
    } catch (error) {
        console.log(error);
    }
    return defaultLocation;
};
