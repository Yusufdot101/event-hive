import { useEffect, useState } from "react";
import {
    changeAttendingStatus,
    deleteEvent,
    EventItem,
    getAttendingStatus,
    getEventAttendees,
    getEventCreator,
    UserItem,
} from "../utilities/event";
import { useAuthStore } from "../store/useAuthStore";
import { useRouter } from "next/navigation";

type Props = {
    event: EventItem;
    handleClose: () => void;
};

const Event = ({ event, handleClose }: Props) => {
    const dateFormat: Intl.DateTimeFormatOptions = {
        year: "numeric",
        month: "long",
        day: "numeric",
        hour: "2-digit",
        minute: "2-digit",
    };

    const [isAttendingEvent, setIsAttendingEvent] = useState(false);
    const [eventAttendees, setEventAttendees] = useState<UserItem[]>();
    const [eventCreator, setEventCreator] = useState<UserItem>();
    const userID = useAuthStore((state) => state.userID);

    useEffect(() => {
        console.log("here", userID);
        (async () => {
            const attendingStatus = await getAttendingStatus(event.ID);
            setIsAttendingEvent(attendingStatus);
            const eventAttendees = await getEventAttendees(event.ID);
            if (!eventAttendees) return;
            setEventAttendees(eventAttendees);

            const eventCreator = await getEventCreator(event.ID);
            if (!eventCreator) return;
            setEventCreator(eventCreator);
        })();
    }, [event.ID, isAttendingEvent]);

    const handleClickAttend = async () => {
        const success = await changeAttendingStatus(
            event.ID,
            isAttendingEvent ? "unattend" : "attend",
        );
        if (!success) return;
        setIsAttendingEvent((prev) => !prev);
    };

    const handleDelete = async () => {
        const success = await deleteEvent(event.ID);
        if (!success) return;
        window.location.replace("/");
    };

    return (
        <div className="absolute left-1/2 top-[120px] -translate-x-1/2 bg-foreground/80 p-[8px] rounded-[4px] text-background w-[80vw] max-w-[800px] min-w-[360px] text-[20px] flex flex-col gap-y-[4px]">
            <div className="flex justify-between items-center text-[20px] bg-foreground/80 px-[12px] py-[4px] rounded-[4px]">
                <span className="font-bold" aria-label="event title">
                    {event.Title}
                </span>
                <div className="flex items-center gap-x-[12px]">
                    {event.CreatorID === userID ? (
                        <svg
                            viewBox="0 0 1024 1024"
                            xmlns="http://www.w3.org/2000/svg"
                            fill="currentColor"
                            className="cursor-pointer hover:text-red-500 h-[20px]"
                            aria-label="delete event"
                            aria-hidden={event.CreatorID !== userID}
                            onClick={handleDelete}
                        >
                            <g id="SVGRepo_bgCarrier" strokeWidth="0"></g>
                            <g
                                id="SVGRepo_tracerCarrier"
                                strokeLinecap="round"
                                strokeLinejoin="round"
                            ></g>
                            <g id="SVGRepo_iconCarrier">
                                <path
                                    fill="currentColor"
                                    d="M160 256H96a32 32 0 0 1 0-64h256V95.936a32 32 0 0 1 32-32h256a32 32 0 0 1 32 32V192h256a32 32 0 1 1 0 64h-64v672a32 32 0 0 1-32 32H192a32 32 0 0 1-32-32V256zm448-64v-64H416v64h192zM224 896h576V256H224v640zm192-128a32 32 0 0 1-32-32V416a32 32 0 0 1 64 0v320a32 32 0 0 1-32 32zm192 0a32 32 0 0 1-32-32V416a32 32 0 0 1 64 0v320a32 32 0 0 1-32 32z"
                                ></path>
                            </g>
                        </svg>
                    ) : undefined}
                    <button
                        onClick={handleClose}
                        className="hover:text-accent duration-300 cursor-pointer"
                        aria-label="close event"
                    >
                        x
                    </button>
                </div>
            </div>

            <div className="flex flex-col gap-y-[4px] bg-foreground/70 px-[12px] py-[4px] rounded-[4px]">
                <span aria-label="event description">{event.Description}</span>

                <div className="flex flex-col">
                    <span aria-label="event start time">
                        Starts at:{" "}
                        {new Date(event.StartsAt).toLocaleString(
                            "en-UK",
                            dateFormat,
                        )}
                    </span>
                    <span aria-label="event end time">
                        Ends at:{" "}
                        {new Date(event.EndsAt).toLocaleString(
                            "en-UK",
                            dateFormat,
                        )}
                    </span>
                </div>

                <div className="flex items-center gap-x-[4px]">
                    <svg
                        width="28px"
                        height="28px"
                        viewBox="0 0 24 24"
                        className="text-foreground"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                        stroke=""
                    >
                        <g id="SVGRepo_bgCarrier" strokeWidth="0"></g>
                        <g
                            id="SVGRepo_tracerCarrier"
                            strokeLinecap="round"
                            strokeLinejoin="round"
                        ></g>
                        <g id="SVGRepo_iconCarrier">
                            <circle
                                cx="12"
                                cy="12"
                                r="9"
                                className="text-background"
                                stroke="currentColor"
                                strokeLinejoin="round"
                            ></circle>
                            <path
                                d="M12 3C12 3 8.5 6 8.5 12C8.5 18 12 21 12 21"
                                className="text-background"
                                stroke="currentColor"
                                strokeLinejoin="round"
                            ></path>
                            <path
                                d="M12 3C12 3 15.5 6 15.5 12C15.5 18 12 21 12 21"
                                className="text-background"
                                stroke="currentColor"
                                strokeLinejoin="round"
                            ></path>{" "}
                            <path
                                d="M3 12H21"
                                className="text-background"
                                stroke="currentColor"
                                strokeLinejoin="round"
                            ></path>
                            <path
                                d="M19.5 7.5H4.5"
                                className="text-background"
                                stroke="currentColor"
                                strokeLinejoin="round"
                            ></path>{" "}
                            <g filter="url(#filter0_d_15_556)">
                                <path
                                    d="M19.5 16.5H4.5"
                                    className="text-background"
                                    stroke="currentColor"
                                    strokeLinejoin="round"
                                ></path>
                            </g>
                        </g>
                    </svg>
                    <span aria-label="event location">{event.Address}</span>
                </div>
            </div>

            <div className="flex flex-col gap-y-[4px] bg-foreground/70 px-[12px] py-[4px] rounded-[4px]">
                <div className="flex gap-x-[4px]">
                    <span>Created by: </span>
                    <span className="bg-accent text-foreground text-center rounded-[4px] px-[4px]">
                        {eventCreator?.Name}
                    </span>
                </div>

                <div className="flex items-center gap-x-[4px]">
                    <span>Attending:</span>{" "}
                    <div className="flex gap-x-[4px] overflow-x-auto">
                        {eventAttendees?.map((attendee) => (
                            <span
                                className="bg-background text-foreground text-center rounded-[4px] px-[4px]"
                                key={attendee.ID}
                            >
                                {attendee.Name}
                            </span>
                        ))}
                    </div>
                </div>
            </div>

            <button
                onClick={handleClickAttend}
                className="bg-foreground/70 text-background rounded-[4px] py-[4px] cursor-pointer hover:bg-muted hover:text-white duration-300"
                aria-label={
                    isAttendingEvent ? "don't attend event" : "attend event"
                }
            >
                {isAttendingEvent ? "Attending Event" : "Attend Event"}
            </button>
        </div>
    );
};

export default Event;
