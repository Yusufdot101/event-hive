"use client";

import { useMap } from "@vis.gl/react-maplibre";
import { CSSProperties } from "react";

type Props = {
    lng: number;
    lat: number;
    children: React.ReactNode;
    style?: CSSProperties;
    onClick?: () => void;
};

const CustomMarker = ({ lng, lat, children, style, onClick }: Props) => {
    const { current: map } = useMap();
    if (!map) return;
    const { x, y } = map?.project([lng, lat]);

    return (
        <div
            onClick={onClick}
            style={{
                position: "absolute",
                left: `${x}px`,
                top: `${y}px`,
                ...style,
            }}
        >
            {children}
        </div>
    );
};

export default CustomMarker;
