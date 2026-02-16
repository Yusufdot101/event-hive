"use client";
import Link from "next/link";
import { useState } from "react";
import { signup } from "../utilities/signup";
import ShowHide from "./ShowHide";

const SignupForm = () => {
    const [username, setUsername] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");

    const [showPassword, setShowPassword] = useState(false);
    const [showConfirmPassword, setShowConfirmPassword] = useState(false);

    const [error, setError] = useState("");

    const handleSubmit = async () => {
        if (password !== confirmPassword) {
            return;
        }
        if (
            !username ||
            !email ||
            !password ||
            !confirmPassword ||
            !(password.length >= 8 && password.length <= 72)
        ) {
            return;
        }
        const success = await signup(
            (error: string) => setError(error),
            username,
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
                <label htmlFor="username">Username</label>
                <input
                    id="username"
                    name="username"
                    type="text"
                    minLength={2}
                    value={username}
                    onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
                        setUsername(e.target.value)
                    }
                    required
                    min={2}
                    className="border-muted-foreground border-1 rounded-[4px] p-[8px] outline-none"
                    placeholder="Choose a username"
                />
            </div>

            <div className="flex flex-col w-full">
                <label htmlFor="email">Email</label>
                <input
                    id="email"
                    name="email"
                    type="email"
                    value={email}
                    onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
                        setEmail(e.target.value)
                    }
                    required
                    className="border-muted-foreground border-1 rounded-[4px] p-[8px] outline-none"
                    placeholder="your.email@example.com"
                />
            </div>

            <div className="flex flex-col w-full">
                <label htmlFor="password">Password</label>
                <div className="relative">
                    <input
                        id="password"
                        name="password"
                        type={showPassword ? "text" : "password"}
                        value={password}
                        onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
                            setPassword(e.target.value)
                        }
                        required
                        minLength={8}
                        maxLength={72}
                        className="border-muted-foreground border-1 rounded-[4px] p-[8px] outline-none w-full pr-[40px]"
                        placeholder="Create a strong password"
                    />
                    <ShowHide
                        show={showPassword}
                        handleClick={() => setShowPassword((prev) => !prev)}
                    />
                </div>
            </div>

            <div className="flex flex-col w-full">
                <label htmlFor="confirmPassword">Confirm Password</label>
                <div className="relative">
                    <input
                        id="confirmPassword"
                        name="confirmPassword"
                        type={showConfirmPassword ? "text" : "password"}
                        value={confirmPassword}
                        onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
                            setConfirmPassword(e.target.value)
                        }
                        required
                        minLength={8}
                        maxLength={72}
                        pattern={password}
                        title="passwords must match"
                        className="border-muted-foreground border-1 rounded-[4px] p-[8px] outline-none w-full pr-[40px]"
                        placeholder="Re-enter your password"
                    />
                    <ShowHide
                        show={showConfirmPassword}
                        handleClick={() =>
                            setShowConfirmPassword((prev) => !prev)
                        }
                    />
                </div>
            </div>

            {error && (
                <div className="bg-red-500 text-white py-[4px] text-center rounded-[4px]">
                    {error}
                </div>
            )}

            <button
                className="bg-foreground text-background rounded-[4px] py-[4px] cursor-pointer hover:bg-muted hover:text-white duration-300 border-1 border-muted-foreground"
                onClick={handleSubmit}
            >
                Create account
            </button>

            <section className="flex gap-x-[4px] w-full justify-center items-center">
                <p className="text-muted-foreground">
                    Already have an account?{" "}
                </p>
                <Link href={"/signin"} className="font-bold">
                    Sign In
                </Link>
            </section>
        </form>
    );
};

export default SignupForm;
