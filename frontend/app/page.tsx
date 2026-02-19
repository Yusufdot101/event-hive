"use client";
import { LngLatLike, Map, MapRef } from "@vis.gl/react-maplibre";
import { useCallback, useEffect, useState } from "react";
import { defaultLocation, getLocation } from "./utilities/api";
import RightClickOptions from "./components/RightClickOptions";

function App() {
    const [viewState, setViewState] = useState({
        longitude: defaultLocation.lng,
        latitude: defaultLocation.lat,
        zoom: 2,
    });
    const [node, setNode] = useState<MapRef | null>(null);
    const callbackRef = useCallback((el: MapRef) => setNode(el), []);

    useEffect(() => {
        if (!node) return;
        (async () => {
            const location = await getLocation();
            if (location !== defaultLocation) {
                node.flyTo({ center: location, zoom: 17 });
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

    const [rightClickOptionsLocation, setRightClickOptionsLocation] =
        useState(defaultLocation);
    const [showRightClickOptions, setShowRightClickOptions] = useState(false);

    return (
        <>
            <Map
                ref={callbackRef}
                initialViewState={viewState}
                onMove={(e) => {
                    setViewState(e.viewState);
                }}
                mapStyle={
                    isDark
                        ? "styles/dark.json"
                        : "https://tiles.openfreemap.org/styles/liberty"
                }
                style={{
                    minHeight: "100svh",
                    width: "100%",
                    marginTop: "-64px",
                }}
                onClick={() => {
                    setShowRightClickOptions(false);
                }}
                onContextMenu={(e) => {
                    setShowRightClickOptions(true);
                    setRightClickOptionsLocation(e.lngLat);
                }}
            >
                <RightClickOptions
                    isShown={showRightClickOptions}
                    lng={rightClickOptionsLocation.lng}
                    lat={rightClickOptionsLocation.lat}
                />
            </Map>
        </>
    );
}

export default App;
