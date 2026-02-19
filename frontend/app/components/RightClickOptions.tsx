"use client";

import { useMap } from "@vis.gl/react-maplibre";

type Props = {
    isShown: boolean;
    lng: number;
    lat: number;
};

const options = [
    {
        text: "Create Event",
        onClick: () => {},
    },
    {
        text: "Browse Events",
        onClick: () => {},
    },
];

const RightClickOptions = ({ isShown, lng, lat }: Props) => {
    const { current: map } = useMap();

    if (!map) return;
    const { x, y } = map?.project([lng, lat]);
    return (
        <ul
            style={{
                position: "absolute",
                left: `${x}px`,
                top: `${y}px`,
            }}
            className={`${isShown ? "" : "hidden"} absolute z-100 text-[12px] bg-foreground/75 text-background rounded-[4px] p-[4px] text-[16px]`}
        >
            {options.map((option) => (
                <li
                    key={`${option.text}-${option.onClick}`}
                    onClick={option.onClick}
                    className="cursor-pointer hover:text-accent duration-300"
                >
                    {option.text}
                </li>
            ))}
        </ul>
    );
};

export default RightClickOptions;
