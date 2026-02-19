"use client";

import { useMap } from "@vis.gl/react-maplibre";

type Props = {
    lng: number;
    lat: number;
    children: React.ReactNode;
};

const CustomMarker = ({ lng, lat, children }: Props) => {
    const { current: map } = useMap();
    if (!map) return;
    const { x, y } = map?.project([lng, lat]);

    return (
        <div
            style={{
                position: "absolute",
                left: `${x}px`,
                top: `${y}px`,
            }}
            className="w-full"
        >
            {children}
        </div>
    );
};

export default CustomMarker;
