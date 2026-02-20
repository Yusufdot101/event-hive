"use client";
import { Map, MapRef } from "@vis.gl/react-maplibre";
import { useCallback, useEffect, useState } from "react";
import {
    defaultLocation,
    getAddressFromLngLat,
    getLocation,
    location,
} from "./utilities/api";
import RightClickOptions from "./components/RightClickOptions";
import CustomMarker from "./components/CustomMarker";

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
                node.flyTo({ center: location, zoom: 15 });
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

    useEffect(() => {
        if (!node) return;
        const handleResize = () => node?.resize();
        window.addEventListener("resize", handleResize);

        return () => {
            window.removeEventListener("resize", handleResize);
        };
    }, [node]);

    const [isSelectingLocation, setIsSelectingLocation] = useState(false);
    const [selectedLocation, setSelectedLocation] = useState<location | null>(
        null,
    );
    const [selectedAddress, setSelectedAddress] = useState("");
    const [isCreatingEvent, setIsCreatingEvent] = useState(false);

    const handleCancel = () => {
        setIsSelectingLocation(false);
        setSelectedLocation(null);
        setSelectedAddress("");
    };

    const handleChooseLocation = () => {
        setIsSelectingLocation(false);
        setIsCreatingEvent(true);
    };

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
                    height: "100svh",
                    width: "100%",
                    marginTop: "-64px",
                }}
                onClick={async (e) => {
                    setIsCreatingEvent(false);
                    if (isSelectingLocation) {
                        setSelectedLocation(e.lngLat);
                        const location = await getAddressFromLngLat(e.lngLat);

                        if (!location) return;
                        const { country, city, street } = location;
                        setSelectedAddress(
                            `${country} ${city ? ", " + city : ""} ${street ? ", " + street : ""}`,
                        );
                    }

                    setShowRightClickOptions(false);
                }}
                onContextMenu={(e) => {
                    if (isSelectingLocation) return;
                    setShowRightClickOptions(true);
                    setRightClickOptionsLocation(e.lngLat);
                    setIsCreatingEvent(false);
                }}
            >
                <CustomMarker
                    lng={rightClickOptionsLocation.lng}
                    lat={rightClickOptionsLocation.lat}
                    style={{ width: "100%" }}
                >
                    <RightClickOptions
                        isShown={showRightClickOptions}
                        handleClose={() => setShowRightClickOptions(false)}
                        handleClickSelectLocation={() =>
                            setIsSelectingLocation(true)
                        }
                        selectedLocation={selectedLocation ?? defaultLocation}
                        selectedAddress={selectedAddress}
                        isCreatingEvent={isCreatingEvent}
                        handleClickCreateEvent={() => setIsCreatingEvent(true)}
                        handleCloseCreateEvent={() => setIsCreatingEvent(false)}
                    />
                </CustomMarker>

                {isSelectingLocation && selectedLocation && (
                    <CustomMarker
                        lng={selectedLocation.lng}
                        lat={selectedLocation.lat}
                        style={{
                            translate: "-30px -52px",
                            width: "fit",
                        }}
                    >
                        <svg
                            className="text-foreground"
                            fill="currentColor"
                            version="1.1"
                            id="Layer_1"
                            xmlns="http://www.w3.org/2000/svg"
                            width="64px"
                            height="64px"
                            viewBox="0 0 100 100"
                            enableBackground="new 0 0 100 100"
                        >
                            <g id="SVGRepo_bgCarrier" strokeWidth="0" />
                            <g
                                id="SVGRepo_tracerCarrier"
                                strokeLinecap="round"
                                strokeLinejoin="round"
                            />
                            <g id="SVGRepo_iconCarrier">
                                <g>
                                    <path d="M50,10.417c-15.581,0-28.201,12.627-28.201,28.201c0,6.327,2.083,12.168,5.602,16.873L45.49,86.823 c0.105,0.202,0.21,0.403,0.339,0.588l0.04,0.069l0.011-0.006c0.924,1.278,2.411,2.111,4.135,2.111c1.556,0,2.912-0.708,3.845-1.799 l0.047,0.027l0.179-0.31c0.264-0.356,0.498-0.736,0.667-1.155L72.475,55.65c3.592-4.733,5.726-10.632,5.726-17.032 C78.201,23.044,65.581,10.417,50,10.417z M49.721,52.915c-7.677,0-13.895-6.221-13.895-13.895c0-7.673,6.218-13.895,13.895-13.895 s13.895,6.222,13.895,13.895C63.616,46.693,57.398,52.915,49.721,52.915z" />
                                </g>
                            </g>
                        </svg>
                    </CustomMarker>
                )}
            </Map>

            {isSelectingLocation && (
                <div className="absolute bottom-[40px] right-[40px] rounded-[4px] overflow-hidden">
                    {selectedAddress && (
                        <p className="bg-foreground text-background text-center px-[8px] py-[4px]">
                            {selectedAddress}
                        </p>
                    )}

                    <button
                        className="cursor-pointer bg-red-500 hover:bg-muted hover:text-white text-nowrap w-full p-[12px] bg-accent text-white py-[4px]  duration-300"
                        onClick={handleCancel}
                    >
                        Cancel
                    </button>

                    <button
                        disabled={selectedLocation === null}
                        className={`${selectedLocation === null ? "opacity-80 cursor-not-allowed" : "cursor-pointer hover:bg-muted hover:text-white"} text-nowrap min-w-[300px] w-full p-[12px] bg-accent text-white py-[4px]  duration-300`}
                        onClick={handleChooseLocation}
                    >
                        Choose location
                    </button>
                </div>
            )}
        </>
    );
}

export default App;
