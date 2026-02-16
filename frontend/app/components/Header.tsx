"use client";
import Link from "next/link";
import { useEffect, useState } from "react";
import { useAuthStore } from "../store/useAuthStore";

function Header() {
    const [hidden, setHidden] = useState(false);

    const isLoggedIn = useAuthStore((state) => state.isLoggedIn);

    const navLinks = [
        {
            text: "Signup",
            href: "/signup",
            hidden: isLoggedIn,
        },
        {
            text: "Logout",
            href: "/logout",
            hidden: !isLoggedIn,
        },
    ];

    useEffect(() => {
        const lastY = window.scrollY;
        const onScroll = () => {
            const y = window.scrollY;
            setHidden(y > lastY && y > 80);
        };

        window.addEventListener("scroll", onScroll);
        return () => window.removeEventListener("scroll", onScroll);
    }, []);

    return (
        <div
            className={`fixed w-full flex flex-col p-[12px] gap-y-[4px] top-0 z-10 transition-transform duration-300 ${hidden ? "-translate-y-full" : "translate-y-0"}`}
        >
            <div className="flex justify-between items-center">
                <a href="/" className="cursor-pointer text-accent text-[20px]">
                    EventHive
                </a>
                <ul className="flex gap-x-[12px] items-center">
                    {navLinks.map((link) => (
                        <li key={`${link.text}-${link.href}`}>
                            <Link
                                href={link.href}
                                className="cursor-pointer hover:text-[110%] duration-300"
                                hidden={link.hidden}
                            >
                                {link.text}
                            </Link>
                        </li>
                    ))}
                </ul>
            </div>
            <hr className="w-full opacity-50 border-[1px]" />
        </div>
    );
}

export default Header;
