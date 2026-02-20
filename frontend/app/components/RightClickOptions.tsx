"use client";

import { location } from "../utilities/api";
import CreateEvent from "./CreateEvent";

type Props = {
    isShown: boolean;
    handleClose: () => void;
    handleClickSelectLocation: () => void;
    selectedLocation: location;
    selectedAddress: string;
    isCreatingEvent: boolean;
    handleClickCreateEvent: () => void;
    handleCloseCreateEvent: () => void;
};

const RightClickOptions = ({
    isShown,
    handleClose,
    handleClickSelectLocation,
    selectedLocation,
    selectedAddress,
    isCreatingEvent,
    handleClickCreateEvent,
    handleCloseCreateEvent,
}: Props) => {
    const options = [
        {
            text: "Create Event",
            onClick: () => {
                handleClickCreateEvent();
                handleClose();
            },
        },
        {
            text: "Browse Events",
            onClick: () => {
                handleClose();
            },
        },
    ];

    return (
        <>
            <ul
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

            {isCreatingEvent ? (
                <CreateEvent
                    handleClose={handleCloseCreateEvent}
                    handleClickSelectLocation={() => {
                        handleCloseCreateEvent();
                        handleClickSelectLocation();
                    }}
                    selectedLocation={selectedLocation}
                    selectedAddress={selectedAddress}
                />
            ) : undefined}
        </>
    );
};

export default RightClickOptions;
