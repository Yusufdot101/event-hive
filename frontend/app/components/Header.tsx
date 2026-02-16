"use client";
import Link from "next/link";
import { useEffect, useState } from "react";
import { useAuthStore } from "../store/useAuthStore";
import { refreshToken } from "../utilities/refreshToken";
import { useRouter } from "next/navigation";
import { logout } from "../utilities/logout";

function Header() {
    const [hidden, setHidden] = useState(false);

    const isLoggedIn = useAuthStore((state) => state.isLoggedIn);

    const router = useRouter();
    const navLinks = [
        {
            text: "Signup",
            hidden: isLoggedIn,
            onClick: () => router.push("/signup"),
        },
        {
            text: "Logout",
            hidden: !isLoggedIn,
            onClick: () => logout(),
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

    useEffect(() => {
        if (!isLoggedIn) {
            refreshToken();
        }
    }, [isLoggedIn]);

    return (
        <div
            className={`fixed w-full flex flex-col p-[12px] gap-y-[4px] top-0 z-10 transition-transform duration-300 ${hidden ? "-translate-y-full" : "translate-y-0"}`}
        >
            <div className="flex justify-between items-center">
                <Link
                    href="/"
                    className="cursor-pointer text-accent text-[20px]"
                >
                    EventHive
                </Link>
                <ul className="flex gap-x-[12px] items-center">
                    {navLinks.map((link) => (
                        <li
                            key={`${link.text}-${link.onClick}`}
                            hidden={link.hidden}
                            className="cursor-pointer hover:text-[110%] duration-300"
                            onClick={link.onClick}
                        >
                            {link.text}
                        </li>
                    ))}
                </ul>
            </div>
            <hr className="w-full opacity-50 border-[1px]" />
        </div>
    );
}

export default Header;
