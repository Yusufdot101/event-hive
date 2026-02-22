import type { EventItem } from "../utilities/event";
import CustomMarker from "./CustomMarker";

type Props = {
    event: EventItem;
    onClick: (event: EventItem) => void;
};

const EventMarker = ({ event, onClick }: Props) => {
    return (
        <CustomMarker
            lng={event.Longitude}
            lat={event.Latitude}
            style={{
                transform: "translate(-30px, -52px)",
                width: "64px",
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                cursor: "pointer",
            }}
            onClick={() => onClick(event)}
        >
            <span className="font-bold text-center">{event.Title}</span>
            <svg
                className="text-accent"
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
    );
};

export default EventMarker;
