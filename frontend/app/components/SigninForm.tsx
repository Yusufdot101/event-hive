"use client";
import Link from "next/link";
import { useState } from "react";
import ShowHide from "./ShowHide";
import { signin } from "../utilities/signin";

const SigninForm = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const [showPassword, setShowPassword] = useState(false);

    const [error, setError] = useState("");

    const handleSubmit = async () => {
        if (
            !email ||
            !password ||
            !(password.length >= 8 && password.length <= 72)
        ) {
            return;
        }
        const success = await signin(
            (error: string) => setError(error),
            email,
            password,
        );

        if (!success) return;
    };

    return (
        <form
            onSubmit={(e) => {
                e.preventDefault();
                handleSubmit();
            }}
            className="flex flex-col gap-y-[8px] justify-center border-1 border-muted-foreground p-[20px] rounded-[8px] w-full max-w-[600px] min-[901]:text-[20px]"
        >
            <div className="flex flex-col w-full">
                <label htmlFor="email">Email</label>
                <input
                    id="email"
                    name="email"
                    type="email"
                    required
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    className="border-muted-foreground border-1 rounded-[4px] p-[8px] outline-none"
                    placeholder="your.email@example.com"
                />
            </div>

            <div className="flex flex-col w-full">
                <div className="flex justify-between">
                    <label htmlFor="password">Password</label>
                    <Link
                        href={"/forgotpassword"}
                        className="font-light text-muted-foreground cursor-pointer hover:text-foreground duration-300"
                    >
                        Forgot Password?
                    </Link>
                </div>
                <div className="relative">
                    <input
                        id="password"
                        name="password"
                        type={showPassword ? "text" : "password"}
                        required
                        minLength={8}
                        maxLength={72}
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        className="border-muted-foreground border-1 rounded-[4px] p-[8px] outline-none w-full pr-[40px]"
                        placeholder="Enter your password"
                    />
                    <ShowHide
                        show={showPassword}
                        handleClick={() => setShowPassword((prev) => !prev)}
                    />
                </div>
            </div>

            {error && (
                <div className="bg-red-500 text-white py-[4px] text-center rounded-[4px]">
                    {error}
                </div>
            )}

            <button className="bg-foreground text-background rounded-[4px] py-[4px] cursor-pointer hover:bg-muted hover:text-white duration-300 border-1 border-muted-foreground">
                Sign In
            </button>

            <section className="flex gap-x-[4px] w-full justify-center items-center">
                <p className="text-muted-foreground">
                    Don&apos;t have an account?{" "}
                </p>
                <Link href={"/signup"} className="font-bold">
                    Sign up
                </Link>
            </section>
        </form>
    );
};

export default SigninForm;
