"use client";
import { Map, MapRef } from "@vis.gl/react-maplibre";
import { useCallback, useEffect, useLayoutEffect, useState } from "react";
import { defaultLocation, getLocation } from "./utilities/api";

function App() {
    const [node, setNode] = useState<MapRef | null>(null);
    const callbackRef = useCallback((el: MapRef) => setNode(el), []);

    useLayoutEffect(() => {
        if (!node) return;
        (async () => {
            const location = await getLocation();
            if (location !== defaultLocation) {
                node.flyTo({ center: location, zoom: 18 });
            }
        })();
    }, [node]);

    const [isDark, setIsDark] = useState(
        typeof window !== "undefined"
            ? window.matchMedia("(prefers-color-scheme: dark)").matches
            : false,
    );

    useEffect(() => {
        const mq = window.matchMedia("(prefers-color-scheme: dark)");

        const update = (e: MediaQueryListEvent) => {
            setIsDark(e.matches);
        };

        mq.addEventListener("change", update);
        return () => mq.removeEventListener("change", update);
    }, []);

    return (
        <Map
            id="map"
            onClick={(e) => {
                console.log("here: ", e.lngLat);
            }}
            ref={callbackRef}
            initialViewState={{
                longitude: defaultLocation.lng,
                latitude: defaultLocation.lat,
                zoom: 2,
            }}
            mapStyle={
                isDark
                    ? "styles/dark.json"
                    : "https://tiles.openfreemap.org/styles/liberty"
            }
            style={{ height: "100svh", width: "100%", marginTop: "-64px" }}
        ></Map>
    );
}

export default App;
