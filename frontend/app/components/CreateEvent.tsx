"use client";
import { useState } from "react";
import { location } from "../utilities/api";
import { createEvent } from "../utilities/event";
import { useRouter } from "next/navigation";

type Props = {
    handleClose: () => void;
    handleClickSelectLocation: () => void;
    selectedLocation: location;
    selectedAddress: string;
};

const today = new Date().toISOString().slice(0, 10);

const CreateEvent = ({
    handleClose,
    handleClickSelectLocation,
    selectedLocation,
    selectedAddress,
}: Props) => {
    const [title, setTitle] = useState("");
    const [description, setDescription] = useState("");
    const [startDate, setStartDate] = useState("");
    const [startTime, setStartTime] = useState("");
    const [endDate, setEndDate] = useState("");
    const [endTime, setEndTime] = useState("");

    const addMinutesToTime = (time: string, duration: number): string => {
        const [h, m] = time.split(":").map(Number);
        const date = new Date(0, 0, 0, h, m);
        date.setMinutes(date.getMinutes() + duration);
        return date.toTimeString().slice(0, 5);
    };

    const handleSubmit = async () => {
        if (
            !title ||
            !description ||
            !startDate ||
            !startTime ||
            !endDate ||
            !endTime ||
            !selectedLocation ||
            !selectedAddress
        ) {
            alert("please fill in all the fields");
            return;
        }
        const success = await createEvent(
            title,
            description,
            startDate,
            startTime,
            endDate,
            endTime,
            selectedAddress,
            selectedLocation,
        );

        if (!success) return;
        window.location.reload();
        handleClose();
    };

    return (
        <form
            onKeyDown={(e) => {
                if (e.code === "Enter") {
                    e.preventDefault();
                }
            }}
            onSubmit={(e) => {
                e.preventDefault();
            }}
            className="bg-background/80 flex flex-col gap-y-[8px] justify-center border-1 border-muted-foreground p-[20px] rounded-[8px] w-full max-w-[600px] min-[901]:text-[20px] absolute top-1/2 left-1/2 -translate-1/2 h-[full] z-10"
        >
            <div className="flex flex-col w-full">
                <label htmlFor="password">Event location</label>
                <div className="flex items-center gap-x-[4px]">
                    <input
                        id="eventLocation"
                        name="eventLocation"
                        type="text"
                        value={selectedAddress}
                        readOnly
                        required
                        className="border-muted-foreground border-1 rounded-[4px] p-[8px] outline-none w-full pr-[40px] cursor-pointer"
                        placeholder="click event location"
                        onClick={handleClickSelectLocation}
                    />
                </div>
            </div>

            <div className="flex flex-col w-full">
                <label htmlFor="title">Title</label>
                <input
                    id="title"
                    name="title"
                    type="text"
                    required
                    value={title}
                    onChange={(e) => {
                        setTitle(e.target.value);
                    }}
                    minLength={2}
                    className="border-muted-foreground border-1 rounded-[4px] p-[8px] outline-none"
                    placeholder="choose event title"
                />
            </div>

            <div className="flex flex-col w-full">
                <label htmlFor="desciption">Description</label>
                <textarea
                    id="desciption"
                    name="desciption"
                    required
                    value={description}
                    onChange={(e) => {
                        setDescription(e.target.value);
                    }}
                    minLength={2}
                    className="border-muted-foreground border-1 rounded-[4px] p-[8px] outline-none min-h-[50px]"
                    placeholder="event desciption"
                />
            </div>

            <div className="flex flex-col w-full">
                <label htmlFor="startDate">Start date and time</label>
                <div className="flex items-center gap-x-[4px]">
                    <input
                        id="startDate"
                        name="startDate"
                        type="date"
                        required
                        value={startDate}
                        onChange={(e) => {
                            setStartDate(e.target.value);
                        }}
                        min={today}
                        className="border-muted-foreground border-1 rounded-[4px] p-[8px] outline-none w-full pr-[40px]"
                        placeholder="Create a strong password"
                    />
                    <input
                        id="startTime"
                        name="startTime"
                        type="time"
                        required
                        value={startTime}
                        onChange={(e) => {
                            setStartTime(e.target.value);
                        }}
                        className="border-muted-foreground border-1 rounded-[4px] p-[8px] outline-none w-full pr-[40px]"
                        placeholder="Create a strong password"
                    />
                </div>
            </div>

            <div className="flex flex-col w-full">
                <label htmlFor="endDate">End date and time</label>
                <div className="flex items-center gap-x-[4px]">
                    <input
                        id="endDate"
                        name="endDate"
                        type="date"
                        required
                        value={endDate}
                        onChange={(e) => {
                            setEndDate(e.target.value);
                        }}
                        min={startDate}
                        className="border-muted-foreground border-1 rounded-[4px] p-[8px] outline-none w-full pr-[40px]"
                        placeholder="Create a strong password"
                    />
                    <input
                        id="endTime"
                        name="endTime"
                        type="time"
                        required
                        value={endTime}
                        onChange={(e) => {
                            setEndTime(e.target.value);
                        }}
                        min={
                            startDate === endDate
                                ? addMinutesToTime(startTime, 15)
                                : ""
                        }
                        className="border-muted-foreground border-1 rounded-[4px] p-[8px] outline-none w-full pr-[40px]"
                        placeholder="Create a strong password"
                    />
                </div>
            </div>

            <div className="flex gap-x-[4px]">
                <button
                    className="w-full bg-red-500 text-white rounded-[4px] py-[4px] cursor-pointer hover:bg-muted hover:text-white duration-300 border-1 border-muted-foreground"
                    onClick={handleClose}
                >
                    Discard
                </button>
                <button
                    className="w-full bg-accent text-white rounded-[4px] py-[4px] cursor-pointer hover:bg-muted hover:text-white duration-300 border-1 border-muted-foreground"
                    onClick={handleSubmit}
                >
                    Create event
                </button>
            </div>
        </form>
    );
};

export default CreateEvent;
